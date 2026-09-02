package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerCostCenterRoutes(r *gin.RouterGroup, co *orchestrator_controllers.OrchestratorController) {
	costCenter := r.Group("/cost_center")
	{
		costCenter.POST("", co.CreateCostCenterHandler)           // creates a new cost center
		costCenter.GET("/:costCenterID", co.GetCostCenterHandler) // fetches a single cost center
		costCenter.GET("", co.ListAllCostCentersHandler)          // lists the org's cost centers

		// The cost center's own departments. Nested under the site because a
		// department is created inside one — the id in the path is the only
		// thing that says which.
		costCenter.POST("/:costCenterID/department", co.CreateCostCenterDepartmentHandler)

	}

}
