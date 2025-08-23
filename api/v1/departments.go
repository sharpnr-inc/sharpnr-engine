package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func DepartmentsRoutes(v1 *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {

	// newDeptRepo := repositories.NewDepartmentRepository(db)
	// newDeptService := deptservices.DepartmentService(newDeptRepo)

	v1.POST("/department/create")
	v1.POST("/department/get-all")
	v1.POST("/department/edit")
	v1.POST("/department/add-employee")

}
