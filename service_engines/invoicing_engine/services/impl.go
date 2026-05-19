package invoice_engine_services

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	dbtx "sharpnr-engine/pkg/dbTx"
	invoice_engine_contracts "sharpnr-engine/service_engines/invoicing_engine/contracts"
	invoice_engine_models "sharpnr-engine/service_engines/invoicing_engine/models"
	tax_engine_contracts "sharpnr-engine/service_engines/tax_engine/contracts"
)

// ─────────────────────────────────────────────────────────────────────────────
// GenerateInvoice
// ─────────────────────────────────────────────────────────────────────────────

func (s *invoiceEngineService) GenerateInvoice(
	ctx context.Context,
	req invoice_engine_contracts.CreateInvoiceRequest,
) (*invoice_engine_models.Invoice, error) {

	if len(req.Items) == 0 {
		return nil, fmt.Errorf("invoice must have at least one line item")
	}

	var invoice *invoice_engine_models.Invoice

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		invoiceID := uuid.New()

		var (
			subTotal float64
			taxTotal float64
			items    []*invoice_engine_models.InvoiceItem
			breakups []*invoice_engine_models.InvoiceTaxBreakup
		)

		for _, itemReq := range req.Items {
			amount := itemReq.Quantity * itemReq.Rate

			taxResult, err := s.taxEngine.CalculateTax(txCtx, tax_engine_contracts.TaxCalculationRequest{
				LegalEntityID: *req.LegalEntityID,
				TaxCategoryID: *itemReq.TaxCategoryID,
				BaseAmount:    amount,
			})
			if err != nil {
				return fmt.Errorf("tax calculation failed for '%s': %w", itemReq.Description, err)
			}

			items = append(items, &invoice_engine_models.InvoiceItem{
				ID:            uuid.New(),
				InvoiceID:     invoiceID,
				Description:   itemReq.Description,
				Quantity:      itemReq.Quantity,
				Rate:          itemReq.Rate,
				Amount:        amount,
				TaxCategoryID: itemReq.TaxCategoryID,
				HSNCode:       itemReq.HSNCode,
				TaxRate:       taxResult.TotalTax,
				TaxAmount:     taxResult.TotalTax,
				TotalAmount:   amount + taxResult.TotalTax,
			})

			subTotal += amount
			taxTotal += taxResult.TotalTax

			for _, component := range taxResult.Components {
				breakups = append(breakups, &invoice_engine_models.InvoiceTaxBreakup{
					ID:        uuid.New(),
					InvoiceID: invoiceID,
					TaxType:   component.Name,
					Rate:      component.Rate,
					Amount:    component.Rate,
					CreatedAt: time.Now(),
				})
			}
		}

		now := time.Now()

		invoice = &invoice_engine_models.Invoice{
			ID:                invoiceID,
			OrganizationID:    req.OrganizationID,
			LegalEntityID:     req.LegalEntityID,
			InvoiceNumber:     generateInvoiceNumber(),
			InvoiceDate:       now,
			CustomerID:        req.CustomerID,
			CustomerName:      req.CustomerName,
			CustomerGST:       req.CustomerGST,
			CustomerStateCode: req.CustomerStateCode,
			SubTotal:          subTotal,
			TaxTotal:          taxTotal,
			Total:             subTotal + taxTotal,
			Status:            "DRAFT",
			PaymentStatus:     "UNPAID",
			Source:            req.Source,
			SourceRefID:       req.SourceRefID,
		}

		if err := s.invoicingRepo.CreateInvoice(txCtx, invoice); err != nil {
			return fmt.Errorf("create invoice: %w", err)
		}

		for _, item := range items {
			if err := s.invoicingRepo.CreateInvoiceItem(txCtx, item); err != nil {
				return fmt.Errorf("create invoice item: %w", err)
			}
		}

		for _, breakup := range breakups {
			// FIXED: Breakup -> BreakUp
			if err := s.invoicingRepo.CreateInvoiceTaxBreakUp(txCtx, breakup); err != nil {
				return fmt.Errorf("create tax breakup: %w", err)
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return invoice, nil
}

// ─────────────────────────────────────────────────────────────────────────────
// IssueInvoice
// ─────────────────────────────────────────────────────────────────────────────

func (s *invoiceEngineService) IssueInvoice(
	ctx context.Context,
	invoiceID uuid.UUID,
) (*invoice_engine_models.Invoice, error) {

	var invoice *invoice_engine_models.Invoice

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		var err error
		invoice, err = s.invoicingRepo.GetInvoiceByID(txCtx, invoiceID)
		if err != nil {
			return fmt.Errorf("invoice not found: %w", err)
		}

		if invoice.Status != "DRAFT" {
			return fmt.Errorf("only DRAFT invoices can be issued, current status: %s", invoice.Status)
		}

		if err := s.invoicingRepo.UpdateInvoiceStatus(txCtx, invoiceID, "ISSUED"); err != nil {
			return fmt.Errorf("update invoice status: %w", err)
		}

		invoice.Status = "ISSUED"
		return nil
	})
	if err != nil {
		return nil, err
	}

	return invoice, nil
}

