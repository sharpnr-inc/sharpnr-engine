package org_signup_repositories

import (
	"context"
	"errors"
	"sharpnr-engine/core/models"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// WithTx executes the provided function within a database transaction.
//
// This method establishes a transactional boundary for the organization
// signup flow. It is typically used after OTP verification to ensure that
// all database operations involved in creating an organization and its
// related entities are executed atomically.
//
// All repository calls made through the `tx` argument participate in the
// same database transaction. If the function returns an error, the
// transaction is rolled back. If the function returns nil, the transaction
// is committed.
//
// The service layer controls the business flow, while this method guarantees
// transactional consistency without exposing the underlying *gorm.DB.
//
// Parameters:
//   - ctx: Context used to control request lifetime, cancellation, and deadlines.
//   - fn: A function that receives a transactional instance of
//     MainOrgSignUpRepositoryInterface.
//
// Returns:
//   - error: Any error returned by the function or raised during transaction
//     execution.
//
// Example:
//
//	err := repo.WithTx(ctx, func(tx MainOrgSignUpRepositoryInterface) error {
//	    if err := tx.CreateOrganization(ctx); err != nil {
//	        return err
//	    }
//
//	    if err := tx.CreateOwnerEmployee(ctx); err != nil {
//	        return err
//	    }
//
//	    if err := tx.CreateDefaultPolicies(ctx, orgID); err != nil {
//	        return err
//	    }
//
//	    return nil
//	})
//
//	if err != nil {
//	    // Transaction rolled back due to error
//	    return err
//	}
func (d *orgMainSignUpRepo) WithTx(
	ctx context.Context,
	fn func(tx MainOrgSignUpRepositoryInterface) error,
) error {

	if fn == nil {
		return errors.New("transaction function cannot be nil")
	}

	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txRepo := &orgMainSignUpRepo{
			db: tx,
		}
		return fn(txRepo)
	})
}

func (r *orgMainSignUpRepo) ExistsByEmail(
	ctx context.Context,
	email string,
) (bool, error) {

	if email == "" {
		return false, errors.New("email is required")
	}

	normalizedEmail := strings.ToLower(strings.TrimSpace(email))
	return r.exists(ctx, "email", normalizedEmail)
}

func (r *orgMainSignUpRepo) ExistsByPhone(
	ctx context.Context,
	phone string,
) (bool, error) {

	if phone == "" {
		return false, errors.New("phone is required")
	}

	// Phone should already be normalized (E.164) at service layer
	return r.exists(ctx, "phone", phone)
}

func (r *orgMainSignUpRepo) ExistsBySlug(
	ctx context.Context,
	slug string,
) (bool, error) {

	if slug == "" {
		return false, errors.New("slug is required")
	}

	normalizedSlug := strings.ToLower(strings.TrimSpace(slug))
	return r.exists(ctx, "slug", normalizedSlug)
}

func (r *orgMainSignUpRepo) exists(
	ctx context.Context,
	column string,
	value string,
) (bool, error) {

	var count int64

	err := r.db.
		WithContext(ctx).
		Model(&models.Organization{}).
		Where(column+" = ?", value).
		Limit(1).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *orgMainSignUpRepo) CreateOrganization(
	ctx context.Context,
	org *models.Organization,
) error {

	if org == nil {
		return errors.New("organization is nil")
	}

	if org.ID == uuid.Nil {
		return errors.New("organization ID is required")
	}

	return r.db.
		WithContext(ctx).
		Create(org).
		Error
}
