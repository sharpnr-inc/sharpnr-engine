package siloEngineControllers

import (
	"errors"
	"fmt"
	"net/http"

	webMiddleware "sharpnr-engine/middlewares/web"
	"sharpnr-engine/shared/utils"
	siloEngineErrs "sharpnr-engine/silo_engine/errs"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// resolveRequestContexts pulls the caller's org and employee out of the session
// the auth middleware stamped onto the gin context. Both are required — there is
// no Silo read that is not scoped to one tenant.
func (s *siloController) resolveRequestContexts(ctx *gin.Context) (uuid.UUID, uuid.UUID, bool) {
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

// failFromSiloError maps the module's sentinel errors onto status codes.
//
// It is the client-facing twin of the orchestrator's mapper of the same name;
// the two are separate because the packages are, and this one only has to cover
// the errors the product API can actually raise.
//
// "Not initialized" and "disabled" are both 409 rather than 404: the org is real
// and the caller is allowed to ask, but the service is in a state that refuses
// the request — and the fix differs, which is why the two stay apart all the way
// to the client.
func failFromSiloError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, siloEngineErrs.ErrSiloServiceNotInitialized),
		errors.Is(err, siloEngineErrs.ErrSiloAccountIsDisabled),
		errors.Is(err, siloEngineErrs.ErrSiloServiceAlreadyStarted):
		utils.Failure(ctx, http.StatusConflict, err)

	case errors.Is(err, siloEngineErrs.ErrTermsNotAccepted):
		utils.Failure(ctx, http.StatusBadRequest, err)

	default:
		utils.Failure(ctx, http.StatusInternalServerError, err)
	}
}
