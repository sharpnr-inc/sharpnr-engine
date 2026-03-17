package org_customer_repo

import (
	"context"
	org_customer_models "sharpnr-engine/core/customer/models"
	dbtx "sharpnr-engine/core/dbTx"

	"github.com/google/uuid"
)

func (r *orgCustomerRepo) IsMobileTaken(ctx context.Context, orgID uuid.UUID, mobile string) (bool, error) {
	var count int64
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Check if any record exists with this mobile AND orgID
	err := db.WithContext(ctx).
		Model(&org_customer_models.Customer{}).
		Where("organization_id = ? AND phone = ?", orgID, mobile).
		Count(&count).
		Error

	return count > 0, err
}
