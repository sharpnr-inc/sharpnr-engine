package sentinelAggregatorService_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	sentinelAggregatorContracts "sharpnr-engine/sentinel_aggregators/contracts"
	sentinelAggregatorService "sharpnr-engine/sentinel_aggregators/service"
	sentinelAggregatorTypes "sharpnr-engine/sentinel_aggregators/types"

	"github.com/google/uuid"
)

// stubProvider registers whatever it is given.
type stubProvider struct {
	regs []sentinelAggregatorContracts.ResourceRegistration
}

func (s stubProvider) Resources() []sentinelAggregatorContracts.ResourceRegistration { return s.regs }

func resource(rtype string, actions []string, attrs ...sentinelAggregatorContracts.AttributeDescriptor) sentinelAggregatorContracts.ResourceRegistration {
	ads := make([]sentinelAggregatorContracts.ActionDescriptor, 0, len(actions))
	for _, a := range actions {
		ads = append(ads, sentinelAggregatorContracts.ActionDescriptor{Name: a, Label: a})
	}
	return sentinelAggregatorContracts.ResourceRegistration{
		Descriptor: sentinelAggregatorContracts.ResourceDescriptor{
			Type: rtype, Label: rtype, Actions: ads, Attributes: attrs,
		},
	}
}

// mustPanic asserts a wiring bug fails the boot, and that the message explains why.
func mustPanic(t *testing.T, wantSubstr string, fn func()) {
	t.Helper()
	defer func() {
		r := recover()
		if r == nil {
			t.Fatalf("expected a bootstrap panic containing %q, got none", wantSubstr)
		}
		msg, _ := r.(string)
		if !strings.Contains(msg, wantSubstr) {
			t.Fatalf("panic message %q does not mention %q", msg, wantSubstr)
		}
	}()
	fn()
}

/* -------------------------------------------------------------------------- */
/*                          The collision that matters                        */
/* -------------------------------------------------------------------------- */

// Cedar action ids are global: an id shared by two resources means a policy
// scoped to that action reaches both. This must never reach production.
func TestDuplicateActionAcrossModulesPanics(t *testing.T) {
	mustPanic(t, "registered by both", func() {
		sentinelAggregatorService.NewSentinelAggregatorService(
			stubProvider{[]sentinelAggregatorContracts.ResourceRegistration{
				resource("Property", []string{"canView"}),
			}},
			stubProvider{[]sentinelAggregatorContracts.ResourceRegistration{
				resource("Invoice", []string{"canView"}), // same id, different module
			}},
		)
	})
}

func TestDuplicateActionWithinOneResourcePanics(t *testing.T) {
	mustPanic(t, "declares action", func() {
		sentinelAggregatorService.NewSentinelAggregatorService(
			stubProvider{[]sentinelAggregatorContracts.ResourceRegistration{
				resource("Property", []string{"canViewProperty", "canViewProperty"}),
			}},
		)
	})
}

// Resource-suffixed names — the convention the panic message recommends — must
// coexist happily.
func TestResourceSuffixedActionsDoNotCollide(t *testing.T) {
	svc := sentinelAggregatorService.NewSentinelAggregatorService(
		stubProvider{[]sentinelAggregatorContracts.ResourceRegistration{
			resource("Property", []string{"canViewProperty", "canEditProperty"}),
			resource("Invoice", []string{"canViewInvoice", "canEditInvoice"}),
		}},
	)
	got, err := svc.ListResources(context.Background(), uuid.New(), uuid.New())
	if err != nil || len(got) != 2 {
		t.Fatalf("want 2 resources, got %d (err %v)", len(got), err)
	}
}

/* -------------------------------------------------------------------------- */
/*                            Other wiring bugs                               */
/* -------------------------------------------------------------------------- */

func TestDuplicateResourceTypePanics(t *testing.T) {
	mustPanic(t, "registered twice", func() {
		sentinelAggregatorService.NewSentinelAggregatorService(
			stubProvider{[]sentinelAggregatorContracts.ResourceRegistration{resource("Property", []string{"a"})}},
			stubProvider{[]sentinelAggregatorContracts.ResourceRegistration{resource("Property", []string{"b"})}},
		)
	})
}

func TestEmptyResourceTypePanics(t *testing.T) {
	mustPanic(t, "empty Type", func() {
		sentinelAggregatorService.NewSentinelAggregatorService(
			stubProvider{[]sentinelAggregatorContracts.ResourceRegistration{resource("", []string{"a"})}},
		)
	})
}

func TestEmptyActionNamePanics(t *testing.T) {
	mustPanic(t, "empty Name", func() {
		sentinelAggregatorService.NewSentinelAggregatorService(
			stubProvider{[]sentinelAggregatorContracts.ResourceRegistration{resource("Property", []string{""})}},
		)
	})
}

