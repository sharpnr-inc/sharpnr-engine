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
		pr.GET("", c.ListAllOrgProperty)
		pr.POST("/create", c.AddProperty)
		pr.GET("/:propertyID", c.GetPropertyDetails)
		pr.GET("/:propertyID/infra_modules", c.GetPropertyInfra) // fetches all the infra of the property
		pr.GET("/:propertyID/business_modules")                  // fetches all the business modules of the property
		pr.POST("/:propertyID/add/business_modules")             // Add business modules in the property
		pr.POST("/:propertyID/add/infra_modules")                // Adds infra modules in the property
		

	}

}
