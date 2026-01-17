package app

import (
	sharpnr_identity_auth_services "sharpnr-engine/core/sharpnr_identity/services"

	"gorm.io/gorm"
)

func bootstrapIdentity(
	db *gorm.DB,
) sharpnr_identity_auth_services.SharpnrIdentityAuthServiceInterface {

	return sharpnr_identity_auth_services.NewSharpnrIdentityAuthService()
}
