package accounting_service

import (
	"context"
	accountingdtos "sharpnr-engine/sharp_accounting_engine/accounting_engine/dto"

	"github.com/google/uuid"
)

func (s *accountingServices) ListAllPaymentOrder(
	ctx context.Context,
	orgID, emplID, legalEntityID uuid.UUID,
) ([]accountingdtos.PaymentOrderResponseDTO, error) {

	// 1. Fetch raw models from the engine/repo
	// Assuming your engine/repo returns []*payment_engine_models.PaymentOrder
	orders, err := s.paymentEngine.GetOrderByLegalEntity(ctx, orgID, emplID, legalEntityID)
	if err != nil {
		return nil, err
	}

	// 2. Map models to DTOs
	response := make([]accountingdtos.PaymentOrderResponseDTO, len(orders))
	for i, order := range orders {
		response[i] = accountingdtos.PaymentOrderResponseDTO{
			ID:               order.ID,
			OrderNumber:      order.OrderNumber,
			LegalEntityID:    order.LegalEntityID,
			PaymentAccountID: order.PaymentAccountID,
			PayerType:        order.PayerType,
			PayerName:        order.PayerName,
			PayerRef:         order.PayerRef,
			Currency:         order.CurrencyCode,

			// Convert smallest units (Paise) to standard units for the DTO
			AmountPaise: order.AmountPaise,
			PaidPaise:   order.PaidPaise,
			TotalAmount: float64(order.AmountPaise) / 100.0,
			PaidAmount:  float64(order.PaidPaise) / 100.0,

			Status:      order.Status,
			Source:      order.Source,
			SourceRefID: order.SourceRefID,
			SourceMeta:  order.SourceMeta,
			ExpiresAt:   order.ExpiresAt,
			CreatedAt:   order.CreatedAt,
		}
	}

	return response, nil
}
