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
		crm.GET("", crmContr.CrmAccountStatusHandler)
	}
}
