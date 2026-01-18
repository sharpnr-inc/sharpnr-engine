package org_legal_entity_repo

import "gorm.io/gorm"

type orgLegalEntityRepo struct {
	db *gorm.DB
}

func NewOrgLegalEntityRepo(db *gorm.DB) *orgLegalEntityRepo {
	return &orgLegalEntityRepo{
		db: db,
	}
}
