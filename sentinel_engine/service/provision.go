package sentinelservice

import (
	"context"
	"fmt"
	dbtx "sharpnr-engine/pkg/dbTx"
	sentinelengine_models "sharpnr-engine/sentinel_engine/models"
	sentinel_utils "sharpnr-engine/sentinel_engine/utils"

	"github.com/cedar-policy/cedar-go"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ProvisionOrganization sets up the default owner role + policy and seeds the
// first employee as owner. Runs in one transaction so an org never ends up
// half-provisioned (a role with no policy, or an owner with no role).
func (s *sentinelService) ProvisionOrganization(ctx context.Context, orgID, firstUserID uuid.UUID) error {
	// Validate the default policy compiles before we commit anything.
	var p cedar.Policy
	if err := p.UnmarshalCedar([]byte(sentinel_utils.OwnerPolicyText)); err != nil {
		return fmt.Errorf("default owner policy invalid: %w", err)
	}

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		dbtx := dbtx.InjectTx(ctx, tx)

		role := &sentinelengine_models.Role{
			OrganizationID: orgID,
			Slug:           sentinel_utils.OwnerRoleSlug,
			Name:           "Owner",
			Description:    "Full access within the organization.",
		}
		if err := s.repo.CreateRole(dbtx, role); err != nil {
			return fmt.Errorf("create owner role: %w", err)
		}

		policy := &sentinelengine_models.SentinelPolicy{
			OrganizationID: orgID,
			Name:           "owner-full-access",
			Description:    "Grants the owner role full access within the org.",
			Policy:         sentinel_utils.OwnerPolicyText,
			Enabled:        true,
		}
		if err := s.repo.CreatePolicy(dbtx, policy); err != nil {
			return fmt.Errorf("create owner policy: %w", err)
		}

		if err := s.repo.AssignUserRole(dbtx, firstUserID, role.ID); err != nil {
			return fmt.Errorf("assign first user as owner: %w", err)
		}
		return nil
	})
}
