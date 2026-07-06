package iamEngineControllers

import (
	"errors"
	"fmt"
	"net/http"

	webMiddleware "sharpnr-engine/middlewares/web"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *iamEngineControllers) GetMyProfile(ctx *gin.Context) {

	orgID, empID, err := c.getOrgIDAndEmpIDFromContext(ctx)
	if err != nil {
		utils.Failure(ctx, http.StatusUnauthorized, utils.CodeUnauthorized, err)
		return
	}

	profile, err := c.iamEngineService.GetProfile(ctx.Request.Context(), orgID, empID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Profile fetched successfully", profile)
}

/* ---------------------------------- <***> --------------------------------- */

func (c *iamEngineControllers) getOrgIDAndEmpIDFromContext(ctx *gin.Context) (uuid.UUID, uuid.UUID, error) {

	orgIDVal, ok := ctx.Get(webMiddleware.OrgUUIDContextKey)
	if !ok {
		return uuid.Nil, uuid.Nil, errors.New("org_id not found in context")
	}

	empIDVal, ok := ctx.Get(webMiddleware.UserIDContextKey)
	if !ok {
		return uuid.Nil, uuid.Nil, errors.New("employee_id not found in context")
	}

	orgID, err := uuid.Parse(fmt.Sprint(orgIDVal))
	if err != nil {
		return uuid.Nil, uuid.Nil, fmt.Errorf("invalid org_id format: %w", err)
	}

	empID, err := uuid.Parse(fmt.Sprint(empIDVal))
	if err != nil {
		return uuid.Nil, uuid.Nil, fmt.Errorf("invalid employee_id format: %w", err)
	}

	return orgID, empID, nil
}
