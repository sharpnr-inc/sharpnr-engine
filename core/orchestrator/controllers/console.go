package orchestrator_controllers

import "github.com/gin-gonic/gin"

func (c *OrchestratorController) GetOrganizationSetupStatus(ctx *gin.Context) {

	// *Because we are using middleware and it set the OrgID and it comes in String form
	// *But we have to convert into the UUID and validate it
	// *Step 1 --> using organization Services we can find the the if organization is Setup is Completed or Not
	// *Then

}
