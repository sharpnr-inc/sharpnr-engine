// Tests for ProvisionOrganization.
//
// Provisioning now goes through the managed-policy path: create the owner role,
// attach the shared "Full Access" system policy, generate the owner's Cedar from
// that attachment, then assign the first user. These tests fake the data layer
// with an embedded-interface stub (overriding only the methods the flow calls)
// and use a real engine built on a nil repo — InvalidateOrg never touches it.
package sentinelservice

import (
	"context"
	"errors"
	"strings"
	"testing"

	dbtx "sharpnr-engine/pkg/dbTx"
	sentinelengine "sharpnr-engine/sentinel_engine/engine"
	sentinelengine_models "sharpnr-engine/sentinel_engine/models"
	sentinel_repo "sharpnr-engine/sentinel_engine/repo"
	sentinel_utils "sharpnr-engine/sentinel_engine/utils"

	"github.com/cedar-policy/cedar-go"
	"github.com/go-openapi/testify/v2/require"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/* ========================================================================== */
/*                                   STUBS                                    */
/* ========================================================================== */

// stubSentinelRepo fakes the data layer. Each method delegates to a function
// field so a test wires up only what its scenario should reach; any unset
// method that gets called hits the nil embedded interface and panics.
type stubSentinelRepo struct {
	sentinel_repo.SentinelRepositoryInterface

	createRole                  func(ctx context.Context, role *sentinelengine_models.Role) error
	attachRolePolicy            func(ctx context.Context, edge *sentinelengine_models.RolePolicy) error
	getRoleByID                 func(ctx context.Context, id uuid.UUID) (sentinelengine_models.Role, error)
	listAttachedManagedPolicies func(ctx context.Context, orgID, roleID uuid.UUID) ([]*sentinelengine_models.ManagedPolicy, error)
	deleteGeneratedPolicies     func(ctx context.Context, orgID, roleID uuid.UUID) error
	createPolicy                func(ctx context.Context, policy *sentinelengine_models.SentinelPolicy) error
	assignUserRole              func(ctx context.Context, orgID, userID, roleID uuid.UUID) error
}

func (s *stubSentinelRepo) CreateRole(ctx context.Context, role *sentinelengine_models.Role) error {
	return s.createRole(ctx, role)
}
func (s *stubSentinelRepo) AttachRolePolicy(ctx context.Context, edge *sentinelengine_models.RolePolicy) error {
	return s.attachRolePolicy(ctx, edge)
}
func (s *stubSentinelRepo) GetRoleByID(ctx context.Context, id uuid.UUID) (sentinelengine_models.Role, error) {
	return s.getRoleByID(ctx, id)
}
func (s *stubSentinelRepo) ListAttachedManagedPolicies(ctx context.Context, orgID, roleID uuid.UUID) ([]*sentinelengine_models.ManagedPolicy, error) {
	return s.listAttachedManagedPolicies(ctx, orgID, roleID)
}
func (s *stubSentinelRepo) DeleteGeneratedPoliciesByRole(ctx context.Context, orgID, roleID uuid.UUID) error {
	return s.deleteGeneratedPolicies(ctx, orgID, roleID)
}
func (s *stubSentinelRepo) CreatePolicy(ctx context.Context, policy *sentinelengine_models.SentinelPolicy) error {
	return s.createPolicy(ctx, policy)
}
func (s *stubSentinelRepo) AssignUserRole(ctx context.Context, orgID, userID, roleID uuid.UUID) error {
	return s.assignUserRole(ctx, orgID, userID, roleID)
}

// stubTxManager runs the callback inline with a nil *gorm.DB. The stub repo
// never touches the tx, so nil is fine; nested WithTx (the regenerate step) just
// runs inline again.
type stubTxManager struct{}

func (stubTxManager) WithTx(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return fn(nil)
}

// fullAccessPolicy is the managed policy the owner role is attached to.
func fullAccessPolicy() *sentinelengine_models.ManagedPolicy {
	return &sentinelengine_models.ManagedPolicy{
		ID:         sentinel_utils.SystemFullAccessPolicyID,
		Name:       "Full Access",
		Effect:     sentinelengine_models.EffectPermit,
		AllActions: true,
		System:     true,
	}
}

func newEngine() *sentinelengine.Sentinel { return sentinelengine.New(nil) }

/* ========================================================================== */
/*                          ProvisionOrganization                             */
/* ========================================================================== */

// Happy path: owner role created, Full Access attached, a wildcard owner policy
// generated and tagged with its source, and the first user assigned the role.
func TestProvisionOrganization_Success(t *testing.T) {
	orgID := uuid.New()
	firstUserID := uuid.New()
	generatedRoleID := uuid.New()

	var gotRole *sentinelengine_models.Role
	var gotEdge *sentinelengine_models.RolePolicy
	var gotPolicy *sentinelengine_models.SentinelPolicy
	var assignOrg, assignUser, assignRole uuid.UUID
	assignCalled := false

	repo := &stubSentinelRepo{
		createRole: func(_ context.Context, role *sentinelengine_models.Role) error {
			role.ID = generatedRoleID
			gotRole = role
			return nil
		},
		attachRolePolicy: func(_ context.Context, edge *sentinelengine_models.RolePolicy) error {
			gotEdge = edge
			return nil
		},
		getRoleByID: func(_ context.Context, id uuid.UUID) (sentinelengine_models.Role, error) {
			return sentinelengine_models.Role{ID: id, OrganizationID: orgID, Slug: sentinel_utils.OwnerRoleSlug}, nil
		},
		listAttachedManagedPolicies: func(_ context.Context, _, _ uuid.UUID) ([]*sentinelengine_models.ManagedPolicy, error) {
			return []*sentinelengine_models.ManagedPolicy{fullAccessPolicy()}, nil
		},
		deleteGeneratedPolicies: func(_ context.Context, _, _ uuid.UUID) error { return nil },
		createPolicy: func(_ context.Context, policy *sentinelengine_models.SentinelPolicy) error {
			gotPolicy = policy
			return nil
		},
		assignUserRole: func(_ context.Context, orgID, userID, roleID uuid.UUID) error {
			assignCalled = true
			assignOrg, assignUser, assignRole = orgID, userID, roleID
			return nil
		},
	}

	svc := &sentinelService{repo: repo, txManager: stubTxManager{}, engine: newEngine()}

	err := svc.ProvisionOrganization(context.Background(), orgID, firstUserID)
	require.NoError(t, err)

	// --- role ---
	require.NotNil(t, gotRole)
	require.Equal(t, orgID, gotRole.OrganizationID)
	require.Equal(t, sentinel_utils.OwnerRoleSlug, gotRole.Slug)
	require.Equal(t, "Owner", gotRole.Name)

	// --- attachment: the shared Full Access system policy on the owner role ---
	require.NotNil(t, gotEdge)
	require.Equal(t, sentinel_utils.SystemFullAccessPolicyID, gotEdge.ManagedPolicyID)
	require.Equal(t, generatedRoleID, gotEdge.RoleID)

	// --- generated policy: wildcard, tenant-scoped, tagged with its source ---
	require.NotNil(t, gotPolicy, "owner policy must be generated")
	require.Equal(t, orgID, gotPolicy.OrganizationID)
	require.True(t, gotPolicy.Enabled)
	require.Contains(t, gotPolicy.Policy, `principal in Role::"owner"`)
	require.Contains(t, gotPolicy.Policy, "action,", "Full Access must compile to a wildcard action")
	require.Contains(t, gotPolicy.Policy, "resource.orgId == principal.orgId")
	require.NotNil(t, gotPolicy.SourceRoleID)
	require.Equal(t, generatedRoleID, *gotPolicy.SourceRoleID)
	require.NotNil(t, gotPolicy.SourceManagedPolicyID)
	require.Equal(t, sentinel_utils.SystemFullAccessPolicyID, *gotPolicy.SourceManagedPolicyID)

	// the generated text must be valid Cedar
	var p cedar.Policy
	require.NoError(t, p.UnmarshalCedar([]byte(gotPolicy.Policy)))

	// --- assignment: same org, first user, the created role's ID ---
	require.True(t, assignCalled)
	require.Equal(t, orgID, assignOrg)
	require.Equal(t, firstUserID, assignUser)
	require.Equal(t, generatedRoleID, assignRole)
}

// CreateRole fails: error propagates and nothing downstream is attempted.
func TestProvisionOrganization_CreateRoleFails_Propagates(t *testing.T) {
	wantErr := errors.New("insert role failed")
	attachCalled, assignCalled := false, false

	repo := &stubSentinelRepo{
		createRole:       func(_ context.Context, _ *sentinelengine_models.Role) error { return wantErr },
		attachRolePolicy: func(_ context.Context, _ *sentinelengine_models.RolePolicy) error { attachCalled = true; return nil },
		assignUserRole:   func(_ context.Context, _, _, _ uuid.UUID) error { assignCalled = true; return nil },
	}
	svc := &sentinelService{repo: repo, txManager: stubTxManager{}, engine: newEngine()}

	err := svc.ProvisionOrganization(context.Background(), uuid.New(), uuid.New())

	require.ErrorIs(t, err, wantErr)
	require.False(t, attachCalled)
	require.False(t, assignCalled)
}

// AttachRolePolicy fails: error propagates and the user is not assigned.
func TestProvisionOrganization_AttachFails_Propagates(t *testing.T) {
	wantErr := errors.New("attach failed")
	assignCalled := false

	repo := &stubSentinelRepo{
		createRole:       func(_ context.Context, _ *sentinelengine_models.Role) error { return nil },
		attachRolePolicy: func(_ context.Context, _ *sentinelengine_models.RolePolicy) error { return wantErr },
		assignUserRole:   func(_ context.Context, _, _, _ uuid.UUID) error { assignCalled = true; return nil },
	}
	svc := &sentinelService{repo: repo, txManager: stubTxManager{}, engine: newEngine()}

	err := svc.ProvisionOrganization(context.Background(), uuid.New(), uuid.New())

	require.ErrorIs(t, err, wantErr)
	require.False(t, assignCalled)
}

// AssignUserRole fails: error propagates (generation already succeeded).
func TestProvisionOrganization_AssignFails_Propagates(t *testing.T) {
	wantErr := errors.New("assign failed")
	orgID := uuid.New()

	repo := &stubSentinelRepo{
		createRole:       func(_ context.Context, role *sentinelengine_models.Role) error { role.ID = uuid.New(); return nil },
		attachRolePolicy: func(_ context.Context, _ *sentinelengine_models.RolePolicy) error { return nil },
		getRoleByID: func(_ context.Context, id uuid.UUID) (sentinelengine_models.Role, error) {
			return sentinelengine_models.Role{ID: id, OrganizationID: orgID, Slug: sentinel_utils.OwnerRoleSlug}, nil
		},
		listAttachedManagedPolicies: func(_ context.Context, _, _ uuid.UUID) ([]*sentinelengine_models.ManagedPolicy, error) {
			return []*sentinelengine_models.ManagedPolicy{fullAccessPolicy()}, nil
		},
		deleteGeneratedPolicies: func(_ context.Context, _, _ uuid.UUID) error { return nil },
		createPolicy:            func(_ context.Context, _ *sentinelengine_models.SentinelPolicy) error { return nil },
		assignUserRole:          func(_ context.Context, _, _, _ uuid.UUID) error { return wantErr },
	}
	svc := &sentinelService{repo: repo, txManager: stubTxManager{}, engine: newEngine()}

	err := svc.ProvisionOrganization(context.Background(), orgID, uuid.New())

	require.ErrorIs(t, err, wantErr)
}

// The generated Full Access policy must be valid Cedar — guards the wildcard
// template independently of the provisioning wiring.
func TestFullAccess_GeneratesValidCedar(t *testing.T) {
	text, ok := buildCedarPolicy(sentinel_utils.OwnerRoleSlug, fullAccessPolicy())
	require.True(t, ok)
	require.True(t, strings.Contains(text, "action,"), "must be a wildcard action scope")
	var p cedar.Policy
	require.NoError(t, p.UnmarshalCedar([]byte(text)))
}

// Compile-time guard: stubTxManager must satisfy the real TxManager interface.
var _ dbtx.TxManager = stubTxManager{}
