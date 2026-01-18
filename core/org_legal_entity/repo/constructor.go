package org_legal_entity_repo

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	org_legal_entity_models "sharpnr-engine/core/org_legal_entity/models"
)

type orgLegalEntityRepo struct {
	db *gorm.DB
}

// AddAddressInLegalEntity implements OrgLegalEntityRepoInterface.
func (o *orgLegalEntityRepo) AddAddressInLegalEntity(ctx context.Context) error {
	panic("unimplemented")
}

// CreateLegalEntity implements OrgLegalEntityRepoInterface.
func (o *orgLegalEntityRepo) CreateLegalEntity(ctx context.Context, legalEntity org_legal_entity_models.OrganizationLegalEntity) error {
	panic("unimplemented")
}

// GetLegalEntityByID implements OrgLegalEntityRepoInterface.
func (o *orgLegalEntityRepo) GetLegalEntityByID(ctx context.Context, legalEntityID uuid.UUID) (org_legal_entity_models.OrganizationLegalEntity, error) {
	panic("unimplemented")
}

// GetLegalEntityByOrgID implements OrgLegalEntityRepoInterface.
func (o *orgLegalEntityRepo) GetLegalEntityByOrgID(ctx context.Context, orgID uuid.UUID) ([]org_legal_entity_models.OrganizationLegalEntity, error) {
	panic("unimplemented")
}

func NewOrgLegalEntityRepo(db *gorm.DB) *orgLegalEntityRepo {
	return &orgLegalEntityRepo{
		db: db,
	}
}
