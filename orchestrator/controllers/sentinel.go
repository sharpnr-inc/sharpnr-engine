package orchestrator_controllers

import (
	"errors"
	"net/http"

	sentinelservice "sharpnr-engine/sentinel_engine/service"
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

// ListManagedPoliciesHandler godoc
// @Summary      List managed policies
// @Description  Returns the managed policies an org can attach to roles (its own + system), with permissions expanded.
// @Tags         Orchestrator.Sentinel
// @Produce      json
// @Success      200 {object} utils.Response "Managed policies"
// @Router       /orchestrator/sentinel/policies [get]
func (c *OrchestratorController) ListManagedPoliciesHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}
	policies, err := c.orchestService.ListManagedPolicies(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, err)
		return
	}
	utils.Success(ctx, http.StatusOK, "Managed policies fetched", policies)
}

// CreateManagedPolicyHandler godoc
// @Summary      Create a managed policy
// @Tags         Orchestrator.Sentinel
// @Accept       json
// @Produce      json
// @Success      201 {object} utils.Response "Created"
// @Router       /orchestrator/sentinel/policies [post]
func (c *OrchestratorController) CreateManagedPolicyHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}
	req, ok := utils.BindJSON[sentinelservice.CreateManagedPolicyInput](ctx)
	if !ok {
		return
	}
	id, err := c.orchestService.CreateManagedPolicy(ctx.Request.Context(), orgID, emplID, *req)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeValidationFailed, err)
		return
	}
	utils.Success(ctx, http.StatusCreated, "Managed policy created", gin.H{"id": id})
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
	req, ok := utils.BindJSON[sentinelservice.CreateRoleInput](ctx)
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
	req, ok := utils.BindJSON[sentinelservice.AttachPoliciesInput](ctx)
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
