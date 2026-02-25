package org_legal_entity_service

import (
	org_legal_entity_repo "sharpnr-engine/core/org_legal_entity/repo"
)

type orgLegalEntity struct {
	legalEntityRepo org_legal_entity_repo.OrgLegalEntityRepoInterface
}

func NewOrgLegalEntityService(legalEntityRepo org_legal_entity_repo.OrgLegalEntityRepoInterface) *orgLegalEntity {
	return &orgLegalEntity{
		legalEntityRepo: legalEntityRepo,
	}
}
