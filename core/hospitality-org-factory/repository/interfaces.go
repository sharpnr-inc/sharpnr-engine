package hospitality_org_factory_repository

import (
	"context"
	hospitality_org_models "sharpnr-engine/core/hospitality-org-factory/models"
)

type HospitalityOrgFactoryRepoInterface interface {
	CreateHospitalityOrg(ctx context.Context, hosOrg *hospitality_org_models.OrganizationHospitality) error

	CreateHospitalityConfig(ctx context.Context, hospitalityConfig *hospitality_org_models.OrganizationHospitalityConfiguration) error
}
