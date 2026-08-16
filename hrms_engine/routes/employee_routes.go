package hr_routes

import (
	"sharpnr-engine/app"
	hrmsEngineControllers "sharpnr-engine/hrms_engine/controllers"

	"github.com/gin-gonic/gin"
)

func registerEmployeeRoutes(hr *gin.RouterGroup, modules app.ModulesProvider) {

	hrmsController := hrmsEngineControllers.NewHrmsControllers(modules.HRMSService())

	employees := hr.Group("/employees")
	{
		employees.GET("", hrmsController.ListAllEmployeesOfOrg) // List employees (filters, pagination)
		employees.POST("/add")
		employees.GET("/:employeeId") // Returns employee details
	}

	profile := employees.Group("/:employeeId/profile")
	{
		profile.GET("") // Get profile
		profile.PUT("") // Update profile
	}
}
