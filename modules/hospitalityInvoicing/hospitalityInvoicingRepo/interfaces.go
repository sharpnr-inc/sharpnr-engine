package hospitalityinvoicingrepo

// HospitalityInvoicingRepositoryInterfaces defines the set of
// repository-level operations for managing invoice data within
// the Hospitality Invoicing module.
//
// This repository acts as the data access layer for all invoice-
// related interactions, including creation, retrieval, updates,
// and linking invoices with Restaurant, Bar, and Hotel sessions.
//
// Implementations of this interface should handle:
//   - Persistent storage of invoices
//   - Querying invoices by ID or session reference
//   - Updating invoice totals, taxes, and payment status
//   - Ensuring consistency across hospitality services
//
// This abstraction decouples the invoicing domain from specific
// database technologies (e.g., GORM, SQL, NoSQL), enabling
// flexible backend implementations while keeping higher-level
// services independent of storage details.
type HospitalityInvoicingRepositoryInterfaces interface {
}
