package orchestrator_routes

import (
	"sharpnr-engine/app"
	"sharpnr-engine/pkg/cookieManager"

	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"
	dbtx "sharpnr-engine/pkg/dbTx"
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
	registerOrgRoutes(org, orchestController)
	registerPropertyRoutes(org, orchestController)
	registerLegalEntityRoutes(org, orchestController)
	registerSetupWizardRoutes(org, orchestController)
	registerCustomerRoutes(org, orchestController)
	paymentOrchestratorRoutes(org, orchestController)
	registerOrgSupplierRoutes(org, orchestController)
	registerInventoryRoutes(org, orchestController)
	registerUtilsRoutes(org, orchestController)

}
