package sharpnr_identity_auth_controllers

import (
	"sharpnr-engine/cookieManager"
	sharpnr_identity_auth_services "sharpnr-engine/core/sharpnr_identity/services"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"
)

type sharpnrIdentityAuthControllers struct {
	cookieMgr cookieManager.Manager
	jwtMgr    jwt_manager.Manager
	IdSvc     sharpnr_identity_auth_services.SharpnrIdentityAuthServiceInterface
}

func NewSharpnrIdentityAuthController(
	cookieMgr cookieManager.Manager,
	jwtMgr jwt_manager.Manager,
	IdSvc sharpnr_identity_auth_services.SharpnrIdentityAuthServiceInterface,
) *sharpnrIdentityAuthControllers {
	return &sharpnrIdentityAuthControllers{
		cookieMgr: cookieMgr,
		jwtMgr:    jwtMgr,
		IdSvc:     IdSvc,
	}
}
