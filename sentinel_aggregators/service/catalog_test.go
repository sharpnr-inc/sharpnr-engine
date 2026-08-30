package sentinelAggregatorService_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	agg "sharpnr-engine/sentinel_aggregators/contracts"
	sentinelAggregatorService "sharpnr-engine/sentinel_aggregators/service"
	aggTypes "sharpnr-engine/sentinel_aggregators/types"

	"github.com/google/uuid"
)

// stubProvider registers whatever it is given.
type stubProvider struct{ regs []agg.ResourceRegistration }

func (p stubProvider) Resources() []agg.ResourceRegistration { return p.regs }

func provide(regs ...agg.ResourceRegistration) stubProvider {
	return stubProvider{regs: regs}
}

// mustPanic asserts that building the catalog fails at bootstrap, and that the
// message says why. Every validation in the aggregator panics on purpose: the
// catalog is assembled once at boot, so a failed start is loud where a silently
// dead permission rule is not.
func mustPanic(t *testing.T, wantSubstring string, providers ...agg.ResourceProvider) {
	t.Helper()
	defer func() {
		r := recover()
		if r == nil {
			t.Fatalf("expected a panic containing %q, got none", wantSubstring)
		}
		msg, _ := r.(string)
		if !strings.Contains(msg, wantSubstring) {
			t.Fatalf("panic %q does not mention %q", msg, wantSubstring)
		}
	}()
	sentinelAggregatorService.NewSentinelAggregatorService(providers...)
}

/* -------------------------------------------------------------------------- */
/*                                  Builder                                   */
/* -------------------------------------------------------------------------- */

// Standard() must generate the four CRUD actions with names that follow
// can<Verb><ResourceType> and the right RequiresInstance flags. Those flags are
// not cosmetic: the policy generator uses them to decide which actions can be
// pinned to an instance and which must stay unscoped.
func TestStandardGeneratesCRUDActions(t *testing.T) {
	svc := sentinelAggregatorService.NewSentinelAggregatorService(
		provide(agg.Standard("Widget", "Widgets").Category("Ops").Build()),
	)

	catalog, err := svc.GetCatalog(context.Background(), uuid.New(), uuid.New())
	if err != nil {
		t.Fatal(err)
	}
	if len(catalog) != 1 {
		t.Fatalf("expected 1 resource, got %d", len(catalog))
	}

	want := map[string]bool{ // action id -> RequiresInstance
		"canViewWidget":   true,
		"canCreateWidget": false,
		"canEditWidget":   true,
		"canDeleteWidget": true,
	}
	if len(catalog[0].Actions) != len(want) {
		t.Fatalf("expected %d actions, got %d", len(want), len(catalog[0].Actions))
	}
	for _, a := range catalog[0].Actions {
		instance, ok := want[a.Name]
		if !ok {
			t.Errorf("unexpected action %q", a.Name)
			continue
		}
		if a.RequiresInstance != instance {
			t.Errorf("%s: RequiresInstance = %v, want %v", a.Name, a.RequiresInstance, instance)
		}
		if a.Label == "" {
			t.Errorf("%s: generated action has no label", a.Name)
		}
	}
}

// An attribute must arrive with its Cedar path and a usable operator set, or
// the builder cannot render a restriction row for it.
func TestAttributesCarryCedarPathAndOperators(t *testing.T) {
	svc := sentinelAggregatorService.NewSentinelAggregatorService(
		provide(agg.Standard("Widget", "Widgets").
			Category("Ops").
			WithAttributes(agg.Attr("city", aggTypes.AttrString, "City")).
			ResolveFacts(func(context.Context, uuid.UUID, []string) (map[string]agg.ResourceFacts, error) {
				return nil, nil
			}).
			Build()),
	)

	catalog, _ := svc.GetCatalog(context.Background(), uuid.New(), uuid.New())
	attr := catalog[0].Attributes[0]
	if attr.Cedar != "resource.city" {
		t.Errorf("Cedar = %q, want resource.city", attr.Cedar)
	}
	if len(attr.Operators) == 0 {
		t.Error("attribute offers no operators, so no row can be built for it")
	}
}

