package base

import "github.com/sharpnr-inc/sharpnr-engine/core/models"

// ModuleInitializer defines the interface for module-specific organization initialization.
// Each module (e.g., Hotel, Salon, Retail) that requires specific setup after
// organization registration should implement this interface.
type ModuleInitializer interface {
	InitializeOrganization(org *models.Organization) error
}
