package org_signup_services

import (
	dbtx "sharpnr-engine/core/dbTx"
	"sharpnr-engine/sere"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"
)

type orgSignUpService struct {
	txManager    dbtx.TxManager
	sere         *sere.SERE
	mailerEngine *sharp_mailer_engine.SharpMailerEngine
}

func NewOrgSignUpService(
	txManager dbtx.TxManager,
	sere *sere.SERE,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
) *orgSignUpService {
	return &orgSignUpService{
		txManager:    txManager,
		sere:         sere,
		mailerEngine: mailerEngine,
	}
}
