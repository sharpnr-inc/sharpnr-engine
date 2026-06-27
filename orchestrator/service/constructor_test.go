package orchestrator_service

// Tests for NewSharpnrOrchestratorService.
//
// The constructor has no logic beyond wiring: it takes 12 collaborators and
// stores each on the orchestratorService struct. So the test is a pure WIRING
// check — pass a distinct value for every parameter, then assert that each
// struct field holds exactly that value (same pointer). This catches the
// classic copy/paste bug where two fields get assigned the same argument.
//
// Notes on the technique:
//   - White-box (same package): we read the unexported fields directly.
//   - Stubs embed their interface (nil) purely to be a distinct, type-correct
//     value to pass in. No methods are ever called, so nil embedding is safe.
//     stubOrgPropertyService / stubOrganizationService / stubTxManager are
//     reused from properties_test.go.
//   - require.Same asserts pointer identity (the value came through untouched),
//     which is stronger than require.Equal for this purpose.

import (
	"testing"

	org_customer_service "sharpnr-engine/core/customer/service"
	legal_entity_service "sharpnr-engine/core/legal_entity/service"
	org_setup_wizard_services "sharpnr-engine/core/org_setup_wizard/services"
	suppliers_service "sharpnr-engine/core/org_supplier/service"
	employee_services "sharpnr-engine/hrms_engine/employees/services"
	payment_engine_service "sharpnr-engine/payment_engine/services"
	sentinelengine "sharpnr-engine/sentinel_engine/engine"
	sentinelservice "sharpnr-engine/sentinel_engine/service"
	"sharpnr-engine/sero"
	accounting_service "sharpnr-engine/sharp_accounting_engine/accounting_engine/services"
	inventory_engine_services "sharpnr-engine/sharp_inventory_engine/services"

	hospitalitydomain "sharpnr-engine/app/hospitality_domain"

	"github.com/go-openapi/testify/v2/require"
)

/* ========================================================================== */
/*                       ADDITIONAL WIRING-ONLY STUBS                         */
/* ========================================================================== */
/*
   These collaborators are only constructed (never called) in this test, so a
   bare interface-embedding stub is enough. The interfaces stubbed in
   properties_test.go (org property, organization, tx manager) are reused.
*/

type stubSetupWizService struct {
	org_setup_wizard_services.OrgSetupWizardV1_ServiceInterface
}

type stubAccountingService struct {
	accounting_service.AccountingServiceInterface
}

type stubLegalEntityService struct {
	legal_entity_service.OrgLegalEntityServiceInterface
}

type stubCustomerService struct {
	org_customer_service.OrgCustomerServiceInterface
}

type stubPaymentService struct {
	payment_engine_service.PaymentServiceInterface
}

type stubInventoryService struct {
	inventory_engine_services.InventoryEngineServicesInterface
}

type stubSupplierService struct {
	suppliers_service.SupplierServiceInterface
}

type stubEmployeeService struct {
	employee_services.EmployeesHRServiceInterface
}

/* ========================================================================== */
/*                     NewSharpnrOrchestratorService                          */
/* ========================================================================== */

func TestNewSharpnrOrchestratorService_WiresAllDependencies(t *testing.T) {
	// --- arrange: one distinct value per constructor parameter ---
	var (
		organizationService = &stubOrganizationService{}
		orgPropertyService  = &stubOrgPropertyService{}
		orgSetupWizService  = &stubSetupWizService{}
		hospitalityDomain   = &hospitalitydomain.HospitalityDomain{}
		accountingService   = &stubAccountingService{}
		legalEntityService  = &stubLegalEntityService{}
		txManager           = &stubTxManager{}
		seroEngine          = &sero.SEROEngine{}
		customerEngine      = &stubCustomerService{}
		paymentsEngine      = &stubPaymentService{}
		inventoryEngine     = &stubInventoryService{}
		supplierService     = &stubSupplierService{}
		employeeHRService   = &stubEmployeeService{}
		// nil repo is fine: this is a wiring test, no authorization call is made.
		sentinelEngine  = sentinelengine.New(nil)
		sentinelService = sentinelservice.NewSentinelService(nil, nil)
	)

	// --- act ---
	svc := NewSharpnrOrchestratorService(
		organizationService,
		orgPropertyService,
		orgSetupWizService,
		hospitalityDomain,
		accountingService,
		legalEntityService,
		txManager,
		seroEngine,
		customerEngine,
		paymentsEngine,
		inventoryEngine,
		supplierService,
		employeeHRService,
		sentinelEngine,
		sentinelService,
	)

	// --- assert: non-nil, and every field wired to the exact value passed ---
	require.NotNil(t, svc)

	require.Same(t, organizationService, svc.organizationService)
	require.Same(t, orgPropertyService, svc.orgPropertyService)
	require.Same(t, orgSetupWizService, svc.orgSetupWizService)
	require.Same(t, hospitalityDomain, svc.hospitalityDomainSvc)
	require.Same(t, accountingService, svc.accountingEngineService)
	require.Same(t, legalEntityService, svc.legalEntityService)
	require.Same(t, txManager, svc.txManager)
	require.Same(t, seroEngine, svc.seroEngine)
	require.Same(t, customerEngine, svc.customerEngine)
	require.Same(t, paymentsEngine, svc.paymentsEngine)
	require.Same(t, inventoryEngine, svc.inventoryEngine)
	require.Same(t, supplierService, svc.supplierService)
	require.Same(t, employeeHRService, svc.employeeHRService)
	require.Same(t, sentinelEngine, svc.sentinelEngine)
	require.Same(t, sentinelService, svc.sentinelService)
}
