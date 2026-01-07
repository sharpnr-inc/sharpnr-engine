package org_signup_repositories

import (
	"context"
	"sharpnr-engine/core/models"

	"github.com/google/uuid"
)

type TempOrgSignUpRepositoryInterface interface {
	// Store temp org with TTL
	Create(
		ctx context.Context,
		tempOrg *models.OrganizationSignUpStaging,
	) error

	// Fetch temp org
	Get(
		ctx context.Context,
		signupID string,
	) (*models.OrganizationSignUpStaging, error)

	// Existence check (used before OTP resend, etc.)
	Exists(
		ctx context.Context,
		signupID string,
	) (bool, models.OrganizationSignUpStaging, error)
	// Remove after success or expiry
	Delete(
		ctx context.Context,
		signupID string,
	) error

	GetByEmailOrPhone(
		ctx context.Context,
		email string,
		phone string,
	) (bool, error)
}

type MainOrgSignUpRepositoryInterface interface {
	// ─────────────── Existence Checks ───────────────

	// Check if an organization already exists by email
	ExistsByEmail(
		ctx context.Context,
		email string,
	) (bool, error)

	// Check if an organization already exists by phone
	ExistsByPhone(
		ctx context.Context,
		phone string,
	) (bool, error)

	// Optional: check by slug / subdomain
	ExistsBySlug(
		ctx context.Context,
		slug string,
	) (bool, error)

	// ─────────────── Creation ───────────────

	// Create the organization (final step after OTP verification)
	CreateOrganization(
		ctx context.Context,
		org *models.Organization,
	) error

	FindOrgBySharpnrID(
		ctx context.Context,
		Id string,
	) (models.OrganizationSignUpStaging, bool, error)

	//_-__-_-__-_-__-_-_-_-_-__-_-___

	CreateOrgSubscription(
		ctx context.Context,
		orgSub models.OrganizationSubscription,
	) error

	GetOrgSubscriptionPlan(
		ctx context.Context,
		orgId uuid.UUID,
	) (*models.OrganizationSubscription, error)

	AddOrgVerificationDetails(
		ctx context.Context,
		orgVerification models.OrganizationVerification,
	) error

	AddOrgOfficeAddress(
		ctx context.Context,
		Addr models.OrganizationOffice,
	) error
}
