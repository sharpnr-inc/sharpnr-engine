package accounting_service

import (
	"context"
	"fmt"
	dbtx "sharpnr-engine/core/dbTx"
	accountingdtos "sharpnr-engine/service_engines/accounting_engine/dto"
	accounting_engine_helpers "sharpnr-engine/service_engines/accounting_engine/helpers"
	accounting_models "sharpnr-engine/service_engines/accounting_engine/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *accountingServices) AddNewPurchaseInvoice(
	ctx context.Context,
	orgID, emplID, legalEntityID uuid.UUID,
	request *accountingdtos.PurchaseInvoiceRequestDto,
) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		accountingBook, err := s.accRepo.GetAccountingBookByLegalEntityID(txCtx, legalEntityID)
		if err != nil {
			return fmt.Errorf("there is was error getting accounting book")
		}

		activefiscalYear, err := s.accRepo.GetOpenFiscalYearByBookAndDate(txCtx, accountingBook.ID, time.Now())
		if err != nil {
			return fmt.Errorf("issue occured during the searching active fiscal year")
		}

		searchedSuppliers, err := s.supplierService.ListAllSuppliersOfLegalEntity(txCtx, orgID, emplID, legalEntityID)
		if err != nil {
			return fmt.Errorf("error occured searching the supplier")
		}

		// Step 3: Use the helper to find the specific supplier from the request
		supplier, err := accounting_engine_helpers.FindSupplierByID(searchedSuppliers, request.SupplierID)
		if err != nil {
			return err
		}

		purchaseInvoice := &accounting_models.PurchaseInvoice{
			ID:               uuid.New(),
			AccountingBookID: accountingBook.ID,
			FiscalYearID:     activefiscalYear.ID,
			SupplierID:       supplier.ID,
			PurchaseOrderID:  nil,

			InvoiceNumber: request.InvoiceNumber,
			OurReference:  request.Reference,
			InvoiceDate:   request.InvoiceDate,
			DueDate:       request.DueDate,

			// Calculations
			SubTotal:    0.00,
			TaxAmount:   0.00,
			TotalAmount: 0.00,
			AmountPaid:  0.00,
			AmountDue:   0.00,

			// accounting linkage
			JournalEntryID:   nil,
			PayableAccountID: request.PayableAccountID,

			// Meta
			Narration: request.Narration,
			CreatedBy: emplID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if request.HasPurchaseOrder() {
			purchaseInvoice.PurchaseOrderID = request.PurchaseOrderID
		}

		if err := s.accRepo.CreatePurchaseInvoice(txCtx, purchaseInvoice); err != nil {
			return err
		}

		// Looping over the purchase lines
		for _, pl := range request.Lines {

			purchaseLines := &accounting_models.PurchaseInvoiceLine{
				ID:                 uuid.New(),
				PurchaseInvoiceID:  purchaseInvoice.ID,
				Description:        pl.Description,
				Quantity:           pl.Qty,
				UnitPrice:          pl.UnitPrice,
				DebitAccountID:     pl.DebitAccountID,
				LineAmount:         pl.Qty * pl.UnitPrice,
				TotalAmount:        pl.Total(),
				InventoryAccountID: pl.WareHouseInventoryAccountID,
				InventoryID:        pl.InventoryID,
				InventoryItemID:    pl.InventoryItemID,
				CreatedAt:          time.Now(),
			}

			if err := s.accRepo.CreatePurchaseInvoiceLine(txCtx, purchaseLines); err != nil {
				return err
			}
		}

		return nil
	})
}
