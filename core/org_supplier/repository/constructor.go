package org_supplier_repository

import "gorm.io/gorm"

type supplierRepository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) *supplierRepository {
	return &supplierRepository{db: db}
}
