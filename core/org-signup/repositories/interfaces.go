package org_signup_repositories

import "context"

type TempOrgSignUpRepositoryInterface interface {
	// Store temp org with TTL
	Create(
		ctx context.Context,
		signupID string,

		ttlSeconds int,
	) error

	// Fetch temp org
	Get(
		ctx context.Context,
		signupID string,
	) error

	// Existence check (used before OTP resend, etc.)
	Exists(
		ctx context.Context,
		signupID string,
	) (bool, error)

	// Remove after success or expiry
	Delete(
		ctx context.Context,
		signupID string,
	) error
}

type MainOrgSignUpRepositoryInterface interface {
	// Transaction boundary (used during OTP verification → org creation)
	WithTx(
		ctx context.Context,
		fn func(tx MainOrgSignUpRepositoryInterface) error,
	) error

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
		// org *Organization,
	) error

	// ─────────────── Post-Creation Hooks (Optional but Useful) ───────────────

	// Assign default owner / super admin
	CreateOwnerEmployee(
		ctx context.Context,
		// owner *Employee,
	) error

	// Initialize default policies / roles
	CreateDefaultPolicies(
		ctx context.Context,
		orgID string,
	) error

	// Initialize modules/submodules selected during signup
	InitializeOrganizationModules(
		ctx context.Context,
		orgID string,
		modules []string,
	) error
}
