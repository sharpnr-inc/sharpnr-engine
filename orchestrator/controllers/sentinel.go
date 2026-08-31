package orchestrator_controllers

import (
	"errors"
	"net/http"
	"strconv"

	sentinelDtos "sharpnr-engine/sentinel/dtos"
	sentinelAggregatorContracts "sharpnr-engine/sentinel_aggregators/contracts"
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
		utils.Failure(ctx, http.StatusInternalServerError, err)
		return
	}
	utils.Success(ctx, http.StatusOK, "Principals fetched", principals)
}

// GetAuthzContextPaletteHandler serves the policy builder's context vocabulary:
// the canonical context.* and principal.* keys conditions can be authored
// against, with types and closed value sets for rendering pickers.
func (c *OrchestratorController) GetAuthzContextPaletteHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}
	palette, err := c.orchestService.GetAuthzContextPalette(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err)
		return
	}
	utils.Success(ctx, http.StatusOK, "Context palette fetched", palette)
}

// ListAuthzResourcesHandler serves the policy builder's resource picker: every
// resource type registered with the sentinel aggregator.
func (c *OrchestratorController) ListAuthzResourcesHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}
	resources, err := c.orchestService.ListAuthzResources(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err)
		return
	}
	utils.Success(ctx, http.StatusOK, "Resources fetched", resources)
}

// GetAuthzResourceDetailsHandler serves one resource's full palette: actions,
// attributes and the org's live instances.
func (c *OrchestratorController) GetAuthzResourceDetailsHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}
	resourceType := ctx.Param("resourceType")

	// The picker searches server-side, and passes ?id= repeatedly to resolve
	// the specific instances a stored policy is scoped to — without that it
	// would have to fetch the whole list just to label four of them.
	filter := sentinelAggregatorContracts.InstanceFilter{
		Query:  ctx.Query("q"),
		IDs:    ctx.QueryArray("id"),
		Cursor: ctx.Query("cursor"),
	}
	if limit, err := strconv.Atoi(ctx.Query("limit")); err == nil && limit > 0 {
		filter.Limit = limit
	}

	details, err := c.orchestService.GetAuthzResourceDetails(ctx.Request.Context(), orgID, emplID, resourceType, filter)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, err)
		return
	}
	utils.Success(ctx, http.StatusOK, "Resource details fetched", details)
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
		utils.Failure(ctx, http.StatusBadRequest, err)
		return
	}
	utils.Success(ctx, http.StatusCreated, "Policy created", gin.H{"id": id})
}

/* ---------------------------------- <***> --------------------------------- */
// ListOrgPoliciesHandler serves the policies registered against the
// organization itself — those that reach every employee. It reads the org
// register (org_sentinel_policies), NOT the org's full policy catalog: a policy
// the org owns but has not registered anywhere reaches nobody and does not
// appear here.
func (c *OrchestratorController) ListOrgPoliciesHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}
	policies, err := c.orchestService.ListOrgPolicies(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		failFromServiceError(ctx, err)
		return
	}
	utils.Success(ctx, http.StatusOK, "Organization policies fetched", policies)
}

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) GetPolicyHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}
	policyID, err := uuid.Parse(ctx.Param("policyID"))
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, err)
		return
	}
	policy, err := c.orchestService.GetPolicy(ctx.Request.Context(), orgID, emplID, policyID)
	if err != nil {
		utils.Failure(ctx, http.StatusNotFound, err)
		return
	}
	utils.Success(ctx, http.StatusOK, "Policy fetched", policy)
}

func (c *OrchestratorController) ListRolesHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}
	roles, err := c.orchestService.ListRoles(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err)
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
		utils.Failure(ctx, http.StatusBadRequest, err)
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
		utils.Failure(ctx, http.StatusInternalServerError, err)
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
		utils.Failure(ctx, http.StatusBadRequest, err)
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
		utils.Failure(ctx, http.StatusBadRequest, errors.New("invalid policyID format"))
		return
	}
	if err := c.orchestService.DetachPolicyFromRole(ctx.Request.Context(), orgID, emplID, roleID, policyID); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err)
		return
	}
	utils.Success(ctx, http.StatusOK, "Policy detached")
}

// parseRoleID extracts and validates the :roleID path parameter.
func parseRoleID(ctx *gin.Context) (uuid.UUID, error) {
	id, err := uuid.Parse(ctx.Param("roleID"))
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, errors.New("invalid roleID format"))
		return uuid.Nil, err
	}
	return id, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) ListOrgSentinelUserGroups(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	res, err := c.orchestService.ListSentinelUserGroups(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err)
	}

	utils.Success(ctx, http.StatusOK, "User Group Fetched successfully", res)
}

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) AddNewSentinelGroupHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	req, ok := utils.BindJSON[sentinelDtos.CreateGroupsRequestDto](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.CreateNewSentinelGroup(ctx.Request.Context(), orgID, emplID, req); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err)
	}

	utils.Success(ctx, http.StatusOK, "Group Created Successfully")

}

// SentinelGroupDetailHandler serves one group: its details, the policies
// registered against it, and its members.
func (c *OrchestratorController) SentinelGroupDetailHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	groupID, available := c.getSentinelGroupIDFromParams(ctx)
	if !available {
		utils.Failure(ctx, http.StatusBadRequest, errors.New("a valid groupID is required"))
		return
	}

	res, err := c.orchestService.GetSentinelGroupDetail(ctx.Request.Context(), orgID, emplID, groupID)
	if err != nil {
		failFromServiceError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Group details fetched", res)
}
