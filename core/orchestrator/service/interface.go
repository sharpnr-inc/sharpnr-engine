package orchestrator_service

import (
	"context"

	"github.com/google/uuid"
)

type OrchestratorServiceInterface interface {
	GetOrganizationSetupStatus(ctx context.Context, orgID uuid.UUID) (bool, error)
}
