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

// GetSetupProgress godoc
// @Summary      Short, punchy summary of the action
// @Description  Detailed explanation of what happens (e.g., database changes, emails sent).
// @Tags         CategoryName (e.g., accounting-ap)
// @Accept       json
// @Produce      json
// @Param        request  body  orchestrator_dtos.SetupProgressResponseDTO true  "Input description"
// @Success      200  {object}  utils.Response{data=your_dtos.ResponseDTO} "Success message"
// @Failure      400  {object}  utils.Response "Validation Error"
// @Failure      401  {object}  utils.Response "Unauthorized"
// @Failure      500  {object}  utils.Response "Internal Server Error"
// @Router       /orchestrator/setup-wizard/progress [get]
func (c *OrchestratorController) GetSetupProgress(ctx *gin.Context) {

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

	// 3️⃣ Service call
	prog, err := c.orchestService.SetUpWizProgress(
		ctx.Request.Context(),
		orgID,
	)
	if err != nil {
		// Ideally map domain errors → HTTP here
		utils.Failure(
			ctx,
			http.StatusInternalServerError,
			utils.CodeInternalError,
			err,
		)
		return
	}

	// 4️⃣ Success response
	utils.Success(
		ctx,
		http.StatusOK,
		"setup progress fetched successfully",
		prog,
	)
}

// Step_one_AddOrgHQ godoc
// @Summary      Short, punchy summary of the action
// @Description  Detailed explanation of what happens (e.g., database changes, emails sent).
// @Tags         CategoryName (e.g., accounting-ap)
// @Accept       json
// @Produce      json
// @Param        request  body      your_dtos.TargetDTO  true  "Input description"
// @Success      200  {object}  utils.Response{data=your_dtos.ResponseDTO} "Success message"
// @Failure      400  {object}  utils.Response "Validation Error"
// @Failure      401  {object}  utils.Response "Unauthorized"
// @Failure      500  {object}  utils.Response "Internal Server Error"
// @Router        /orchestrator/setup-wizard/step1 [post]
func (c *OrchestratorController) Step_one_AddOrgHQ(ctx *gin.Context) {

	_, exist := ctx.Get(webMiddleware.UserIDContextKey)
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "user not authenticated",
		})
		return
	}

	orgUUIDRaw, exists := ctx.Get(webMiddleware.OrgUUIDContextKey)
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "organization not found in context",
		})
		return
	}

	orgID, ok := orgUUIDRaw.(uuid.UUID)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "invalid org uuid type",
		})
		return
	}

	var req *orchestrator_dtos.SetupOrganizationHQAddressReqDTO
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "organization not found in context",
		})
		return
	}

	if err := c.orchestService.AddOrganizationHQAddress(ctx.Request.Context(), orgID, req); err != nil {
		utils.Success(ctx, 200, "Organization HQ Added")
	}

}

// Step_two_AddProperty godoc
// @Summary      Short, punchy summary of the action
// @Description  Detailed explanation of what happens (e.g., database changes, emails sent).
// @Tags         CategoryName (e.g., accounting-ap)
// @Accept       json
// @Produce      json
// @Param        request  body      your_dtos.TargetDTO  true  "Input description"
// @Success      200  {object}  utils.Response{data=your_dtos.ResponseDTO} "Success message"
// @Failure      400  {object}  utils.Response "Validation Error"
// @Failure      401  {object}  utils.Response "Unauthorized"
// @Failure      500  {object}  utils.Response "Internal Server Error"
// @Router        /orchestrator/setup-wizard/step2 [post]
func (c *OrchestratorController) Step_two_AddProperty(ctx *gin.Context) {

	_, exist := ctx.Get(webMiddleware.UserIDContextKey)
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "user not authenticated",
		})
		return
	}

	orgUUIDRaw, exists := ctx.Get(webMiddleware.OrgUUIDContextKey)
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "organization not found in context",
		})
		return
	}

	_, ok := orgUUIDRaw.(uuid.UUID)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "invalid org uuid type",
		})
		return
	}

	var req *orchestrator_dtos.SetupOrgPropertyReqDTO
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "organization not found in context",
		})
		return
	}

}

// Step_two_AddProperty godoc
// @Summary      Short, punchy summary of the action
// @Description  Detailed explanation of what happens (e.g., database changes, emails sent).
// @Tags         CategoryName (e.g., accounting-ap)
// @Accept       json
// @Produce      json
// @Param        request  body      your_dtos.TargetDTO  true  "Input description"
// @Success      200  {object}  utils.Response{data=your_dtos.ResponseDTO} "Success message"
// @Failure      400  {object}  utils.Response "Validation Error"
// @Failure      401  {object}  utils.Response "Unauthorized"
// @Failure      500  {object}  utils.Response "Internal Server Error"
// @Router       /api/v1/your-path [post]
func (c *OrchestratorController) Step_three_LegalEntity(ctx *gin.Context) {

	_, exist := ctx.Get(webMiddleware.UserIDContextKey)
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "user not authenticated",
		})
		return
	}

	orgUUIDRaw, exists := ctx.Get(webMiddleware.OrgUUIDContextKey)
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "organization not found in context",
		})
		return
	}

	_, ok := orgUUIDRaw.(uuid.UUID)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "invalid org uuid type",
		})
		return
	}

	var req *orchestrator_dtos.SetupLegalEntityReqDTO
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "organization not found in context",
		})
		return
	}

}
