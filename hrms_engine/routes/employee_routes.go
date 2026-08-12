package hr_routes

import (
	"sharpnr-engine/app"
	employee_controllers "sharpnr-engine/hrms_engine/employees/controllers"

	"github.com/gin-gonic/gin"
)

func registerEmployeeRoutes(hr *gin.RouterGroup, modules app.ModulesProvider) {

	hrmsEmployeeService := modules.HRMSEmployeeService()
	newEmpController := employee_controllers.NewEmployeeController(hrmsEmployeeService)

	employees := hr.Group("/employees")
	{
		employees.GET("", newEmpController.ListAllEmployeesOfOrg) // List employees (filters, pagination)
		employees.POST("/add")
		employees.GET("/:employeeId") // Returns employee details
	}

	profile := employees.Group("/:employeeId/profile")
	{
		profile.GET("") // Get profile
		profile.PUT("") // Update profile
	}
}
