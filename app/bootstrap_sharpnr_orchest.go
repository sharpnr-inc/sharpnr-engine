package app

import (
	hospitalitydomain "sharpnr-engine/app/hospitality_domain"
	org_customer_service "sharpnr-engine/core/customer/service"
	dbtx "sharpnr-engine/core/dbTx"
	legal_entity_service "sharpnr-engine/core/legal_entity/service"
	org_property_services "sharpnr-engine/core/org_property/services"
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"
	suppliers_service "sharpnr-engine/core/org_supplier/service"
	organization_services "sharpnr-engine/core/organization/services"
	hospitality_orchestration_service "sharpnr-engine/modules/hospitality/orchestration/service"
	orchestrator_service "sharpnr-engine/orchestrator/service"
	payment_engine_service "sharpnr-engine/payment_engine/services"
	"sharpnr-engine/sero"
	accounting_service "sharpnr-engine/service_engines/accounting_engine/services"
	inventory_engine_services "sharpnr-engine/service_engines/inventory_engine/services"
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
	)
}
