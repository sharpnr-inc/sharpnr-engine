package v1

import (
	employeecontrollers "sharpnr-engine/core/controllers/employee_controllers"
	employeerepo "sharpnr-engine/core/repositories/employee_repo"
	organizationrepositories "sharpnr-engine/core/repositories/organization_repo"
	employeeservice "sharpnr-engine/core/services/employee_service"
	"sharpnr-engine/shared/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func EmployeeRoutes(v1 *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	newOrgRepo := organizationrepositories.NewOrganizationRepo(db, redis)
	newEmployeeRepo := employeerepo.EmployeeRepo(db)
	newEmployeeService := employeeservice.EmployeeService(newOrgRepo, newEmployeeRepo)
	newEmployeeController := employeecontrollers.EmployeeController(newEmployeeService)

	v1.POST("/emp/add", middlewares.OrgRegOtpVerificationMiddleware(), newEmployeeController.AddEmployeeInOrg)
	v1.GET("/emp/dashboard")
}
