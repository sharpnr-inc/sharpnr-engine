package orchestrator_controllers

import (
	"errors"
	"net/http"
	webMiddleware "sharpnr-engine/middlewares/web"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *OrchestratorController) ListAllOrgProperty(ctx *gin.Context) {
	// 1️⃣ Auth check
	if _, exists := ctx.Get(webMiddleware.UserIDContextKey); !exists {
		utils.Failure(
			ctx,
			http.StatusUnauthorized,
			utils.CodeUnauthorized,
			errors.New("user not authenticated"),
		)
		return
	}

	// 2️⃣ Org context check
	orgUUIDRaw, exists := ctx.Get(webMiddleware.OrgUUIDContextKey)
	if !exists {
		utils.Failure(
			ctx,
			http.StatusUnauthorized,
			utils.CodeUnauthorized,
			errors.New("organization not found in context"),
		)
		return
	}

	orgID, ok := orgUUIDRaw.(uuid.UUID)
	if !ok {
		utils.Failure(
			ctx,
			http.StatusInternalServerError,
			utils.CodeInternalError,
			errors.New("invalid org uuid type"),
		)
		return
	}

	res, err := c.orchestService.GetAllPropertiesOfOrg(ctx.Request.Context(), orgID)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeConflict, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Properties fetch successfully", res)

}
