package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

// registerSentinelRoutes mounts the authorization (Sentinel) endpoints used by
// the frontend to author Cedar policies, manage roles, and inspect the
// privileges a role grants.
//
//	/sentinel
//	├── GET  /principals                            principal palette: types + roles + groups
//	├── GET  /contexts                              context palette: authorable context.*/principal.* keys
//	├── /resources
//	│   ├── GET                                     resource palette: every registered resource type
//	│   └── GET  /:resourceType                     one resource's actions, attributes + live instances
//	├── /policies
//	│   ├── GET                                     list the org's policies
//	│   ├── POST                                    create a policy (raw Cedar text)
//	│   └── GET  /:policyID                         one policy
//	└── /roles
//	    ├── GET                                     list roles
//	    ├── POST                                    create a role (+ optional initial policies)
//	    └── /:roleID
//	        ├── GET                                 role + its privileges
//	        └── /policies
//	            ├── POST                            attach policies
//	            └── DELETE /:policyID               detach a policy
func registerSentinelRoutes(rg *gin.RouterGroup, c *orchestrator_controllers.OrchestratorController) {
	s := rg.Group("/sentinel")
	{
		s.GET("/principals", c.ListPrincipalsHandler)
		s.GET("/contexts", c.GetAuthzContextPaletteHandler)

		resources := s.Group("/resources")
		{
			resources.GET("", c.ListAuthzResourcesHandler)
			resources.GET("/:resourceType", c.GetAuthzResourceDetailsHandler)
		}

		policies := s.Group("/policies")
		{
			policies.GET("", c.ListPoliciesHandler)
			policies.POST("", c.CreatePolicyHandler)
			policies.GET("/:policyID", c.GetPolicyHandler)
		}

		roles := s.Group("/roles")
		{
			roles.GET("", c.ListRolesHandler)
			roles.POST("", c.CreateRoleHandler)

			role := roles.Group("/:roleID")
			{
				role.GET("", c.GetRolePrivilegesHandler)
				role.POST("/policies", c.AttachPoliciesToRoleHandler)
				role.DELETE("/policies/:policyID", c.DetachPolicyFromRoleHandler)
			}
		}
	}
}
