package orchestrator_routes

import (
	"sharpnr-engine/app"
	"sharpnr-engine/cookieManager"
	dbtx "sharpnr-engine/core/dbTx"
	orchestrator_controllers "sharpnr-engine/core/orchestrator/controllers"
	webMiddleware "sharpnr-engine/middlewares/web"
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

	org := webV1.Group("/orchestrator", webMiddleware.WebAuthMiddleware(jwtMgr))

	registerConsoleRoutes(org, orchestController)
	registerSetupWizardRoutes(org, orchestController)
	registerLegalEntityRoutes()

}
