package hospitalityinvoicingrepo

import "gorm.io/gorm"

// hospitalityInvoicingRepo provides the GORM-based implementation of
// the HospitalityInvoicingRepositoryInterfaces interface.
//
// This repository is responsible for all database interactions related to
// the Hospitality Invoice module, including:
//   - Creating invoices
//   - Updating invoice totals or payment status
//   - Retrieving invoices by ID or session reference
//   - Querying invoice lists for reporting or audits
//
// The repository abstracts away direct DB operations and ensures the service
// layer remains fully decoupled from the underlying persistence engine.
// Any changes in database technology (GORM, SQL, NoSQL, etc.) require only
// adjustments in this repository layer, not in business logic.
type hospitalityInvoicingRepo struct {
	db *gorm.DB
}

// NewHospitalityInvoicingRepo creates and returns a new instance of
// hospitalityInvoicingRepo using the provided GORM database connection.
//
// This function follows dependency injection principles, enabling the
// invoicing module to operate independently of specific configuration
// details of the database layer.
//
// Parameters:
//   - db: GORM database connection used for executing invoice-related queries.
//
// Returns:
//
//	*hospitalityInvoicingRepo — a fully initialized repository instance.
func NewHospitalityInvoicingRepo(db *gorm.DB) *hospitalityInvoicingRepo {
	return &hospitalityInvoicingRepo{
		db: db,
	}
}
