package invoice_engine_models

import "github.com/google/uuid"

type InvoiceCustomer struct {
	ID         uuid.UUID
	CustomerID uuid.UUID // this bind with customer engine from customers.customer
	InvoiceID  uuid.UUID
}
