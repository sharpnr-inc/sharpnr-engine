package v1

import (
	employeecontrollers "sharpnr-engine/core/controllers/employeeControllers"
	employeerepo "sharpnr-engine/core/repositories/employee_repo"
	organizationrepositories "sharpnr-engine/core/repositories/organization_repo"
	employeeservice "sharpnr-engine/core/services/employee_service"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// These are Org-Core routes
func CoreV1Routes(v1 *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {

	newOrgRepo := organizationrepositories.InitializeNewOrgRegistration(db, redis)
	newEmployeeRepo := employeerepo.EmployeeRepo(db)
	newEmployeeService := employeeservice.EmployeeService(newOrgRepo, newEmployeeRepo)
	newEmployeeController := employeecontrollers.EmployeeController(newEmployeeService)

	v1.GET("/org/dashboard")
	v1.GET("/org/get-employee", newEmployeeController.AddEmployeeInOrg)
	v1.POST("/org/add-employee")
	v1.POST("/org/update-employee")
	v1.GET("/org/profile")
}
