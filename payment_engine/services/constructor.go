package payment_engine_service

import (
	"encoding/json"
	"fmt"
	dbtx "sharpnr-engine/core/dbTx"
	payment_gateway "sharpnr-engine/payment_engine/gateway"
	payment_engine_repo "sharpnr-engine/payment_engine/repo"
	"sharpnr-engine/sero"
	"sharpnr-engine/sharp_mailer_engine"
)

type paymentEngine struct {
	paymentRepo     payment_engine_repo.PaymentEngineRepoInterface
	gatewayRegistry *payment_gateway.ProviderRegistry // ← add this
	mailerEngine    *sharp_mailer_engine.SharpMailerEngine
	txManager       dbtx.TxManager
	seroEngine      *sero.SEROEngine
}

func NewPaymentService(
	paymentRepo payment_engine_repo.PaymentEngineRepoInterface,
	gatewayRegistry *payment_gateway.ProviderRegistry,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	txManager dbtx.TxManager,
	seroEngine *sero.SEROEngine,
) *paymentEngine {
	return &paymentEngine{
		paymentRepo:     paymentRepo,
		gatewayRegistry: gatewayRegistry,
		mailerEngine:    mailerEngine,
		txManager:       txManager,
		seroEngine:      seroEngine,
	}
}

func encryptConfig(config map[string]string) (string, error) {
	// TODO: replace with AES/KMS
	raw, err := json.Marshal(config)
	if err != nil {
		return "", fmt.Errorf("encryptConfig: %w", err)
	}
	return string(raw), nil
}

func decryptConfig(encrypted string) (map[string]string, error) {
	// TODO: replace with AES/KMS
	var result map[string]string
	if err := json.Unmarshal([]byte(encrypted), &result); err != nil {
		return nil, fmt.Errorf("decryptConfig: %w", err)
	}
	return result, nil
}
