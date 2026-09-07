package hr_routes

import (
	"sharpnr-engine/app"
	hrmsEngineControllers "sharpnr-engine/hrms_engine/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterHRRoutes(v1 *gin.RouterGroup, modules app.ModulesProvider) {
	hrmsController := hrmsEngineControllers.NewHrmsControllers(modules.HRMSService())

	hr := v1.Group("/sharp-hr")

	{
		// Is HRMS on for this org, and is the account live? The client asks this
		// before anything else, so it answers serviceStarted=false rather than
		// 404 for an org that never started the service.
		//
		// Provisioning lives on the orchestrator (POST
		// /orchestrator/sharp-hr/start) — it is an admin action on the org, not
		// an HR product call, which is why it is not mirrored here. That is also
		// why the bare GET/POST stubs that used to sit on this group are gone:
		// registered with no handler, they answered an empty 200 instead of
		// failing loudly.
		hr.GET("/status", hrmsController.HrmsAccountStatusHandler)
	}
	registerEmployeeRoutes(hr, modules)
}
