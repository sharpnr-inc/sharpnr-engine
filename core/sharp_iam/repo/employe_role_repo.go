package sharpIam_repo

import (
	"context"
	sharp_iam_models "sharpnr-engine/core/sharp_iam/models"

	"github.com/google/uuid"
)

func (r *sharpIamRepository) AssignRoleToEmployee(ctx context.Context, employeeID uuid.UUID, roleID uuid.UUID) error {
	panic("unimplemented")
}

func (r *sharpIamRepository) RemoveRoleFromEmployee(ctx context.Context, employeeID uuid.UUID, roleID uuid.UUID) error {
	panic("unimplemented")
}

func (r *sharpIamRepository) ListRolesForEmployee(ctx context.Context, employeeID uuid.UUID) ([]*sharp_iam_models.SharpIamRoles, error) {
	panic("unimplemented")
}