// A mistyped attribute type renders a broken picker with no error anywhere —
// catch it at boot instead.
func TestInvalidAttributeTypePanics(t *testing.T) {
	mustPanic(t, "valid types are", func() {
		sentinelAggregatorService.NewSentinelAggregatorService(
			stubProvider{[]sentinelAggregatorContracts.ResourceRegistration{
				resource("Property", []string{"canViewProperty"},
					sentinelAggregatorContracts.AttributeDescriptor{Name: "active", Type: "boolean"}), // not "bool"
			}},
		)
	})
}

func TestDuplicateAttributePanics(t *testing.T) {
	mustPanic(t, "declares attribute", func() {
		sentinelAggregatorService.NewSentinelAggregatorService(
			stubProvider{[]sentinelAggregatorContracts.ResourceRegistration{
				resource("Property", []string{"canViewProperty"},
					sentinelAggregatorContracts.AttributeDescriptor{Name: "city", Type: sentinelAggregatorTypes.AttrString},
					sentinelAggregatorContracts.AttributeDescriptor{Name: "city", Type: sentinelAggregatorTypes.AttrString}),
			}},
		)
	})
}

/* -------------------------------------------------------------------------- */
/*                                   Reads                                    */
/* -------------------------------------------------------------------------- */

func TestGetResourceDetailsUsesTheBoundLister(t *testing.T) {
	called := 0
	reg := resource("Property", []string{"canViewProperty"})
	reg.ListInstances = func(ctx context.Context, orgID, emplID uuid.UUID) ([]sentinelAggregatorContracts.ResourceInstance, error) {
		called++
		return []sentinelAggregatorContracts.ResourceInstance{{ID: "p1", Label: "Mall"}}, nil
	}
	svc := sentinelAggregatorService.NewSentinelAggregatorService(stubProvider{[]sentinelAggregatorContracts.ResourceRegistration{reg}})

	view, err := svc.GetResourceDetails(context.Background(), uuid.New(), uuid.New(), "Property")
	if err != nil {
		t.Fatal(err)
	}
	if called != 1 {
		t.Fatalf("lister called %d times, want 1", called)
	}
	if !view.HasInstances {
		t.Error("HasInstances must be true when a lister is registered")
	}
	if len(view.Instances) != 1 || view.Instances[0].Label != "Mall" {
		t.Fatalf("instances not returned: %+v", view.Instances)
	}
	if len(view.Actions) != 1 {
		t.Fatalf("descriptor not carried into the view: %+v", view.ResourceDescriptor)
	}
}

// nil lister = "no selectable instances", distinct from "none in this org".
func TestNilListerMeansTypeLevelOnly(t *testing.T) {
	svc := sentinelAggregatorService.NewSentinelAggregatorService(
		stubProvider{[]sentinelAggregatorContracts.ResourceRegistration{resource("Report", []string{"canViewReport"})}},
	)
	view, err := svc.GetResourceDetails(context.Background(), uuid.New(), uuid.New(), "Report")
	if err != nil {
		t.Fatal(err)
	}
	if view.HasInstances {
		t.Error("HasInstances must be false when no lister is registered")
	}
	if view.Instances == nil {
		t.Error("Instances must serialize as [] not null")
	}
}

func TestUnknownResourceTypeIsTyped(t *testing.T) {
	svc := sentinelAggregatorService.NewSentinelAggregatorService(
		stubProvider{[]sentinelAggregatorContracts.ResourceRegistration{resource("Property", []string{"canViewProperty"})}},
	)
	_, err := svc.GetResourceDetails(context.Background(), uuid.New(), uuid.New(), "Nope")
	if !errors.Is(err, sentinelAggregatorService.ErrUnknownResourceType) {
		t.Fatalf("want ErrUnknownResourceType (so controllers answer 404), got %v", err)
	}
}

// Catalog order must be stable across calls so the FE picker does not reshuffle.
func TestCatalogOrderIsRegistrationOrder(t *testing.T) {
	svc := sentinelAggregatorService.NewSentinelAggregatorService(
		stubProvider{[]sentinelAggregatorContracts.ResourceRegistration{
			resource("Zebra", []string{"canViewZebra"}),
			resource("Apple", []string{"canViewApple"}),
			resource("Mango", []string{"canViewMango"}),
		}},
	)
	want := []string{"Zebra", "Apple", "Mango"}
	for i := 0; i < 5; i++ {
		got, _ := svc.ListResources(context.Background(), uuid.New(), uuid.New())
		for j, w := range want {
			if got[j].Type != w {
				t.Fatalf("call %d: position %d = %q, want %q", i, j, got[j].Type, w)
			}
		}
	}
}
