package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerSetupWizardRoutes(webV1 *gin.RouterGroup, controller *orchestrator_controllers.OrchestratorController) {

	setupWiz := webV1.Group("/setup-wiz")

	setupWiz.GET("/progress", controller.GetSetupProgress)
	setupWiz.POST("/sasetup", controller.HandleSuperAdminSetupCompletion)
	setupWiz.POST("/step3", controller.Step_three_LegalEntity)

}
