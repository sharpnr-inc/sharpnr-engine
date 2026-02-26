package orchestrator_controllers

import (
	"net/http"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *OrchestratorController) GetPropertiesOfRestaurant(ctx *gin.Context) {

	orgID, empID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	result, err := c.orchestService.GetPropertiesAndRestaurant(ctx.Request.Context(), orgID, empID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Properties and restaurants fetched successfully", result)
}

func (c *OrchestratorController) GetAllRestaurants(ctx *gin.Context) {
	orgID, empID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	results, err := c.orchestService.GetAllRestaurant(ctx.Request.Context(), orgID, empID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "restaurants fetched successfully", results)
}

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) GetOrgLauncher(ctx *gin.Context) {

	orgID, empID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	res, err := c.orchestService.GetSharpnrLauncherData(ctx.Request.Context(), orgID, empID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Launcher Fetched Successfully", res)

}
