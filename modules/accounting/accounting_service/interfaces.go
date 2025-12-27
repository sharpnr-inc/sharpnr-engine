package accounting_service

import (
	"context"
	accountingdtos "sharpnr-engine/modules/accounting/accounting_dtos"
	"time"

	"github.com/google/uuid"
)

type AccountingServiceInterface interface {

	/***********************************************
	 * All the Service Methods for Accounting Books *
	***********************************************/

	GetAccountingBooksOfOrg(
		ctx context.Context,
		orgID uuid.UUID,
	) (*accountingdtos.GetAccountBooksOfOrg_Response_dto, error)

	// -----------------------------------------------------------------------//

	/**********************************************
	 * All the Service methods for Accounting Groups *
	**********************************************/

	//Service method For fetching All the Accounting Groups in Accounting Book
	// It takes OrgId and BookID of Account Groups we want
	GetAccountingGroupOfBook(
		ctx context.Context,
		orgID, bookId uuid.UUID,
	) (*accountingdtos.GetAccountingGroupOfBook_Response_dto, error)

	// -----------------------------------------------------------------------//

	/************************************
	 * All Service Methods for Accounts *
	************************************/

	CreateAccount(
		ctx context.Context,
		orgID uuid.UUID,
		req accountingdtos.CreateAccountRequest,
	) error

	// -----------------------------------------------------------------------//

	/***********************************************
	 * All the Service methods for Account Mapping *
	 ***********************************************/

	GetAccountMapping(
		ctx context.Context,
		orgID uuid.UUID,
	) ([]accountingdtos.AccountMappingDTO, error)

	CreateAccountMapping(
		ctx context.Context,
		orgId uuid.UUID,
		req accountingdtos.CreateAccountMappingRequest_dto,
	) error

	// -----------------------------------------------------------------------//

	/*******************************
	 * Transactions related method *
	*******************************/
	GetTransactions(
		ctx context.Context,
		orgID uuid.UUID,
		bookID uuid.UUID,
		from time.Time,
		to time.Time,
		limit int,
		offset int,
	) ([]accountingdtos.JournalEntryListDTO, int64, error)
}
