package invoice_engine_services

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	invoice_engine_contracts "sharpnr-engine/service_engines/invoicing_engine/contracts"
	invoice_engine_models "sharpnr-engine/service_engines/invoicing_engine/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *invoiceEngineService) GenerateInvoice(ctx context.Context, request invoice_engine_contracts.CreateInvoiceRequest) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		now := time.Now()

		if err := s.invoicingRepo.CreateInvoice(txCtx, &invoice_engine_models.Invoice{
			ID:                uuid.New(),
			OrganizationID:    request.OrganizationID,
			LegalEntityID:     request.LegalEntityID,
			InvoiceNumber:     "222",
			InvoiceDate:       now,
			DueDate:           &now,
			CustomerID:        request.CustomerID,
			CustomerName:      request.CustomerName,
			CustomerGST:       request.CustomerGST,
			CustomerStateCode: request.CustomerStateCode,
		}); err != nil {
			return err
		}

		return nil
	})

}

func (s *invoiceEngineService) AddInvoiceItems(
	ctx context.Context,
	invoiceitems invoice_engine_contracts.InvoiceItemRequest,
) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)
		itemModel := &invoice_engine_models.InvoiceItem{
			ID:            uuid.New(),
			InvoiceID:     invoiceitems.InvoiceID,
			Description:   invoiceitems.Description,
			Quantity:      invoiceitems.Quantity,
			Rate:          invoiceitems.Rate,
			TaxCategoryID: invoiceitems.TaxCategoryID,
			HSNCode:       invoiceitems.HSNCode,
		}

		if err := s.invoicingRepo.CreateInvoiceItem(txCtx, itemModel); err != nil {
			return err
		}

		return nil
	})
}

// Method Freezes the Invoice and Basically Changing the status
func (s *invoiceEngineService) IssueInvoice(ctx context.Context, invoiceID uuid.UUID) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)
		if err := s.invoicingRepo.UpdateInvoiceStatus(txCtx, invoiceID, "ISSUES"); err != nil {
			return err
		}

		return nil
	})
}
