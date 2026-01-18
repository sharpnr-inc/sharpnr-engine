package app

import (
	org_legal_entity_repo "sharpnr-engine/core/org_legal_entity/repo"
	org_legal_entity_service "sharpnr-engine/core/org_legal_entity/service"

	"gorm.io/gorm"
)

func bootstrapOrgLegalEntity(db *gorm.DB) org_legal_entity_service.OrgLegalEntityServiceInterface {

	orgLegalEntityService := org_legal_entity_repo.NewOrgLegalEntityRepo(db)
	return org_legal_entity_service.NewOrgLegalEntityService(orgLegalEntityService)
}
