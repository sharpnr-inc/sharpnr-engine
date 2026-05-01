package hr_routes

import (
	"sharpnr-engine/app"
	employee_controllers "sharpnr-engine/modules/hr/employees/controllers"

	"github.com/gin-gonic/gin"
)

func registerEmployeeRoutes(hr *gin.RouterGroup, modules app.ModulesProvider) {

	employeeService := modules.EmployeesHRService()
	newEmpController := employee_controllers.NewEmployeeController(employeeService)

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
