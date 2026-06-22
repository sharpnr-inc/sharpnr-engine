// Tests for the Sentinel authorization engine.
//
// Conventions follow the project template (see sentinel_engine/service tests):
//   - White-box (same package) so we can reach the policy cache directly and
//     build a Sentinel with a fake Repository — no real DB, no Cedar mocking.
//   - The Repository is faked with a function-field stub. Each test wires only
//     the methods its scenario reaches; unset methods return empty results, so a
//     test that forgets to allow access still gets a deterministic deny.
//   - Real Cedar policies (text form) are compiled by the engine exactly as in
//     production, so these tests exercise the genuine permit/deny semantics.
//   - Assert with testify require; engine errors are wrapped with %w (ErrorIs).
package sentinelengine

import (
	"context"
	"errors"
	"testing"

	sentinel_actions "sharpnr-engine/sentinel_engine/actions"
	sentinelengine_models "sharpnr-engine/sentinel_engine/models"
	sentinel_utils "sharpnr-engine/sentinel_engine/utils"

	"github.com/cedar-policy/cedar-go/types"
	"github.com/go-openapi/testify/v2/require"
	"github.com/google/uuid"
)

/* ========================================================================== */
/*                                   STUBS                                    */
/* ========================================================================== */

// fakeRepo fakes the data layer the engine depends on. Each method delegates to
// a function field; a nil field yields an empty result (no policies / no roles /
// no groups), which makes "deny by default" the natural state of an unconfigured
// test.
type fakeRepo struct {
	listPolicies func(ctx context.Context, orgID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error)
	userRoles    func(ctx context.Context, orgID, userID uuid.UUID) ([]string, error)
	userGroups   func(ctx context.Context, orgID, userID uuid.UUID) ([]GroupMembership, error)
}

func (f *fakeRepo) ListEnabledPolicies(ctx context.Context, orgID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
	if f.listPolicies == nil {
		return nil, nil
	}
	return f.listPolicies(ctx, orgID)
}

func (f *fakeRepo) UserRoleSlugs(ctx context.Context, orgID, userID uuid.UUID) ([]string, error) {
	if f.userRoles == nil {
		return nil, nil
	}
	return f.userRoles(ctx, orgID, userID)
}

func (f *fakeRepo) UserGroupClosure(ctx context.Context, orgID, userID uuid.UUID) ([]GroupMembership, error) {
	if f.userGroups == nil {
		return nil, nil
	}
	return f.userGroups(ctx, orgID, userID)
}

// Compile-time guard: fakeRepo must satisfy the engine's Repository interface.
var _ Repository = (*fakeRepo)(nil)

/* ========================================================================== */
/*                                  HELPERS                                   */
/* ========================================================================== */

// policy builds an enabled SentinelPolicy row with a fresh ID for the given org.
func policy(orgID uuid.UUID, text string) *sentinelengine_models.SentinelPolicy {
	return &sentinelengine_models.SentinelPolicy{
		ID:             uuid.New(),
		OrganizationID: orgID,
		Policy:         text,
		Enabled:        true,
	}
}

// rolesReturning returns a userRoles func that always yields slugs.
func rolesReturning(slugs ...string) func(context.Context, uuid.UUID, uuid.UUID) ([]string, error) {
	return func(context.Context, uuid.UUID, uuid.UUID) ([]string, error) { return slugs, nil }
}

// groupsReturning returns a userGroups func that always yields the memberships.
func groupsReturning(gs ...GroupMembership) func(context.Context, uuid.UUID, uuid.UUID) ([]GroupMembership, error) {
	return func(context.Context, uuid.UUID, uuid.UUID) ([]GroupMembership, error) { return gs, nil }
}

// docInput builds a standard AuthorizeInput against a "Document" resource.
func docInput(orgID, userID uuid.UUID, action sentinel_actions.Name) AuthorizeInput {
	return AuthorizeInput{
		OrgID:  orgID,
		UserID: userID,
		Action: string(action),
		Resource: Resource{
			Type: "Document",
			ID:   "doc-1",
		},
	}
}

