package accounting_service

import (
	"context"
	accountingdtos "sharpnr-engine/modules/accounting/accounting_dtos"

	"github.com/google/uuid"
)

type AccountingServiceInterface interface {
	// Books

	GetAccountingBooksOfOrg(
		ctx context.Context,
		orgID uuid.UUID,
	) (*accountingdtos.GetAccountBooksOfOrg_Response_dto, error)

	GetAccountingGroupOfBook(
		ctx context.Context,
		orgID, bookId uuid.UUID,
	) (*accountingdtos.GetAccountingGroupOfBook_Response_dto, error)
}
