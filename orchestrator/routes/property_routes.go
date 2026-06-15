package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

/*
TODO :- change the add modules to business-opperations
*/

func registerPropertyRoutes(r *gin.RouterGroup, c *orchestrator_controllers.OrchestratorController) {

	pr := r.Group("/property")
	{
		pr.POST("/create", c.AddProperty)
		pr.GET("", c.ListAllOrgProperty)
		pr.GET("/:propertyID", c.GetPropertyDetails)
		pr.POST("/:propertyID/add-modules", c.AddModulesInProperty)
		pr.POST("/:propertyID/add-infra", c.AddInfraInProperty)
		pr.GET("/:propertyID/infra", c.GetPropertyInfra)

	}

}
