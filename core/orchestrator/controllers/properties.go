package orchestrator_controllers

import (
	"errors"
	"net/http"
	orchestrator_dtos "sharpnr-engine/core/orchestrator/dtos"
	webMiddleware "sharpnr-engine/middlewares/web"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *OrchestratorController) ListAllOrgProperty(ctx *gin.Context) {
	// 1️⃣ Auth check
	if _, exists := ctx.Get(webMiddleware.UserIDContextKey); !exists {
		utils.Failure(
			ctx,
			http.StatusUnauthorized,
			utils.CodeUnauthorized,
			errors.New("user not authenticated"),
		)
		return
	}

	// 2️⃣ Org context check
	orgUUIDRaw, exists := ctx.Get(webMiddleware.OrgUUIDContextKey)
	if !exists {
		utils.Failure(
			ctx,
			http.StatusUnauthorized,
			utils.CodeUnauthorized,
			errors.New("organization not found in context"),
		)
		return
	}

	orgID, ok := orgUUIDRaw.(uuid.UUID)
	if !ok {
		utils.Failure(
			ctx,
			http.StatusInternalServerError,
			utils.CodeInternalError,
			errors.New("invalid org uuid type"),
		)
		return
	}

	res, err := c.orchestService.GetAllPropertiesOfOrg(ctx.Request.Context(), orgID)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeConflict, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Properties fetch successfully", res)

}

func (c *OrchestratorController) GetPropertyDetails(ctx *gin.Context) {
	// 1. User Authentication Check
	if _, exists := ctx.Get(webMiddleware.UserIDContextKey); !exists {
		utils.Failure(ctx, http.StatusUnauthorized, utils.CodeUnauthorized, errors.New("user not authenticated"))
		return
	}

	// 2. Org context check
	orgUUIDRaw, exists := ctx.Get(webMiddleware.OrgUUIDContextKey)
	if !exists {
		utils.Failure(ctx, http.StatusUnauthorized, utils.CodeUnauthorized, errors.New("organization not found in context"))
		return
	}

	orgID, ok := orgUUIDRaw.(uuid.UUID)
	if !ok {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, errors.New("invalid org uuid type"))
		return
	}

	// 3. Param Parsing
	propertyId := ctx.Param("propertyID")
	pId, err := uuid.Parse(propertyId)
	if err != nil {
		// Changed to StatusBadRequest (400) because the user sent a bad UUID string
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeValidationFailed, errors.New("invalid propertyID format"))
		return
	}

	// 4. Service Call
	res, err := c.orchestService.GetPropertyDetail(ctx.Request.Context(), orgID, pId)
	if err != nil {
		// IMPORTANT: Don't just return. Tell the user what went wrong.
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, err)
		return
	}

	// 5. Success Response
	utils.Success(ctx, http.StatusOK, "Property details fetch successfully", res)
}

func (c *OrchestratorController) AddDomainInProperty(ctx *gin.Context) {

	// 1. User Authentication Check
	if _, exists := ctx.Get(webMiddleware.UserIDContextKey); !exists {
		utils.Failure(ctx, http.StatusUnauthorized, utils.CodeUnauthorized, errors.New("user not authenticated"))
		return
	}

	// 2. Org context check
	orgUUIDRaw, exists := ctx.Get(webMiddleware.OrgUUIDContextKey)
	if !exists {
		utils.Failure(ctx, http.StatusUnauthorized, utils.CodeUnauthorized, errors.New("organization not found in context"))
		return
	}

	orgID, ok := orgUUIDRaw.(uuid.UUID)
	if !ok {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, errors.New("invalid org uuid type"))
		return
	}

	// 3. Param Parsing
	propertyId := ctx.Param("propertyID")
	pId, err := uuid.Parse(propertyId)
	if err != nil {
		// Changed to StatusBadRequest (400) because the user sent a bad UUID string
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeValidationFailed, errors.New("invalid propertyID format"))
		return
	}

	var req *orchestrator_dtos.AddDomainToPropertyReq_DTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeConflict, err)
		return
	}

	if err := c.orchestService.ConfigurePropertyToDomain(ctx.Request.Context(), orgID, pId, req); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeUnauthorized, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Domain Added Successfully")

}

// ------------------------------------------------------------------------------//

func (c *OrchestratorController) AddModulesInProperty(ctx *gin.Context) {
	// 1. User Authentication Check
	if _, exists := ctx.Get(webMiddleware.UserIDContextKey); !exists {
		utils.Failure(ctx, http.StatusUnauthorized, utils.CodeUnauthorized, errors.New("user not authenticated"))
		return
	}

	// 2. Org context check
	orgUUIDRaw, exists := ctx.Get(webMiddleware.OrgUUIDContextKey)
	if !exists {
		utils.Failure(ctx, http.StatusUnauthorized, utils.CodeUnauthorized, errors.New("organization not found in context"))
		return
	}

	orgID, ok := orgUUIDRaw.(uuid.UUID)
	if !ok {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, errors.New("invalid org uuid type"))
		return
	}

	// 3. Param Parsing
	propertyId := ctx.Param("propertyID")
	pId, err := uuid.Parse(propertyId)
	if err != nil {
		// Changed to StatusBadRequest (400) because the user sent a bad UUID string
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeValidationFailed, errors.New("invalid propertyID format"))
		return
	}

	var req *orchestrator_dtos.PropertyModuleConfigurationDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeConflict, err)
		return
	}

	if err := c.orchestService.AddModulesInProperty(ctx.Request.Context(), orgID, pId, req); err != nil {
		return
	}

}
