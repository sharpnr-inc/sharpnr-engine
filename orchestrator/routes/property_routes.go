package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerPropertyRoutes(r *gin.RouterGroup, c *orchestrator_controllers.OrchestratorController) {

	pr := r.Group("/property")
	{
		pr.GET("/utils/domains", c.PropertyDomainTypes)
		pr.POST("/create", c.AddProperty)
		pr.GET("", c.ListAllOrgProperty)
		pr.GET("/:propertyID", c.GetPropertyDetails)
		pr.POST("/:propertyID/add-domain", c.AddDomainInProperty)
		pr.POST("/:propertyID/add-modules", c.AddModulesInProperty)
		pr.POST("/:propertyID/configure-restaurant", c.ConfigurePropertiesRestaurant)
		pr.POST("/:propertyID/configure-hotel", c.ConfigurePropertyHotel)
		pr.GET("/:propertyID/infra", c.GetPropertyInfra)
		pr.GET("/allinfra")

	}

}
