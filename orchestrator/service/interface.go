package orchestrator_service

import (
	"context"
	org_customer_contracts "sharpnr-engine/core/customer/contracts"
	org_customer_models "sharpnr-engine/core/customer/models"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"
	supplier_contracts "sharpnr-engine/core/org_supplier/contracts"
	suppliers_dtos "sharpnr-engine/core/org_supplier/dtos"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
	payment_engine_contracts "sharpnr-engine/payment_engine/contracts"
	sentinelDtos "sharpnr-engine/sentinel/dtos"
	sentinelengine_models "sharpnr-engine/sentinel/models"
	sentinelAggregatorService "sharpnr-engine/sentinel_aggregators/service"
	sentinelContexts "sharpnr-engine/sentinel_contexts"
	accountingdtos "sharpnr-engine/sharp_accounting_engine/accounting_engine/dto"
	inventory_dtos "sharpnr-engine/sharp_inventory_engine/dtos"
	inventory_engine_models "sharpnr-engine/sharp_inventory_engine/models"

	"github.com/google/uuid"
)

type OrchestratorServiceInterface interface {
	OrganizationOrchestrationInterface
	OrganizationPropertyOrchestrationInterface
	LegalEntityOrchestrationServiceInterface
	CustomerOrchestratorServiceInterface
	PaymentOrchestratorServiceInterface
	SuppliersAndVendorsOrchestratorServiceInterface
	InventoryOrchestratorServiceInterface
	OrchestratorSetupWizardInterface
	EmployeeOrchestrationInterface
	SentinelOrchestrationInterface
}

/* ---------------------------------- <***> --------------------------------- */

// SentinelOrchestrationInterface exposes the authorization authoring + display
// surface (Cedar policies, roles, privileges) to the frontend. Everything
// delegates to the Sentinel service; the orchestrator is just the HTTP facade.
type SentinelOrchestrationInterface interface {
	// Policy-builder palettes. Principals: types + roles + groups.
	ListPrincipals(ctx context.Context, orgID, emplID uuid.UUID) (sentinelDtos.PrincipalsView, error)

	// Resources palette, served by the sentinel aggregator: which resource
	// types exist, and per type its actions, attributes and live instances.
	ListAuthzResources(ctx context.Context, orgID, emplID uuid.UUID) ([]sentinelAggregatorService.ResourceSummaryView, error)
	GetAuthzResourceDetails(ctx context.Context, orgID, emplID uuid.UUID, resourceType string) (sentinelAggregatorService.ResourceDetailsView, error)

	// Context palette: the canonical context.*/principal.* keys the engine
	// stamps on every evaluation, for the builder's condition editor.
	GetAuthzContextPalette(ctx context.Context, orgID, emplID uuid.UUID) (sentinelContexts.Palette, error)

	ListPolicies(ctx context.Context, orgID, emplID uuid.UUID, scope sentinelengine_models.PolicyScope) ([]sentinelDtos.PolicyView, error)
	CreatePolicy(ctx context.Context, orgID, emplID uuid.UUID, in sentinelDtos.AddPolicyRequestDto) (uuid.UUID, error)

	ListRoles(ctx context.Context, orgID, emplID uuid.UUID) ([]sentinelDtos.RoleView, error)
	CreateRole(ctx context.Context, orgID, emplID uuid.UUID, in sentinelDtos.CreateRoleInput) (uuid.UUID, error)
	GetRolePrivileges(ctx context.Context, orgID, emplID, roleID uuid.UUID) (sentinelDtos.RolePrivilegesView, error)

	AttachPoliciesToRole(ctx context.Context, orgID, emplID, roleID uuid.UUID, in sentinelDtos.AttachPoliciesInput) error
	DetachPolicyFromRole(ctx context.Context, orgID, emplID, roleID, policyID uuid.UUID) error
}

type OrganizationOrchestrationInterface interface {
}

/* ---------------------------------- <***> --------------------------------- */

type EmployeeOrchestrationInterface interface {
	GetAllEmployeesOfOrg(ctx context.Context, orgID, emplID uuid.UUID) ([]orchestrator_dtos.EmployeeOfOrganization, error)
}

/* ---------------------------------- <***> --------------------------------- */

// Property Operations in Orchestraor services interface
type OrganizationPropertyOrchestrationInterface interface {

	// General Property service methods
	GetAllPropertiesOfOrg(ctx context.Context, orgID, emplID uuid.UUID) (*orchestrator_dtos.PropertiesOfOrganization_ResDto, error)
	GetPropertyDetail(ctx context.Context, orgID, emplID, propertyID uuid.UUID) (*orchestrator_dtos.PropertyDetails_res_dto, error)
	CreateProperty(ctx context.Context, orgID, emplID uuid.UUID, req *orchestrator_dtos.CreatePropertyReq_dto) error

	// Business Modules
	AddBusinessOperationModulesInProperty(ctx context.Context, orgID, emplID, propertyID uuid.UUID, req orchestrator_dtos.AddBusinessOpModulesRequestDto) error
	ListAllBusinessOpsModules(ctx context.Context, orgID, emplID, propertyID uuid.UUID) (orchestrator_dtos.PropertyBusinessOperationModulesResponseDto, error)

	// Infra Modules
	AddInfraModulesInProperty(ctx context.Context, orgID, emplID, propertyID uuid.UUID, request orchestrator_dtos.AddInfraModulesRequestDto) error
	ListAllInfraModulesOfProperty(ctx context.Context, orgID, emplID, propertyID uuid.UUID) (orchestrator_dtos.PropertyInfraModulesResponseDto, error)
}

