package interfaces

import (
	"github.com/google/uuid"
	"github.com/sharpnr-inc/sharpnr-engine/models"
)

// OrganizationRepository defines methods for interacting with the organization storage layer.
type OrganizationRepository interface {
	// Create inserts a new organization record into the database.
	Create(org *models.Organization) error

	// Update updates fields of an existing organization identified by UUID.
	Update(uuid uuid.UUID, updates map[string]interface{}) error

	// FindByUUID returns the organization by its internal UUID.
	FindByUUID(uuid uuid.UUID) (*models.Organization, error)

	// FindBySharpnrOrgID returns the organization by its public-facing ID.
	FindBySharpnrOrgID(id string) (*models.Organization, error)

	ExistsBySharpnrOrgID(id string) (bool, error)
	FindByEmail(email string, org *models.Organization) error
}
