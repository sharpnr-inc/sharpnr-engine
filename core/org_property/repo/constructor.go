package org_property_repo

import "gorm.io/gorm"

type organizationPropertyRepo struct {
	db *gorm.DB
}

func NewOrganizationPropertyRepo(db *gorm.DB) *organizationPropertyRepo {
	return &organizationPropertyRepo{
		db: db,
	}
}
