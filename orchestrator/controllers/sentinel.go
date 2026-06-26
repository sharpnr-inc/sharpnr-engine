package orchestrator_controllers

import (
	"errors"
	"net/http"

	sentinelDtos "sharpnr-engine/sentinel_engine/dtos"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Sentinel (authorization) endpoints: the policy authoring + role/privilege
// surface, and the resource-centric policy builder. All handlers resolve the
// tenant context the same way as the rest of the orchestrator and delegate to
// the service.

// ListPrincipalsHandler godoc
// @Summary      List the principal palette
// @Description  Returns the principal entity types (User/Role/Group) plus the org's roles and groups the policy builder can scope a principal to.
// @Tags         Orchestrator.Sentinel
// @Produce      json
// @Success      200 {object} utils.Response "Principals"
// @Router       /orchestrator/sentinel/principals [get]
func (c *OrchestratorController) ListPrincipalsHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}
	principals, err := c.orchestService.ListPrincipals(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, err)
		return
	}
	utils.Success(ctx, http.StatusOK, "Principals fetched", principals)
}

// ListResourcesHandler godoc
// @Summary      List resource types
// @Description  Returns the resource types a policy can be scoped to, each with its actions and attributes (the policy builder's resource picker).
// @Tags         Orchestrator.Sentinel
// @Produce      json
// @Success      200 {object} utils.Response "Resources"
// @Router       /orchestrator/sentinel/resources [get]
func (c *OrchestratorController) ListResourcesHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}
	resources := c.orchestService.ListResources(ctx.Request.Context(), orgID, emplID)
	utils.Success(ctx, http.StatusOK, "Resources fetched", resources)
}

// GetResourceActionsHandler godoc
// @Summary      List a resource type's actions
// @Description  Returns the actions applicable to the given resource type.
// @Tags         Orchestrator.Sentinel
// @Produce      json
// @Param        type path string true "resource entity type, e.g. Property"
// @Success      200 {object} utils.Response "Actions"
// @Router       /orchestrator/sentinel/resources/{type}/actions [get]
func (c *OrchestratorController) GetResourceActionsHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}
	actions, err := c.orchestService.GetResourceActions(ctx.Request.Context(), orgID, emplID, ctx.Param("type"))
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeValidationFailed, err)
		return
	}
	utils.Success(ctx, http.StatusOK, "Resource actions fetched", actions)
}

// GetResourceAttributesHandler godoc
// @Summary      List a resource type's attributes
// @Description  Returns the ABAC attributes of the given resource type (for `when { resource.<attr> ... }`).
// @Tags         Orchestrator.Sentinel
// @Produce      json
// @Param        type path string true "resource entity type, e.g. Property"
// @Success      200 {object} utils.Response "Attributes"
// @Router       /orchestrator/sentinel/resources/{type}/attributes [get]
func (c *OrchestratorController) GetResourceAttributesHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}
	attrs, err := c.orchestService.GetResourceAttributes(ctx.Request.Context(), orgID, emplID, ctx.Param("type"))
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeValidationFailed, err)
		return
	}
	utils.Success(ctx, http.StatusOK, "Resource attributes fetched", attrs)
}

// ListResourceInstancesHandler godoc
// @Summary      List a resource type's live instances
// @Description  Returns the org's live instances of the given resource type (e.g. its properties), for `resource in <Type>::"<id>"`.
// @Tags         Orchestrator.Sentinel
// @Produce      json
// @Param        type path string true "resource entity type, e.g. Property"
// @Success      200 {object} utils.Response "Instances"
// @Router       /orchestrator/sentinel/resources/{type}/instances [get]
func (c *OrchestratorController) ListResourceInstancesHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}
	instances, err := c.orchestService.ListResourceInstances(ctx.Request.Context(), orgID, emplID, ctx.Param("type"))
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeValidationFailed, err)
		return
	}
	utils.Success(ctx, http.StatusOK, "Resource instances fetched", instances)
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
