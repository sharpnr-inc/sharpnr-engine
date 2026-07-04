package sentinelEngine

import (
	"context"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"
	"github.com/google/uuid"
)

// fakeStore is an in-memory PolicyStore: each user's effective policies are
// preassembled, exactly as GetAllPoliciesOfUser would return them (org-scoped
// union role/group/user attachments).
type fakeStore struct {
	byUser map[uuid.UUID][]PolicyRecord
	calls  int
}

func (f *fakeStore) GetAllPoliciesOfUser(_ context.Context, _, userID uuid.UUID) ([]PolicyRecord, error) {
	f.calls++
	return f.byUser[userID], nil
}

func TestEngineDecisions(t *testing.T) {
	orgID := uuid.New()
	owner := uuid.New()   // has the seeded full-access policy via the Owner role
	manager := uuid.New() // has one instance-scoped policy via a group
	outsider := uuid.New()
	propertyID := uuid.New().String()

	fullAccess := PolicyRecord{ID: uuid.New(), Text: "permit (principal, action, resource);"}
	viewOneProperty := PolicyRecord{ID: uuid.New(), Text: `permit (principal, action == Action::"canViewProperty", resource == Property::"` + propertyID + `");`}
	// Org-scoped policy with a resource-attribute condition — applies to everyone.
	viewMumbai := PolicyRecord{ID: uuid.New(), Text: `permit (principal, action == Action::"canViewProperty", resource is Property) when { resource.city == "Mumbai" };`}
	noCreate := PolicyRecord{ID: uuid.New(), Text: `forbid (principal, action == Action::"canCreateNewProperty", resource);`}

	store := &fakeStore{byUser: map[uuid.UUID][]PolicyRecord{
		owner:    {fullAccess, viewMumbai},
		manager:  {viewOneProperty, viewMumbai},
		outsider: {viewMumbai},
	}}
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
		{"outsider is denied by default", AccessRequest{OrgID: orgID, UserID: outsider, Action: "canEditPropertyDetails", Resource: Resource{Type: "Property", ID: propertyID}}, false},
		{"attribute condition allows matching city", AccessRequest{OrgID: orgID, UserID: outsider, Action: "canViewProperty", Resource: Resource{Type: "Property", ID: propertyID}, ResourceAttributes: map[string]any{"city": "Mumbai"}}, true},
		{"attribute condition denies other city", AccessRequest{OrgID: orgID, UserID: outsider, Action: "canViewProperty", Resource: Resource{Type: "Property", ID: propertyID}, ResourceAttributes: map[string]any{"city": "Pune"}}, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := engine.IsAuthorized(ctx, tc.req)
			require.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}

	t.Run("forbid overrides permit", func(t *testing.T) {
		// Give the owner a forbid alongside full access (e.g. attached via a
		// group); Cedar's forbid must win over any permit.
		store.byUser[owner] = []PolicyRecord{fullAccess, noCreate}
		engine.InvalidateOrganization(orgID)

		got, err := engine.IsAuthorized(ctx, AccessRequest{OrgID: orgID, UserID: owner, Action: "canCreateNewProperty", Resource: Resource{Type: "Property"}})
		require.NoError(t, err)
		assert.False(t, got)
	})
}

func TestEngineFailsClosedOnBadPolicy(t *testing.T) {
	userID := uuid.New()
	store := &fakeStore{byUser: map[uuid.UUID][]PolicyRecord{
		userID: {{ID: uuid.New(), Text: "this is not cedar"}},
	}}
	engine := NewSentinelEngine(store)

	_, err := engine.IsAuthorized(context.Background(), AccessRequest{
		OrgID: uuid.New(), UserID: userID, Action: "x", Resource: Resource{Type: "Property"},
	})
	require.Error(t, err, "an uncompilable stored policy must fail the evaluation")
}

func TestEngineValidatesRequest(t *testing.T) {
	engine := NewSentinelEngine(&fakeStore{})
	ctx := context.Background()

	_, err := engine.IsAuthorized(ctx, AccessRequest{UserID: uuid.New(), Action: "x", Resource: Resource{Type: "Property"}})
	assert.Error(t, err, "missing orgID must be rejected")

	_, err = engine.IsAuthorized(ctx, AccessRequest{OrgID: uuid.New(), UserID: uuid.New(), Resource: Resource{Type: "Property"}})
	assert.Error(t, err, "missing action must be rejected")
}

func TestPerUserCacheAndInvalidation(t *testing.T) {
	orgID := uuid.New()
	user := uuid.New()
	store := &fakeStore{byUser: map[uuid.UUID][]PolicyRecord{}}
	engine := NewSentinelEngine(store)
	ctx := context.Background()

	req := AccessRequest{OrgID: orgID, UserID: user, Action: "canViewProperty", Resource: Resource{Type: "Property"}}

	got, err := engine.IsAuthorized(ctx, req)
	require.NoError(t, err)
	assert.False(t, got, "no policies -> deny")
	assert.Equal(t, 1, store.calls)

	// Second evaluation within the TTL is served from the per-user cache.
	_, err = engine.IsAuthorized(ctx, req)
	require.NoError(t, err)
	assert.Equal(t, 1, store.calls, "cached policy set should be reused")

	// A new attachment appears; the stale cache still denies until the org is
	// invalidated (which the authoring service does after every policy write).
	store.byUser[user] = []PolicyRecord{{ID: uuid.New(), Text: "permit (principal, action, resource);"}}
	got, err = engine.IsAuthorized(ctx, req)
	require.NoError(t, err)
	assert.False(t, got, "stale cache should still deny")

	engine.InvalidateOrganization(orgID)
	got, err = engine.IsAuthorized(ctx, req)
	require.NoError(t, err)
	assert.True(t, got, "after invalidation the new policy must apply")
	assert.Equal(t, 2, store.calls, "only the initial load and the post-invalidation reload hit the store")
}
