package orchestrator_routes

import (
	"sharpnr-engine/app"
	"sharpnr-engine/cookieManager"
	dbtx "sharpnr-engine/core/dbTx"
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

	org := webV1.Group("/orchestrator")

	registerPropertyRoutes(org, orchestController)
	registerLegalEntityRoutes(org, orchestController)
	registerSetupWizardRoutes(org, orchestController)
	registerCustomerRoutes(org, orchestController)
	paymentOrchestratorRoutes(org, orchestController)
	registerOrgSupplierRoutes(org, orchestController)
	registerInventoryRoutes(org, orchestController)

}
