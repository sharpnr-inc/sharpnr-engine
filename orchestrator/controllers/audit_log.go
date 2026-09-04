package orchestrator_controllers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	auditEngineDtos "sharpnr-engine/audit_log_engine/dto"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Query-string validation errors for the audit history read.
var (
	errInvalidLimit        = errors.New("limit must be a non-negative number")
	errInvalidOffset       = errors.New("offset must be a non-negative number")
	errInvalidResourceID   = errors.New("resourceId must be a valid uuid")
	errInvalidActorID      = errors.New("actorId must be a valid uuid")
	errInvalidResourceType = errors.New("resourceType is required")
)

/* ---------------------------------- <***> --------------------------------- */

// ListOrgAuditLogsHandler godoc
// @Summary      List the organization's audit history
// @Description  Returns one page of the organization's event history, newest first, with the filters the history screen offers. limit defaults to 50 and is capped at 200; page through with offset. The response carries the pagination block describing the page that was actually served.
// @Tags         Orchestrator.AuditLogs
// @Accept       json
// @Produce      json
// @Param        Cookie        header    string  true   "Session cookie containing org_id and empl_id"
// @Param        limit         query     int     false  "Page size (default 50, max 200)"
// @Param        offset        query     int     false  "Rows to skip (default 0)"
// @Param        module        query     string  false  "Emitting module, e.g. departments"
// @Param        action        query     string  false  "Event name, e.g. department.created"
// @Param        resourceType  query     string  false  "Resource the event is about, e.g. org_department"
// @Param        resourceId    query     string  false  "Resource id (UUID)"
// @Param        actorId       query     string  false  "Employee who performed the action (UUID)"
// @Success      200           {object}  utils.Response{data=auditEngineDtos.AuditLogPageResponseDto} "Audit logs fetched successfully"
// @Failure      400           {object}  utils.Response "Invalid query parameters"
// @Failure      401           {object}  utils.Response "Unauthorized"
// @Failure      500           {object}  utils.Response "Internal server error"
// @Router       /orchestrator/audit-logs [get]
func (c *OrchestratorController) ListOrgAuditLogsHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	query, ok := auditLogQueryFromRequest(ctx)
	if !ok {
		return
	}

	page, err := c.orchestService.ListOrgAuditLogs(ctx.Request.Context(), orgID, emplID, query)
	if err != nil {
		failFromServiceError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Audit logs fetched successfully", page)
}

/* ---------------------------------- <***> --------------------------------- */

// GetResourceAuditHistoryHandler godoc
// @Summary      Get one entity's audit trail
// @Description  Returns everything that ever happened to a single entity — one department, one customer — newest first. Not paginated: a single entity's trail is bounded by how often it was touched.
// @Tags         Orchestrator.AuditLogs
// @Accept       json
// @Produce      json
// @Param        Cookie        header    string  true  "Session cookie containing org_id and empl_id"
// @Param        resourceType  path      string  true  "Resource type, e.g. org_department"
// @Param        resourceID    path      string  true  "Resource ID (UUID)"
// @Success      200           {object}  utils.Response{data=[]auditEngineDtos.AuditLogEntry} "Audit trail fetched successfully"
// @Failure      400           {object}  utils.Response "Invalid resource type or id"
// @Failure      401           {object}  utils.Response "Unauthorized"
// @Failure      500           {object}  utils.Response "Internal server error"
// @Router       /orchestrator/audit-logs/resource/{resourceType}/{resourceID} [get]
func (c *OrchestratorController) GetResourceAuditHistoryHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	resourceType := strings.TrimSpace(ctx.Param("resourceType"))
	if resourceType == "" {
		utils.Failure(ctx, http.StatusBadRequest, errInvalidResourceType)
		return
	}

	resourceID, ok := pathUUID(ctx, "resourceID")
	if !ok {
		return
	}

	logs, err := c.orchestService.GetResourceAuditHistory(ctx.Request.Context(), orgID, emplID, resourceType, resourceID)
	if err != nil {
		failFromServiceError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Audit trail fetched successfully", logs)
}

/* ---------------------------------- <***> --------------------------------- */

// auditLogQueryFromRequest reads the history filters off the query string.
//
// A malformed value is a 400 rather than a silently ignored filter: dropping a
// filter would answer with somebody else's rows — more history than the caller
// asked to see, which is the wrong direction to fail in for an audit trail.
func auditLogQueryFromRequest(ctx *gin.Context) (auditEngineDtos.ListLogsQuery, bool) {
	query := auditEngineDtos.ListLogsQuery{
		Module:       strings.TrimSpace(ctx.Query("module")),
		Action:       strings.TrimSpace(ctx.Query("action")),
		ResourceType: strings.TrimSpace(ctx.Query("resourceType")),
	}

	if raw := ctx.Query("limit"); raw != "" {
		limit, err := strconv.Atoi(raw)
		if err != nil || limit < 0 {
			utils.Failure(ctx, http.StatusBadRequest, errInvalidLimit)
			return query, false
		}
		query.Limit = limit
	}

	if raw := ctx.Query("offset"); raw != "" {
		offset, err := strconv.Atoi(raw)
		if err != nil || offset < 0 {
			utils.Failure(ctx, http.StatusBadRequest, errInvalidOffset)
			return query, false
		}
		query.Offset = offset
	}

	if raw := ctx.Query("resourceId"); raw != "" {
		resourceID, err := uuid.Parse(raw)
		if err != nil {
			utils.Failure(ctx, http.StatusBadRequest, errInvalidResourceID)
			return query, false
		}
		query.ResourceID = &resourceID
	}

	if raw := ctx.Query("actorId"); raw != "" {
		actorID, err := uuid.Parse(raw)
		if err != nil {
			utils.Failure(ctx, http.StatusBadRequest, errInvalidActorID)
			return query, false
		}
		query.ActorID = &actorID
	}

	return query, true
}