const editorPolicy = `permit (
  principal in Role::"editor",
  action == Action::"editDocument",
  resource
) when { resource.orgId == principal.orgId };`

/* ========================================================================== */
/*                                IsAuthorized                                */
/* ========================================================================== */

// The owner policy (full access within the tenant) permits an action when the
// user is assigned the owner role.
func TestIsAuthorized_OwnerRoleAllows(t *testing.T) {
	orgID, userID := uuid.New(), uuid.New()
	ownerPol := `permit (principal in Role::"owner", action, resource)
		when { resource.orgId == principal.orgId };`

	p := policy(orgID, ownerPol)
	repo := &fakeRepo{
		listPolicies: func(context.Context, uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
			return []*sentinelengine_models.SentinelPolicy{p}, nil
		},
		userRoles: rolesReturning("owner"),
	}

	res, err := New(repo).IsAuthorized(context.Background(), docInput(orgID, userID, sentinel_actions.EditDocument))

	require.NoError(t, err)
	require.True(t, res.Allowed)
	require.Empty(t, res.Errors)
	// The deciding policy's ID (== the row ID) is surfaced as a reason.
	require.Contains(t, res.Reasons, p.ID.String())
}

// End-to-end against the real provisioning defaults: an org provisioned with
// the bundled OwnerPolicyText + OwnerRoleSlug grants its owner full access for
// any action. This ties the engine directly to what ProvisionOrganization
// writes, so a future edit to either default that breaks evaluation fails here.
func TestIsAuthorized_BundledOwnerDefaultsGrantFullAccess(t *testing.T) {
	orgID, userID := uuid.New(), uuid.New()
	repo := &fakeRepo{
		listPolicies: func(context.Context, uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
			return []*sentinelengine_models.SentinelPolicy{policy(orgID, sentinel_utils.OwnerPolicyText)}, nil
		},
		userRoles: rolesReturning(sentinel_utils.OwnerRoleSlug),
	}
	eng := New(repo)

	// The owner role is unrestricted by action, so every defined action allows.
	for _, action := range []sentinel_actions.Name{
		sentinel_actions.ViewDocument,
		sentinel_actions.EditDocument,
		sentinel_actions.DeleteDocument,
		sentinel_actions.ViewRepo,
	} {
		res, err := eng.IsAuthorized(context.Background(), docInput(orgID, userID, action))
		require.NoError(t, err)
		require.True(t, res.Allowed, "owner must be allowed to %s", action)
	}

	// A non-owner with no other policy is denied, confirming the grant is
	// scoped to the owner role and not a blanket allow.
	denyRepo := &fakeRepo{
		listPolicies: func(context.Context, uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
			return []*sentinelengine_models.SentinelPolicy{policy(orgID, sentinel_utils.OwnerPolicyText)}, nil
		},
		userRoles: rolesReturning("viewer"),
	}
	res, err := New(denyRepo).IsAuthorized(context.Background(), docInput(orgID, userID, sentinel_actions.ViewDocument))
	require.NoError(t, err)
	require.False(t, res.Allowed, "non-owner must not get owner access")
}

// Default deny: with no policies (and no roles), every request is refused.
func TestIsAuthorized_NoPoliciesDenies(t *testing.T) {
	orgID, userID := uuid.New(), uuid.New()

	res, err := New(&fakeRepo{}).IsAuthorized(context.Background(), docInput(orgID, userID, sentinel_actions.ViewDocument))

	require.NoError(t, err)
	require.False(t, res.Allowed)
	require.Empty(t, res.Reasons)
}

