package orchestrator_service

import (
	"context"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"

	"github.com/google/uuid"
)

type OrchestratorServiceInterface interface {
	GetOrganizationStatus(ctx context.Context, orgID uuid.UUID) (bool, error)

	// SetupWiz
	AddOrganizationHQAddress(ctx context.Context, orgID uuid.UUID, req *orchestrator_dtos.SetupOrganizationHQAddressReqDTO) error
	StepTwoProperty(ctx context.Context, orgID uuid.UUID, req *orchestrator_dtos.SetupOrgPropertyReqDTO) error
	StepThreeLegalEntity(ctx context.Context, orgID uuid.UUID, req *orchestrator_dtos.SetupLegalEntityReqDTO) error
	SetUpWizProgress(ctx context.Context, orgID uuid.UUID) (orchestrator_dtos.SetupProgressResponseDTO, error)

	// Property
	GetAllPropertiesOfOrg(ctx context.Context, orgID uuid.UUID) (*orchestrator_dtos.PropertiesOfOrganization_ResDto, error)
	ConfigurePropertyToDomain(
		ctx context.Context,
		orgID uuid.UUID,
		propertyID uuid.UUID,
		req *orchestrator_dtos.AddDomainToPropertyReq_DTO,
	) error

	AddModulesInProperty(
		ctx context.Context,
		orgID uuid.UUID,
		propertyID uuid.UUID,
		req *orchestrator_dtos.PropertyModuleConfigurationDTO,
	) (*orchestrator_dtos.ModuleOperationResult, error)
	ConfigureRestaurant(
		ctx context.Context,
		orgID uuid.UUID,
		propertyID uuid.UUID,
		req *orchestrator_dtos.ConfigurePropertysRestaurant_Req_dto,
	) error

	ConfigureHotel(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID, req *orchestrator_dtos.ConfigureHotel_Request_dto) error

	GetPropertyDetail(ctx context.Context, orgID, propertyID uuid.UUID) (*orchestrator_dtos.PropertyDetails_res_dto, error)
	CreateProperty(ctx context.Context, orgID uuid.UUID, req *orchestrator_dtos.CreatePropertyReq_dto) error

	GetPropertiesAndRestaurant(
		ctx context.Context,
		orgID uuid.UUID,
		empID uuid.UUID,
	) (*orchestrator_dtos.PropertyOfResturants, error)

	//---------------Legal-Entity----------------------------------//

	CreateLegalEntity(ctx context.Context, orgID uuid.UUID, req *orchestrator_dtos.AddLegalEntity_REQ_DTO) error
	UpdateLegalEntity()
	GetAllLegalEntity(ctx context.Context, orgID uuid.UUID) ([]*orchestrator_dtos.LegalEntity_Response_dto, error)
}
