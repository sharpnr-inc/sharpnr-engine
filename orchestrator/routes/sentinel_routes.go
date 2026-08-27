package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerSentinelRoutes(rg *gin.RouterGroup, c *orchestrator_controllers.OrchestratorController) {
	s := rg.Group("/sentinel")
	{
		s.GET("/principals", c.ListPrincipalsHandler)       // list all the principal i.e Users of organization
		s.GET("/contexts", c.GetAuthzContextPaletteHandler) // TODO: list all the context available for now this not ready yet

		// This is the endpoint for the Resources
		resources := s.Group("/resources")
		{
			resources.GET("", c.ListAuthzResourcesHandler)                    // list all the Resources that is available to the organization
			resources.GET("/:resourceType", c.GetAuthzResourceDetailsHandler) // TODO:based on the resourceType
		}

		// Org Polices Endpoints
		org_policies := s.Group("/orgpolicies")
		{
			org_policies.GET("", c.ListPoliciesHandler)        // this endpoint list all polices attached straight to organization
			org_policies.POST("", c.CreatePolicyHandler)       // this endpoint add/create policies straight to organization
			org_policies.GET("/:policyID", c.GetPolicyHandler) // to fetch the Policy Details
		}

		// Roles Endpoint
		roles := s.Group("/roles")
		{
			roles.GET("", c.ListRolesHandler)   // list all roles owned by the organization.
			roles.POST("", c.CreateRoleHandler) // Creates/Add new roles in the organization.

			role := roles.Group("/:roleID")
			{
				role.GET("", c.GetRolePrivilegesHandler)                          // this fetches Role Details, like user who are into this Role
				role.POST("/policies", c.AttachPoliciesToRoleHandler)             // Add polices in the Role
				role.DELETE("/policies/:policyID", c.DetachPolicyFromRoleHandler) // Remove the policy that is attached to the role
				role.GET("/user/:userID")                                         // this will add user in this role
				role.DELETE("/user/:userID")                                      // this will remove the user from the role
			}
		}

		// Groups Endpoint
		groups := s.Group("/groups")
		{
			groups.GET("")  // list all the groups in the organization
			groups.POST("") // Create/Add new Groups in the organization

			group := groups.Group("/:groupID")
			{
				group.GET("")                      // list details of Groups Like Users in the Groups, etc
				group.PATCH("")                    // Edit Group Details
				group.POST("/polices")             // Add new/polices
				group.POST("/polices/:policyID")   // Edit policy
				group.DELETE("/polices/:policyID") // Remove the policy from the Group
				group.GET("/user/:userID")         // Add user in the Groups
				group.DELETE("/user/:userID")      // Remove the user from the Group

			}
		}
	}
}
