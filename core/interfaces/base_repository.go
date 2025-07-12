package interfaces

import "github.com/sharpnr-inc/sharpnr-engine/core/models"

type BaseRepository interface {
}

type DepartmentRepository interface {
	Create(dept *models.Department) error
	// Add other department repository methods as needed
}


type TerminalRepository interface {
	Create(terminal *models.Terminal) error
	// Add other terminal repository methods as needed
}