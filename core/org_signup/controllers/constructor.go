package org_signup_controllers

import (
	"sharpnr-engine/cookieManager"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"
)

type orgSignUpController struct {
	cookies cookieManager.Manager
	jwtMgr  jwt_manager.Manager
}

func NewOrgSignController(
	cookies cookieManager.Manager,
	jwtMgr jwt_manager.Manager,
) *orgSignUpController {
	return &orgSignUpController{
		cookies: cookies,
		jwtMgr:  jwtMgr,
	}
}
