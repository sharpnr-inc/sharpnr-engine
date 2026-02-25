package sharpnr_identity_auth_controllers

import (
	"errors"
	"fmt"
	webMiddleware "sharpnr-engine/middlewares/web"

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