/* -------------------------------------------------------------------------- */
/*                            Bootstrap validation                            */
/* -------------------------------------------------------------------------- */

// The rule the whole contract exists for: anything declared must be answerable.
// An attribute with no resolver produces a condition that reads as active and
// matches nothing, because Cedar skips a policy whose attribute is missing.
func TestAttributesWithoutResolverPanic(t *testing.T) {
	mustPanic(t, "ResolveFacts",
		provide(agg.Standard("Widget", "Widgets").
			Category("Ops").
			WithAttributes(agg.Attr("city", aggTypes.AttrString, "City")).
			Build()),
	)
}

func TestHousedInWithoutResolverPanics(t *testing.T) {
	mustPanic(t, "ResolveContainer",
		provide(
			agg.Standard("Container", "Containers").Category("Ops").Build(),
			agg.ResourceRegistration{
				Descriptor: agg.ResourceDescriptor{
					Type: "Widget", Label: "Widgets", Category: "Ops",
					Actions: []agg.ActionDescriptor{{Name: "canViewWidget", Label: "View"}},
				},
				HousedIn: "Container",
			},
		),
	)
}

func TestHousedInUnknownTypePanics(t *testing.T) {
	mustPanic(t, "which no module registered",
		provide(agg.Standard("Widget", "Widgets").
			Category("Ops").
			HousedIn("Nowhere", func(context.Context, uuid.UUID, []string) (map[string]string, error) {
				return nil, nil
			}).
			Build()),
	)
}

func TestMissingCategoryPanics(t *testing.T) {
	mustPanic(t, "Category", provide(agg.Standard("Widget", "Widgets").Build()))
}

// Cedar action ids are a flat global namespace with no resource qualifier, so
// two modules registering the same id means a policy meant for one silently
// reaches the other. The aggregator is the only place that sees both.
func TestCrossResourceActionCollisionPanics(t *testing.T) {
	dup := agg.ResourceDescriptor{
		Type: "Beta", Label: "Betas", Category: "Ops",
		Actions: []agg.ActionDescriptor{{Name: "canViewAlpha", Label: "View"}},
	}
	mustPanic(t, "registered by both",
		provide(
			agg.Standard("Alpha", "Alphas").Category("Ops").Build(),
			agg.ResourceRegistration{Descriptor: dup},
		),
	)
}

func TestActionNamingConventionEnforced(t *testing.T) {
	mustPanic(t, "does not contain the resource type",
		provide(agg.Custom("Widget", "Widgets").
			Category("Ops").
			WithActions(agg.Action("canFrobulate", "Frobulate")).
			Build()),
	)
}

// Two ids predate the convention and are already embedded in stored policy
// text; renaming them would silently break live policies, so they are exempt
// rather than fixed.
func TestGrandfatheredActionsAllowed(t *testing.T) {
	svc := sentinelAggregatorService.NewSentinelAggregatorService(
		provide(agg.Standard("Property", "Properties").
			Category("Ops").
			WithActions(agg.Action("canAddBusinessModules", "Set up business services").OnInstance()).
			Build()),
	)
	if err := svc.ValidateActions([]string{"canAddBusinessModules"}); err != nil {
		t.Fatalf("grandfathered action rejected: %v", err)
	}
}

/* -------------------------------------------------------------------------- */
/*                              ValidateActions                               */
/* -------------------------------------------------------------------------- */

// This is what makes an over-broad grant unrepresentable rather than merely
// forbidden: only ids a module declared can reach stored policy text.
func TestValidateActionsRejectsUnknown(t *testing.T) {
	svc := sentinelAggregatorService.NewSentinelAggregatorService(
		provide(agg.Standard("Widget", "Widgets").Category("Ops").Build()),
	)
	if err := svc.ValidateActions([]string{"canViewWidget"}); err != nil {
		t.Fatalf("known action rejected: %v", err)
	}
	err := svc.ValidateActions([]string{"canViewWidget", "canDoAnythingIWant"})
	if !errors.Is(err, sentinelAggregatorService.ErrUnknownAction) {
		t.Fatalf("expected ErrUnknownAction, got %v", err)
	}
}

