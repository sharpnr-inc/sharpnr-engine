package sentinelEngine

import (
	"context"
	"testing"

	sentinelContexts "sharpnr-engine/sentinel_contexts"

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

func TestTenantBoundary(t *testing.T) {
	orgA := uuid.New()
	orgB := uuid.New()
	ownerOfA := uuid.New()
	propertyOfB := uuid.New().String()

	// The owner of org A holds a full-access permit — the widest grant possible.
	store := &fakeStore{byUser: map[uuid.UUID][]PolicyRecord{
		ownerOfA: {{ID: uuid.New(), Text: "permit (principal, action, resource);"}},
	}}
	engine := NewSentinelEngine(store)
	ctx := context.Background()

	t.Run("resource owned by another org is denied structurally", func(t *testing.T) {
		got, err := engine.IsAuthorized(ctx, AccessRequest{
			OrgID:  orgA,
			UserID: ownerOfA,
			Action: "canViewProperty",
			// The call site resolved the instance and found it belongs to org B.
			Resource: Resource{Type: "Property", ID: propertyOfB, OrgID: orgB},
		})
		require.NoError(t, err)
		assert.False(t, got, "no policy may cross the tenant boundary")
	})

	t.Run("resource owned by the request's own org is allowed", func(t *testing.T) {
		got, err := engine.IsAuthorized(ctx, AccessRequest{
			OrgID:    orgA,
			UserID:   ownerOfA,
			Action:   "canViewProperty",
			Resource: Resource{Type: "Property", ID: uuid.New().String(), OrgID: orgA},
		})
		require.NoError(t, err)
		assert.True(t, got)
	})

	t.Run("authored orgId guard evaluates via auto-stamped attributes", func(t *testing.T) {
		// A policy written with the explicit tenant guard: the engine stamps
		// principal.orgId and resource.orgId itself, so the condition evaluates
		// instead of erroring on a missing attribute.
		guarded := uuid.New()
		store.byUser[guarded] = []PolicyRecord{{
			ID:   uuid.New(),
			Text: "permit (principal, action, resource) when { resource.orgId == principal.orgId };",
		}}

		got, err := engine.IsAuthorized(ctx, AccessRequest{
			OrgID:    orgA,
			UserID:   guarded,
			Action:   "canViewProperty",
			Resource: Resource{Type: "Property", ID: uuid.New().String(), OrgID: orgA},
		})
		require.NoError(t, err)
		assert.True(t, got, "same-org request must satisfy the guard")

		got, err = engine.IsAuthorized(ctx, AccessRequest{
			OrgID:    orgA,
			UserID:   guarded,
			Action:   "canViewProperty",
			Resource: Resource{Type: "Property", ID: propertyOfB, OrgID: orgB},
		})
		require.NoError(t, err)
		assert.False(t, got, "cross-org request must be denied")
	})
}

func TestContextStamping(t *testing.T) {
	orgID := uuid.New()
	ctxBase := context.Background()

	newEngineWithPolicy := func(userID uuid.UUID, text string) *SentinelEngine {
		return NewSentinelEngine(&fakeStore{byUser: map[uuid.UUID][]PolicyRecord{
			userID: {{ID: uuid.New(), Text: text}},
		}})
	}
	baseReq := func(userID uuid.UUID) AccessRequest {
		return AccessRequest{OrgID: orgID, UserID: userID, Action: "canViewProperty", Resource: Resource{Type: "Property"}}
	}

	t.Run("context.channel gates by injected channel", func(t *testing.T) {
		userID := uuid.New()
		engine := newEngineWithPolicy(userID, `permit (principal, action, resource) when { context.channel == "mobile" };`)

		mobileCtx := sentinelContexts.Inject(ctxBase, sentinelContexts.AuthzContext{
			Request: sentinelContexts.RequestContext{Channel: sentinelContexts.ChannelMobile},
		})
		got, err := engine.IsAuthorized(mobileCtx, baseReq(userID))
		require.NoError(t, err)
		assert.True(t, got, "mobile request must satisfy the channel condition")

		webCtx := sentinelContexts.Inject(ctxBase, sentinelContexts.AuthzContext{
			Request: sentinelContexts.RequestContext{Channel: sentinelContexts.ChannelWeb},
		})
		got, err = engine.IsAuthorized(webCtx, baseReq(userID))
		require.NoError(t, err)
		assert.False(t, got, "web request must not satisfy the channel condition")

		// No injection at all: canonical keys are still stamped (zero values),
		// so the condition evaluates to false instead of erroring the policy.
		got, err = engine.IsAuthorized(ctxBase, baseReq(userID))
		require.NoError(t, err)
		assert.False(t, got)
	})

	t.Run("context.mfaVerified gates by injected MFA state", func(t *testing.T) {
		userID := uuid.New()
		engine := newEngineWithPolicy(userID, `permit (principal, action, resource) when { context.mfaVerified };`)

		got, err := engine.IsAuthorized(ctxBase, baseReq(userID))
		require.NoError(t, err)
		assert.False(t, got, "MFA not verified -> deny")

		mfaCtx := sentinelContexts.Inject(ctxBase, sentinelContexts.AuthzContext{
			Request: sentinelContexts.RequestContext{MFAVerified: true},
		})
		got, err = engine.IsAuthorized(mfaCtx, baseReq(userID))
		require.NoError(t, err)
		assert.True(t, got)
	})

	t.Run("principal.department gates by injected user context", func(t *testing.T) {
		userID := uuid.New()
		engine := newEngineWithPolicy(userID, `permit (principal, action, resource) when { principal.department == "Finance" };`)

		financeCtx := sentinelContexts.Inject(ctxBase, sentinelContexts.AuthzContext{
			User: sentinelContexts.UserContext{Department: "Finance"},
		})
		got, err := engine.IsAuthorized(financeCtx, baseReq(userID))
		require.NoError(t, err)
		assert.True(t, got)

		hrCtx := sentinelContexts.Inject(ctxBase, sentinelContexts.AuthzContext{
			User: sentinelContexts.UserContext{Department: "HR"},
		})
		got, err = engine.IsAuthorized(hrCtx, baseReq(userID))
		require.NoError(t, err)
		assert.False(t, got)
	})

	t.Run("caller extras pass through; canonical keys win on collision", func(t *testing.T) {
		userID := uuid.New()
		engine := newEngineWithPolicy(userID, `permit (principal, action, resource) when { context.amount > 1000 && context.channel == "web" };`)

		webCtx := sentinelContexts.Inject(ctxBase, sentinelContexts.AuthzContext{
			Request: sentinelContexts.RequestContext{Channel: sentinelContexts.ChannelWeb},
		})

		req := baseReq(userID)
		req.Context = map[string]any{
			"amount":  5000,
			"channel": "mobile", // spoof attempt — the injected (engine-stamped) channel must win
		}
		got, err := engine.IsAuthorized(webCtx, req)
		require.NoError(t, err)
		assert.True(t, got, "extras merge in, but the canonical channel stays web")

		req.Context["amount"] = 100
		got, err = engine.IsAuthorized(webCtx, req)
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
