package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerOrgSupplierRoutes(r *gin.RouterGroup, c *orchestrator_controllers.OrchestratorController) {
	sr := r.Group("/suppliers")
	sr.GET("/list", c.ListAllSupplierHandler)
	sr.POST("/add-new", c.CreateNewSupplierHandler)
	sr.POST("/:suppliersID/add-lel", c.AddLegalEntityLinkToSupplierHandler)
}
