package orchestrator_controllers

import (
	"errors"
	"net/http"
	property_constants "sharpnr-engine/core/org_property/constants"
	webMiddleware "sharpnr-engine/middlewares/web"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
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
	orgID, _, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return // Response already sent by helper
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

	orgID, _, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return // Response already sent by helper
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

	// 1. Capture both the result and the error
	result, err := c.orchestService.AddModulesInProperty(ctx.Request.Context(), orgID, pId, req)

	if err != nil {
		// Handle specific errors (like the "Only one module" constraint)
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeConflict, err)
		return
	}

	// 2. Pass the result object back in the success response
	// This gives the frontend the 'Created', 'Skipped', and 'Message' fields
	utils.Success(ctx, http.StatusOK, "Result", result)

}

func (c *OrchestratorController) AddProperty(ctx *gin.Context) {

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

	OrgID, ok := orgUUIDRaw.(uuid.UUID)
	if !ok {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, errors.New("invalid org uuid type"))
		return
	}

	var req *orchestrator_dtos.CreatePropertyReq_dto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	if err := c.orchestService.CreateProperty(ctx.Request.Context(), OrgID, req); err != nil {
		utils.Success(ctx, http.StatusCreated, "Property Added Successfully")
		return
	}

}

func (c *OrchestratorController) PropertyDomainTypes(ctx *gin.Context) {
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

	_, ok := orgUUIDRaw.(uuid.UUID)
	if !ok {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, errors.New("invalid org uuid type"))
		return
	}

	utils.Success(ctx, http.StatusOK, "Property Domain Fetched successfully", gin.H{
		"domains": property_constants.PropertyDomains,
	})

}

func (c *OrchestratorController) ConfigurePropertiesRestaurant(ctx *gin.Context) {
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
	var req *orchestrator_dtos.ConfigurePropertysRestaurant_Req_dto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeConflict, err)
		return
	}

	if err := c.orchestService.ConfigureRestaurant(ctx.Request.Context(), orgID, pId, req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeConflict, err)
		return
	}

	utils.Success(ctx, 201, "Configured successfully")

}

func (c *OrchestratorController) ConfigurePropertyHotel(ctx *gin.Context) {
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

	var req *orchestrator_dtos.ConfigureHotel_Request_dto

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeConflict, err)
		return
	}

	if err := c.orchestService.ConfigureHotel(ctx.Request.Context(), orgID, pId, req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeConflict, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Configuration Complete")
}
