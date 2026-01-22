package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/core/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerSetupWizardRoutes(webV1 *gin.RouterGroup, controller *orchestrator_controllers.OrchestratorController) {

	setupWiz := webV1.Group("/setup-wiz")

	setupWiz.GET("/progress", controller.GetSetupProgress)

}
