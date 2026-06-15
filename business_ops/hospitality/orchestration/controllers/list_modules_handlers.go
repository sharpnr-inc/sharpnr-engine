package hospitality_orchestration_controllers

import (
	"net/http"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *hospitalityOrchestrationController) ListAllHotelsHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	res, err := c.hospitalityOrchestService.ListAllOrgsHotels(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Something went wrong Try again", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "List of Hotel Fetched Successfully", res)

}

/* -------------------------------- ********* ------------------------------- */

func (c *hospitalityOrchestrationController) ListAllRestaurantHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	res, err := c.hospitalityOrchestService.ListAllOrgsHotels(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Something went wrong Try again", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "List of Hotel Fetched Successfully", res)

}

/* -------------------------------- ********* ------------------------------- */

func (c *hospitalityOrchestrationController) ListAllBarHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	res, err := c.hospitalityOrchestService.ListAllOrgsHotels(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Something went wrong Try again", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "List of Hotel Fetched Successfully", res)

}

/* -------------------------------- ********* ------------------------------- */

func (c *hospitalityOrchestrationController) ListAllInfraHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	res, err := c.hospitalityOrchestService.ListAllOrgsHotels(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Something went wrong Try again", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "List of Hotel Fetched Successfully", res)

}
