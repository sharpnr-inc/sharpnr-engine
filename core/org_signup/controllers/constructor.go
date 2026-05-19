package org_signup_controllers

import (
	org_signup_services "sharpnr-engine/core/org_signup/services"
	"sharpnr-engine/pkg/cookieManager"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"
)

type orgSignUpController struct {
	cookies   cookieManager.Manager
	jwtMgr    jwt_manager.Manager
	signUpSvc org_signup_services.OrganizationSignUpServiceInterfaces
}

func NewOrgSignController(
	cookies cookieManager.Manager,
	jwtMgr jwt_manager.Manager,
	signUpSvc org_signup_services.OrganizationSignUpServiceInterfaces,
) *orgSignUpController {
	return &orgSignUpController{
		cookies:   cookies,
		jwtMgr:    jwtMgr,
		signUpSvc: signUpSvc,
	}
}