// A user lacking the role required by the only policy is denied.
func TestIsAuthorized_MissingRoleDenies(t *testing.T) {
	orgID, userID := uuid.New(), uuid.New()
	repo := &fakeRepo{
		listPolicies: func(context.Context, uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
			return []*sentinelengine_models.SentinelPolicy{policy(orgID, editorPolicy)}, nil
		},
		userRoles: rolesReturning("viewer"), // not "editor"
	}

	res, err := New(repo).IsAuthorized(context.Background(), docInput(orgID, userID, sentinel_actions.EditDocument))

	require.NoError(t, err)
	require.False(t, res.Allowed)
}

// The action gate is honored: the editor policy permits editDocument only, so a
// viewDocument request by an editor is denied.
func TestIsAuthorized_WrongActionDenies(t *testing.T) {
	orgID, userID := uuid.New(), uuid.New()
	repo := &fakeRepo{
		listPolicies: func(context.Context, uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
			return []*sentinelengine_models.SentinelPolicy{policy(orgID, editorPolicy)}, nil
		},
		userRoles: rolesReturning("editor"),
	}

	res, err := New(repo).IsAuthorized(context.Background(), docInput(orgID, userID, sentinel_actions.ViewDocument))

	require.NoError(t, err)
	require.False(t, res.Allowed)
}

// Direct group membership grants access through a Group-scoped policy.
func TestIsAuthorized_DirectGroupAllows(t *testing.T) {
	orgID, userID := uuid.New(), uuid.New()
	groupPol := `permit (principal in Group::"engineering", action, resource)
		when { resource.orgId == principal.orgId };`
	repo := &fakeRepo{
		listPolicies: func(context.Context, uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
			return []*sentinelengine_models.SentinelPolicy{policy(orgID, groupPol)}, nil
		},
		userGroups: groupsReturning(GroupMembership{Slug: "engineering", Direct: true}),
	}

	res, err := New(repo).IsAuthorized(context.Background(), docInput(orgID, userID, sentinel_actions.ViewDocument))

	require.NoError(t, err)
	require.True(t, res.Allowed)
}

// A group reached only transitively (the user is a direct member of a child
// group whose parent is the policy's group) still grants access.
func TestIsAuthorized_ParentGroupAllowsTransitively(t *testing.T) {
	orgID, userID := uuid.New(), uuid.New()
	parentPol := `permit (principal in Group::"engineering", action, resource)
		when { resource.orgId == principal.orgId };`
	repo := &fakeRepo{
		listPolicies: func(context.Context, uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
			return []*sentinelengine_models.SentinelPolicy{policy(orgID, parentPol)}, nil
		},
		// User is a direct member of "frontend"; "engineering" is its parent and
		// appears in the closure as a non-direct ancestor.
		userGroups: groupsReturning(
			GroupMembership{Slug: "frontend", ParentSlug: "engineering", Direct: true},
			GroupMembership{Slug: "engineering", Direct: false},
		),
	}

	res, err := New(repo).IsAuthorized(context.Background(), docInput(orgID, userID, sentinel_actions.ViewDocument))

	require.NoError(t, err)
	require.True(t, res.Allowed)
}

// A non-direct (ancestor-only) membership with no path from the user must NOT
// grant access — this guards the Direct flag handling in buildEntities.
func TestIsAuthorized_NonDirectGroupDenies(t *testing.T) {
	orgID, userID := uuid.New(), uuid.New()
	groupPol := `permit (principal in Group::"admins", action, resource)
		when { resource.orgId == principal.orgId };`
	repo := &fakeRepo{
		listPolicies: func(context.Context, uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
			return []*sentinelengine_models.SentinelPolicy{policy(orgID, groupPol)}, nil
		},
		userGroups: groupsReturning(GroupMembership{Slug: "admins", Direct: false}),
	}

	res, err := New(repo).IsAuthorized(context.Background(), docInput(orgID, userID, sentinel_actions.ViewDocument))

	require.NoError(t, err)
	require.False(t, res.Allowed)
}

