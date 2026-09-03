package cablesEngineControllers

import (
	"errors"
	"fmt"
	"net/http"

	cablesEngineErrs "sharpnr-engine/cables_engine/errs"
	webMiddleware "sharpnr-engine/middlewares/web"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *cablesEngineController) resolveRequestContexts(ctx *gin.Context) (uuid.UUID, uuid.UUID, bool) {
	// 1. Extract and Parse Employee ID
	empIDRaw, exists := ctx.Get(webMiddleware.UserIDContextKey)
	if !exists {
		utils.Failure(ctx, http.StatusUnauthorized, errors.New("user not authenticated"))
		return uuid.Nil, uuid.Nil, false
	}

	empID, err := uuid.Parse(fmt.Sprint(empIDRaw))
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, errors.New("invalid employee uuid format"))
		return uuid.Nil, uuid.Nil, false
	}

	// 2. Extract and Parse Org ID
	orgUUIDRaw, exists := ctx.Get(webMiddleware.OrgUUIDContextKey)
	if !exists {
		utils.Failure(ctx, http.StatusUnauthorized, errors.New("organization not found in context"))
		return uuid.Nil, uuid.Nil, false
	}

	orgID, err := uuid.Parse(fmt.Sprint(orgUUIDRaw))
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, errors.New("invalid org uuid format"))
		return uuid.Nil, uuid.Nil, false
	}

	return orgID, empID, true
}

/* ---------------------------------- <***> --------------------------------- */

// failFromCablesError maps the module's sentinel errors onto status codes.
//
// It is the client-facing twin of the orchestrator's mapper of the same name;
// the two are separate because the packages are, and this one only has to cover
// the errors the product API can actually raise.
func failFromCablesError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, cablesEngineErrs.ErrNoChannelFound):
		utils.Failure(ctx, http.StatusNotFound, err)
	case errors.Is(err, cablesEngineErrs.ErrCablesServiceNotInitialized),
		errors.Is(err, cablesEngineErrs.ErrCablesAccountIsDisabled):
		utils.Failure(ctx, http.StatusConflict, err)
	default:
		utils.Failure(ctx, http.StatusInternalServerError, err)
	}
}
