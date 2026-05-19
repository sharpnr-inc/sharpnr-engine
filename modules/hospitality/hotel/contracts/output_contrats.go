package hotel_contracts

import (
	"time"

	"github.com/google/uuid"
)

type Hotels struct {
	ID             uuid.UUID
	Name           string
	OrganizationID uuid.UUID
	PropertyID     uuid.UUID
	LegalEntityID  *uuid.UUID

	IsActive      bool
	IsConfigured  bool
	LastUpdatedAt time.Time
}
