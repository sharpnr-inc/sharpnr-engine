package hospitality_org_factory_repository

import (
	"context"
	hospitality_org_models "sharpnr-engine/core/hospitality-org-factory/models"
)

func (c *hospitalityOrgFactoryRepository) CreateHospitalityOrg(ctx context.Context, hosOrg *hospitality_org_models.OrganizationHospitality) error {
	return c.db.WithContext(ctx).Create(hosOrg).Error
}

func (c *hospitalityOrgFactoryRepository) CreateHospitalityConfig(ctx context.Context, hospitalityConfig *hospitality_org_models.OrganizationHospitalityConfiguration) error {
	return c.db.WithContext(ctx).Create(hospitalityConfig).Error
}
