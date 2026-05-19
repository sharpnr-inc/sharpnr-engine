package hospitality_orchestration_routes

import (
	"sharpnr-engine/app"
	hospitality_orchestration_controllers "sharpnr-engine/modules/hospitality/orchestration/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterOrchestrationRoutes(r *gin.RouterGroup, m app.ModulesProvider) {

	hoc := hospitality_orchestration_controllers.NewOrchestrationController(m.HospitalityOrchestration())

	orchestration := r.Group("/orchestration")
	{
		orchestration.GET("/hotels", hoc.ListAllHotelsHandler)
		orchestration.GET("/restuarant", hoc.ListAllRestaurantHandler)
		orchestration.GET("/bar", hoc.ListAllBarHandler)
		orchestration.GET("/infra", hoc.ListAllInfraHandler)
	}
}
