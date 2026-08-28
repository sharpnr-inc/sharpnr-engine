package sentinelservice_test

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	models "sharpnr-engine/sentinel/models"
	sentinel_repo "sharpnr-engine/sentinel/repo"
	sentinelEngine "sharpnr-engine/sentinel_engine"

	"github.com/google/uuid"
)

// Roles and groups must be identified by ID, not name: every attach endpoint
// takes an id, groups have no unique index on (org, name), and two employees
// can share a display name.
func TestPrincipalRefsCarryIDsNotNames(t *testing.T) {
	svc, repo := newService(t)
	ctx := context.Background()
	org := uuid.New()

	role := &models.Role{OrganizationID: org, Name: "Auditor"}
	if err := repo.CreateRole(ctx, role); err != nil {
		t.Fatal(err)
	}
	group := &models.Group{OrganizationID: org, Name: "Stakeholders"}
	if err := repo.CreateGroup(ctx, group); err != nil {
		t.Fatal(err)
	}

	view, err := svc.GetPrincipals(ctx, org)
	if err != nil {
		t.Fatal(err)
	}
	if len(view.Roles) != 1 || view.Roles[0].Value != role.ID.String() {
		t.Errorf("role ref value = %q, want the id %s", view.Roles[0].Value, role.ID)
	}
	if view.Roles[0].Label != "Auditor" {
		t.Errorf("role ref label = %q, want the name", view.Roles[0].Label)
	}
	if len(view.Groups) != 1 || view.Groups[0].Value != group.ID.String() {
		t.Errorf("group ref value = %q, want the id %s", view.Groups[0].Value, group.ID)
	}

	// the value must round-trip into the attach call the builder will make
	roleID, err := uuid.Parse(view.Roles[0].Value)
	if err != nil {
		t.Fatalf("role ref value is not a uuid: %v", err)
	}
	policy := &models.SentinelPolicy{OrganizationID: org, Name: "p-" + uuid.NewString(),
		Policy: "permit(principal,action,resource);", Enabled: true}
	if err := repo.CreatePolicy(ctx, policy); err != nil {
		t.Fatal(err)
	}
	if err := repo.AttachPolicyToRole(ctx, org, roleID, policy.ID); err != nil {
		t.Fatalf("the value the palette returned could not be used to attach: %v", err)
	}
}

// The sentinel service still cannot resolve users — Users stays empty here and
// is filled by the orchestrator.
func TestGetPrincipalsLeavesUsersToTheOrchestrator(t *testing.T) {
	svc, _ := newService(t)
	view, err := svc.GetPrincipals(context.Background(), uuid.New())
	if err != nil {
		t.Fatal(err)
	}
	if len(view.Users) != 0 {
		t.Errorf("sentinel has no access to the employee module; Users must be empty, got %d", len(view.Users))
	}
	raw, _ := json.Marshal(view)
	var decoded map[string]json.RawMessage
	_ = json.Unmarshal(raw, &decoded)
	for _, key := range []string{"roles", "groups", "users"} {
		raw, ok := decoded[key]
		if !ok {
			t.Errorf("the palette must always carry a %q key", key)
			continue
		}
		if string(raw) == "null" {
			t.Errorf("%q serialized as null; the FE would crash mapping over it", key)
		}
	}
	t.Logf("sentinel half of the palette: %s", raw)
}

// The palette advertises User with `==`. Now that it hands back real employee
// ids, does that form actually match?
func TestUserEqualsPolicyMatchesWithThePaletteID(t *testing.T) {
	dsnRepo := func() sentinel_repo.SentinelRepositoryInterface { _, r := newService(t); return r }()
	engine := sentinelEngine.NewSentinelEngine(dsnRepo.(sentinelEngine.PolicyStore))
	ctx := context.Background()
	org, user := uuid.New(), uuid.New()

	policy := &models.SentinelPolicy{
		OrganizationID: org, Name: "p-" + uuid.NewString(),
		Policy:  `permit(principal == User::"` + user.String() + `", action, resource);`,
		Enabled: true,
	}
	if err := dsnRepo.CreatePolicy(ctx, policy); err != nil {
		t.Fatal(err)
	}
	if err := dsnRepo.AttachPolicyToUser(ctx, org, user, policy.ID); err != nil {
		t.Fatal(err)
	}

	allowed, err := engine.IsAuthorized(ctx, sentinelEngine.AccessRequest{
		OrgID: org, UserID: user, Action: "canViewProperty",
		Resource: sentinelEngine.Resource{Type: "Property", ID: uuid.New().String()},
	})
	if err != nil {
		t.Fatalf("engine: %v", err)
	}
	if !allowed {
		t.Error(`principal == User::"<id>" did not match — the palette advertises "==" for User`)
	}

	// and it must NOT match a different employee
	other, err := engine.IsAuthorized(ctx, sentinelEngine.AccessRequest{
		OrgID: org, UserID: uuid.New(), Action: "canViewProperty",
		Resource: sentinelEngine.Resource{Type: "Property", ID: uuid.New().String()},
	})
	if err != nil {
		t.Fatal(err)
	}
	if other {
		t.Error(`principal == User::"<id>" matched a DIFFERENT employee`)
	}
}

