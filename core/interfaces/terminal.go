package interfaces

import (
	"github.com/google/uuid"
)

type TerminalService interface {
	InitializeDefaultAccountingTerminals(orgUUID uuid.UUID, deptID uuid.UUID) error
	// Add other terminal-related methods as needed
}
