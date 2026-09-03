package cablesEngineRoutes

import (
	"sharpnr-engine/app"
	cablesEngineControllers "sharpnr-engine/cables_engine/controllers"

	"github.com/gin-gonic/gin"
)

func ResgisterCablesEngineRoutes(
	v1 *gin.RouterGroup,
	modules app.ModulesProvider,

) {

	cablesController := cablesEngineControllers.NewCablesEngineController(modules.CablesEngineService())

	cables := v1.Group("/cables")
	{
		cables.GET("/status", cablesController.GetCablesAccountStatus)
	}

}
