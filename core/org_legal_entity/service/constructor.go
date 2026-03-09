package org_legal_entity_service

import (
	dbtx "sharpnr-engine/core/dbTx"
	org_legal_entity_repo "sharpnr-engine/core/org_legal_entity/repo"
	"sharpnr-engine/sharp_mailer_engine"
)

type orgLegalEntity struct {
	legalEntityRepo org_legal_entity_repo.OrgLegalEntityRepoInterface
	txManager       dbtx.TxManager
	mailerEngine    *sharp_mailer_engine.SharpMailerEngine
}

func NewOrgLegalEntityService(
	legalEntityRepo org_legal_entity_repo.OrgLegalEntityRepoInterface,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
) *orgLegalEntity {
	return &orgLegalEntity{
		legalEntityRepo: legalEntityRepo,
		txManager:       txManager,
		mailerEngine:    mailerEngine,
	}
}
