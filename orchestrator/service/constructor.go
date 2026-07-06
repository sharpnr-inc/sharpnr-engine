package orchestrator_service

import (
	hospitalitydomain "sharpnr-engine/app/hospitality_domain"
	org_customer_service "sharpnr-engine/core/customer/service"
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	legal_entity_service "sharpnr-engine/core/legal_entity/service"
	org_property_services "sharpnr-engine/core/org_property/services"
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"
	suppliers_service "sharpnr-engine/core/org_supplier/service"
	organization_services "sharpnr-engine/core/organization/services"

	payment_engine_service "sharpnr-engine/payment_engine/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	sentinelservice "sharpnr-engine/sentinel/service"
	sentinelAggregatorService "sharpnr-engine/sentinel_aggregators/service"
	sentinelEngine "sharpnr-engine/sentinel_engine"
	"sharpnr-engine/sero"
	accounting_service "sharpnr-engine/sharp_accounting_engine/accounting_engine/services"
	inventory_engine_services "sharpnr-engine/sharp_inventory_engine/services"
)

type orchestratorService struct {
	organizationService     organization_services.OrganizationServiceInterface
	orgPropertyService      org_property_services.OrganizationPropertyServicesInterfaces
	orgSetupWizService      org_setup_wizard_services.OrgSetupWizardV1_ServiceInterface
	hospitalityDomainSvc    *hospitalitydomain.HospitalityDomain
	accountingEngineService accounting_service.AccountingServiceInterface
	legalEntityService      legal_entity_service.OrgLegalEntityServiceInterface
	txManager               dbtx.TxManager
	seroEngine              *sero.SEROEngine
	customerEngine          org_customer_service.OrgCustomerServiceInterface
	paymentsEngine          payment_engine_service.PaymentServiceInterface
	inventoryEngine         inventory_engine_services.InventoryEngineServicesInterface
	supplierService         suppliers_service.SupplierServiceInterface
	employeeHRService       coreEmployeeServices.EmployeeHRServiceInterface
	employeeService         coreEmployeeServices.EmployeeOrchestratorServiceInterface
	sentinelService         sentinelservice.SentinelServiceInterface
	sentinelAggregator      sentinelAggregatorService.SentinelAggregatorServiceInterface
	sentinelEngine          sentinelEngine.SentinelEngineInterface
}

func NewSharpnrOrchestratorService(
	organizationService organization_services.OrganizationServiceInterface,
	orgPropertyService org_property_services.OrganizationPropertyServicesInterfaces,
	orgSetupWizService org_setup_wizard_services.OrgSetupWizardV1_ServiceInterface,
	hospitalityDomainSvc *hospitalitydomain.HospitalityDomain,
	accountingEngineService accounting_service.AccountingServiceInterface,
	legalEntityService legal_entity_service.OrgLegalEntityServiceInterface,
	txManager dbtx.TxManager,
	seroEngine *sero.SEROEngine,
	customerEngine org_customer_service.OrgCustomerServiceInterface,
	paymentsEngine payment_engine_service.PaymentServiceInterface,
	inventoryEngine inventory_engine_services.InventoryEngineServicesInterface,
	supplierService suppliers_service.SupplierServiceInterface,
	employeeHRService coreEmployeeServices.EmployeeHRServiceInterface,
	employeeService coreEmployeeServices.EmployeeOrchestratorServiceInterface,

	sentinelService sentinelservice.SentinelServiceInterface,
	sentinelAggregator sentinelAggregatorService.SentinelAggregatorServiceInterface,
	sentinelAuthzEngine sentinelEngine.SentinelEngineInterface,
) *orchestratorService {
	return &orchestratorService{
		organizationService:     organizationService,
		orgPropertyService:      orgPropertyService,
		orgSetupWizService:      orgSetupWizService,
		hospitalityDomainSvc:    hospitalityDomainSvc,
		accountingEngineService: accountingEngineService,
		legalEntityService:      legalEntityService,
		txManager:               txManager,
		seroEngine:              seroEngine,
		customerEngine:          customerEngine,
		paymentsEngine:          paymentsEngine,
		inventoryEngine:         inventoryEngine,
		supplierService:         supplierService,
		employeeHRService:       employeeHRService,
		employeeService:         employeeService,
		sentinelService:         sentinelService,
		sentinelAggregator:      sentinelAggregator,
		sentinelEngine:          sentinelAuthzEngine,
	}
}
