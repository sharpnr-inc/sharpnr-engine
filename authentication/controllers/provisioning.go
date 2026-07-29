package authenticationEngineControllers

import (
	"net/http"
	authenticationEngineDtos "sharpnr-engine/authentication/dtos"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *authenticationControllers) VerifyEmployeeProvision(ctx *gin.Context) {
	// we take out the token and verify it if it is expired or not

	req, ok := utils.BindJSON[authenticationEngineDtos.VerifyEmployeeProvisionRequestDto](ctx)
	if !ok {
		utils.Failure(ctx, http.StatusBadRequest, "Something went wrong", nil)
		return
	}

	_, err := c.jwtMgr.VerifyEmployeeProvisioningToken(req.Token)
	if err != nil {
		utils.Failure(ctx, http.StatusUnauthorized, "Invalid or expired provisioning token", err)
		return
	}

}
