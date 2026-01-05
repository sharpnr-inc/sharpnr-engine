package org_signup_controllers

import (
	"sharpnr-engine/cookieManager"
	org_signup_services "sharpnr-engine/core/org-signup/services"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"
)

type orgSignUpController struct {
	orgSvc  org_signup_services.OrgSignUpServiceInterface
	cookies cookieManager.Manager
	jwtMgr  jwt_manager.Manager
}

func NewOrgSignController(
	orgSvc org_signup_services.OrgSignUpServiceInterface,
	cookies cookieManager.Manager,
	jwtMgr jwt_manager.Manager,
) *orgSignUpController {
	return &orgSignUpController{
		orgSvc:  orgSvc,
		cookies: cookies,
		jwtMgr:  jwtMgr,
	}
}
