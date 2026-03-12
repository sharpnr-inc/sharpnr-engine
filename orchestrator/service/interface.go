package orchestrator_service

import (
	"context"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"

	"github.com/google/uuid"
)

type OrchestratorServiceInterface interface {
	OrganizationOrchestartionInterface
	OrganizationPropertyOrchestrationInterface
	LegalEntityOrchestrationServiceInterface
	/* -------------------------------------------------------------------------- */
	/*                       Setupwizard related operations                       */
	/* -------------------------------------------------------------------------- */
	AddOrganizationHQAddress(ctx context.Context, orgID uuid.UUID, req *orchestrator_dtos.SetupOrganizationHQAddressReqDTO) error
	StepTwoProperty(ctx context.Context, orgID uuid.UUID, req *orchestrator_dtos.SetupOrgPropertyReqDTO) error
	StepThreeLegalEntity(ctx context.Context, orgID uuid.UUID, req *orchestrator_dtos.SetupLegalEntityReqDTO) error
	SetUpWizProgress(ctx context.Context, orgID uuid.UUID) (orchestrator_dtos.SetupProgressResponseDTO, error)

	/* -------------------------------------------------------------------------- */
	/*                                 UI related                                 */
	/* -------------------------------------------------------------------------- */

	GetSharpnrLauncherData(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID) (orchestrator_dtos.OrgLauncherResponse, error)
}

type OrganizationPropertyOrchestrationInterface interface {
	ConfigurePropertyToDomain(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID, req *orchestrator_dtos.AddDomainToPropertyReq_DTO) error
	AddModulesInProperty(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID, req *orchestrator_dtos.PropertyModuleConfigurationDTO) (*orchestrator_dtos.ModuleOperationResult, error)
	AddInfraInProperty(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID, req *orchestrator_dtos.PropertyInfraConfiguration) error
	HospitalityPropertyOrchestrationInterface
}

type HospitalityPropertyOrchestrationInterface interface {
	GetAllRestaurant(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID) ([]orchestrator_dtos.AllRestaurantsResponseDto, error)
	GetAllKitchen(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID) ([]orchestrator_dtos.AllKitchenResponseDto, error)
	GetPropertyDetail(ctx context.Context, orgID, propertyID uuid.UUID) (*orchestrator_dtos.PropertyDetails_res_dto, error)
	CreateProperty(ctx context.Context, orgID uuid.UUID, req *orchestrator_dtos.CreatePropertyReq_dto) error
	GetPropertyInfra(ctx context.Context, propertyID uuid.UUID, orgID uuid.UUID, infraType *string) (*orchestrator_dtos.ListAllInfraOfProperty_Res_dto, error)
}

type OrganizationOrchestartionInterface interface {
	GetOrganizationStatus(ctx context.Context, orgID uuid.UUID) (*string, error)
	GetAllPropertiesOfOrg(ctx context.Context, orgID uuid.UUID) (*orchestrator_dtos.PropertiesOfOrganization_ResDto, error)
}

/* ---------------------------------- <***> --------------------------------- */

/* -------------------------------------------------------------------------- */
/*                     Orchestration for the Legal Entity                     */
/* -------------------------------------------------------------------------- */

type LegalEntityOrchestrationServiceInterface interface {
	CreateLegalEntity(ctx context.Context, orgID uuid.UUID, req *orchestrator_dtos.AddLegalEntity_REQ_DTO) error
	GetAllLegalEntity(ctx context.Context, orgID uuid.UUID) ([]*orchestrator_dtos.LegalEntity_Response_dto, error)

	// Tax Registration related OP
	AddTaxRegsitrationInLegalEntity(
		ctx context.Context,
		orgID, emplID, legalEntity uuid.UUID,
		req legal_entity_contracts.CreateLegalEntityTaxRegistrationRequestDto,
	) error

	GetAllTaxRegistrationOfLegalEntity(
		ctx context.Context,
		orgID, emplID, legalEntity uuid.UUID,
	) ([]legal_entity_contracts.LegalEntityTaxRegistrationResponseDto, error)

	AddTaxCategoryInLegalEntiy(
		ctx context.Context,
		orgID, emplID, legalEntity uuid.UUID,
		req legal_entity_contracts.CreateLegalEntityTaxCategory,
	) error

	ListAllTaxCategoryOfLegalEntity(
		ctx context.Context,
		orgID, emplID, legalEntity uuid.UUID,
	) (interface{}, error)
}

/* ---------------------------------- <***> --------------------------------- */