// ─────────────────────────────────────────────────────────────────────────────
// AddInvoiceItems
// ─────────────────────────────────────────────────────────────────────────────

func (s *invoiceEngineService) AddInvoiceItems(
	ctx context.Context,
	invoiceID uuid.UUID,
	items []invoice_engine_contracts.InvoiceItemRequest,
) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		invoice, err := s.invoicingRepo.GetInvoiceByID(txCtx, invoiceID)
		if err != nil {
			return fmt.Errorf("invoice not found: %w", err)
		}
		if invoice.Status != "DRAFT" {
			return fmt.Errorf("cannot add items to a %s invoice", invoice.Status)
		}

		for _, itemReq := range items {
			item := &invoice_engine_models.InvoiceItem{
				ID:            uuid.New(),
				InvoiceID:     invoiceID,
				Description:   itemReq.Description,
				Quantity:      itemReq.Quantity,
				Rate:          itemReq.Rate,
				Amount:        itemReq.Quantity * itemReq.Rate,
				TaxCategoryID: itemReq.TaxCategoryID,
				HSNCode:       itemReq.HSNCode,
			}
			if err := s.invoicingRepo.CreateInvoiceItem(txCtx, item); err != nil {
				return fmt.Errorf("create invoice item: %w", err)
			}
		}

		return nil
	})
}

// ─────────────────────────────────────────────────────────────────────────────
// CancelInvoice
// ─────────────────────────────────────────────────────────────────────────────

func (s *invoiceEngineService) CancelInvoice(
	ctx context.Context,
	invoiceID uuid.UUID,
	reason string,
) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		invoice, err := s.invoicingRepo.GetInvoiceByID(txCtx, invoiceID)
		if err != nil {
			return fmt.Errorf("invoice not found: %w", err)
		}
		if invoice.Status == "PAID" {
			return fmt.Errorf("cannot cancel a paid invoice — initiate a refund instead")
		}
		if invoice.Status == "CANCELLED" {
			return fmt.Errorf("invoice is already cancelled")
		}

		return s.invoicingRepo.UpdateInvoiceStatus(txCtx, invoiceID, "CANCELLED")
	})
}

// ─────────────────────────────────────────────────────────────────────────────
// RecordPayment
// ─────────────────────────────────────────────────────────────────────────────

func (s *invoiceEngineService) RecordPayment(
	ctx context.Context,
	req invoice_engine_contracts.RecordPaymentRequest,
) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		invoice, err := s.invoicingRepo.GetInvoiceByID(txCtx, req.InvoiceID)
		if err != nil {
			return fmt.Errorf("invoice not found: %w", err)
		}
		if invoice.Status == "CANCELLED" {
			return fmt.Errorf("cannot record payment on a cancelled invoice")
		}

		link := &invoice_engine_models.InvoicePaymentLink{
			ID:             uuid.New(),
			InvoiceID:      req.InvoiceID,
			PaymentRef:     req.PaymentRef,
			PaymentMethod:  req.PaymentMethod,
			CurrencyCode:   "INR",
			AllocatedPaise: req.AllocatedPaise,
			Status:         "APPLIED",
			PaidAt:         req.PaidAt,
			CreatedAt:      time.Now(),
		}

		// FIXED: CreatePaymentLink -> CreateInvoicePaymentLink
		if err := s.invoicingRepo.CreateInvoicePaymentLink(txCtx, link); err != nil {
			return fmt.Errorf("create payment link: %w", err)
		}

		newPaidPaise := invoice.PaidPaise + req.AllocatedPaise
		newPaymentStatus := resolveInvoicePaymentStatus(invoice.Total, float64(newPaidPaise)/100)

		if err := s.invoicingRepo.UpdateInvoicePaidAmount(
			txCtx,
			req.InvoiceID,
			newPaidPaise,
			newPaymentStatus,
		); err != nil {
			return fmt.Errorf("update invoice paid amount: %w", err)
		}

		if newPaymentStatus == "PAID" {
			if err := s.invoicingRepo.UpdateInvoiceStatus(txCtx, req.InvoiceID, "PAID"); err != nil {
				return fmt.Errorf("update invoice status to PAID: %w", err)
			}
		}

		return nil
	})
}

// ─────────────────────────────────────────────────────────────────────────────
// ReversePayment
// ─────────────────────────────────────────────────────────────────────────────

