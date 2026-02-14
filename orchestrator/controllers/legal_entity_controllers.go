package orchestrator_controllers

import (
	"net/http"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */
/*           All controllers legal entity management will be defined          */
/* -------------------------------------------------------------------------- */

/* ---------------------------- Add Legal Entity ---------------------------- */

func (c *OrchestratorController) AddLegalEntity(ctx *gin.Context) {

	orgID, _, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	var req *orchestrator_dtos.AddLegalEntity_REQ_DTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeConflict, err)
		return
	}

	if err := c.orchestService.CreateLegalEntity(ctx, orgID, req); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Legal Entity Created successfully")
}

/* -------------------------- list all Legal Entity ------------------------- */

func (c *OrchestratorController) ListAllLegalEntity(ctx *gin.Context) {

}
