package hotel_events

import (
	"log"
	"sharpnr-engine/modules/accounting/accounting_repo"
	hotelevents "sharpnr-engine/modules/hospitality/hotel/hotel_events"
	"sharpnr-engine/sere"
)

type Hospitality_ChargesCreatedEventHandler struct {
	accRepo accounting_repo.AccountingRepositoryInterface
}

func NewHospitality_ChargesCreatedEventHandler(
	accRepo accounting_repo.AccountingRepositoryInterface,
) *Hospitality_ChargesCreatedEventHandler {
	return &Hospitality_ChargesCreatedEventHandler{accRepo: accRepo}
}

func (h *Hospitality_ChargesCreatedEventHandler) HandleCharges(e sere.Event) error {
	event := e.Payload.(hotelevents.ChargesCreatedInBooking)
	// ctx := e.Ctx

	// 2. Log that data was received (using your request to "say data received")
	log.Printf("[Accounting Handler] Data Received for Booking: %s", event.BookingID)
	log.Printf("Payload Details -> Total: %v, Type: %v, OrgID: %s",
		event.TotalAmount,
		event.ChargesType,
		event.OrgID,
	)

	// 3. Logic: Map this charge to your Accounting Ledger
	// Example: h.repo.CreateJournalEntry(ctx, ...)

	return nil

}
