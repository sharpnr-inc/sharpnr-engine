package sharpIam_services

import (
	"context"
	sharp_iam_models "sharpnr-engine/core/sharp_iam/models"
	sharpIam_repo "sharpnr-engine/core/sharp_iam/repo"

	"github.com/cedar-policy/cedar-go"
	"github.com/google/uuid"
)

type SharpIamService struct {
	repo   sharpIam_repo.SharpIamRepositoryInterface
	engine *SharpIAM
}

func NewSharpIamService(repo sharpIam_repo.SharpIamRepositoryInterface, engine *SharpIAM) *SharpIamService {
	return &SharpIamService{
		repo:   repo,
		engine: engine,
	}
}

// Authorize is the entry point for your business logic
func (s *SharpIamService) Authorize(
	ctx context.Context,
	empID uuid.UUID,
	action string,
	resType string,
	resID uuid.UUID,
) (cedar.Decision, error) {

	// 1. Fetch current roles for the employee
	roles, err := s.repo.ListRolesForEmployee(ctx, empID)
	if err != nil {
		return false, err
	}

	var roleNames []string
	for _, r := range roles {
		roleNames = append(roleNames, r.RoleName)
	}

	// 2. Fetch the resource hierarchy (parents like OrgID, PropertyID)
	// This ensures Cedar can evaluate "when { resource in Org::'...' }"
	parents, err := s.repo.GetResourceParents(resType, resID)
	if err != nil {
		parents = []uuid.UUID{}
	}

	// 3. Pass everything to the engine
	return s.engine.IsAuthorized(
		ctx,
		empID,
		roleNames,
		action,
		resType,
		resID,
		parents,
	)
}

func (s *SharpIamService) CreatePolicy(ctx context.Context, policy *sharp_iam_models.SharpIamPolicies) error {
	return s.repo.CreatePolicy(ctx, policy)
}
