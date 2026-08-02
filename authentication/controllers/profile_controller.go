package authenticationEngineControllers

import (
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *authenticationControllers) GetMyProfile(ctx *gin.Context) {

	orgID, empID, err := c.getOrgIDAndEmpIDFromContext(ctx)
	if err != nil {
		utils.Failure(ctx, 500, err)
		return
	}

	profile, err := c.authSvc.GetProfile(ctx.Request.Context(), orgID, empID)
	if err != nil {
		utils.Failure(ctx, 500, err)
		return
	}

	utils.Success(ctx, 200, "Profile fetched successfully", profile)

}
