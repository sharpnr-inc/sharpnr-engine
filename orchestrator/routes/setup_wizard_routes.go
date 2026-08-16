package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerSetupWizardRoutes(webV1 *gin.RouterGroup, controller *orchestrator_controllers.OrchestratorController) {

	setupWiz := webV1.Group("/setup-wiz")

	setupWiz.GET("/progress", controller.SetupWizGetSetupProgress)

	/* -------------------------------------------------------------------------- */
	/*              these are routes for the setup wiz version 1.0.0              */
	/* -------------------------------------------------------------------------- */
	setupWiz.POST("/v1/sasetup", controller.SetupWizHandleSuperAdminSetupCompletion)
	setupWiz.POST("/v1/org_details", controller.SetupWiz_UpdateOrganizationName)
	setupWiz.POST("/v1/property_setup", controller.SetupWizOperation_AddProperty)

}
