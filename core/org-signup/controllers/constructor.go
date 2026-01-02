package org_signup_controllers

import org_signup_services "sharpnr-engine/core/org-signup/services"

type orgSignUpController struct {
	orgSvc org_signup_services.OrgSignUpServiceInterface
}

func NewOrgSignController(orgSvc org_signup_services.OrgSignUpServiceInterface) *orgSignUpController {
	return &orgSignUpController{
		orgSvc: orgSvc,
	}
}