/* -------------------------------------------------------------------------- */
/*                                 Ancestry                                   */
/* -------------------------------------------------------------------------- */

// Room -> Hotel -> Property. The aggregator must return BOTH ancestors, not
// just the direct one: Cedar's `in` is transitive only over entities present in
// the EntityMap, so a gap in the chain makes a property-scoped policy deny a
// room silently while one-hop checks still pass.
func TestResolveAncestryWalksTheWholeChain(t *testing.T) {
	roomToHotel := map[string]string{"room-1": "hotel-1", "room-2": "hotel-2"}
	hotelToProperty := map[string]string{"hotel-1": "prop-1", "hotel-2": "prop-1"}

	svc := sentinelAggregatorService.NewSentinelAggregatorService(
		provide(
			agg.Standard("Property", "Properties").Category("Ops").Build(),
			agg.Standard("Hotel", "Hotels").Category("Ops").
				HousedIn("Property", lookup(hotelToProperty)).Build(),
			agg.Standard("Room", "Rooms").Category("Ops").
				HousedIn("Hotel", lookup(roomToHotel)).Build(),
		),
	)

	got, err := svc.ResolveAncestry(context.Background(), uuid.New(), "Room", []string{"room-1", "room-2"})
	if err != nil {
		t.Fatal(err)
	}

	want := []sentinelAggregatorService.ResourceRef{
		{Type: "Hotel", ID: "hotel-1"},
		{Type: "Property", ID: "prop-1"},
	}
	if len(got["room-1"]) != len(want) {
		t.Fatalf("room-1 ancestry = %v, want %v", got["room-1"], want)
	}
	for i, ref := range want {
		if got["room-1"][i] != ref {
			t.Errorf("room-1 ancestry[%d] = %v, want %v", i, got["room-1"][i], ref)
		}
	}

	// Both rooms share prop-1 through different hotels — the batching must not
	// merge or lose either one's chain.
	if len(got["room-2"]) != 2 || got["room-2"][0].ID != "hotel-2" || got["room-2"][1].ID != "prop-1" {
		t.Errorf("room-2 ancestry = %v", got["room-2"])
	}
}

// A top-level type has no chain. Organization is stamped by the engine from the
// request's org, so no walk ever has to reach it.
func TestResolveAncestryOfTopLevelTypeIsEmpty(t *testing.T) {
	svc := sentinelAggregatorService.NewSentinelAggregatorService(
		provide(agg.Standard("Property", "Properties").Category("Ops").Build()),
	)
	got, err := svc.ResolveAncestry(context.Background(), uuid.New(), "Property", []string{"prop-1"})
	if err != nil {
		t.Fatal(err)
	}
	if len(got["prop-1"]) != 0 {
		t.Errorf("expected no ancestors, got %v", got["prop-1"])
	}
}

// An instance whose container cannot be resolved — deleted, or invisible to the
// caller — must be absent rather than an error. A missing ancestry means
// container-scoped policies do not match, which is the safe direction.
func TestResolveAncestryToleratesMissingContainer(t *testing.T) {
	svc := sentinelAggregatorService.NewSentinelAggregatorService(
		provide(
			agg.Standard("Property", "Properties").Category("Ops").Build(),
			agg.Standard("Hotel", "Hotels").Category("Ops").
				HousedIn("Property", lookup(map[string]string{"hotel-1": "prop-1"})).Build(),
		),
	)
	got, err := svc.ResolveAncestry(context.Background(), uuid.New(), "Hotel", []string{"hotel-1", "hotel-gone"})
	if err != nil {
		t.Fatal(err)
	}
	if len(got["hotel-1"]) != 1 {
		t.Errorf("hotel-1 lost its ancestry: %v", got["hotel-1"])
	}
	if len(got["hotel-gone"]) != 0 {
		t.Errorf("unresolvable instance got ancestry: %v", got["hotel-gone"])
	}
}

func lookup(m map[string]string) agg.ContainerResolver {
	return func(_ context.Context, _ uuid.UUID, ids []string) (map[string]string, error) {
		out := make(map[string]string, len(ids))
		for _, id := range ids {
			if parent, ok := m[id]; ok {
				out[id] = parent
			}
		}
		return out, nil
	}
}
