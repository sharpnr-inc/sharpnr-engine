package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

// These are the org-level department routes: the departments an organization
// manages directly, as opposed to the departments that hang off a cost center.
// Every handler resolves the organization from the session, so nothing here
// takes an organization id.
func registerDepartmentRoutes(r *gin.RouterGroup, co *orchestrator_controllers.OrchestratorController) {
	dept := r.Group("/department")
	{
		dept.POST("", co.CreateOrgDepartmentHandler)                 // creates a new department
		dept.GET("/:departmentID", co.GetOrgDepartmentHandler)       // fetches a single department
		dept.GET("", co.ListOrgAllDepartmentHandler)                 // lists the org's departments
		dept.PUT("/:departmentID", co.UpdateOrgDepartmentHandler)    // partially updates a department
		dept.DELETE("/:departmentID", co.DeleteOrgDepartmentHandler) // soft-deletes a department
	}
}
