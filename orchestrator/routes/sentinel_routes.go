package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

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

		org_policies := s.Group("/policies")
		{
			org_policies.GET("", c.ListPoliciesHandler)
			org_policies.POST("", c.CreatePolicyHandler)
			org_policies.GET("/:policyID", c.GetPolicyHandler)
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
