package sharpnr_identity_auth_controllers

import sharpnr_identity_auth_services "sharpnr-engine/core/sharpnrIdentityAuth/services"

type sharpnrIdentityAuthControllers struct {
	IdSvc sharpnr_identity_auth_services.SharpnrIdentityAuthServiceInterface
}

func NewSharpnrIdentityAuthController(
	IdSvc sharpnr_identity_auth_services.SharpnrIdentityAuthServiceInterface,
) *sharpnrIdentityAuthControllers {
	return &sharpnrIdentityAuthControllers{
		IdSvc: IdSvc,
	}
}