func (s *invoiceEngineService) ReversePayment(
	ctx context.Context,
	req invoice_engine_contracts.ReversePaymentRequest,
) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		// FIXED: GetPaymentLinkByRef -> GetInvoicePaymentLinkByRef
		link, err := s.invoicingRepo.GetInvoicePaymentLinkByRef(txCtx, req.InvoiceID, req.PaymentRef)
		if err != nil {
			return fmt.Errorf("payment link not found: %w", err)
		}
		if link.Status == "REVERSED" {
			return fmt.Errorf("payment link is already reversed")
		}

		now := time.Now()
		// FIXED: UpdatePaymentLinkStatus -> UpdateInvoicePaymentLinkStatus
		if err := s.invoicingRepo.UpdateInvoicePaymentLinkStatus(txCtx, link.ID, "REVERSED", &now); err != nil {
			return fmt.Errorf("reverse payment link: %w", err)
		}

		invoice, err := s.invoicingRepo.GetInvoiceByID(txCtx, req.InvoiceID)
		if err != nil {
			return err
		}

		newPaidPaise := invoice.PaidPaise - link.AllocatedPaise
		if newPaidPaise < 0 {
			newPaidPaise = 0
		}

		newPaymentStatus := resolveInvoicePaymentStatus(invoice.Total, float64(newPaidPaise)/100)

		return s.invoicingRepo.UpdateInvoicePaidAmount(txCtx, req.InvoiceID, newPaidPaise, newPaymentStatus)
	})
}

// ─────────────────────────────────────────────────────────────────────────────
// Queries
// ─────────────────────────────────────────────────────────────────────────────

func (s *invoiceEngineService) GetInvoiceByID(
	ctx context.Context,
	invoiceID uuid.UUID,
) (*invoice_engine_models.Invoice, error) {

	invoice, err := s.invoicingRepo.GetInvoiceByID(ctx, invoiceID)
	if err != nil {
		return nil, fmt.Errorf("invoice not found: %w", err)
	}
	return invoice, nil
}

func (s *invoiceEngineService) GetInvoiceWithDetails(
	ctx context.Context,
	invoiceID uuid.UUID,
) (*invoice_engine_contracts.InvoiceWithDetails, error) {

	invoice, err := s.invoicingRepo.GetInvoiceByID(ctx, invoiceID)
	if err != nil {
		return nil, fmt.Errorf("invoice not found: %w", err)
	}

	items, err := s.invoicingRepo.GetInvoiceItems(ctx, invoiceID)
	if err != nil {
		return nil, fmt.Errorf("get invoice items: %w", err)
	}

	// FIXED: GetInvoiceTaxBreakups -> GetInvoiceTaxBreakups (Matches interface)
	breakups, err := s.invoicingRepo.GetInvoiceTaxBreakups(ctx, invoiceID)
	if err != nil {
		return nil, fmt.Errorf("get tax breakups: %w", err)
	}

	// FIXED: GetPaymentLinks -> GetPaymentLinks (Matches interface)
	paymentLinks, err := s.invoicingRepo.GetPaymentLinks(ctx, invoiceID)
	if err != nil {
		return nil, fmt.Errorf("get payment links: %w", err)
	}

	return &invoice_engine_contracts.InvoiceWithDetails{
		Invoice:      *invoice,
		Items:        items,
		TaxBreakups:  breakups,
		PaymentLinks: paymentLinks,
	}, nil
}

func (s *invoiceEngineService) GetInvoicesBySource(
	ctx context.Context,
	source string,
	sourceRefID uuid.UUID,
) ([]*invoice_engine_models.Invoice, error) {

	invoices, err := s.invoicingRepo.GetInvoicesBySource(ctx, source, sourceRefID)
	if err != nil {
		return nil, fmt.Errorf("get invoices by source: %w", err)
	}
	return invoices, nil
}

func (s *invoiceEngineService) GetOutstandingAmount(
	ctx context.Context,
	invoiceID uuid.UUID,
) (int64, error) {

	invoice, err := s.invoicingRepo.GetInvoiceByID(ctx, invoiceID)
	if err != nil {
		return 0, fmt.Errorf("invoice not found: %w", err)
	}

	totalPaise := int64(invoice.Total * 100)
	outstanding := totalPaise - invoice.PaidPaise

	if outstanding < 0 {
		return 0, nil
	}
	return outstanding, nil
}

// ─────────────────────────────────────────────────────────────────────────────
// Helpers
// ─────────────────────────────────────────────────────────────────────────────

func resolveInvoicePaymentStatus(total float64, paidAmount float64) string {
	switch {
	case paidAmount <= 0:
		return "UNPAID"
	case paidAmount > total:
		return "OVERPAID"
	case paidAmount >= total:
		return "PAID"
	default:
		return "PARTIALLY_PAID"
	}
}

func generateInvoiceNumber() string {
	return fmt.Sprintf("INV-%d", time.Now().UnixMilli())
}
