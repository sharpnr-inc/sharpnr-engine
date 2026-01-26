package orchestrator_service

import (
	"context"
	orchestrator_dtos "sharpnr-engine/core/orchestrator/dtos"

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
	// UpdateProperty()
}
