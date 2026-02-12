package sharpnr_identity_auth_controllers

import (
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *sharpnrIdentityAuthControllers) GetMyProfile(ctx *gin.Context) {

	orgID, empID, err := c.getOrgIDAndEmpIDFromContext(ctx)
	if err != nil {
		utils.Failure(ctx, 500, utils.CodeInternalError, err)
		return
	}

	profile, err := c.IdSvc.GetProfile(ctx.Request.Context(), orgID, empID)
	if err != nil {
		utils.Failure(ctx, 500, utils.CodeInternalError, err)
		return
	}

	utils.Success(ctx, 200, "Profile fetched successfully", profile)

}


func (c *sharpnrIdentityAuthControllers) GetUserProfile(ctx *gin.Context) {
	
}