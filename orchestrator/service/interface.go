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
	inventory_dtos "sharpnr-engine/service_engines/inventory_engine/dtos"
	inventory_engine_models "sharpnr-engine/service_engines/inventory_engine/models"

	"github.com/google/uuid"
)

type OrchestratorServiceInterface interface {
	OrganizationOrchestartionInterface
	OrganizationPropertyOrchestrationInterface
	LegalEntityOrchestrationServiceInterface
	CustomerOrchestratorServiceInterface
	PaymentOrchestratorServiceInterface
	SuppliersAndVendorsOrchestratorServiceInterface
	InventoryOrchestratorServiceInterface
	OrchestratorSetupWizardInterface
}

type OrganizationPropertyOrchestrationInterface interface {
	ConfigurePropertyToDomain(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID, req *orchestrator_dtos.AddDomainToPropertyReq_DTO) error
	AddModulesInProperty(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID, req *orchestrator_dtos.PropertyModuleConfigurationDTO) (*orchestrator_dtos.ModuleOperationResult, error)
	AddInfraInProperty(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID, req *orchestrator_dtos.PropertyInfraConfiguration) error
	HospitalityPropertyOrchestrationInterface
}

type HospitalityPropertyOrchestrationInterface interface {
	GetPropertyDetail(ctx context.Context, orgID, propertyID uuid.UUID) (*orchestrator_dtos.PropertyDetails_res_dto, error)
	CreateProperty(ctx context.Context, orgID uuid.UUID, req *orchestrator_dtos.CreatePropertyReq_dto) error
	GetPropertyInfra(ctx context.Context, propertyID uuid.UUID, orgID uuid.UUID, infraType *string) (*orchestrator_dtos.ListAllInfraOfProperty_Res_dto, error)
}

type OrganizationOrchestartionInterface interface {
	GetAllPropertiesOfOrg(ctx context.Context, orgID uuid.UUID) (*orchestrator_dtos.PropertiesOfOrganization_ResDto, error)
}

/* ---------------------------------- <***> --------------------------------- */

/* -------------------------------------------------------------------------- */
/*                     Orchestration for the Legal Entity                     */
/* -------------------------------------------------------------------------- */

type LegalEntityOrchestrationServiceInterface interface {
	GetAllLegalEntity(ctx context.Context, orgID, emplID uuid.UUID) ([]legal_entity_contracts.LegalEntityResponseDto, error)
	CreateLegalEntity(ctx context.Context, orgID, emplID uuid.UUID, req *legal_entity_contracts.NewLegalEntityRequestDto) error
	GetDetailsOfLegalEntity(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) (legal_entity_contracts.LegalEntityDetails, error)
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
