package orchestrator_routes

import (
	"sharpnr-engine/app"

	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterOrchestratorRoutes(
	webV1 *gin.RouterGroup,
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
	registerEmployeeRoutes(org, orchestController)
	registerSentinelRoutes(org, orchestController)
	registerAuditLogEngineRoutes(org, orchestController)
	registerDepartmentRoutes(org, orchestController)
	registerCostCenterRoutes(org, orchestController)
	registerCablesOrchestratorRoutes(org, orchestController)
	registerCrmOrchestratorRoutes(org, orchestController)
	registerCortexOrchestratorRoutes(org, orchestController)
	registerSiloOrchestratorRoutes(org, orchestController)

}
