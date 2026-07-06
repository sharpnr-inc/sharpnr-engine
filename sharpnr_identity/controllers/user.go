package sharpnr_identity_auth_controllers

import (
	"net/http"
	"sharpnr-engine/pkg/cookieManager"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *sharpnrIdentityAuthControllers) VerifyAuthentication(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	utils.Success(ctx, http.StatusOK, "", gin.H{
		"authenticated": true,
		"user_id":       emplID,
		"org_id":        orgID,
	})

}

func (c *sharpnrIdentityAuthControllers) UpdateMyProfile(ctx *gin.Context) {

}

func (c *sharpnrIdentityAuthControllers) GetUser(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	res, err := c.IdSvc.GetAuthenticatedUserDetails(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		utils.Failure(ctx, http.StatusUnauthorized, "unauthorised", err)
	}

	utils.Success(ctx, http.StatusOK, "Success", res)
}

func (c *sharpnrIdentityAuthControllers) LogoutUser(ctx *gin.Context) {

	// Clear auth cookies
	c.cookieMgr.Clear(ctx, cookieManager.Auth)
	c.cookieMgr.Clear(ctx, cookieManager.CSRF)
	c.cookieMgr.Clear(ctx, cookieManager.Signup)

	// Send success response
	utils.Success(ctx, http.StatusOK, "User logged out successfully", nil)
}
