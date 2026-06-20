// Package orchestrator_routes wires the orchestrator's HTTP endpoints onto a
// Gin router. Routes are organized hierarchically around a property: a property
// owns business modules (hospitality, and later healthcare), infra modules, and
// resources. Each of those domains nests its own sub-resources beneath it.
//
// The full property route tree:
//
//	/property
//	├── GET                                              list all properties for the org
//	├── POST   /create                                   create a property
//	└── /:propertyID
//	    ├── GET                                          property details
//	    ├── /business_modules
//	    │   ├── GET                                      list business modules on the property
//	    │   ├── POST                                     add a business module to the property
//	    │   └── /hospitality
//	    │       ├── /hotel/:hotelID
//	    │       │   ├── GET                              hotel details
//	    │       │   ├── PATCH /general-details           update general details
//	    │       │   ├── POST  /kitchen                   add a preferred kitchen
//	    │       │   └── POST  /legal-entity              add a preferred legal entity
//	    │       ├── /restaurant/:restaurantID
//	    │       │   ├── GET                              restaurant details
//	    │       │   ├── PATCH /gen-details               update general details
//	    │       │   └── POST  /add-tables                add tables
//	    │       └── /bar/:barID
//	    │           └── GET                              bar details
//	    ├── /infra_modules
//	    │   ├── GET                                      list infra modules on the property
//	    │   ├── POST                                     add an infra module to the property
//	    │   └── /hospitality
//	    │       └── /kitchen/:kitchenID
//	    │           └── GET                              kitchen details
//	    └── /resources
//	        ├── GET                                      list resources on the property
//	        ├── POST                                     add a resource to the property
//	        └── /hospitality
//	            └── /menu/:menuID
//	                └── GET                              menu details
//
// TODO: rename "business_modules" to "business-operations" across these routes
// and their handlers.
package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

// todo is a placeholder handler for routes whose real handlers are not yet
// implemented. It exists for two reasons:
//
//  1. Gin requires every registered route to have at least one handler. A bare
//     pr.GET("") with no handler panics with "there must be at least one
//     handler" under gin.New(), and silently degrades to a no-op 200 under
//     gin.Default() (where Logger/Recovery middleware fills the chain). Pointing
//     unfinished routes at todo makes the behavior explicit instead.
//  2. It returns 501 Not Implemented, so callers hitting an unfinished endpoint
//     get a clear, correct signal rather than an empty success.
//
// Replace each todo reference with the real controller method as it lands.
func todo(ctx *gin.Context) {
	ctx.JSON(501, gin.H{"error": "not implemented"})
}

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
					hotel.GET("", todo)                   // fetch hotel details
					hotel.PATCH("/general-details", todo) // update general details
					hotel.POST("/kitchen", todo)          // add a preferred kitchen
					hotel.POST("/legal-entity", todo)     // add a preferred legal entity
				}

				// Restaurant — scoped by :restaurantID.
				restaurant := hospitality.Group("/restaurant/:restaurantID")
				{
					restaurant.GET("", todo)               // fetch restaurant details
					restaurant.PATCH("/gen-details", todo) // update general details
					restaurant.POST("/add-tables", todo)   // add tables
				}

				// Bar — scoped by :barID.
				bar := hospitality.Group("/bar/:barID")
				{
					bar.GET("", todo) // fetch bar details
				}
			}

			// Business Operation Module :-> Healthcare (not ready yet)
			// healthcare := prop.Group("/business_modules/healthcare")
			// {
			// }

			/* ------------------------------ Infra Modules ----------------------------- */
			// Physical / operational infrastructure backing the business modules.

			prop.GET("/infra_modules", c.GetInfraModulesOfPropertyHandler)  // list infra modules
			prop.POST("/infra_modules", c.AddInfraModulesOfPropertyHandler) // attach an infra module

			// General infra (not ready yet).
			// generalInfra := prop.Group("/infra_modules/general")
			// {
			// }

			// Hospitality infra.
			infraHospitality := prop.Group("/infra_modules/hospitality")
			{
				// Kitchen — scoped by :kitchenID.
				kitchen := infraHospitality.Group("/kitchen/:kitchenID")
				{
					kitchen.GET("", todo) // fetch kitchen details
					// further kitchen routes
				}
			}

			/* -------------------------------- Resources ------------------------------- */
			// Reusable assets owned by the property (menus, etc.).

			prop.GET("/resources", todo)  // list resources
			prop.POST("/resources", todo) // add a resource

			// General resources (not ready yet).
			// generalResources := prop.Group("/resources/general")
			// {
			// }

			// Hospitality resources.
			resourcesHospitality := prop.Group("/resources/hospitality")
			{
				// Menu — scoped by :menuID.
				menu := resourcesHospitality.Group("/menu/:menuID")
				{
					menu.GET("", todo) // fetch menu details
					// further menu routes
				}
			}
		}
	}
}
