package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

// registerStockwaveOrchestratorRoutes replaces registerInventoryRoutes.
//
// The old /inventory-account group is gone with the tables it read: its
// handlers were empty stubs, and migration 000062 dropped
// inventory.inventory_accounts along with the rest of 000013's schema. The
// group comes back as the stock structures are redesigned.
func registerStockwaveOrchestratorRoutes(r *gin.RouterGroup, c *orchestrator_controllers.OrchestratorController) {
	stockwave := r.Group("/stockwave")
	{
		// Has the org taken stockwave, and is the account live? The admin shell
		// asks this before offering the start action, so it answers
		// serviceStarted=false rather than 404 for an org that never started it.
		stockwave.GET("", c.GetStockwaveAccountStatusHandler)

		// Provisioning lives here rather than on the stockwave module routes: it
		// is an admin action on the organization, not a stockwave product call.
		stockwave.POST("/start", c.StartStockwaveServiceHandler)
	}
}
