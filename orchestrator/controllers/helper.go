package orchestrator_controllers

import (
	"errors"
	"fmt"
	"net/http"
	webMiddleware "sharpnr-engine/middlewares/web"
	sentinelServiceErrs "sharpnr-engine/sentinel/errs"
	sentinelEngine "sharpnr-engine/sentinel_engine"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// failFromServiceError maps a service-layer error to the right HTTP failure:
// a Sentinel deny becomes 403, a "not found in this organization" becomes 404,
// anything else stays 500. Handlers whose service path is guarded by the
// authorization engine respond through this helper.
//
// The not-found cases matter for tenant safety: the repo now scopes every read
// by org, so another tenant's group or policy arrives here as not-found. Letting
// that surface as a 500 would both mislead the caller and hint that the id
// exists somewhere.
func failFromServiceError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, sentinelEngine.ErrAccessDenied):
		utils.Failure(ctx, http.StatusForbidden, err)
	case errors.Is(err, sentinelServiceErrs.GroupNotFoundInOrgErr),
		errors.Is(err, sentinelServiceErrs.RoleNotFoundInOrgErr),
		errors.Is(err, sentinelServiceErrs.PolicyNotFoundInOrgErr),
		errors.Is(err, gorm.ErrRecordNotFound):
		utils.Failure(ctx, http.StatusNotFound, err)
	default:
		utils.Failure(ctx, http.StatusInternalServerError, err)
	}
}

func (c *OrchestratorController) resolveRequestContexts(ctx *gin.Context) (uuid.UUID, uuid.UUID, bool) {
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

// pathUUID extracts and parses the named UUID path parameter. On a malformed
// value it writes the 400 response itself and returns ok=false — callers just
// return.
func pathUUID(ctx *gin.Context, name string) (uuid.UUID, bool) {
	id, err := uuid.Parse(ctx.Param(name))
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest,
			fmt.Errorf("invalid %s format", name))

		return uuid.Nil, false
	}
	return id, true
}

func (c *OrchestratorController) getEmployeeIDFromParams(ctx *gin.Context) (uuid.UUID, bool) {
	panic("")
}

func (c *OrchestratorController) getPropertyIDFromParams(ctx *gin.Context) (propertyID uuid.UUID, isThere bool) {
	return utils.ExtractUUIDParam(ctx, "propertyID")
}

func (c *OrchestratorController) getRestaurntIDFromParams(ctx *gin.Context) (restaurantID uuid.UUID, isThere bool) {
	return utils.ExtractUUIDParam(ctx, "restaurantID")
}

func (c *OrchestratorController) getSentinelGroupIDFromParams(ctx *gin.Context) (sentinelGroupID uuid.UUID, isThere bool) {
	return utils.ExtractUUIDParam(ctx, "groupID")
}
