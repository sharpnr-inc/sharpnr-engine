package sentinelEngine

import (
	"context"
	"testing"

	"github.com/google/uuid"
)

// fakeStore is an in-memory PolicyStore for exercising the engine without a DB.
type fakeStore struct {
	policies []PolicyRecord
	parents  map[uuid.UUID]PrincipalParents
}

func (f *fakeStore) ListEnabledPolicies(_ context.Context, _ uuid.UUID) ([]PolicyRecord, error) {
	return f.policies, nil
}

func (f *fakeStore) GetPrincipalParents(_ context.Context, _, userID uuid.UUID) (PrincipalParents, error) {
	return f.parents[userID], nil
}

func TestEngineDecisions(t *testing.T) {
	orgID := uuid.New()
	owner := uuid.New()
	manager := uuid.New()
	outsider := uuid.New()
	propertyID := uuid.New().String()

	store := &fakeStore{
		policies: []PolicyRecord{
			// Owner full access — what ProvisionOrganization seeds.
			{ID: uuid.New(), Text: `permit (principal in Role::"Owner", action, resource);`},
			// Managers may view exactly one property instance.
			{ID: uuid.New(), Text: `permit (principal in Group::"Managers", action == Action::"canViewProperty", resource == Property::"` + propertyID + `");`},
			// Nobody in Group::"Contractors" may create properties, ever.
			{ID: uuid.New(), Text: `forbid (principal in Group::"Contractors", action == Action::"canCreateNewProperty", resource);`},
			// Attribute condition: viewing is allowed org-wide only in Mumbai.
			{ID: uuid.New(), Text: `permit (principal, action == Action::"canViewProperty", resource is Property) when { resource.city == "Mumbai" };`},
		},
		parents: map[uuid.UUID]PrincipalParents{
			owner:   {Roles: []string{"Owner"}, Groups: []string{"Stakeholders", "Contractors"}},
			manager: {Groups: []string{"Managers"}},
		},
	}
	engine := NewSentinelEngine(store)
	ctx := context.Background()

	cases := []struct {
		name string
		req  AccessRequest
		want bool
	}{
		{"owner can do anything (type-level)", AccessRequest{OrgID: orgID, UserID: owner, Action: "canCreateNewProperty", Resource: Resource{Type: "Property"}}, true},
		{"manager can view the granted instance", AccessRequest{OrgID: orgID, UserID: manager, Action: "canViewProperty", Resource: Resource{Type: "Property", ID: propertyID}}, true},
		{"manager cannot view another instance", AccessRequest{OrgID: orgID, UserID: manager, Action: "canViewProperty", Resource: Resource{Type: "Property", ID: uuid.New().String()}}, false},
		{"manager cannot create properties", AccessRequest{OrgID: orgID, UserID: manager, Action: "canCreateNewProperty", Resource: Resource{Type: "Property"}}, false},
		{"outsider is denied everything", AccessRequest{OrgID: orgID, UserID: outsider, Action: "canViewProperty", Resource: Resource{Type: "Property", ID: propertyID}}, false},
		{"attribute condition allows matching city", AccessRequest{OrgID: orgID, UserID: outsider, Action: "canViewProperty", Resource: Resource{Type: "Property", ID: propertyID}, ResourceAttributes: map[string]any{"city": "Mumbai"}}, true},
		{"attribute condition denies other city", AccessRequest{OrgID: orgID, UserID: outsider, Action: "canViewProperty", Resource: Resource{Type: "Property", ID: propertyID}, ResourceAttributes: map[string]any{"city": "Pune"}}, false},
		// forbid must beat the owner's blanket permit
		{"forbid overrides permit", AccessRequest{OrgID: orgID, UserID: owner, Action: "canCreateNewProperty", Resource: Resource{Type: "Property", ID: ""}}, false},
	}

	// The owner is also in Contractors only for the forbid-overrides-permit case;
	// run that case last after flipping membership, so first drop it for the
	// blanket-permit case.
	store.parents[owner] = PrincipalParents{Roles: []string{"Owner"}, Groups: []string{"Stakeholders"}}

	for _, tc := range cases[:len(cases)-1] {
		t.Run(tc.name, func(t *testing.T) {
			got, err := engine.IsAuthorized(ctx, tc.req)
			if err != nil {
				t.Fatalf("IsAuthorized: %v", err)
			}
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}

	// Membership changes must apply without cache invalidation (parents are read
	// per request).
	store.parents[owner] = PrincipalParents{Roles: []string{"Owner"}, Groups: []string{"Stakeholders", "Contractors"}}
	last := cases[len(cases)-1]
	t.Run(last.name, func(t *testing.T) {
		got, err := engine.IsAuthorized(ctx, last.req)
		if err != nil {
			t.Fatalf("IsAuthorized: %v", err)
		}
		if got != last.want {
			t.Fatalf("got %v, want %v", got, last.want)
		}
	})
}

func TestEngineFailsClosedOnBadPolicy(t *testing.T) {
	store := &fakeStore{
		policies: []PolicyRecord{{ID: uuid.New(), Text: `this is not cedar`}},
		parents:  map[uuid.UUID]PrincipalParents{},
	}
	engine := NewSentinelEngine(store)

	_, err := engine.IsAuthorized(context.Background(), AccessRequest{
		OrgID: uuid.New(), UserID: uuid.New(), Action: "x", Resource: Resource{Type: "Property"},
	})
	if err == nil {
		t.Fatal("expected an error for an uncompilable stored policy")
	}
}

func TestInvalidateOrganizationDropsCache(t *testing.T) {
	orgID := uuid.New()
	user := uuid.New()
	store := &fakeStore{
		policies: []PolicyRecord{},
		parents:  map[uuid.UUID]PrincipalParents{},
	}
	engine := NewSentinelEngine(store)
	ctx := context.Background()

	req := AccessRequest{OrgID: orgID, UserID: user, Action: "canViewProperty", Resource: Resource{Type: "Property"}}
	if got, err := engine.IsAuthorized(ctx, req); err != nil || got {
		t.Fatalf("expected deny with no policies, got %v err %v", got, err)
	}

	// Add a permit-everything policy; the cached (empty) set must still deny
	// until the org is invalidated.
	store.policies = []PolicyRecord{{ID: uuid.New(), Text: `permit (principal, action, resource);`}}
	if got, _ := engine.IsAuthorized(ctx, req); got {
		t.Fatal("cache should still hold the old policy set")
	}
	engine.InvalidateOrganization(orgID)
	if got, err := engine.IsAuthorized(ctx, req); err != nil || !got {
		t.Fatalf("expected allow after invalidation, got %v err %v", got, err)
	}
}
