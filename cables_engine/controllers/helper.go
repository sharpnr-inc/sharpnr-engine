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
	case errors.Is(err, cablesEngineErrs.ErrNoChannelFound),
		errors.Is(err, cablesEngineErrs.ErrConversationNotFound),
		errors.Is(err, cablesEngineErrs.ErrMessageNotFound),
		errors.Is(err, cablesEngineErrs.ErrMemberNotFound):
		utils.Failure(ctx, http.StatusNotFound, err)

	// Membership is the authorization model, so "not a member" is the deny.
	// Deliberately the same answer a caller gets for a conversation that does
	// not exist would be a 404 — the two are kept apart here only because the
	// client renders them differently, never to reveal which is which for a
	// conversation the caller could not otherwise see.
	case errors.Is(err, cablesEngineErrs.ErrNotAConversationMember),
		errors.Is(err, cablesEngineErrs.ErrChannelNotJoinable):
		utils.Failure(ctx, http.StatusForbidden, err)

	case errors.Is(err, cablesEngineErrs.ErrCablesServiceNotInitialized),
		errors.Is(err, cablesEngineErrs.ErrCablesAccountIsDisabled),
		errors.Is(err, cablesEngineErrs.ErrCablesServiceAlreadyStarted),
		errors.Is(err, cablesEngineErrs.ErrChannelAlreadyExists),
		errors.Is(err, cablesEngineErrs.ErrChannelArchived):
		utils.Failure(ctx, http.StatusConflict, err)

	case errors.Is(err, cablesEngineErrs.ErrChannelNameRequired),
		errors.Is(err, cablesEngineErrs.ErrInvalidChannelType),
		errors.Is(err, cablesEngineErrs.ErrInvalidChannelVisibility),
		errors.Is(err, cablesEngineErrs.ErrMessageEmpty),
		errors.Is(err, cablesEngineErrs.ErrMessageTooLong),
		errors.Is(err, cablesEngineErrs.ErrCannotDirectMessageSelf),
		errors.Is(err, cablesEngineErrs.ErrTermsNotAccepted):
		utils.Failure(ctx, http.StatusBadRequest, err)

	default:
		utils.Failure(ctx, http.StatusInternalServerError, err)
	}
}

/* ---------------------------------- <***> --------------------------------- */

// resolveUUIDParam reads a uuid out of the path. A malformed one is the
// caller's mistake, so it is a 400 rather than a lookup that cannot match.
func (c *cablesEngineController) resolveUUIDParam(ctx *gin.Context, name string) (uuid.UUID, bool) {
	id, err := uuid.Parse(ctx.Param(name))
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, fmt.Errorf("invalid %s", name))
		return uuid.Nil, false
	}
	return id, true
}

func (c *cablesEngineController) resolveConversationID(ctx *gin.Context) (uuid.UUID, bool) {
	return c.resolveUUIDParam(ctx, "conversationID")
}
