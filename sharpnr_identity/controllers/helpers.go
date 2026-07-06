package sharpnr_identity_auth_controllers

import (
	"errors"
	"fmt"
	"net/http"
	webMiddleware "sharpnr-engine/middlewares/web"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *sharpnrIdentityAuthControllers) getOrgIDAndEmpIDFromContext(ctx *gin.Context) (uuid.UUID, uuid.UUID, error) {

	// 1. Get OrgID from context
	orgIDVal, ok := ctx.Get(webMiddleware.OrgUUIDContextKey)
	if !ok {
		return uuid.Nil, uuid.Nil, errors.New("org_id not found in context")
	}

	// 2. Get EmpID from context
	empIDVal, ok := ctx.Get(webMiddleware.UserIDContextKey)
	if !ok {
		return uuid.Nil, uuid.Nil, errors.New("employee_id not found in context")
	}

	// 3. Parse OrgID (assuming it's a string)
	orgID, err := uuid.Parse(fmt.Sprint(orgIDVal))
	if err != nil {
		return uuid.Nil, uuid.Nil, fmt.Errorf("invalid org_id format: %w", err)
	}

	// 4. Parse EmpID (assuming it's a string)
	empID, err := uuid.Parse(fmt.Sprint(empIDVal))
	if err != nil {
		return uuid.Nil, uuid.Nil, fmt.Errorf("invalid employee_id format: %w", err)
	}

	return orgID, empID, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (c *sharpnrIdentityAuthControllers) resolveRequestContexts(ctx *gin.Context) (uuid.UUID, uuid.UUID, bool) {
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
