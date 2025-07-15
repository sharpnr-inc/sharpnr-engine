package interfaces

import "github.com/sharpnr-inc/sharpnr-engine/core/models"

// CoreOrganizationRepository defines the interface for interacting with the core organization data store.
// It provides a set of methods for creating, retrieving, and managing organizations,
// forming an abstraction layer between the business logic and the database.
type CoreOrganizationRepository interface {
	// Create persists a new organization record in the database.
	// It takes an Organization model pointer as input and is responsible for setting it up.
	// Returns an error if the creation fails.
	Create(org *models.Organization) error

	// Update modifies an existing organization's details in the database.
	// It takes an Organization model pointer with the updated data.
	// Returns an error if the update operation fails.
	Update(org *models.Organization) error

	// FindByUUID retrieves a single organization from the database using its unique UUID.
	// It returns a pointer to the Organization model and an error if not found or if a database error occurs.
	FindByUUID(uuid string) (*models.Organization, error)

	// FindBySharpnrOrgID retrieves a single organization by its custom Sharpnr Org ID.
	// It returns a pointer to the Organization model and an error if not found or if a database error occurs.
	FindBySharpnrOrgID(id string) (bool, error)

	// FindByMail retrieves a single organization from the database using the owner's email address.
	// This is commonly used during login or recovery processes.
	// It returns a pointer to the Organization model and an error if not found or if a database error occurs.
	FindByMail(email string) (*models.Organization, error)
}
