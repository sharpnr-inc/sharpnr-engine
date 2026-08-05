package orchestrator_controllers

import (
	"errors"
	"net/http"
	orgSetUpWizardErrs "sharpnr-engine/core/org_setup_wizard/errs"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
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
func (c *OrchestratorController) SetupWizGetSetupProgress(ctx *gin.Context) {

	orgID, _, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	prog, err := c.orchestService.SetUpWizProgress(
		ctx.Request.Context(),
		orgID,
	)
	if err != nil {
		// Ideally map domain errors → HTTP here
		utils.Failure(
			ctx,
			http.StatusInternalServerError,

			err)

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

func (s *OrchestratorController) SetupWizHandleSuperAdminSetupCompletion(ctx *gin.Context) {
	orgID, emplID, ok := s.resolveRequestContexts(ctx)
	if !ok {
		return // Error response already sent in resolveRequestContexts
	}

	req, ok := utils.BindJSON[orchestrator_dtos.SetupSuperAdminDetailsReqDTO](ctx)
	if !ok {
		return
	}

	if err := s.orchestService.SetUpSuperAdminDetails(ctx.Request.Context(), orgID, emplID, req); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err)
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
func (c *OrchestratorController) SetupWizOperation_AddProperty(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	// 3. Bind JSON Request (Fixed: bind to variable address)
	var req orchestrator_dtos.SetupOrgPropertyReqDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(ctx, err)
		return
	}

	// 4. Call Service Layer
	if err := c.orchestService.SetupWiz_AddProperty(ctx.Request.Context(), orgID, emplID, &req); err != nil {
		// A malformed step or missing property details is the caller's mistake,
		// so it must not come back as an internal error.
		switch {
		case errors.Is(err, orgSetUpWizardErrs.InvalidStepIDErr),
			errors.Is(err, orgSetUpWizardErrs.PropertyDetailsRequiredErr):
			utils.Failure(ctx, http.StatusBadRequest, err)
		case errors.Is(err, orgSetUpWizardErrs.StepDoesntExistErr):
			utils.Failure(ctx, http.StatusNotFound, err)
		case errors.Is(err, orgSetUpWizardErrs.InvalidStepStatusErr):
			utils.Failure(ctx, http.StatusConflict, err)
		default:
			utils.Failure(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	// 5. Success Response
	utils.Success(ctx, http.StatusOK, "Property Added Successfully")
}

func (c *OrchestratorController) SetupWiz_UpdateOrganizationName(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	// 3. Bind JSON Request (Fixed: initialized the DTO)
	var req orchestrator_dtos.SetupWizUpdateOrgNameRequestDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(ctx, err)
		return
	}

	if err := c.orchestService.SetupWiz_UpdateOrgName(ctx.Request.Context(), orgID, emplID, &req); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Organization Name Added Successfully")

}
