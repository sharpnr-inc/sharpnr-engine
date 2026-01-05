package org_signup_controllers

import (
	"sharpnr-engine/cookieManager"
	org_signup_services "sharpnr-engine/core/org-signup/services"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_orgSignUpController_VerifyOtpForSignup(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		orgSvc  org_signup_services.OrgSignUpServiceInterface
		cookies cookieManager.Manager
		jwtMgr  jwt_manager.Manager
		// Named input parameters for target function.
		ctx *gin.Context
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewOrgSignController(tt.orgSvc, tt.cookies, tt.jwtMgr)
			c.VerifyOtpForSignup(tt.ctx)
		})
	}
}
