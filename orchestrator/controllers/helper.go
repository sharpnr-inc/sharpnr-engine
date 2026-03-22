package orchestrator_controllers

import (
	"errors"
	"fmt"
	"net/http"
	webMiddleware "sharpnr-engine/middlewares/web"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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

func (c *OrchestratorController) getPropertyFromParams(ctx *gin.Context) (uuid.UUID, error) {

	// 3. Param Parsing
	propertyId := ctx.Param("propertyID")
	pId, err := uuid.Parse(propertyId)
	if err != nil {
		// Changed to StatusBadRequest (400) because the user sent a bad UUID string
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeValidationFailed, errors.New("invalid propertyID format"))
		return uuid.Nil, err
	}

	return pId, err
}

func (c *OrchestratorController) getLegalEntityIDFromParams(ctx *gin.Context) (uuid.UUID, error) {

	// 3. Param Parsing
	legalEntityID := ctx.Param("legalEntityID")
	pId, err := uuid.Parse(legalEntityID)
	if err != nil {
		// Changed to StatusBadRequest (400) because the user sent a bad UUID string
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeValidationFailed, errors.New("invalid legalEntityID format"))
		return uuid.Nil, err
	}

	return pId, err
}

func (c *OrchestratorController) getPaymentAccountIDFromParams(ctx *gin.Context) (uuid.UUID, error) {

	// 3. Param Parsing
	paymentAccountID := ctx.Param("paymentAccountID")
	pId, err := uuid.Parse(paymentAccountID)
	if err != nil {
		// Changed to StatusBadRequest (400) because the user sent a bad UUID string
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeValidationFailed, errors.New("invalid payent account id format"))
		return uuid.Nil, err
	}

	return pId, err
}
