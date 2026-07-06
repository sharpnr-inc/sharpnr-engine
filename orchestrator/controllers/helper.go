package orchestrator_controllers

import (
	"errors"
	"fmt"
	"net/http"
	webMiddleware "sharpnr-engine/middlewares/web"
	sentinelEngine "sharpnr-engine/sentinel_engine"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// failFromServiceError maps a service-layer error to the right HTTP failure:
// a Sentinel deny becomes 403, anything else stays 500. Handlers whose service
// path is guarded by the authorization engine respond through this helper.
func failFromServiceError(ctx *gin.Context, err error) {
	if errors.Is(err, sentinelEngine.ErrAccessDenied) {
		utils.Failure(ctx, http.StatusForbidden, utils.CodeForbidden, err)
		return
	}
	utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, err)
}

func (c *OrchestratorController) resolveRequestContexts(ctx *gin.Context) (uuid.UUID, uuid.UUID, bool) {
	// 1. Extract and Parse Employee ID
	empIDRaw, exists := ctx.Get(webMiddleware.UserIDContextKey)
	if !exists {
		utils.Failure(ctx, http.StatusUnauthorized, utils.CodeUnauthorized, errors.New("user not authenticated"))
		return uuid.Nil, uuid.Nil, false
	}

	empID, err := uuid.Parse(fmt.Sprint(empIDRaw))
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, errors.New("invalid employee uuid format"))
		return uuid.Nil, uuid.Nil, false
	}

	// 2. Extract and Parse Org ID
	orgUUIDRaw, exists := ctx.Get(webMiddleware.OrgUUIDContextKey)
	if !exists {
		utils.Failure(ctx, http.StatusUnauthorized, utils.CodeUnauthorized, errors.New("organization not found in context"))
		return uuid.Nil, uuid.Nil, false
	}

	orgID, err := uuid.Parse(fmt.Sprint(orgUUIDRaw))
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, errors.New("invalid org uuid format"))
		return uuid.Nil, uuid.Nil, false
	}

	return orgID, empID, true
}

// pathUUID extracts and parses the named UUID path parameter. On a malformed
// value it writes the 400 response itself and returns ok=false — callers just
// return.
func pathUUID(ctx *gin.Context, name string) (uuid.UUID, bool) {
	id, err := uuid.Parse(ctx.Param(name))
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeValidationFailed,
			fmt.Errorf("invalid %s format", name))
		return uuid.Nil, false
	}
	return id, true
}
