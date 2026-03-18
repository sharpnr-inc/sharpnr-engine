package payment_engine_service

import (
	dbtx "sharpnr-engine/core/dbTx"
	payment_engine_repo "sharpnr-engine/payment_engine/repo"
	"sharpnr-engine/sero"
	"sharpnr-engine/sharp_mailer_engine"
)

type paymentEngine struct {
	paymentRepo  payment_engine_repo.PaymentEngineRepoInterface
	mailerEngine *sharp_mailer_engine.SharpMailerEngine
	txManager    dbtx.TxManager
	seroEngine   *sero.SEROEngine
}

func NewPaymentService(
	paymentRepo payment_engine_repo.PaymentEngineRepoInterface,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	txManager dbtx.TxManager,
	seroEngine *sero.SEROEngine,
) *paymentEngine {
	return &paymentEngine{
		paymentRepo:  paymentRepo,
		mailerEngine: mailerEngine,
		txManager:    txManager,
		seroEngine:   seroEngine,
	}
}
