package app

import (
	hospitalitydomain "sharpnr-engine/app/hospitality_domain"
	auditEngineService "sharpnr-engine/audit_log_engine/service"
	cablesEngineService "sharpnr-engine/cables_engine/service"
	orgCostCenterService "sharpnr-engine/core/cost_center/service"
	org_customer_service "sharpnr-engine/core/customer/service"
	orgDepartmentService "sharpnr-engine/core/departments/service"
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	legal_entity_service "sharpnr-engine/core/legal_entity/service"
	org_property_services "sharpnr-engine/core/org_property/services"
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"
	suppliers_service "sharpnr-engine/core/org_supplier/service"
	organization_services "sharpnr-engine/core/organization/services"
	hospitality_orchestration_service "sharpnr-engine/property_business_ops/hospitality/orchestration/service"
	sharpCrmEngineService "sharpnr-engine/sharpCRM_engine/service"

	orchestrator_service "sharpnr-engine/orchestrator/service"
	payment_engine_service "sharpnr-engine/payment_engine/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	sentinelservice "sharpnr-engine/sentinel/service"
	sentinelAggregatorService "sharpnr-engine/sentinel_aggregators/service"
	sentinelEngine "sharpnr-engine/sentinel_engine"
	"sharpnr-engine/sero"
	accounting_service "sharpnr-engine/sharp_accounting_engine/accounting_engine/services"
	stockwaveEngineServices "sharpnr-engine/stockwave_engine/services"
)

func bootstrapSharpnrOrchestrator(
	organizationService organization_services.OrganizationServiceInterface,
	orgPropertyService org_property_services.OrganizationPropertyServicesInterfaces,
	orgSetupWizService org_setup_wizard_services.OrgSetupWizardV1_ServiceInterface,
	hospitalityDomainSvc *hospitalitydomain.HospitalityDomain,
	accountingEngineService accounting_service.AccountingServiceInterface,
	legalEntityService legal_entity_service.OrgLegalEntityServiceInterface,
	orgCustomerService org_customer_service.OrgCustomerServiceInterface,
	txManager dbtx.TxManager,
	seroEngine *sero.SEROEngine,
	paymentsEngine payment_engine_service.PaymentServiceInterface,
	inventoryEngine stockwaveEngineServices.InventoryEngineServicesInterface,
	supplierService suppliers_service.SupplierServiceInterface,
	employeeHRService coreEmployeeServices.EmployeeHRServiceInterface,
	employeeService coreEmployeeServices.EmployeeOrchestratorServiceInterface,
	sentinelService sentinelservice.SentinelServiceInterface,
	sentinelAggregator sentinelAggregatorService.SentinelAggregatorServiceInterface,
	sentinelAuthzEngine sentinelEngine.SentinelEngineInterface,
	auditLogService auditEngineService.AuditLogServiceInterface,
	orgDepartmentService orgDepartmentService.OrgDepartmentServiceInterface,
	orgCostCenterService orgCostCenterService.OrgCostCenterServiceInterface,
	cablesEngineService cablesEngineService.CablesEngineServiceInterface,
	sharpCrmEngineService sharpCrmEngineService.SharpCrmEngineServiceInterface,

) orchestrator_service.OrchestratorServiceInterface {
	return orchestrator_service.NewSharpnrOrchestratorService(
		organizationService,
		orgPropertyService,
		orgSetupWizService,
		hospitalityDomainSvc,
		accountingEngineService,
		legalEntityService,
		txManager,
		seroEngine,
		orgCustomerService,
		paymentsEngine,
		inventoryEngine,
		supplierService,
		employeeHRService,
		employeeService,
		sentinelService,
		sentinelAggregator,
		sentinelAuthzEngine,
		auditLogService,
		orgDepartmentService,
		orgCostCenterService,
		cablesEngineService,
		sharpCrmEngineService,
	)
}

func bootstrapHospitalityDomainOrchestration(
	orgPropertyService org_property_services.OrganizationPropertyServicesInterfaces,
	hospitalityDomainSvc *hospitalitydomain.HospitalityDomain,
	legalEntityService legal_entity_service.OrgLegalEntityServiceInterface,
) hospitality_orchestration_service.OrchestrationServiceInterface {
	return hospitality_orchestration_service.NewHospitalityOrchestrationService(
		orgPropertyService,
		hospitalityDomainSvc.HotelService(),
		hospitalityDomainSvc.RestaurantService(),
		legalEntityService,
		hospitalityDomainSvc.KitchenService(),
	)
}
