package orchestrator_service

import (
	"context"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
	payment_engine_contracts "sharpnr-engine/payment_engine/contracts"

	"github.com/google/uuid"
)

func (c *orchestratorService) AddPaymentAccount(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, req orchestrator_dtos.CreatePaymentAccountRequestDto) error {

	if err := req.ValidateReq(); err != nil {
		return err
	}

	return nil
}

func (c *orchestratorService) GetAllPaymentAccount(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID) ([]payment_engine_contracts.PaymentAccountResponseDto, error) {
	panic("Unimplemented")
}
