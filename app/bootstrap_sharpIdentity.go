package app

import (
	sharpnr_identity_auth_repositories "sharpnr-engine/core/sharpnr-identity/repositories"
	sharpnr_identity_auth_services "sharpnr-engine/core/sharpnr-identity/services"

	"gorm.io/gorm"
)

func bootstrapIdentity(
	db *gorm.DB,
	hr *hrBootstrap,
) sharpnr_identity_auth_services.SharpnrIdentityAuthServiceInterface {

	identityRepo := sharpnr_identity_auth_repositories.NewSharpnrIdentityAuthRepository(db)

	return sharpnr_identity_auth_services.NewSharpnrIdentityAuthService(
		hr.EmpIdentityRepo,
		hr.EmpAuthRepo,
		identityRepo,
	)
}
