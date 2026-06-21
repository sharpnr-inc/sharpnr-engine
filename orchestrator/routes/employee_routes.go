package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerEmployeeRoutes(r *gin.RouterGroup, co *orchestrator_controllers.OrchestratorController) {
	emr := r.Group("/empl")
	{
		emr.GET("", co.FetchAllEmployeesHandler) // fetches all the employees
	}
}
