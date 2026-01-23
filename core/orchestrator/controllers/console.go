package orchestrator_controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetOrganizationSetupStatus godoc
// @Summary      Short, punchy summary of the action
// @Description  Detailed explanation of what happens (e.g., database changes, emails sent).
// @Tags         Orchestrator --> Console
// @Accept       json
// @Produce      json
// @Param        request  body      your_dtos.TargetDTO  true  "Input description"
// @Success      200  {object}  utils.Response{data=your_dtos.ResponseDTO} "Success message"
// @Failure      400  {object}  utils.Response "Validation Error"
// @Failure      401  {object}  utils.Response "Unauthorized"
// @Failure      500  {object}  utils.Response "Internal Server Error"
// @Router       /orchestrator/setup-status [get]
func (c *OrchestratorController) GetOrganizationSetupStatus(ctx *gin.Context) {

	// *Because we are using middleware and it set the OrgID and it comes in String form
	// *But we have to convert into the UUID and validate it
	// *Step 1 --> using organization Services we can find the the if organization is Setup is Completed or Not
	// *!! The DTO is Remaining the in swagger

	val, exists := ctx.Get("OrgID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: OrgID missing"})
		return
	}

	// 2. Convert and Validate UUID
	orgID, err := uuid.Parse(val.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OrgID format"})
		return
	}

	// 3. Call Service (using the Request Context)
	// We pass ctx.Request.Context() so timeouts/cancellations work.
	done, err := c.orchestService.GetOrganizationSetupStatus(ctx.Request.Context(), orgID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Service failure"})
		return
	}

	// 4. Final Response
	ctx.JSON(http.StatusOK, gin.H{
		"organization_id": orgID,
		"setup_completed": done,
		"message":         "Dummy status retrieved successfully",
	})

}