// Caller-supplied resource attributes are evaluated by the policy.
func TestIsAuthorized_ResourceAttributeGate(t *testing.T) {
	orgID, userID := uuid.New(), uuid.New()
	attrPol := `permit (principal in Role::"editor", action, resource)
		when { resource.confidential == false };`
	repo := &fakeRepo{
		listPolicies: func(context.Context, uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
			return []*sentinelengine_models.SentinelPolicy{policy(orgID, attrPol)}, nil
		},
		userRoles: rolesReturning("editor"),
	}
	eng := New(repo)

	in := docInput(orgID, userID, sentinel_actions.EditDocument)
	in.Resource.Attrs = map[string]types.Value{"confidential": types.Boolean(false)}
	allowRes, err := eng.IsAuthorized(context.Background(), in)
	require.NoError(t, err)
	require.True(t, allowRes.Allowed)

	in.Resource.Attrs = map[string]types.Value{"confidential": types.Boolean(true)}
	denyRes, err := eng.IsAuthorized(context.Background(), in)
	require.NoError(t, err)
	require.False(t, denyRes.Allowed)
}

// Request context is passed through to policy evaluation.
func TestIsAuthorized_ContextGate(t *testing.T) {
	orgID, userID := uuid.New(), uuid.New()
	ctxPol := `permit (principal in Role::"editor", action, resource)
		when { context.mfa == true };`
	repo := &fakeRepo{
		listPolicies: func(context.Context, uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
			return []*sentinelengine_models.SentinelPolicy{policy(orgID, ctxPol)}, nil
		},
		userRoles: rolesReturning("editor"),
	}
	eng := New(repo)

	in := docInput(orgID, userID, sentinel_actions.EditDocument)
	in.Context = map[string]types.Value{"mfa": types.Boolean(true)}
	allowRes, err := eng.IsAuthorized(context.Background(), in)
	require.NoError(t, err)
	require.True(t, allowRes.Allowed)

	in.Context = map[string]types.Value{"mfa": types.Boolean(false)}
	denyRes, err := eng.IsAuthorized(context.Background(), in)
	require.NoError(t, err)
	require.False(t, denyRes.Allowed)
}

/* ========================================================================== */
/*                          FAIL-CLOSED ERROR PATHS                           */
/* ========================================================================== */

// A malformed stored policy is an integrity error: the whole org fails closed
// with an error and no allow.
func TestIsAuthorized_MalformedPolicyFailsClosed(t *testing.T) {
	orgID, userID := uuid.New(), uuid.New()
	repo := &fakeRepo{
		listPolicies: func(context.Context, uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
			return []*sentinelengine_models.SentinelPolicy{policy(orgID, "this is not cedar")}, nil
		},
		userRoles: rolesReturning("owner"),
	}

	res, err := New(repo).IsAuthorized(context.Background(), docInput(orgID, userID, sentinel_actions.EditDocument))

	require.Error(t, err)
	require.False(t, res.Allowed)
}

// A repo failure loading policies fails closed.
func TestIsAuthorized_ListPoliciesErrorFailsClosed(t *testing.T) {
	wantErr := errors.New("db down")
	repo := &fakeRepo{
		listPolicies: func(context.Context, uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
			return nil, wantErr
		},
	}

	res, err := New(repo).IsAuthorized(context.Background(), docInput(uuid.New(), uuid.New(), sentinel_actions.ViewDocument))

	require.ErrorIs(t, err, wantErr)
	require.False(t, res.Allowed)
}

// A repo failure loading the user's roles fails closed.
func TestIsAuthorized_UserRolesErrorFailsClosed(t *testing.T) {
	orgID := uuid.New()
	wantErr := errors.New("roles query failed")
	repo := &fakeRepo{
		listPolicies: func(context.Context, uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
			return []*sentinelengine_models.SentinelPolicy{policy(orgID, editorPolicy)}, nil
		},
		userRoles: func(context.Context, uuid.UUID, uuid.UUID) ([]string, error) { return nil, wantErr },
	}

	res, err := New(repo).IsAuthorized(context.Background(), docInput(orgID, uuid.New(), sentinel_actions.EditDocument))

	require.ErrorIs(t, err, wantErr)
	require.False(t, res.Allowed)
}

