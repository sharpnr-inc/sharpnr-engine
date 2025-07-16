package interfaces

import (
	"github.com/google/uuid"
	"github.com/sharpnr-inc/sharpnr-engine/core/models"
)

// This File Contains both Interfaces for the Teminal Like Service and Repository

// This is the Interface for the Terminal Service
// User have Implemne this ....
type TerminalService interface {
	InitializeDefaultAccountingTerminals(orgUUID uuid.UUID, deptID uuid.UUID) error
	// Add other terminal-related methods as needed
}

//----------------------------------------------------------------------------------------------------------------------------------------------------//

// This is TerminalRepository Service
// Which is Used For the Creating Terminal
type TerminalRepository interface {
	Create(terminal *models.Terminal) error
	// Add other terminal repository methods as needed
}