/* -------------------------------------------------------------------------- */
/*        Every advertised principal type must actually be evaluable          */
/* -------------------------------------------------------------------------- */

// grantsWith registers a policy against the user directly and asks the engine
// whether it grants. The direct register isolates the policy TEXT as the only
// variable — attachment is guaranteed to reach the user.
func grantsWith(t *testing.T, policyText string) bool {
	t.Helper()
	_, repo := newService(t)
	engine := sentinelEngine.NewSentinelEngine(repo.(sentinelEngine.PolicyStore))
	ctx := context.Background()
	org, user := uuid.New(), uuid.New()

	policy := &models.SentinelPolicy{
		OrganizationID: org, Name: "p-" + uuid.NewString(),
		Policy:  strings.ReplaceAll(policyText, "{USER}", user.String()),
		Enabled: true,
	}
	if err := repo.CreatePolicy(ctx, policy); err != nil {
		t.Fatal(err)
	}
	if err := repo.AttachPolicyToUser(ctx, org, user, policy.ID); err != nil {
		t.Fatal(err)
	}
	allowed, err := engine.IsAuthorized(ctx, sentinelEngine.AccessRequest{
		OrgID: org, UserID: user, Action: "canViewProperty",
		Resource: sentinelEngine.Resource{Type: "Property", ID: uuid.New().String()},
	})
	if err != nil {
		t.Fatalf("engine: %v", err)
	}
	return allowed
}

// The palette must advertise only User — the sole entity the engine builds.
//
// Role and Group were previously listed with the `in` operator, which produced
// policies that compile, store, and display as active while matching nothing.
// If this fails because a type was re-added, first build that entity in
// sentinelEngine.buildEntities; the test below shows what happens otherwise.
func TestPaletteAdvertisesOnlyEvaluableTypes(t *testing.T) {
	svc, _ := newService(t)
	view, err := svc.GetPrincipals(context.Background(), uuid.New())
	if err != nil {
		t.Fatal(err)
	}
	if len(view.Types) != 1 || view.Types[0].Type != "User" {
		t.Fatalf("principal types = %+v, want only User", view.Types)
	}

	// Every operator advertised for User must produce a policy that grants.
	forms := map[string]string{
		"is": `permit(principal is User, action, resource);`,
		"==": `permit(principal == User::"{USER}", action, resource);`,
	}
	for _, op := range view.Types[0].Operators {
		text, known := forms[op]
		if !known {
			t.Errorf("operator %q is advertised for User but this test has no form for it", op)
			continue
		}
		if !grantsWith(t, text) {
			t.Errorf("DEAD FORM: operator %q is advertised but %q grants nothing", op, text)
		}
	}
}

// Documents why Role/Group are not principal types. If someone re-adds them to
// the palette, this test still passes — but it is here to show the reasoning,
// and it will start FAILING the day the entity hierarchy is actually built,
// which is the signal that the palette may safely advertise them again.
func TestRoleAndGroupPrincipalFormsRemainUnevaluable(t *testing.T) {
	if grantsWith(t, `permit(principal in Role::"Owner", action, resource);`) {
		t.Error("Role entities now resolve — buildEntities must have changed; " +
			"the principal palette may now advertise Role with `in`")
	}
	if grantsWith(t, `permit(principal in Group::"Stakeholders", action, resource);`) {
		t.Error("Group entities now resolve — buildEntities must have changed; " +
			"the principal palette may now advertise Group with `in`")
	}
}
