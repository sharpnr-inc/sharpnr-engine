package accounting_service

import (
	"context"
	accountingdtos "sharpnr-engine/modules/accounting/accounting_dtos"

	"github.com/google/uuid"
)

func (s *accountingServices) GetAccountMapping(
	ctx context.Context,
	orgID uuid.UUID,
	req accountingdtos.CreateAccountMapping_Request_dto,
) error {
	panic("unimplemented")
}
