package hospitalitymodels

import (
	"sharpnr-engine/core/models"

	"github.com/google/uuid"
)

type Kitchen struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	Organization   models.Organization
}
