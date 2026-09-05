package sharpCrmEngineRoutes

import (
	"sharpnr-engine/app"
	sharpCrmControllers "sharpnr-engine/sharpCRM_engine/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterSharpCrmEngine(r *gin.RouterGroup, modules app.ModulesProvider) {

	crmContr := sharpCrmControllers.NewSharpCrmControllers(modules.SharpCrmEngineService())

	crm := r.Group("/sharpcrm")
	{
		// Is SharpCRM on for this org, and is the account live? The client asks
		// this before anything else, so it answers serviceStarted=false rather
		// than 404 for an org that never started the service.
		//
		// Provisioning lives on the orchestrator (POST
		// /orchestrator/sharpcrm/start) — it is an admin action on the org, not
		// a CRM product call, which is why it is not mirrored here.
		crm.GET("/status", crmContr.CrmAccountStatusHandler)
	}
}
