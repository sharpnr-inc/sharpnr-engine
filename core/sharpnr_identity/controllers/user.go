package sharpnr_identity_auth_controllers

import (
	"net/http"
	"sharpnr-engine/cookieManager"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *sharpnrIdentityAuthControllers) GetMe(ctx *gin.Context) {

}

func (c *sharpnrIdentityAuthControllers) UpdateMyProfile(ctx *gin.Context) {

}

func (c *sharpnrIdentityAuthControllers) LogoutUser(ctx *gin.Context) {

	// Clear auth cookies
	c.cookieMgr.Clear(ctx, cookieManager.Auth)
	c.cookieMgr.Clear(ctx, cookieManager.CSRF)
	c.cookieMgr.Clear(ctx, cookieManager.Signup)

	// Send success response
	utils.Success(ctx, http.StatusOK, "User logged out successfully", nil)
}
