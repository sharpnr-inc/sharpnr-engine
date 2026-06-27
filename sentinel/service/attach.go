package sentinelservice

import (
	"context"
	"fmt"

	dbtx "sharpnr-engine/pkg/dbTx"
	sentinelengine_models "sharpnr-engine/sentinel/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AttachPolicies groups a set of existing policies under a role by inserting
// role_policies edges, in one transaction. The policy rows themselves are
// untouched; this is the organizational record the privileges view reads.
func (s *sentinelService) AttachPolicies(ctx context.Context, orgID, emplID, roleID uuid.UUID, policyIDs []uuid.UUID) error {
	if len(policyIDs) == 0 {
		return nil
	}
	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)
		for _, policyID := range policyIDs {
			rp := &sentinelengine_models.RolePolicies{
				ID:               uuid.New(),
				OrganizationID:   orgID,
				RoleID:           roleID,
				SentinelPolicyID: policyID,
			}
			if err := s.repo.CreateRolePolicy(txCtx, rp); err != nil {
				return fmt.Errorf("attach policy %s to role %s: %w", policyID, roleID, err)
			}
		}
		return nil
	})
}

// DetachPolicy removes one policy's grouping under a role. The policy row is not
// deleted, only the role_policies edge.
func (s *sentinelService) DetachPolicy(ctx context.Context, orgID, roleID, policyID uuid.UUID) error {
	if err := s.repo.DeleteRolePolicy(ctx, orgID, roleID, policyID); err != nil {
		return fmt.Errorf("detach policy %s from role %s: %w", policyID, roleID, err)
	}
	return nil
}
