package orchestrator_dtos

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

type CreatePaymentAccountRequestDto struct {
	LegalEntityID    uuid.UUID `validate:"required"`
	AccountReference string    `validate:"required"`
	Currency         string    `validate:"required,len=3"`
}

func (r CreatePaymentAccountRequestDto) ValidateReq() error {
	if r.LegalEntityID == uuid.Nil {
		return errors.New("legal_entity_id is required")
	}

	if strings.TrimSpace(r.AccountReference) == "" {
		return errors.New("account_reference is required")
	}

	if strings.TrimSpace(r.Currency) == "" {
		return errors.New("currency is required")
	}

	if len(r.Currency) != 3 {
		return errors.New("currency must be ISO 4217 format (e.g. INR, USD)")
	}

	return nil
}
