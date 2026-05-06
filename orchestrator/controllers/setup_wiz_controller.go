package orchestrator_controllers

import (
	"errors"
	"net/http"
	webMiddleware "sharpnr-engine/middlewares/web"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
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
// @Success      200  {object}  utils.Response "Complete" "Success message"
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

/* ---------------------------------- <***> --------------------------------- */

func (s *OrchestratorController) HandleSuperAdminSetupCompletion(ctx *gin.Context) {
	orgID, emplID, ok := s.resolveRequestContexts(ctx)
	if !ok {
		return // Error response already sent in resolveRequestContexts
	}

	req, ok := utils.BindJSON[orchestrator_dtos.SetupSuperAdminDetailsReqDTO](ctx)
	if !ok {
		return
	}

	if err := s.orchestService.SetUpSuperAdminDetails(ctx.Request.Context(), orgID, emplID, req); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Super admin details saved successfully")
}

/* ---------------------------------- <***> --------------------------------- */

// Step_two_AddProperty godoc
// @Summary      Setup Organization Property
// @Description  Second step of the onboarding wizard: Configures the core property attributes for the organization.
// @Tags         core-setup
// @Accept       json
// @Produce      json
// @Param        request  body  orchestrator_dtos.SetupOrgPropertyReqDTO  true  "Property Details"
// @Success      200  {object}  utils.Response "Property Added Successfully"
// @Failure      400  {object}  utils.Response "Validation Error"
// @Failure      401  {object}  utils.Response "Unauthorized"
// @Failure      500  {object}  utils.Response "Internal Server Error"
// @Router       /orchestrator/setup-wizard/step2 [post]
func (c *OrchestratorController) Step_two_AddProperty(ctx *gin.Context) {

	// 1. Authenticated User Check
	_, exist := ctx.Get(webMiddleware.UserIDContextKey)
	if !exist {
		utils.Failure(ctx, http.StatusUnauthorized, utils.CodeUnauthorized, nil)
		return
	}

	// 2. Organization Context Check
	orgUUIDRaw, exists := ctx.Get(webMiddleware.OrgUUIDContextKey)
	if !exists {
		utils.Failure(ctx, http.StatusUnauthorized, utils.CodeUnauthorized, nil)
		return
	}

	orgID, ok := orgUUIDRaw.(uuid.UUID)
	if !ok {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, nil)
		return
	}

	// 3. Bind JSON Request (Fixed: bind to variable address)
	var req orchestrator_dtos.SetupOrgPropertyReqDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(ctx, err)
		return
	}

	// 4. Call Service Layer
	if err := c.orchestService.StepTwoProperty(ctx.Request.Context(), orgID, &req); err != nil {
		// Use your stable CodeInternalError or a specific domain code
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, err)
		return
	}

	// 5. Success Response
	utils.Success(ctx, http.StatusOK, "Property Added Successfully")
}

// Step_three_LegalEntity godoc
// @Summary      Setup Legal Entity
// @Description  Third step of the onboarding wizard to define the organization's legal and physical presence.
// @Tags         core-setup
// @Accept       json
// @Produce      json
// @Param        request body orchestrator_dtos.SetupLegalEntityReqDTO true "Legal Entity Details"
// @Success      200 {object} utils.Response "Legal entity configured successfully"
// @Failure      400 {object} utils.Response "Validation or binding error"
// @Failure      401 {object} utils.Response "Unauthorized"
// @Failure      500 {object} utils.Response "Internal server error"
// @Router       /orchestrator/setup-wiz/step3 [post]
func (c *OrchestratorController) Step_three_LegalEntity(ctx *gin.Context) {

	// 1. Check Authentication
	_, exist := ctx.Get(webMiddleware.UserIDContextKey)
	if !exist {
		utils.Failure(ctx, http.StatusUnauthorized, utils.CodeUnauthorized, nil)
		return
	}

	// 2. Retrieve OrgID from Context
	orgUUIDRaw, exists := ctx.Get(webMiddleware.OrgUUIDContextKey)
	if !exists {
		utils.Failure(ctx, http.StatusUnauthorized, utils.CodeUnauthorized, nil)
		return
	}

	orgID, ok := orgUUIDRaw.(uuid.UUID)
	if !ok {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, nil)
		return
	}

	// 3. Bind JSON Request (Fixed: initialized the DTO)
	var req orchestrator_dtos.SetupLegalEntityReqDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(ctx, err)
		return
	}

	// 4. Call Service Layer
	err := c.orchestService.StepThreeLegalEntity(ctx.Request.Context(), orgID, &req)
	if err != nil {
		// Map service error to appropriate failure
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, err)
		return
	}

	// 5. Success Wrap
	utils.Success(ctx, http.StatusOK, "Legal entity details saved successfully")
}
