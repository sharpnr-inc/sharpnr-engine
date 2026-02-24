package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registersharnprLauncherRoutes(r *gin.RouterGroup, oc *orchestrator_controllers.OrchestratorController) {

	r.GET("/launcher", oc.GetOrgLauncher)
}
