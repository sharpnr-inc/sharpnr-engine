package SharpIam_services

import (
	sharp_iam_models "sharpnr-engine/core/sharp_iam/models"
	sharpIam_repo "sharpnr-engine/core/sharp_iam/repo"
)

type SharpIamServiceInterface interface {
	// Policy Management
	CreatePolicy(policy *sharp_iam_models.SharpIamPolicies) error
}

type sharpIamService struct {
	repo   sharpIam_repo.SharpIamRepositoryInterface
	schema []byte
}

func NewSharpIamService(repo sharpIam_repo.SharpIamRepositoryInterface, schema []byte) *sharpIamService {
	return &sharpIamService{
		repo:   repo,
		schema: schema,
	}
}
