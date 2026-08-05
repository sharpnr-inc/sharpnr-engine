package authenticationEngineControllers

import (
	"net/http"
	"sharpnr-engine/pkg/cookieManager"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *authenticationControllers) VerifyAuthentication(ctx *gin.Context) {

	_, _, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	utils.Success(ctx, http.StatusAccepted, "User is authenticated")
}

/* ---------------------------------- <***> --------------------------------- */

func (c *authenticationControllers) GetUser(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	res, err := c.authSvc.GetAuthenticatedUserDetails(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		utils.Failure(ctx, http.StatusUnauthorized, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Success", res)
}

/* ---------------------------------- <***> --------------------------------- */

func (c *authenticationControllers) LogoutUser(ctx *gin.Context) {

	// Clear auth cookies
	c.cookieMgr.Clear(ctx, cookieManager.Auth)
	c.cookieMgr.Clear(ctx, cookieManager.CSRF)
	c.cookieMgr.Clear(ctx, cookieManager.Signup)

	// Send success response
	utils.Success(ctx, http.StatusOK, "User logged out successfully", nil)
}
