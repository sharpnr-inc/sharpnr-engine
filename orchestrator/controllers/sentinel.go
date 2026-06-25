package orchestrator_controllers

import (
	"errors"
	"net/http"

	sentinelDtos "sharpnr-engine/sentinel_engine/dtos"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Sentinel (authorization) endpoints: the managed-policy authoring + role/
// privilege surface the frontend uses to build roles and show what they grant.
// All handlers resolve the tenant context the same way as the rest of the
// orchestrator and delegate to the service.

// ListPermissionCatalogHandler godoc
// @Summary      List the permission catalog
// @Description  Returns every permission the UI can put into a managed policy, grouped by category.
// @Tags         Orchestrator.Sentinel
// @Produce      json
// @Success      200 {object} utils.Response "Catalog"
// @Router       /orchestrator/sentinel/permissions [get]
func (c *OrchestratorController) ListPermissionCatalogHandler(ctx *gin.Context) {
	if _, _, ok := c.resolveRequestContexts(ctx); !ok {
		return
	}
	catalog := c.orchestService.ListPermissionCatalog(ctx.Request.Context())
	utils.Success(ctx, http.StatusOK, "Permission catalog fetched", catalog)
}

// ListPoliciesHandler godoc
// @Summary      List policies
// @Description  Returns the org's Cedar policies (raw policy text + metadata).
// @Tags         Orchestrator.Sentinel
// @Produce      json
// @Success      200 {object} utils.Response "Policies"
// @Router       /orchestrator/sentinel/policies [get]
func (c *OrchestratorController) ListPoliciesHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}
	policies, err := c.orchestService.ListPolicies(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, err)
		return
	}
	utils.Success(ctx, http.StatusOK, "Policies fetched", policies)
}

// CreatePolicyHandler godoc
// @Summary      Create a policy
// @Description  Persists a raw Cedar policy authored in the frontend (revalidated server-side).
// @Tags         Orchestrator.Sentinel
// @Accept       json
// @Produce      json
// @Success      201 {object} utils.Response "Created"
// @Router       /orchestrator/sentinel/policies [post]
func (c *OrchestratorController) CreatePolicyHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}
	req, ok := utils.BindJSON[sentinelDtos.AddPolicyRequestDto](ctx)
	if !ok {
		return
	}
	id, err := c.orchestService.CreatePolicy(ctx.Request.Context(), orgID, emplID, *req)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeValidationFailed, err)
		return
	}
	utils.Success(ctx, http.StatusCreated, "Policy created", gin.H{"id": id})
}

// ListRolesHandler godoc
// @Summary      List roles
// @Description  Roles for the org (e.g. the Add-User dropdown).
// @Tags         Orchestrator.Sentinel
// @Produce      json
// @Success      200 {object} utils.Response "Roles"
// @Router       /orchestrator/sentinel/roles [get]
func (c *OrchestratorController) ListRolesHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}
	roles, err := c.orchestService.ListRoles(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, err)
		return
	}
	utils.Success(ctx, http.StatusOK, "Roles fetched", roles)
}

// CreateRoleHandler godoc
// @Summary      Create a role
// @Description  Creates a role and optionally attaches initial managed policies, generating its Cedar.
// @Tags         Orchestrator.Sentinel
// @Accept       json
// @Produce      json
// @Success      201 {object} utils.Response "Created"
// @Router       /orchestrator/sentinel/roles [post]
func (c *OrchestratorController) CreateRoleHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}
	req, ok := utils.BindJSON[sentinelDtos.CreateRoleInput](ctx)
	if !ok {
		return
	}
	id, err := c.orchestService.CreateRole(ctx.Request.Context(), orgID, emplID, *req)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeValidationFailed, err)
		return
	}
	utils.Success(ctx, http.StatusCreated, "Role created", gin.H{"id": id})
}

// GetRolePrivilegesHandler godoc
// @Summary      Get a role's privileges
// @Description  Returns the role, its attached managed policies, and the union of permissions they grant.
// @Tags         Orchestrator.Sentinel
// @Produce      json
// @Success      200 {object} utils.Response "Role privileges"
// @Router       /orchestrator/sentinel/roles/{roleID} [get]
func (c *OrchestratorController) GetRolePrivilegesHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}
	roleID, err := parseRoleID(ctx)
	if err != nil {
		return
	}
	view, err := c.orchestService.GetRolePrivileges(ctx.Request.Context(), orgID, emplID, roleID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, err)
		return
	}
	utils.Success(ctx, http.StatusOK, "Role privileges fetched", view)
}

// AttachPoliciesToRoleHandler godoc
// @Summary      Attach managed policies to a role
// @Tags         Orchestrator.Sentinel
// @Accept       json
// @Produce      json
// @Success      200 {object} utils.Response "Attached"
// @Router       /orchestrator/sentinel/roles/{roleID}/policies [post]
func (c *OrchestratorController) AttachPoliciesToRoleHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}
	roleID, err := parseRoleID(ctx)
	if err != nil {
		return
	}
	req, ok := utils.BindJSON[sentinelDtos.AttachPoliciesInput](ctx)
	if !ok {
		return
	}
	if err := c.orchestService.AttachPoliciesToRole(ctx.Request.Context(), orgID, emplID, roleID, *req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeValidationFailed, err)
		return
	}
	utils.Success(ctx, http.StatusOK, "Policies attached")
}

// DetachPolicyFromRoleHandler godoc
// @Summary      Detach a managed policy from a role
// @Tags         Orchestrator.Sentinel
// @Produce      json
// @Success      200 {object} utils.Response "Detached"
// @Router       /orchestrator/sentinel/roles/{roleID}/policies/{policyID} [delete]
func (c *OrchestratorController) DetachPolicyFromRoleHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}
	roleID, err := parseRoleID(ctx)
	if err != nil {
		return
	}
	policyID, err := uuid.Parse(ctx.Param("policyID"))
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeValidationFailed, errors.New("invalid policyID format"))
		return
	}
	if err := c.orchestService.DetachPolicyFromRole(ctx.Request.Context(), orgID, emplID, roleID, policyID); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, err)
		return
	}
	utils.Success(ctx, http.StatusOK, "Policy detached")
}

// parseRoleID extracts and validates the :roleID path parameter.
func parseRoleID(ctx *gin.Context) (uuid.UUID, error) {
	id, err := uuid.Parse(ctx.Param("roleID"))
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeValidationFailed, errors.New("invalid roleID format"))
		return uuid.Nil, err
	}
	return id, nil
}
