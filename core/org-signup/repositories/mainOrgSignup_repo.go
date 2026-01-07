package org_signup_repositories

import (
	"context"
	"errors"
	"sharpnr-engine/core/models"
	"strings"

	"github.com/google/uuid"
)

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

// AddOrgOfficeAddress implements MainOrgSignUpRepositoryInterface.
func (d *orgMainSignUpRepo) AddOrgOfficeAddress(ctx context.Context, Addr models.OrganizationOffice) error {
	return d.db.WithContext(ctx).Create(Addr).Error
}

// AddOrgVerificationDetails implements MainOrgSignUpRepositoryInterface.
func (d *orgMainSignUpRepo) AddOrgVerificationDetails(ctx context.Context, orgVerification models.OrganizationVerification) error {
	return d.db.WithContext(ctx).Create(orgVerification).Error
}

// CreateOrgSubscription implements MainOrgSignUpRepositoryInterface.
func (d *orgMainSignUpRepo) CreateOrgSubscription(ctx context.Context, orgSub models.OrganizationSubscription) error {
	return d.db.WithContext(ctx).Create(orgSub).Error
}

// FindOrgBySharpnrID implements MainOrgSignUpRepositoryInterface.
func (d *orgMainSignUpRepo) FindOrgBySharpnrID(ctx context.Context, Id string) (models.OrganizationSignUpStaging, bool, error) {
	panic("unimplemented")
}

// GetOrgSubscriptionPlan implements MainOrgSignUpRepositoryInterface.
func (d *orgMainSignUpRepo) GetOrgSubscriptionPlan(ctx context.Context, orgId uuid.UUID) (*models.OrganizationSubscription, error) {
	panic("unimplemented")
}
