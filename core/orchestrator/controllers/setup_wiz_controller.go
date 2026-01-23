package orchestrator_controllers

import (
	"net/http"
	orchestrator_dtos "sharpnr-engine/core/orchestrator/dtos"
	webMiddleware "sharpnr-engine/middlewares/web"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *OrchestratorController) GetSetupProgress(ctx *gin.Context) {

}

func (c *OrchestratorController) Step_one_AddOrgHQ(ctx *gin.Context) {

	_, exist := ctx.Get(webMiddleware.UserIDContextKey)
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "user not authenticated",
		})
		return
	}

	orgUUIDRaw, exists := ctx.Get(webMiddleware.OrgUUIDContextKey)
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "organization not found in context",
		})
		return
	}

	_, ok := orgUUIDRaw.(uuid.UUID)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "invalid org uuid type",
		})
		return
	}

	var req *orchestrator_dtos.SetupOrganizationHQAddressReqDTO
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "organization not found in context",
		})
		return
	}

}

func (c *OrchestratorController) Step_two_AddProperty(ctx *gin.Context) {

}

func (c *OrchestratorController) Step_three_LegalEntity(ctx *gin.Context) {

}
