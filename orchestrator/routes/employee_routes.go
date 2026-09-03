package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerEmployeeRoutes(r *gin.RouterGroup, co *orchestrator_controllers.OrchestratorController) {
	emr := r.Group("/empl")
	{
		emr.GET("", co.FetchAllEmployeesHandler) // fetches all the employees
		emr.GET("/:employeeID")                  // fetches a single employee by ID
		// adds a new superAdmin/owner employee — owners only. Temporarily creates
		// the account directly from the payload's password instead of inviting.
		emr.POST("/add_super_admin", co.CreateSuperAdmin)
		emr.POST("/add-employee-bulk", co.AddEmployeeInBulk) // add employee in the Bulk in the organization

	}
}
