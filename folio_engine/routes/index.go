package folioEngineRoutes

import (
	"sharpnr-engine/app"
	folioEngineControllers "sharpnr-engine/folio_engine/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterFolioEngineRoutes(r *gin.RouterGroup, modules app.ModulesProvider) {

	folioContr := folioEngineControllers.NewFolioEngineController(modules.FolioEngineService())

	folio := r.Group("/folio")
	{
		// Is Folio on for this org, and is the account live? The client asks
		// this before anything else, so it answers serviceStarted=false rather
		// than 404 for an org that never started the service.
		//
		// Provisioning lives on the orchestrator (POST
		// /orchestrator/folio/start) — it is an admin action on the org, not a
		// Folio product call, which is why it is not mirrored here.
		folio.GET("/status", folioContr.FolioAccountStatusHandler)
	}

}
