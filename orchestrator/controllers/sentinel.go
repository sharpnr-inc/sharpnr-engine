package orchestrator_controllers

import (
	"errors"
	"net/http"

	sentinelDtos "sharpnr-engine/sentinel/dtos"
	sentinelengine_models "sharpnr-engine/sentinel/models"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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

func (c *OrchestratorController) ListResourcesHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}
	resources := c.orchestService.ListResources(ctx.Request.Context(), orgID, emplID)
	utils.Success(ctx, http.StatusOK, "Resources fetched", resources)
}

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

func (c *OrchestratorController) ListPoliciesHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}
	scope := sentinelengine_models.PolicyScope(ctx.Query("scope"))
	policies, err := c.orchestService.ListPolicies(ctx.Request.Context(), orgID, emplID, scope)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeValidationFailed, err)
		return
	}
	utils.Success(ctx, http.StatusOK, "Policies fetched", policies)
}

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
