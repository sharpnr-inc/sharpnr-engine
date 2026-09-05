package sharpCrmEngineService

import (
	"context"
	sharpCrmEngineDtos "sharpnr-engine/sharpCRM_engine/dtos"

	"github.com/google/uuid"
)

type SharpCrmEngineServiceInterface interface {
	// this will start the service and also add the settings table
	StartCrmServiceForOrg(ctx context.Context, orgID, emplID uuid.UUID, req sharpCrmEngineDtos.StartCRMServiceRequestDto) error
}
