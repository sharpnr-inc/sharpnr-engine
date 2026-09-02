package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerDepartmentRoutes(r *gin.RouterGroup, co *orchestrator_controllers.OrchestratorController) {
	dept := r.Group("/department")
	{
		dept.POST("") // creates a new department
		dept.GET("/:departmentID")
		dept.GET("")
		dept.PUT("/:departmentID")
		dept.DELETE("/:departmentID")
	}
}
