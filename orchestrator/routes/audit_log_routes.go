package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

// These are the audit log routes: the organization's own history screen. The
// organization is resolved from the session, so nothing here takes an org id —
// an employee can only ever read their own tenant's history.
func registerAuditLogEngineRoutes(rg *gin.RouterGroup, orc *orchestrator_controllers.OrchestratorController) {

	auditLogs := rg.Group("/audit-logs")
	{
		auditLogs.GET("", orc.ListOrgAuditLogsHandler) // one filtered, paginated page of the org history

		// The full trail of a single entity — everything that ever happened to
		// this department, this customer, this property.
		auditLogs.GET("/resource/:resourceType/:resourceID", orc.GetResourceAuditHistoryHandler)
	}
}
