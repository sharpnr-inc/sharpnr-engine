package accounting_service

import (
	"context"
	"fmt"
	dbtx "sharpnr-engine/pkg/dbTx"
	accountingconstants "sharpnr-engine/sharp_accounting_engine/accounting_engine/constants"
	accountingdtos "sharpnr-engine/sharp_accounting_engine/accounting_engine/dto"
	accounting_engine_events "sharpnr-engine/sharp_accounting_engine/accounting_engine/events"
	accounting_engine_helpers "sharpnr-engine/sharp_accounting_engine/accounting_engine/helpers"
	accounting_models "sharpnr-engine/sharp_accounting_engine/accounting_engine/models"
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

		// ── 1. Resolve accounting book ────────────────────────────────────────

		accountingBook, err := s.accRepo.GetAccountingBookByLegalEntityID(txCtx, legalEntityID)
		if err != nil {
			return fmt.Errorf("error getting accounting book: %w", err)
		}

		// ── 2. Resolve active fiscal year ─────────────────────────────────────

		activeFiscalYear, err := s.accRepo.GetOpenFiscalYearByBookAndDate(txCtx, accountingBook.ID, time.Now())
		if err != nil {
			return fmt.Errorf("error finding active fiscal year: %w", err)
		}
		if activeFiscalYear == nil {
			return fmt.Errorf("no open fiscal year found for accounting book")
		}

		// ── 3. Resolve supplier ───────────────────────────────────────────────

		searchedSuppliers, err := s.supplierService.ListAllSuppliersOfLegalEntity(txCtx, orgID, emplID, legalEntityID)
		if err != nil {
			return fmt.Errorf("error searching suppliers: %w", err)
		}

		supplier, err := accounting_engine_helpers.FindSupplierByID(searchedSuppliers, request.SupplierID)
		if err != nil {
			return err
		}

		// ── 4. Duplicate invoice number check (unique per supplier) ───────────

		existing, _ := s.accRepo.GetPurchaseInvoiceBySupplierAndInvoiceNumber(
			txCtx, supplier.ID, request.InvoiceNumber,
		)
		if existing != nil {
			return fmt.Errorf(
				"invoice number %q already exists for supplier %s — duplicates are not allowed",
				request.InvoiceNumber,
				supplier.ID,
			)
		}

		// ── 5. Compute totals from request lines ──────────────────────────────

		var subTotal, taxTotal float64
		for _, pl := range request.Lines {
			subTotal += pl.Subtotal()
			taxTotal += pl.TaxAmount()
		}
		grandTotal := subTotal + taxTotal

		// ── 6. Create PurchaseInvoice header (status = DRAFT) ─────────────────

		purchaseInvoice := &accounting_models.PurchaseInvoice{
			ID:               uuid.New(),
			AccountingBookID: accountingBook.ID,
			FiscalYearID:     activeFiscalYear.ID,
			SupplierID:       supplier.ID,
			PurchaseOrderID:  nil,

			InvoiceNumber: request.InvoiceNumber,
			OurReference:  request.Reference,
			InvoiceDate:   request.InvoiceDate,
			DueDate:       request.DueDate,

			// ✅ Status set on the struct — not separately after the fact
			Status:        accountingconstants.PurchaseInvoiceStatusDraft,
			PaymentStatus: accountingconstants.PurchaseInvoicePaymentStatusUnpaid,

			PayableAccountID: request.PayableAccountID,

			// ✅ grandTotal used for TotalAmount and AmountDue
			SubTotal:    subTotal,
			TaxAmount:   taxTotal,
			TotalAmount: grandTotal,
			AmountPaid:  0,
			AmountDue:   grandTotal,

			JournalEntryID: nil, // linked after journal entry is created
			Narration:      request.Narration,
			CreatedBy:      emplID,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		if request.HasPurchaseOrder() {
			purchaseInvoice.PurchaseOrderID = request.PurchaseOrderID
		}

		if err := s.accRepo.CreatePurchaseInvoice(txCtx, purchaseInvoice); err != nil {
			return fmt.Errorf("error creating purchase invoice: %w", err)
		}

		// ── 7. Build invoice lines ────────────────────────────────────────────

		// ✅ Lines are appended to the slice — not discarded
		var invoiceLines []*accounting_models.PurchaseInvoiceLine

		for _, pl := range request.Lines {
			line := &accounting_models.PurchaseInvoiceLine{
				ID:                 uuid.New(),
				PurchaseInvoiceID:  purchaseInvoice.ID,
				Description:        pl.Description,
				Quantity:           pl.Qty,
				UnitPrice:          pl.UnitPrice,
				DebitAccountID:     pl.DebitAccountID,
				TaxCategoryID:      pl.TaxCategoryID, // ✅ was missing
				LineAmount:         pl.Subtotal(),
				TaxAmount:          pl.TaxAmount(),
				TotalAmount:        pl.TotalAmount(),
				InventoryAccountID: pl.WareHouseInventoryAccountID,
				InventoryID:        pl.InventoryID,
				InventoryItemID:    pl.InventoryItemID,
				CreatedAt:          time.Now(),
			}
			invoiceLines = append(invoiceLines, line)
		}

		// ✅ Batch insert — one DB round-trip instead of N
		if err := s.accRepo.CreatePurchaseInvoiceLines(txCtx, invoiceLines); err != nil {
			return fmt.Errorf("error creating invoice lines: %w", err)
		}

		// ── 8. Validate lines before journaling (Option A — hard reject) ──────
		//
		// Every line MUST have DebitAccountID set and TotalAmount > 0.
		// If any line fails the whole tx rolls back — nothing persists.

		if err := accounting_engine_helpers.ValidatePurchaseInvoiceLines(invoiceLines); err != nil {
			return err
		}

		// ── 9. Build journal entry lines ──────────────────────────────────────
		//
		// For each invoice line:
		//   DR  line.DebitAccountID    line.TotalAmount   (COGS / Inventory / Expense)
		// One combined credit:
		//   CR  PayableAccountID       grandTotal         (Accounts Payable)

		var journalLines []accountingdtos.JournalLineExternalRequestDto

		for _, line := range invoiceLines {
			journalLines = append(journalLines, accountingdtos.JournalLineExternalRequestDto{
				AccountID: line.DebitAccountID,
				Debit:     line.TotalAmount,
				Credit:    0,
				PartyID:   &supplier.ID,
			})
		}

		// single credit leg — Accounts Payable
		journalLines = append(journalLines, accountingdtos.JournalLineExternalRequestDto{
			AccountID: purchaseInvoice.PayableAccountID,
			Debit:     0,
			Credit:    grandTotal,
			PartyID:   &supplier.ID,
		})

		// ── 10. Post journal entry (same tx via txCtx) ────────────────────────

		journalEntryRequest := accountingdtos.JournalEntryExternalRequestDto{
			AccountingBookID: accountingBook.ID,
			Reference:        purchaseInvoice.OurReference,
			Narration: fmt.Sprintf(
				"Purchase Invoice %s from %s",
				purchaseInvoice.InvoiceNumber,
				supplier.Name,
			),
			EntryType: "NORNAL",
			Lines:     journalLines,
		}

		// ✅ journalEntryID captured — not discarded with _
		journalEntryID, err := s.CreateJournalEntry(txCtx, orgID, emplID, legalEntityID, journalEntryRequest)
		if err != nil {
			return fmt.Errorf("error creating journal entry: %w", err)
		}

		// ── 11. Link journal entry + mark invoice as POSTED ───────────────────

		// ✅ SetPurchaseInvoiceJournalEntry called — was completely missing before
		if err := s.accRepo.SetPurchaseInvoiceJournalEntry(txCtx, purchaseInvoice.ID, journalEntryID); err != nil {
			return fmt.Errorf("error linking journal entry to invoice: %w", err)
		}

		// ✅ Status update happens AFTER journal entry succeeds — not before
		if err := s.accRepo.UpdatePurchaseInvoiceStatus(
			txCtx,
			purchaseInvoice.ID,
			"POSTED",
		); err != nil {
			return fmt.Errorf("error updating invoice status to POSTED: %w", err)
		}

		// ── 12. Filter inventory lines + publish event ────────────────────────
		//
		// Only lines with InventoryAccountID + InventoryID set are included.
		// The inventory subscriber handles stock qty update + price history.
		// Published AFTER all DB ops succeed — if publish fails, tx rolls back.

		itemsToTransferToInventory, err := accounting_engine_helpers.FliterItemToTransferedToInventory(request.Lines)
		if err != nil {
			return err
		}

		eventPayload := accounting_engine_events.PurchaceInvoiceItemTransterToInventoryEventPayload{
			SupplierID:     supplier.ID,
			OrganizationID: orgID,
			EmployeeID:     emplID,
			Items:          itemsToTransferToInventory,
		}

		if err := s.seroEngine.Publish(
			txCtx,
			accounting_engine_events.PurchaseInvoiceCreated,
			eventPayload,
		); err != nil {
			return fmt.Errorf("error publishing PurchaseInvoiceCreated event: %w", err)
		}

		return nil
	})
}
