package invoice_engine_services

import (
	"context"
	invoice_engine_contracts "sharpnr-engine/service_engines/invoicing_engine/contracts"
)

func (s *invoiceEngineService) GenerateInvoice(ctx context.Context, request invoice_engine_contracts.CreateInvoiceRequest) (interface{}, error) {
	panic("unimplemented")
}

func (s *invoiceEngineService) addInvoiceItems(ctx context.Context, items []invoice_engine_contracts.InvoiceItemRequest) error {
	panic("unimplemented")
}
