package siloEngineControllers

import (
	"net/http"

	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

// SiloAccountStatusHandler godoc
// @Summary      Silo service status
// @Description  Reports whether the caller's organization has started the Silo (storage and data) service and whether that account is switched on. An organization that never started it answers serviceStarted=false with a null id rather than a 404 — the client asks this before it knows. Starting the service is an orchestrator call, not this one.
// @Tags         Silo
// @Accept       json
// @Produce      json
// @Param        Cookie  header    string  true  "Session cookie containing org_id and empl_id"
// @Success      200     {object}  utils.Response{data=siloEngineDtos.SiloAccountStatusResponse} "Silo status fetched successfully"
// @Failure      401     {object}  utils.Response "Unauthorized"
// @Failure      500     {object}  utils.Response "Internal server error"
// @Router       /silo/status [get]
func (s *siloController) SiloAccountStatusHandler(ctx *gin.Context) {
	orgID, emplID, ok := s.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	status, err := s.siloService.GetSiloAccountStatus(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		failFromSiloError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Silo status fetched successfully", status)
}
