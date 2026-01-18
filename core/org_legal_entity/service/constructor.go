package org_legal_entity_service

import (
	org_legal_entity_repo "sharpnr-engine/core/org_legal_entity/repo"
)

type orgLegalEntity struct {
	orgLegalEntitySvc org_legal_entity_repo.OrgLegalEntityRepoInterface
}

func NewOrgLegalEntityService(orgLegalEntitySvc org_legal_entity_repo.OrgLegalEntityRepoInterface) *orgLegalEntity {
	return &orgLegalEntity{
		orgLegalEntitySvc: orgLegalEntitySvc,
	}
}
