package orchestrator_routes

import (
	organizationSentinalActions "sharpnr-engine/core/organization/actions"
	webMiddleware "sharpnr-engine/middlewares/web"
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"
	sentinel_actions "sharpnr-engine/sentinel_engine/actions"
	sentinelengine "sharpnr-engine/sentinel_engine/engine"

	"github.com/gin-gonic/gin"
)

// registerSentinelRoutes mounts the authorization (Sentinel) endpoints used by
// the frontend to author Cedar policies, manage roles, and inspect the
// privileges a role grants.
//
//	/sentinel
//	├── GET  /principals                            principal palette: types + roles + groups
//	├── /resources
//	│   ├── GET                                     resource types (each with actions + attributes)
//	│   └── /:type
//	│       ├── GET /actions                        the type's actions
//	│       ├── GET /attributes                     the type's ABAC attributes
//	│       └── GET /instances                      the org's live instances of the type
//	├── /policies
//	│   ├── GET                                     list the org's policies
//	│   └── POST                                    create a policy (raw Cedar text)
//	└── /roles
//	    ├── GET                                     list roles
//	    ├── POST                                    create a role (+ optional initial policies)
//	    └── /:roleID
//	        ├── GET                                 role + its privileges
//	        └── /policies
//	            ├── POST                            attach policies
//	            └── DELETE /:policyID               detach a policy
func registerSentinelRoutes(rg *gin.RouterGroup, c *orchestrator_controllers.OrchestratorController, engine *sentinelengine.Sentinel) {
	// needs gates a route on a Sentinel-admin action against the Organization
	// resource (resource id = the caller's org). The owner passes via the
	// full-access policy seeded at provisioning; everyone else is denied unless
	// the owner granted them the action.
	needs := func(action sentinel_actions.Name) gin.HandlerFunc {
		return webMiddleware.RequirePermission(engine, action, sentinelengine.EntityOrganization, "")
	}

	s := rg.Group("/sentinel")
	{
		s.GET("/principals", needs(organizationSentinalActions.ViewSentinel), c.ListPrincipalsHandler)

		resources := s.Group("/resources")
		{
			resources.GET("", needs(organizationSentinalActions.ViewSentinel), c.ListResourcesHandler)

			rt := resources.Group("/:type")
			{
				rt.GET("/actions", needs(organizationSentinalActions.ViewSentinel), c.GetResourceActionsHandler)
				rt.GET("/attributes", needs(organizationSentinalActions.ViewSentinel), c.GetResourceAttributesHandler)
				rt.GET("/instances", needs(organizationSentinalActions.ViewSentinel), c.ListResourceInstancesHandler)
			}
		}

		policies := s.Group("/policies")
		{
			policies.GET("", needs(organizationSentinalActions.ViewSentinel), c.ListPoliciesHandler)
			policies.POST("", needs(organizationSentinalActions.CreatePolicy), c.CreatePolicyHandler)
		}

		roles := s.Group("/roles")
		{
			roles.GET("", needs(organizationSentinalActions.ViewSentinel), c.ListRolesHandler)
			roles.POST("", needs(organizationSentinalActions.CreateRole), c.CreateRoleHandler)

			role := roles.Group("/:roleID")
			{
				role.GET("", needs(organizationSentinalActions.ViewSentinel), c.GetRolePrivilegesHandler)
				role.POST("/policies", needs(organizationSentinalActions.AttachPolicy), c.AttachPoliciesToRoleHandler)
				role.DELETE("/policies/:policyID", needs(organizationSentinalActions.DetachPolicy), c.DetachPolicyFromRoleHandler)
			}
		}
	}
}
