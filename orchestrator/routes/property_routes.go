package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

// registerPropertyRoutes mounts the /property route tree onto r.
//
// All endpoints beyond the org-level list/create are scoped to a single
// property via the /:propertyID path parameter, and split across three domains:
// business modules, infra modules, and resources. See the package doc for the
// complete tree.
func registerPropertyRoutes(r *gin.RouterGroup, c *orchestrator_controllers.OrchestratorController) {
	pr := r.Group("/property")
	{
		// Org-level: not scoped to any one property.
		pr.GET("", c.ListAllOrgProperty)  // list every property in the org
		pr.POST("/create", c.AddProperty) // create a new property

		// Everything below is scoped to a single property by :propertyID.
		prop := pr.Group("/:propertyID")
		{
			prop.GET("", c.GetPropertyDetails) // fetch details for one property

			/* ---------------------------- Business Modules ---------------------------- */
			// A property's operational lines of business (hospitality, healthcare, ...).

			prop.GET("/business_modules", c.GetBusinessOperationModulesOfPropertyHandler)  // list business modules
			prop.POST("/business_modules", c.AddBusinessOperationModulesInPropertyHandler) // attach a business module

			// Business Operation Module :-> Hospitality
			hospitality := prop.Group("/business_modules/hospitality")
			{
				// Hotel — scoped by :hotelID.
				hotel := hospitality.Group("/hotel/:hotelID")
				{
					hotel.GET("")                   // fetch hotel details
					hotel.PATCH("/general-details") // update general details
					hotel.POST("/kitchen")          // add a preferred kitchen
					hotel.POST("/legal-entity")     // add a preferred legal entity
				}

				// Restaurant — scoped by :restaurantID.
				restaurant := hospitality.Group("/restaurant/:restaurantID")
				{
					restaurant.GET("")               // fetch restaurant details
					restaurant.PATCH("/gen-details") // update general details
					restaurant.POST("/add-tables")   // add tables
				}

				// Bar — scoped by :barID.
				bar := hospitality.Group("/bar/:barID")
				{
					bar.GET("") // fetch bar details
				}
			}

			/* ------------------------------ Infra Modules ----------------------------- */
			// Physical / operational infrastructure backing the business modules.

			prop.GET("/infra_modules", c.GetInfraModulesOfPropertyHandler)  // list infra modules
			prop.POST("/infra_modules", c.AddInfraModulesOfPropertyHandler) // attach an infra module

			// Hospitality infra.
			infraHospitality := prop.Group("/infra_modules/hospitality")
			{
				// Kitchen — scoped by :kitchenID.
				kitchen := infraHospitality.Group("/kitchen/:kitchenID")
				{
					kitchen.GET("") // fetch kitchen details
				}
			}

			/* -------------------------------- Resources ------------------------------- */
			// Reusable assets owned by the property (menus, etc.).

			prop.GET("/resources")  // list resources
			prop.POST("/resources") // add a resource

			// Hospitality resources.
			resourcesHospitality := prop.Group("/resources/hospitality")
			{
				// Menu — scoped by :menuID.
				menu := resourcesHospitality.Group("/menu/:menuID")
				{
					menu.GET("") // fetch menu details
				}
			}
		}
	}
}
