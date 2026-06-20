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

		// business modules
		pr.GET("/:propertyID/business_modules", c.GetBusinessOperationModulesOfPropertyHandler) // fetches all the Business operation modules in the property
		pr.POST("/:propertyID/business_modules", c.AddBusinessOperationModulesInPropertyHandler)

		/* -------------------------------------------------------------------------- */
		/*                 Business Operation Module :-> Hospitality                 */
		/* -------------------------------------------------------------------------- */

		// All the Routes that belongs the hospitality domain
		hospitality := pr.Group("/:propertyID/business_modules/hospitality/")
		{

			// All the Routes belongs to the Hotel
			hotelroutes := hospitality.Group("hotel/:hotelID")
			{
				hotelroutes.GET("")                   // fetches the Hotel Details
				hotelroutes.PATCH("/general-details") // updates the General Details
				hotelroutes.POST("/kitchen")          // Add the prefrerred kitchen
				hotelroutes.POST("/legal-entity")     // Add the Preferred legal Entity
			}

			// All the routes belongs the Restaurant modules
			restaurantRoutes := hospitality.Group("restaurant/:restaurantID")
			{
				restaurantRoutes.GET("")
			}

			//All thr Routes beloag to the Bar modules
			barroutes := hospitality.Group("bar/:barID")
			{
				barroutes.GET("")
			}

		}

		/* ---------------------------------- <***> --------------------------------- */

		// infra modules
		pr.GET("/:propertyID/infra_modules", c.GetInfraModulesOfPropertyHandler)
		pr.POST("/:propertyID/infra_modules", c.AddInfraModulesOfPropertyHandler)

	}
}
