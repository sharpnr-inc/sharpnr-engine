package org_customer_repo

import (
	"context"
	"errors"
	org_customer_models "sharpnr-engine/core/customer/models"
	dbtx "sharpnr-engine/core/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *orgCustomerRepo) CreateCustomer(ctx context.Context, customer *org_customer_models.Customer) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Ensure we use the request context for tracing and timeouts
	db = db.WithContext(ctx)

	// GORM's Create method handles the INSERT and
	// populates hooks (BeforeCreate, etc.)
	if err := db.Create(customer).Error; err != nil {
		return err
	}

	return nil

}

func (r *orgCustomerRepo) CreateCustomerMeta(
	ctx context.Context,
	customerMeta *org_customer_models.CustomerOrgMeta,
) error {
	// 1. Extract transaction from context or fallback to the repository's base DB
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// 2. Ensure we use the request context for tracing and timeouts
	// 3. Use .Create() to persist the metadata model
	if err := db.WithContext(ctx).Create(customerMeta).Error; err != nil {
		// Log or wrap error as needed for your organization's standards
		return err
	}

	return nil
}

func (r *orgCustomerRepo) GetCustomersByOrgID(
	ctx context.Context,
	orgID uuid.UUID,
) ([]org_customer_models.Customer, error) {
	var customers []org_customer_models.Customer

	// Extract transaction or use default DB
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Use WithContext for tracing/cancellation
	// .Where() handles the filtering by OrganizationID
	// .Find() populates the slice
	err := db.WithContext(ctx).
		Where("organization_id = ?", orgID).
		Find(&customers).
		Error

	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *orgCustomerRepo) GetCustomerByEmail(
	ctx context.Context,
	customerEmail string,
) (*org_customer_models.Customer, error) {
	var customer org_customer_models.Customer

	// Extract transaction or use default DB
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// .Where() filters by the email column
	// .First() adds a "LIMIT 1" and orders by primary key
	err := db.WithContext(ctx).
		Where("email = ?", customerEmail).
		First(&customer).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Or return a custom "NotFoundError"
		}
		return nil, err
	}

	return &customer, nil
}

func (r *orgCustomerRepo) SearchCustomers(
	ctx context.Context,
	orgID uuid.UUID,
	query string,
) ([]org_customer_models.Customer, error) {
	var customers []org_customer_models.Customer

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Prepare the search term for partial matching
	searchTerm := "%" + query + "%"

	// We wrap the OR conditions in a sub-query group to maintain the OrgID filter
	err := db.WithContext(ctx).
		Where("organization_id = ?", orgID).
		Where(
			db.Where("name ILIKE ?", searchTerm).
				Or("email ILIKE ?", searchTerm).
				Or("phone ILIKE ?", searchTerm),
		).
		Limit(10). // Limit results for better performance in debounced search
		Find(&customers).
		Error

	return customers, err
}

func (r *orgCustomerRepo) UpdateCustomer(
	ctx context.Context,
	orgID, customerID uuid.UUID,
	updates interface{},
) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// We use .Where to ensure the customer belongs to the organization
	// This prevents cross-tenant updates
	result := db.WithContext(ctx).
		Model(&org_customer_models.Customer{}).
		Where("id = ? AND organization_id = ?", customerID, orgID).
		Updates(updates)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("customer not found or no changes made")
	}

	return nil
}