// A repo failure loading the user's group closure fails closed.
func TestIsAuthorized_UserGroupsErrorFailsClosed(t *testing.T) {
	orgID := uuid.New()
	wantErr := errors.New("groups query failed")
	repo := &fakeRepo{
		listPolicies: func(context.Context, uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
			return []*sentinelengine_models.SentinelPolicy{policy(orgID, editorPolicy)}, nil
		},
		userGroups: func(context.Context, uuid.UUID, uuid.UUID) ([]GroupMembership, error) { return nil, wantErr },
	}

	res, err := New(repo).IsAuthorized(context.Background(), docInput(orgID, uuid.New(), sentinel_actions.EditDocument))

	require.ErrorIs(t, err, wantErr)
	require.False(t, res.Allowed)
}

/* ========================================================================== */
/*                            CACHING / ISOLATION                             */
/* ========================================================================== */

// Compiled policies are cached per org: a second authorization reuses the
// compiled set and does not hit ListEnabledPolicies again.
func TestPolicyCache_CompilesOncePerOrg(t *testing.T) {
	orgID, userID := uuid.New(), uuid.New()
	var compileCount int
	repo := &fakeRepo{
		listPolicies: func(context.Context, uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
			compileCount++
			return []*sentinelengine_models.SentinelPolicy{policy(orgID, editorPolicy)}, nil
		},
		userRoles: rolesReturning("editor"),
	}
	eng := New(repo)

	_, err := eng.IsAuthorized(context.Background(), docInput(orgID, userID, sentinel_actions.EditDocument))
	require.NoError(t, err)
	_, err = eng.IsAuthorized(context.Background(), docInput(orgID, userID, sentinel_actions.EditDocument))
	require.NoError(t, err)

	require.Equal(t, 1, compileCount, "policies should be compiled once and cached")
}

// InvalidateOrg drops the cached set so the next authorization recompiles.
func TestInvalidateOrg_ForcesRecompile(t *testing.T) {
	orgID, userID := uuid.New(), uuid.New()
	var compileCount int
	repo := &fakeRepo{
		listPolicies: func(context.Context, uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
			compileCount++
			return []*sentinelengine_models.SentinelPolicy{policy(orgID, editorPolicy)}, nil
		},
		userRoles: rolesReturning("editor"),
	}
	eng := New(repo)

	_, err := eng.IsAuthorized(context.Background(), docInput(orgID, userID, sentinel_actions.EditDocument))
	require.NoError(t, err)

	eng.InvalidateOrg(orgID)

	_, err = eng.IsAuthorized(context.Background(), docInput(orgID, userID, sentinel_actions.EditDocument))
	require.NoError(t, err)

	require.Equal(t, 2, compileCount, "InvalidateOrg should force a recompile")
}

// Policies are scoped per org: a policy compiled for one org never leaks into
// another org's decision. The same editor user is allowed in org A (which has
// the editor policy) and denied in org B (which has none).
func TestPolicyCache_PerOrgIsolation(t *testing.T) {
	orgA, orgB, userID := uuid.New(), uuid.New(), uuid.New()
	repo := &fakeRepo{
		listPolicies: func(_ context.Context, orgID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
			if orgID == orgA {
				return []*sentinelengine_models.SentinelPolicy{policy(orgA, editorPolicy)}, nil
			}
			return nil, nil // org B has no policies
		},
		userRoles: rolesReturning("editor"),
	}
	eng := New(repo)

	resA, err := eng.IsAuthorized(context.Background(), docInput(orgA, userID, sentinel_actions.EditDocument))
	require.NoError(t, err)
	require.True(t, resA.Allowed, "org A grants the editor")

	resB, err := eng.IsAuthorized(context.Background(), docInput(orgB, userID, sentinel_actions.EditDocument))
	require.NoError(t, err)
	require.False(t, resB.Allowed, "org B has no policy, must deny")
}
