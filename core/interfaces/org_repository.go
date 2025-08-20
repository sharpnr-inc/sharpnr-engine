package interfaces

import "sharpnr-engine/core/models"

// ========================================
// DATA FLOW: Redis + PostgreSQL
// ========================================

/*
REGISTRATION FLOW:
1. User submits registration → Redis (temporary, 24h expiry)
2. User verifies email/OTP → Move from Redis → PostgreSQL (permanent)

DATA DISTRIBUTION:
┌─────────────────┐    ┌──────────────────────┐
│     REDIS       │    │    POSTGRESQL        │
│   (Temporary)   │    │    (Permanent)       │
├─────────────────┤    ├──────────────────────┤
│ • Pending       │    │ • Verified users     │
│   registrations │    │ • Organizations      │
│ • Session data  │    │ • User profiles      │
│ • OTP codes     │    │ • Auth credentials   │
│ • Cache         │    │ • Application data   │
│ • Rate limiting │    │ • Audit logs         │
└─────────────────┘    └──────────────────────┘
      TTL: 24h              Permanent Storage
*/

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

// CoreOrganizationRepository defines the interface for interacting with the core organization data store.
// It provides a set of methods for creating, retrieving, and managing organizations,
// forming an abstraction layer between the business logic and the database.
type CoreInitialOrganizationRepository interface {
	// So First of All we need to Find the If email/Phone or either one of Them Already Exist in the main DB
	// In both DB
	FindOrgByEmailInMainDB()

	//
	FindOrgByEmailInRedisDB()
	// Then We need to
	CreateTempOrg()
	//
	MoveTempOrgToMainDB()
}
