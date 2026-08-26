package sentinelservice

import (
	"context"

	sentinelengine_models "sharpnr-engine/sentinel/models"

	"github.com/google/uuid"
)

// ProvisionOrganization implements [SentinelServiceInterface]: seed a new org's
// authorization defaults — the Owner role with its full-access policy, the
// founding employee's owner assignment, and the Stakeholders group. Idempotent:
// if the Owner role already exists the org is considered provisioned. Runs in
// one transaction (joining the caller's, e.g. signup OTP verification) so an
// org is never left half-provisioned.
func (s *sentinelService) ProvisionOrganization(ctx context.Context, orgID, firstUserID uuid.UUID) error {
	provisioned, err := s.repo.RoleExists(ctx, orgID, OwnerRoleName)
	if err != nil {
		return err
	}
	if provisioned {
		return nil
	}

	err = s.withTx(ctx, func(txCtx context.Context) error {
		role := &sentinelengine_models.Role{
			ID:             uuid.New(),
			OrganizationID: orgID,
			Name:           OwnerRoleName,
			Description:    OwnerRoleDescription,
		}
		if err := s.repo.CreateRole(txCtx, role); err != nil {
			return err
		}

		policy := &sentinelengine_models.SentinelPolicy{
			ID:             uuid.New(),
			OrganizationID: orgID,
			Name:           OwnerPolicyName,
			Description:    OwnerPolicyDescription,
			Policy:         ownerPolicyText(),
			Enabled:        true,
		}
		if err := s.repo.CreatePolicy(txCtx, policy); err != nil {
			return err
		}

		edge := &sentinelengine_models.RolePolicies{
			OrganizationID:   orgID,
			RoleID:           role.ID,
			SentinelPolicyID: policy.ID,
		}
		if err := s.repo.AssignPolicyToRole(txCtx, edge); err != nil {
			return err
		}

		if err := s.repo.AssignRoleToUser(txCtx, orgID, firstUserID, role.ID); err != nil {
			return err
		}

		group := &sentinelengine_models.Group{
			ID:             uuid.New(),
			OrganizationID: orgID,
			Name:           StakeholdersGroupName,
			Description:    StakeholdersGroupDescription,
		}
		if err := s.repo.CreateGroup(txCtx, group); err != nil {
			return err
		}
		return s.repo.AssignUserToGroup(txCtx, orgID, firstUserID, group.ID)
	})
	if err != nil {
		return err
	}

	s.invalidateEngine(orgID)
	return nil
}
