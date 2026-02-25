package orchestrator_routes

import (
	"sharpnr-engine/app"
	"sharpnr-engine/cookieManager"
	dbtx "sharpnr-engine/core/dbTx"
	webMiddleware "sharpnr-engine/middlewares/web"
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"

	"github.com/gin-gonic/gin"
)

func RegisterOrchestratorRoutes(
	webV1 *gin.RouterGroup,
	txManager dbtx.TxManager,
	cookieMgr cookieManager.Manager,
	jwtMgr jwt_manager.Manager,
	sharpModules app.ModulesProvider,
) {

	orchestService := sharpModules.SharpnrOrchestrator()
	orchestController := orchestrator_controllers.NewOrchestratorController(orchestService)

	org := webV1.Group("/orchestrator", webMiddleware.WebAuthMiddleware(jwtMgr, cookieMgr))

	registerConsoleRoutes(org, orchestController)
	registerSetupWizardRoutes(org, orchestController)
	registerHospitalityRoutes(org, orchestController)
	registersharnprLauncherRoutes(org, orchestController)

}