/* ---------------------------------- <***> --------------------------------- */

/* -------------------------------------------------------------------------- */
/*                     Orchestration for the Legal Entity                     */
/* -------------------------------------------------------------------------- */

type LegalEntityOrchestrationServiceInterface interface {
	GetAllLegalEntity(ctx context.Context, orgID, emplID uuid.UUID) ([]legal_entity_contracts.LegalEntityResponseDto, error)
	CreateLegalEntity(ctx context.Context, orgID, emplID uuid.UUID, req *legal_entity_contracts.NewLegalEntityRequestDto) error
	GetDetailsOfLegalEntity(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) (legal_entity_contracts.LegalEntityDetails, error)

	AddContactsInLegalEntity(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID, req *legal_entity_contracts.CreateLegalEntityContactsRequestDto) error

	AddAddressInLegalEntity(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID, req *legal_entity_contracts.CreateLegalEntityAddressRequest) error
	AddBankAccountDetailsInLegalEntity(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID, req *legal_entity_contracts.AddBankAccountInLegalEntityRequestDto) error
	AddIdentifiersInLegalEntity(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID, req *legal_entity_contracts.AddIdentifiersInLegalEntityRequestDto) error
	AddLicenseOrComplianceInLegalEntity(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID, req legal_entity_contracts.AddNewLicenseAndComplianceRequestDto) error
	AddTaxRegistrationInLegalEntity(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID, req legal_entity_contracts.AddNewTaxRegistrationInLegalEntityRequestDto) error

	/* -------------------------------- acconting ------------------------------- */
	GetAccountingBooksOfLegalEntity(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) (accountingdtos.AccountingBookOfLegalEntity_Response_dto, error)
	GetAllAccountBalancesOfLegalEntity(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) ([]accountingdtos.AccountGroupBalanceResponseDto, error)
}

/* ---------------------------------- <***> --------------------------------- */

type CustomerOrchestratorServiceInterface interface {
	AddCustomer(ctx context.Context, orgID, emplID uuid.UUID, req org_customer_contracts.AddCustomerRequestDto) error
	SearchCustomer(ctx context.Context, orgID, emplID uuid.UUID, searchQuery string) ([]org_customer_models.Customer, error)
}

type PaymentOrchestratorServiceInterface interface {
	AddPaymentAccount(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, req payment_engine_contracts.CreatePaymentAccountRequest) error
	GetAllPaymentAccount(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID) ([]payment_engine_contracts.PaymentAccountResponseDto, error)
	AddGatewayConfigInPaymentAccount(ctx context.Context, orgID, emplID uuid.UUID, req payment_engine_contracts.AddGatewayConfigRequest) error
	GetPaymentAccountDetails(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, paymentAccountID uuid.UUID) ([]payment_engine_contracts.GatewayConfigResponse, error)
}

type SuppliersAndVendorsOrchestratorServiceInterface interface {
	AddNewSupplier(ctx context.Context, orgID, emplID uuid.UUID, request *suppliers_dtos.CreateSuppliersAPIRequestdto) error
	ListAllOrganizationSuppliers(ctx context.Context, orgID, emplID uuid.UUID) ([]suppliers_dtos.SuppliersResponseDetails, error)
	AddLegalEntityLinksInSuppliers(ctx context.Context, orgID, emplID, suppliersID uuid.UUID, request *supplier_contracts.AddLegalEntityLinkInSuppliersRequestDto) error
}

type InventoryOrchestratorServiceInterface interface {
	AddNewInventoryAccount(ctx context.Context, orgID, emplID uuid.UUID, request *inventory_dtos.NewInventoryAccountRequestDto) error
	ListAllInventoryAccount(ctx context.Context, orgID, emplID uuid.UUID) ([]inventory_engine_models.InventoryAccount, error)
	GetInventoryAccountDetails(ctx context.Context, orgID, emplID, InventoryAccountID uuid.UUID) ([]inventory_engine_models.Inventory, error)
}

/* --------------------------------- <****> --------------------------------- */

// All methods related to the setup wizard orchestration will be defined in this interface. This includes methods for each step of the setup wizard, as well as a method to retrieve the current
// progress of the setup wizard for a given organization.
type OrchestratorSetupWizardInterface interface {
	SetupWiz_AddProperty(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, req *orchestrator_dtos.SetupOrgPropertyReqDTO) error

	// This method retrieves the current progress of the setup wizard for a given organization. It returns a structured response containing the overall status, current step, progress percentage,
	// completion status, and details of each step in the wizard.
	SetUpWizProgress(ctx context.Context, orgID uuid.UUID) (orchestrator_dtos.SetupProgressResponseDTO, error)

	// This methods updates the superadmin details during setup
	// we already created the super admin during the signup
	SetUpSuperAdminDetails(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, req *orchestrator_dtos.SetupSuperAdminDetailsReqDTO) error
	SetupWiz_UpdateOrgName(ctx context.Context, orgID, emplID uuid.UUID, request *orchestrator_dtos.SetupWizUpdateOrgNameRequestDto) error
}

/* --------------------------------- <****> --------------------------------- */
