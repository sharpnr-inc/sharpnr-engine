// Tests for ProvisionOrganization.
//
// Conventions follow the project template (see orchestrator/service tests):
//   - White-box (same package) so we can build sentinelService directly with
//     only the fields the method touches — no real DB.
//   - The repo is faked with an embedded-interface stub: it satisfies the full
//     SentinelRepositoryInterface for free, and we override only the three
//     methods this flow calls via function fields. Any unexpected call hits the
//     nil embedded interface and panics — a deliberate tripwire.
//   - stubTxManager runs the transactional callback inline.
//   - Assert with testify require; errors are wrapped with %w so we use ErrorIs.
package sentinelservice

import (
	"context"
	"errors"
	"testing"

	dbtx "sharpnr-engine/pkg/dbTx"
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
// field so a test wires up only what its scenario should reach.
type stubSentinelRepo struct {
	sentinel_repo.SentinelRepositoryInterface

	createRole     func(ctx context.Context, role *sentinelengine_models.Role) error
	createPolicy   func(ctx context.Context, policy *sentinelengine_models.SentinelPolicy) error
	assignUserRole func(ctx context.Context, orgID, userID, roleID uuid.UUID) error
}

func (s *stubSentinelRepo) CreateRole(ctx context.Context, role *sentinelengine_models.Role) error {
	return s.createRole(ctx, role)
}

func (s *stubSentinelRepo) CreatePolicy(ctx context.Context, policy *sentinelengine_models.SentinelPolicy) error {
	return s.createPolicy(ctx, policy)
}

func (s *stubSentinelRepo) AssignUserRole(ctx context.Context, orgID, userID, roleID uuid.UUID) error {
	return s.assignUserRole(ctx, orgID, userID, roleID)
}

// stubTxManager runs the callback inline with a nil *gorm.DB. The stub repo
// never touches the tx, so nil is fine. Mirrors the production contract: the
// callback's return value is propagated to the caller.
type stubTxManager struct{}

func (stubTxManager) WithTx(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return fn(nil)
}

/* ========================================================================== */
/*                          ProvisionOrganization                             */
/* ========================================================================== */

// Happy path: the owner role and its full-access policy are created with the
// expected fields, and the first user is assigned to the *created* role's ID
// within the same tenant.
func TestProvisionOrganization_Success(t *testing.T) {
	orgID := uuid.New()
	firstUserID := uuid.New()
	// Simulates the DB-generated primary key written back onto the role on Create.
	generatedRoleID := uuid.New()

	var gotRole *sentinelengine_models.Role
	var gotPolicy *sentinelengine_models.SentinelPolicy
	var assignOrg, assignUser, assignRole uuid.UUID
	createPolicyCalled := false
	assignCalled := false

	repo := &stubSentinelRepo{
		createRole: func(_ context.Context, role *sentinelengine_models.Role) error {
			role.ID = generatedRoleID // emulate Postgres gen_random_uuid() write-back
			gotRole = role
			return nil
		},
		createPolicy: func(_ context.Context, policy *sentinelengine_models.SentinelPolicy) error {
			createPolicyCalled = true
			gotPolicy = policy
			return nil
		},
		assignUserRole: func(_ context.Context, orgID, userID, roleID uuid.UUID) error {
			assignCalled = true
			assignOrg, assignUser, assignRole = orgID, userID, roleID
			return nil
		},
	}

	svc := &sentinelService{repo: repo, txManager: stubTxManager{}}

	err := svc.ProvisionOrganization(context.Background(), orgID, firstUserID)
	require.NoError(t, err)

	// --- role ---
	require.NotNil(t, gotRole)
	require.Equal(t, orgID, gotRole.OrganizationID)
	require.Equal(t, sentinel_utils.OwnerRoleSlug, gotRole.Slug)
	require.Equal(t, "Owner", gotRole.Name)

	// --- policy ---
	require.True(t, createPolicyCalled, "owner policy must be created")
	require.NotNil(t, gotPolicy)
	require.Equal(t, orgID, gotPolicy.OrganizationID)
	require.Equal(t, sentinel_utils.OwnerPolicyText, gotPolicy.Policy)
	require.True(t, gotPolicy.Enabled, "owner policy must be enabled")

	// --- assignment: same org, first user, and the role's generated ID ---
	require.True(t, assignCalled, "first user must be assigned the owner role")
	require.Equal(t, orgID, assignOrg)
	require.Equal(t, firstUserID, assignUser)
	require.Equal(t, generatedRoleID, assignRole, "assignment must use the created role's ID")
}

// The bundled default owner policy must be valid Cedar. This guards against a
// future edit to OwnerPolicyText that would only fail at provisioning time.
func TestProvisionOrganization_DefaultPolicyCompiles(t *testing.T) {
	var p cedar.Policy
	err := p.UnmarshalCedar([]byte(sentinel_utils.OwnerPolicyText))
	require.NoError(t, err)
}

// CreateRole fails: error propagates and neither the policy nor the assignment
// is attempted (transaction would roll back).
func TestProvisionOrganization_CreateRoleFails_Propagates(t *testing.T) {
	wantErr := errors.New("insert role failed")
	policyCalled, assignCalled := false, false

	repo := &stubSentinelRepo{
		createRole: func(_ context.Context, _ *sentinelengine_models.Role) error { return wantErr },
		createPolicy: func(_ context.Context, _ *sentinelengine_models.SentinelPolicy) error {
			policyCalled = true
			return nil
		},
		assignUserRole: func(_ context.Context, _, _, _ uuid.UUID) error {
			assignCalled = true
			return nil
		},
	}
	svc := &sentinelService{repo: repo, txManager: stubTxManager{}}

	err := svc.ProvisionOrganization(context.Background(), uuid.New(), uuid.New())

	require.ErrorIs(t, err, wantErr)
	require.False(t, policyCalled, "policy must not be created after role creation fails")
	require.False(t, assignCalled, "user must not be assigned after role creation fails")
}

// CreatePolicy fails: error propagates and the assignment is not attempted.
func TestProvisionOrganization_CreatePolicyFails_Propagates(t *testing.T) {
	wantErr := errors.New("insert policy failed")
	assignCalled := false

	repo := &stubSentinelRepo{
		createRole:   func(_ context.Context, _ *sentinelengine_models.Role) error { return nil },
		createPolicy: func(_ context.Context, _ *sentinelengine_models.SentinelPolicy) error { return wantErr },
		assignUserRole: func(_ context.Context, _, _, _ uuid.UUID) error {
			assignCalled = true
			return nil
		},
	}
	svc := &sentinelService{repo: repo, txManager: stubTxManager{}}

	err := svc.ProvisionOrganization(context.Background(), uuid.New(), uuid.New())

	require.ErrorIs(t, err, wantErr)
	require.False(t, assignCalled, "user must not be assigned after policy creation fails")
}

// AssignUserRole fails: error propagates.
func TestProvisionOrganization_AssignFails_Propagates(t *testing.T) {
	wantErr := errors.New("assign failed")

	repo := &stubSentinelRepo{
		createRole:     func(_ context.Context, _ *sentinelengine_models.Role) error { return nil },
		createPolicy:   func(_ context.Context, _ *sentinelengine_models.SentinelPolicy) error { return nil },
		assignUserRole: func(_ context.Context, _, _, _ uuid.UUID) error { return wantErr },
	}
	svc := &sentinelService{repo: repo, txManager: stubTxManager{}}

	err := svc.ProvisionOrganization(context.Background(), uuid.New(), uuid.New())

	require.ErrorIs(t, err, wantErr)
}

// Compile-time guard: stubTxManager must satisfy the real TxManager interface.
var _ dbtx.TxManager = stubTxManager{}
