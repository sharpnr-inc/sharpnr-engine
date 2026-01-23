package orchestrator_service

import (
	"context"
	orchestrator_dtos "sharpnr-engine/core/orchestrator/dtos"

	"github.com/google/uuid"
)

type OrchestratorServiceInterface interface {
	GetOrganizationSetupStatus(ctx context.Context, orgID uuid.UUID) (bool, error)

	// SetupWiz
	AddOrganizationHQAddress(ctx context.Context, orgID uuid.UUID, req *orchestrator_dtos.SetupOrganizationHQAddressReqDTO) error
}
