package authenticationEngineControllers

import (
	"fmt"
	"net/http"
	authenticationEngineErrs "sharpnr-engine/authentication/errs"
	webMiddleware "sharpnr-engine/middlewares/web"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *authenticationControllers) getOrgIDAndEmpIDFromContext(ctx *gin.Context) (uuid.UUID, uuid.UUID, error) {

	// 1. Get OrgID from context
	orgIDVal, ok := ctx.Get(webMiddleware.OrgUUIDContextKey)
	if !ok {
		return uuid.Nil, uuid.Nil, authenticationEngineErrs.OrgIDNotInContextErr
	}

	// 2. Get EmpID from context
	empIDVal, ok := ctx.Get(webMiddleware.UserIDContextKey)
	if !ok {
		return uuid.Nil, uuid.Nil, authenticationEngineErrs.EmployeeIDNotInContextErr
	}

	// 3. Parse OrgID (assuming it's a string)
	orgID, err := uuid.Parse(fmt.Sprint(orgIDVal))
	if err != nil {
		return uuid.Nil, uuid.Nil,
			fmt.Errorf("%w: %v", authenticationEngineErrs.InvalidOrgIDFormatErr, err)
	}

	// 4. Parse EmpID (assuming it's a string)
	empID, err := uuid.Parse(fmt.Sprint(empIDVal))
	if err != nil {
		return uuid.Nil, uuid.Nil,
			fmt.Errorf("%w: %v", authenticationEngineErrs.InvalidEmployeeIDFormatErr, err)
	}

	return orgID, empID, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (c *authenticationControllers) resolveRequestContexts(ctx *gin.Context) (uuid.UUID, uuid.UUID, bool) {
	// 1. Extract and Parse Employee ID
	empIDRaw, exists := ctx.Get(webMiddleware.UserIDContextKey)
	if !exists {
		utils.Failure(ctx, http.StatusUnauthorized, authenticationEngineErrs.UserNotAuthenticatedErr)
		return uuid.Nil, uuid.Nil, false
	}

	empID, err := uuid.Parse(fmt.Sprint(empIDRaw))
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, authenticationEngineErrs.InvalidEmployeeIDFormatErr)
		return uuid.Nil, uuid.Nil, false
	}

	// 2. Extract and Parse Org ID
	orgUUIDRaw, exists := ctx.Get(webMiddleware.OrgUUIDContextKey)
	if !exists {
		utils.Failure(ctx, http.StatusUnauthorized, authenticationEngineErrs.OrgIDNotInContextErr)
		return uuid.Nil, uuid.Nil, false
	}

	orgID, err := uuid.Parse(fmt.Sprint(orgUUIDRaw))
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, authenticationEngineErrs.InvalidOrgIDFormatErr)
		return uuid.Nil, uuid.Nil, false
	}

	return orgID, empID, true
}
