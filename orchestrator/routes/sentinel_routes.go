package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

// registerSentinelRoutes mounts the authorization (Sentinel) endpoints used by
// the frontend to author managed policies, manage roles, and inspect the
// privileges a role grants.
//
//	/sentinel
//	├── GET  /permissions                          the permission catalog (picker source)
//	├── /policies
//	│   ├── GET                                     list managed policies (org + system)
//	│   └── POST                                    create a managed policy
//	└── /roles
//	    ├── GET                                     list roles
//	    ├── POST                                    create a role (+ optional initial policies)
//	    └── /:roleID
//	        ├── GET                                 role + its privileges
//	        └── /policies
//	            ├── POST                            attach managed policies
//	            └── DELETE /:policyID               detach a managed policy
func registerSentinelRoutes(rg *gin.RouterGroup, c *orchestrator_controllers.OrchestratorController) {
	s := rg.Group("/sentinel")
	{
		s.GET("/permissions", c.ListPermissionCatalogHandler)

		policies := s.Group("/policies")
		{
			policies.GET("", c.ListManagedPoliciesHandler)
			policies.POST("", c.CreateManagedPolicyHandler)
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
