package orchestrator_service

import (
	hospitalitydomain "sharpnr-engine/app/hospitality_domain"
	auditEngineService "sharpnr-engine/audit_log_engine/service"
	org_customer_service "sharpnr-engine/core/customer/service"
	coreEmployeeServices "sharpnr-engine/core/employees/services"
	legal_entity_service "sharpnr-engine/core/legal_entity/service"
	org_property_services "sharpnr-engine/core/org_property/services"
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"
	suppliers_service "sharpnr-engine/core/org_supplier/service"
	organization_services "sharpnr-engine/core/organization/services"
	orchestrator_authz "sharpnr-engine/orchestrator/service/authz"
	property_orchestrator "sharpnr-engine/orchestrator/service/property"
	propertyHospitalityOrchestrator "sharpnr-engine/orchestrator/service/property/hospitality"

	payment_engine_service "sharpnr-engine/payment_engine/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	sentinelservice "sharpnr-engine/sentinel/service"
	sentinelAggregatorService "sharpnr-engine/sentinel_aggregators/service"
	sentinelEngine "sharpnr-engine/sentinel_engine"
	"sharpnr-engine/sero"
	accounting_service "sharpnr-engine/sharp_accounting_engine/accounting_engine/services"
	inventory_engine_services "sharpnr-engine/sharp_inventory_engine/services"
)

// orchestratorService is the composed facade: one embedded orchestrator per
// domain, each defined in its own file with only the dependencies it needs.
// Method promotion makes this struct satisfy OrchestratorServiceInterface.
type orchestratorService struct {
	*property_orchestrator.PropertyOrchestrator
	*setupWizardOrchestrator
	*legalEntityOrchestrator
	*customerOrchestrator
	*paymentsOrchestrator
	*suppliersOrchestrator
	*inventoryOrchestrator
	*employeeOrchestrator
	*sentinelOrchestrator
	*auditOrchestraterService

	// Parked: wired at bootstrap but not consumed by any flow yet. Move onto
	// the owning domain orchestrator when the first flow lands.
	seroEngine *sero.SEROEngine
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
	auditLogEngineService auditEngineService.AuditLogServiceInterface,
) *orchestratorService {

	// One guard shared by every domain orchestrator (rule 1 in doc.go).
	guard := orchestrator_authz.New(sentinelAuthzEngine)

	return &orchestratorService{
		// The property tree lives in its own package, one sub-package per
		// business domain under it. See orchestrator/service/property/doc.go.
		PropertyOrchestrator: property_orchestrator.New(
			guard,
			txManager,
			orgPropertyService,
			organizationService,
			propertyHospitalityOrchestrator.New(guard, hospitalityDomainSvc),
		),
		setupWizardOrchestrator: &setupWizardOrchestrator{
			txManager:           txManager,
			orgSetupWizService:  orgSetupWizService,
			organizationService: organizationService,
		},
		legalEntityOrchestrator: &legalEntityOrchestrator{
			txManager:               txManager,
			legalEntityService:      legalEntityService,
			accountingEngineService: accountingEngineService,
		},
		customerOrchestrator: &customerOrchestrator{
			txManager:      txManager,
			customerEngine: customerEngine,
		},
		paymentsOrchestrator: &paymentsOrchestrator{
			txManager:      txManager,
			paymentsEngine: paymentsEngine,
		},
		suppliersOrchestrator: &suppliersOrchestrator{
			txManager:          txManager,
			supplierService:    supplierService,
			legalEntityService: legalEntityService,
		},
		inventoryOrchestrator: &inventoryOrchestrator{
			txManager:       txManager,
			inventoryEngine: inventoryEngine,
		},
		employeeOrchestrator: &employeeOrchestrator{
			employeeHRService: employeeHRService,
			employeeService:   employeeService,
			sentinelService:   sentinelService,
		},
		sentinelOrchestrator: &sentinelOrchestrator{
			guard:              guard,
			sentinelService:    sentinelService,
			sentinelAggregator: sentinelAggregator,
		},

		auditOrchestraterService: &auditOrchestraterService{
			auditLogService: auditLogEngineService,
		},

		seroEngine: seroEngine,
	}
}
