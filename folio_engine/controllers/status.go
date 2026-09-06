package folioEngineControllers

import (
	"net/http"

	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

// FolioAccountStatusHandler godoc
// @Summary      Folio service status
// @Description  Reports whether the caller's organization has started the Folio (docs, sheets, slides and whiteboards) service and whether that account is switched on. An organization that never started it answers serviceStarted=false with a null id rather than a 404 — the client asks this before it knows. Starting the service is an orchestrator call, not this one.
// @Tags         Folio
// @Accept       json
// @Produce      json
// @Param        Cookie  header    string  true  "Session cookie containing org_id and empl_id"
// @Success      200     {object}  utils.Response{data=folioEngineDtos.FolioAccountStatusResponse} "Folio status fetched successfully"
// @Failure      401     {object}  utils.Response "Unauthorized"
// @Failure      500     {object}  utils.Response "Internal server error"
// @Router       /folio/status [get]
func (f *folioController) FolioAccountStatusHandler(ctx *gin.Context) {
	orgID, emplID, ok := f.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	status, err := f.folioService.GetFolioAccountStatus(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		failFromFolioError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Folio status fetched successfully", status)
}
