# Core Package - Interfaces Documentation

## Folder: `core/package`

### Package: Interfaces

This package contains interfaces for Core functionality, providing abstraction layers for organization management.

---

## Files Overview

1. **[org_service.go](org_service.go)** - Service layer interface for organization operations
2. **[org_repository.go](org_repository.go)** - Repository layer interface for organization data access

---

## CoreOrganisationService Interface

**File:** `org_service.go`

**Package:** `interfaces`

### Description
Service layer interface that defines the contract for organization-related business operations.

### Dependencies
```go
import dto "sharpnr-engine/core/dtos"
```

### Interface Definition
```go
type CoreOrganisationService interface {
	// RegisterOrganization handles the registration of a new organization
	// Parameters:
	//   - req: OrganizationRegistrationRequestDTO containing registration details
	// Returns:
	//   - *dto.OrganizationRegistrationResponseDTO: Response data upon successful registration
	//   - string: Additional message or identifier
	//   - error: Error information if registration fails
	RegisterOrganization(req dto.OrganizationRegistrationRequestDTO) (*dto.OrganizationRegistrationResponseDTO, string, error)
}
```

### Methods

#### `RegisterOrganization`
- **Purpose**: Registers a new organization in the system
- **Input**: `dto.OrganizationRegistrationRequestDTO` - Contains organization registration details
- **Output**: 
  - `*dto.OrganizationRegistrationResponseDTO` - Registration response data
  - `string` - Additional response message or identifier
  - `error` - Error if registration fails

---

## CoreOrganizationRepository Interface

**File:** `org_repository.go`

**Package:** `interfaces`

### Description
Repository layer interface that defines the contract for organization data persistence operations. This interface provides an abstraction layer between the business logic and the database, ensuring loose coupling and testability.

### Dependencies
```go
import "sharpnr-engine/core/models"
```

### Interface Definition
```go
type CoreOrganizationRepository interface {
	Create(org *models.Organization) error
	Update(org *models.Organization) error
	FindByUUID(uuid string) (*models.Organization, error)
	FindBySharpnrOrgID(id string) (bool, error)
	FindByMail(email string) (*models.Organization, error)
}
```

### Methods

#### `Create`
- **Purpose**: Persists a new organization record in the database
- **Input**: `*models.Organization` - Organization model pointer with data to be saved
- **Output**: `error` - Returns error if creation fails
- **Behavior**: Responsible for setting up the organization record in the database

#### `Update`
- **Purpose**: Modifies an existing organization's details in the database
- **Input**: `*models.Organization` - Organization model pointer with updated data
- **Output**: `error` - Returns error if update operation fails
- **Behavior**: Updates existing organization record with new information

#### `FindByUUID`
- **Purpose**: Retrieves a single organization using its unique UUID
- **Input**: `string` - UUID of the organization to find
- **Output**: 
  - `*models.Organization` - Pointer to the found organization
  - `error` - Error if not found or database error occurs
- **Use Case**: Primary key lookup for organization retrieval

#### `FindBySharpnrOrgID`
- **Purpose**: Retrieves organization existence status by custom Sharpnr Org ID
- **Input**: `string` - Custom Sharpnr Organization ID
- **Output**: 
  - `bool` - Indicates if organization exists
  - `error` - Error if database error occurs
- **Use Case**: Checking organization existence by custom identifier

#### `FindByMail`
- **Purpose**: Retrieves organization using the owner's email address
- **Input**: `string` - Owner's email address
- **Output**: 
  - `*models.Organization` - Pointer to the found organization
  - `error` - Error if not found or database error occurs
- **Use Case**: Commonly used during login or recovery processes

---

## Architecture Notes

### Service Layer (`CoreOrganisationService`)
- Handles business logic and orchestration
- Validates input data and business rules
- Coordinates between different components
- Returns structured DTOs for API responses

### Repository Layer (`CoreOrganizationRepository`)
- Provides data access abstraction
- Handles database operations
- Ensures data consistency and integrity
- Works with domain models rather than DTOs

### Separation of Concerns
- **Service**: Business logic and workflow orchestration
- **Repository**: Data persistence and retrieval
- **DTOs**: Data transfer between layers
- **Models**: Domain entities and business objects

---

## Usage Example

```go
// Service implementation would use repository
type organizationService struct {
    repo interfaces.CoreOrganizationRepository
}

func (s *organizationService) RegisterOrganization(req dto.OrganizationRegistrationRequestDTO) (*dto.OrganizationRegistrationResponseDTO, string, error) {
    // Convert DTO to model
    org := &models.Organization{
        // Map fields from req
    }
    
    // Use repository to persist
    if err := s.repo.Create(org); err != nil {
        return nil, "", err
    }
    
    // Return response DTO
    return &dto.OrganizationRegistrationResponseDTO{
        // Map response fields
    }, "Organization registered successfully", nil
}
```