package orchestrator

import (
	"sharpnr-engine/app"
	"sharpnr-engine/cookieManager"
	dbtx "sharpnr-engine/core/dbTx"
	orchestrator_controllers "sharpnr-engine/core/orchestrator/controllers"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"
	"sharpnr-engine/sere"

	"github.com/gin-gonic/gin"
)

func RegisterOrchestratorRoutes(
	webV1 *gin.RouterGroup,
	txManager dbtx.TxManager,
	sere *sere.SERE,
	cookieMgr cookieManager.Manager,
	jwtMgr jwt_manager.Manager,
	sharpModules app.ModulesProvider,
) {

	orchestService := sharpModules.SharpnrOrchestrator()
	orchestController := orchestrator_controllers.NewOrchestratorController(orchestService)

	org := webV1.Group("/orchestrator")
	{
		orgHospitality := org.Group("/hospitality")
		{
			orgHospitality.GET("/overview", orchestController.GetOrgOverview)

			hotelConfigs := orgHospitality.Group("/hotel")
			{
				hotelConfigs.GET("/overview")
			}
		}
	}

}
