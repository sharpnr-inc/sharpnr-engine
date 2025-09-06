package interfaces

import (
	"context"
	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/core/models"

	"github.com/google/uuid"
)

//////////////////////////
// ORGANIZATION REGISTRATION
//////////////////////////

// OrgRegistrationService defines methods for registering an organization
type OrganizationSignUpService interface {
	// InitiateOrgRegistration is the first step to start the registration process
	InitiateOrgRegistration(ctx context.Context, req dto.Org_Signup_REQUEST_DTO) (*dto.Org_Signup_RESPONSE_DTO, string, error)

	// VerifyOrgRegistration verifies the OTP/token sent to the organization
	VerifyOrgRegistration(ctx context.Context, token, code string) (*dto.Org_Verify_Otp_RESPONSE_DTO, bool, error)

	// Optional: CompleteOrgRegistration finalizes the registration
	// CompleteOrgRegistration(token string) (*dto.OrgRegistrationComplete_Response_Dto, error)
}

// OrganizationRepository defines methods for organization signup & lookup.
type OrganizationRepository interface {
	// -----------------------------
	// Redis / Temporary Signup Data
	// -----------------------------

	// EmailOrPhoneExistsInRedis checks if email or phone exists in Redis.
	EmailOrPhoneExistsInRedis(ctx context.Context, email, phone string) (bool, error)

	// CreateOrgInRedisDB saves temporary org registration data in Redis (OTP, etc.).
	CreateOrgInRedisDB(ctx context.Context, tempOrg *models.TempOrgRegistrationData) error

	// GetOrgFromRedisDB retrieves temporary org registration data by email.
	GetOrgFromRedisDB(ctx context.Context, email string) (*models.TempOrgRegistrationData, error)

	// DeleteOrgFromRedisDB removes temporary org registration data after successful OTP verification.
	DeleteOrgFromRedisDB(ctx context.Context, email string) error

	// -----------------------------
	// Main DB / Permanent Org Data
	// -----------------------------

	// EmailOrPhoneExistsInMainDB checks if email or phone exists in the main database.
	EmailOrPhoneExistsInMainDB(ctx context.Context, email, phone string) (bool, error)

	// CreateOrgInMainDB saves the organization permanently in the main database.
	CreateOrgInMainDB(ctx context.Context, org *models.Organization) error

	// FindBySharpnrOrgID checks if an organization exists by SharpnrOrgID.
	FindOrgBySharpnrID(ctx context.Context, id string) (*models.Organization, bool, error)

	// GetOrgUUID_By_SharpnrID retrieves the UUID of the organization by SharpnrOrgID.
	GetOrgUUID_By_SharpnrID(sharpnrID string) (uuid.UUID, error)
}
