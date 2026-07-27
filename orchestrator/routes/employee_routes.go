package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerEmployeeRoutes(r *gin.RouterGroup, co *orchestrator_controllers.OrchestratorController) {
	emr := r.Group("/empl")
	{
		emr.GET("", co.FetchAllEmployeesHandler)             // fetches all the employees
		emr.GET("/:employeeID")                              // fetches a single employee by ID
		emr.POST("/add_super_admin", co.CreateSuperAdmin)    // adds a new superAdmin/owner employee
		emr.POST("/add-employee-bulk", co.AddEmployeeInBulk) // add employee in the Bulk in the organization
	
	}
}
