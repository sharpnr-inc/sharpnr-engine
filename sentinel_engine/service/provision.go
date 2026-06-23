package sentinelservice

import (
	"context"
	"fmt"
	dbtx "sharpnr-engine/pkg/dbTx"
	sentinelengine_models "sharpnr-engine/sentinel_engine/models"
	sentinel_utils "sharpnr-engine/sentinel_engine/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ProvisionOrganization sets up a new org's authorization defaults: it creates
// the "Owner" role, attaches the shared system "Full Access" managed policy to
// it, generates the owner's Cedar from that attachment, and seeds the first
// employee as owner. Runs in one transaction so an org never ends up
// half-provisioned (a role with no policy, or an owner with no role).
//
// The owner grant goes through the same managed-policy path as every other role
// (attach -> generate), so it appears in the policy catalog and in the owner
// role's privileges — no special-cased hand-authored policy row.
func (s *sentinelService) ProvisionOrganization(ctx context.Context, orgID, firstUserID uuid.UUID) error {
	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		role := &sentinelengine_models.Role{
			ID:             uuid.New(),
			OrganizationID: orgID,
			Slug:           sentinel_utils.OwnerRoleSlug,
			Name:           "Owner",
			Description:    "Full access within the organization.",
		}
		if err := s.repo.CreateRole(txCtx, role); err != nil {
			return fmt.Errorf("create owner role: %w", err)
		}

		// Attach the shared "Full Access" system managed policy (seeded in
		// migration 000032) to the owner role.
		edge := &sentinelengine_models.RolePolicy{
			OrganizationID:  orgID,
			RoleID:          role.ID,
			ManagedPolicyID: sentinel_utils.SystemFullAccessPolicyID,
			AttachedBy:      firstUserID,
		}
		if err := s.repo.AttachRolePolicy(txCtx, edge); err != nil {
			return fmt.Errorf("attach full-access policy to owner role: %w", err)
		}

		// Generate the owner's Cedar from the attachment. Runs in this same tx
		// (WithTx joins the ambient tx); InvalidateOrg is a no-op for a brand-new
		// org that has nothing cached yet.
		if err := s.RegenerateRolePolicies(txCtx, orgID, role.ID); err != nil {
			return fmt.Errorf("generate owner policies: %w", err)
		}

		if err := s.repo.AssignUserRole(txCtx, orgID, firstUserID, role.ID); err != nil {
			return fmt.Errorf("assign first user as owner: %w", err)
		}
		return nil
	})
}
