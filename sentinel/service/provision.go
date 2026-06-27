package sentinelservice

import (
	"context"
	"fmt"
	dbtx "sharpnr-engine/pkg/dbTx"
	sentinelDefaultpolicies "sharpnr-engine/sentinel/default_policies"
	sentinelengine_models "sharpnr-engine/sentinel/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ProvisionOrganization sets up a new org's authorization defaults in one
// transaction so an org is never left half-provisioned:
//
//  1. Create the "Owner" role (slug "owner", which the default policy targets).
//  2. Create the owner's full-access policy from the default Cedar text.
//  3. Attach that policy to the owner role (role_policies bookkeeping).
//  4. Assign the founding employee to the owner role.
//  5. Seed the default top-level "Stakeholders" group (empty; no policies yet).
func (s *sentinelService) ProvisionOrganization(ctx context.Context, orgID, firstUserID uuid.UUID) error {
	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		role := &sentinelengine_models.Role{
			ID:             uuid.New(),
			OrganizationID: orgID,
			Slug:           ownerRoleSlug,
			Name:           ownerRoleName,
			Description:    ownerRoleDesc,
		}
		if err := s.repo.CreateRole(txCtx, role); err != nil {
			return fmt.Errorf("create owner role: %w", err)
		}

		policy := &sentinelengine_models.SentinelPolicy{
			ID:             uuid.New(),
			OrganizationID: orgID,
			Name:           ownerPolicyName,
			Description:    ownerPolicyDesc,
			Policy:         string(sentinelDefaultpolicies.OwnerPolcy),
			Enabled:        true,
			Scope:          sentinelengine_models.PolicyScopeRole, // attached to the owner role
		}
		if err := s.repo.CreatePolicy(txCtx, policy); err != nil {
			return fmt.Errorf("create owner policy: %w", err)
		}

		rp := &sentinelengine_models.RolePolicies{
			ID:               uuid.New(),
			OrganizationID:   orgID,
			RoleID:           role.ID,
			SentinelPolicyID: policy.ID,
		}
		if err := s.repo.CreateRolePolicy(txCtx, rp); err != nil {
			return fmt.Errorf("attach owner policy to owner role: %w", err)
		}

		if err := s.repo.AssignUserRole(txCtx, orgID, firstUserID, role.ID); err != nil {
			return fmt.Errorf("assign first user as owner: %w", err)
		}

		group := &sentinelengine_models.Group{
			ID:             uuid.New(),
			OrganizationID: orgID,
			Slug:           stakeholdersGroupSlug,
			Name:           stakeholdersGroupName,
			Description:    stakeholdersGroupDesc,
		}
		if err := s.repo.CreateGroup(txCtx, group); err != nil {
			return fmt.Errorf("create default stakeholders group: %w", err)
		}
		return nil
	})
}
