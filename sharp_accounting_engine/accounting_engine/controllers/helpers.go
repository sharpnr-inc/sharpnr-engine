package accounting_controllers

import (
	"errors"
	"fmt"
	"net/http"
	webMiddleware "sharpnr-engine/middlewares/web"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *AccountingController) resolveRequestContext(ctx *gin.Context) (orgID uuid.UUID, employeeID uuid.UUID, ok bool) {
	// 1. Resolve and Parse Employee/User ID
	userIDRaw, exists := ctx.Get(webMiddleware.UserIDContextKey)
	if !exists {
		utils.Failure(ctx, http.StatusUnauthorized, errors.New("user not authenticated"))
		return uuid.Nil, uuid.Nil, false
	}

	// Convert interface to string and parse
	var err error
	employeeID, err = uuid.Parse(fmt.Sprintf("%v", userIDRaw))
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, errors.New("invalid user uuid format"))
		return uuid.Nil, uuid.Nil, false
	}

	// 2. Resolve and Parse Org ID
	orgUUIDRaw, exists := ctx.Get(webMiddleware.OrgUUIDContextKey)
	if !exists {
		utils.Failure(ctx, http.StatusUnauthorized, errors.New("organization not found in context"))
		return uuid.Nil, uuid.Nil, false
	}

	// Convert interface to string and parse
	orgID, err = uuid.Parse(fmt.Sprintf("%v", orgUUIDRaw))
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, errors.New("invalid org uuid format"))
		return uuid.Nil, uuid.Nil, false
	}

	return orgID, employeeID, true
}

/* ---------------------------------- <***> --------------------------------- */

// resolveLegalEntityID extracts and parses the "legalEntityID" from the URL parameters.
func (c *AccountingController) resolveLegalEntityID(ctx *gin.Context) (uuid.UUID, bool) {
	legalEntityIDStr := ctx.Param("legalEntityID")
	parsedID, err := uuid.Parse(legalEntityIDStr)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, errors.New("invalid legalEntityID format"))
		return uuid.Nil, false
	}
	return parsedID, true
}

func (c *AccountingController) resolveAccountID(ctx *gin.Context) (uuid.UUID, bool) {
	accountIDStr := ctx.Param("accountID")
	parsedID, err := uuid.Parse(accountIDStr)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, errors.New("invalid accountID format"))
		return uuid.Nil, false
	}
	return parsedID, true
}

func (c *AccountingController) resolveJournalEntryID(ctx *gin.Context) (uuid.UUID, bool) {
	entryIDStr := ctx.Param("entryID")
	parsedID, err := uuid.Parse(entryIDStr)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, errors.New("invalid entryID format"))
		return uuid.Nil, false
	}
	return parsedID, true
}
