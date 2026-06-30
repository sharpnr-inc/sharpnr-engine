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
func registerSentinelRoutes(rg *gin.RouterGroup, c *orchestrator_controllers.OrchestratorController) {
	s := rg.Group("/sentinel")
	{
		s.GET("/principals", c.ListPrincipalsHandler)

		policies := s.Group("/policies")
		{
			policies.GET("", c.ListPoliciesHandler)
			policies.POST("", c.CreatePolicyHandler)
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
