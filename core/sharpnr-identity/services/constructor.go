package sharpnr_identity_auth_services

import sharpnr_identity_auth_repositories "sharpnr-engine/core/sharpnr-identity/repositories"

type sharpnrIdentityAuthService struct {
	identitySvc sharpnr_identity_auth_repositories.SharpnrIdentityAuthRepositoryInterface
}

func NewSharpnrIdentityAuthService(
	identitySvc sharpnr_identity_auth_repositories.SharpnrIdentityAuthRepositoryInterface,
) *sharpnrIdentityAuthService {
	return &sharpnrIdentityAuthService{
		identitySvc: identitySvc,
	}
}
