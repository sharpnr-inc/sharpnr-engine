package org_customer_repo

import "gorm.io/gorm"

type orgCustomerRepo struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *orgCustomerRepo {
	return &orgCustomerRepo{
		db: db,
	}
}
