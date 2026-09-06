package folioEngineService

import (
	"context"
	folioEngineDtos "sharpnr-engine/folio_engine/dtos"

	"github.com/google/uuid"
)

type FolioEngineServiceInterface interface {
	// StartFolioServiceForOrg provisions the org's one Folio account together
	// with its settings row.
	StartFolioServiceForOrg(ctx context.Context, orgID, emplID uuid.UUID, req folioEngineDtos.StartFolioServiceRequestDto) error

	// GetFolioAccountStatus is the read every other Folio flow is gated on: has
	// this org started Folio, and is that account switched on.
	GetFolioAccountStatus(ctx context.Context, orgID, emplID uuid.UUID) (folioEngineDtos.FolioAccountStatusResponse, error)
}
