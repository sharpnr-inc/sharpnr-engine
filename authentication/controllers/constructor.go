package authenticationEngineControllers

import (
	authenticationEngineServices "sharpnr-engine/authentication/services"
	"sharpnr-engine/pkg/cookieManager"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"
)

type authenticationControllers struct {
	cookieMgr cookieManager.Manager
	jwtMgr    jwt_manager.Manager
	authSvc   authenticationEngineServices.AuthenticationServiceInterface
}

func NewAuthenticationController(
	cookieMgr cookieManager.Manager,
	jwtMgr jwt_manager.Manager,
	authSvc authenticationEngineServices.AuthenticationServiceInterface,
) *authenticationControllers {
	return &authenticationControllers{
		cookieMgr: cookieMgr,
		jwtMgr:    jwtMgr,
		authSvc:   authSvc,
	}
}
