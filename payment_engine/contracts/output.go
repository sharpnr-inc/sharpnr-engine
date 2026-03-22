package payment_engine_contracts

import "github.com/google/uuid"

type PaymentAccountResponseDto struct {
	Id            uuid.UUID
	LegalEntityID uuid.UUID
	Name          string
	Currency      string
	IsActive      bool
}
