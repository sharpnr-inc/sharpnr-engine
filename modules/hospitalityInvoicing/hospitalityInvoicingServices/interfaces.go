package hospitalityinvoicingservices

// HospitalityInvoicingServiceInterfaces defines the business logic
// layer for the Hospitality Invoicing module.
//
// This service is responsible for implementing all domain-level
// operations related to invoice creation, calculation, and lifecycle
// management across the hospitality ecosystem (Restaurant, Bar, Hotel).
//
// Implementations of this interface should handle high-level logic such as:
//   - Generating invoices for restaurant table sessions, bar sessions,
//     and hotel bookings
//   - Calculating subtotal, taxes, service charges, and final totals
//   - Linking invoices to customers and payment records
//   - Updating invoice status after payment confirmation (including webhook flows)
//   - Providing invoice data to controllers, handlers, and external services
//
// This abstraction ensures that HTTP handlers/controllers interact only
// with business rules, while the underlying storage or payment systems
// remain fully decoupled.
type HospitalityInvoicingServiceInterfaces interface {
}
