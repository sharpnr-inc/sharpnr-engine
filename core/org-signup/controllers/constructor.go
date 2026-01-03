package org_signup_controllers

import (
	"sharpnr-engine/cookieManager"
	org_signup_services "sharpnr-engine/core/org-signup/services"
)

type orgSignUpController struct {
	orgSvc  org_signup_services.OrgSignUpServiceInterface
	cookies cookieManager.Manager
}

func NewOrgSignController(
	orgSvc org_signup_services.OrgSignUpServiceInterface,
	cookies cookieManager.Manager,
) *orgSignUpController {
	return &orgSignUpController{
		orgSvc:  orgSvc,
		cookies: cookies,
	}
}
