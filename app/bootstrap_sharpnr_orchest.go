package app

import (
	hospitalitydomain "sharpnr-engine/app/hospitality_domain"
	hospitality_orchestration_service "sharpnr-engine/business_ops/hospitality/orchestration/service"
	org_customer_service "sharpnr-engine/core/customer/service"
	legal_entity_service "sharpnr-engine/core/legal_entity/service"
	org_property_services "sharpnr-engine/core/org_property/services"
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"
	suppliers_service "sharpnr-engine/core/org_supplier/service"
	organization_services "sharpnr-engine/core/organization/services"
	employee_services "sharpnr-engine/hrms_engine/employees/services"
	orchestrator_service "sharpnr-engine/orchestrator/service"
	payment_engine_service "sharpnr-engine/payment_engine/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	sentinelengine "sharpnr-engine/sentinel_engine/engine"
	sentinelservice "sharpnr-engine/sentinel_engine/service"
	"sharpnr-engine/sero"
	accounting_service "sharpnr-engine/sharp_accounting_engine/accounting_engine/services"
	inventory_engine_services "sharpnr-engine/sharp_inventory_engine/services"
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
	inventoryEngine inventory_engine_services.InventoryEngineServicesInterface,
	supplierService suppliers_service.SupplierServiceInterface,
	employeeHRService employee_services.EmployeesHRServiceInterface,
	sentinelEngine *sentinelengine.Sentinel,
	sentinelService sentinelservice.SentinelService,
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
		sentinelEngine,
		sentinelService,
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
