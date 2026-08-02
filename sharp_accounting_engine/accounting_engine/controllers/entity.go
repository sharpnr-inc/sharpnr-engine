package accounting_controllers

import (
	"net/http"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *AccountingController) GetAllLegalEntities(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContext(ctx)
	if !ok {
		return // Error response already sent in resolveRequestContext
	}

	response, err := c.accSvc.GetLegalEntities(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err)
	}

	utils.Success(ctx, http.StatusOK, "legal entities fetched successfully", response)
}
