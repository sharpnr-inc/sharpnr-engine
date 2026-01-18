package hotel_events

import (
	"encoding/json"
	"errors"
	"sharpnr-engine/modules/accounting/accounting_models"
	hotelevents "sharpnr-engine/modules/hospitality/hotel/hotel_events"
	"sharpnr-engine/sere"

	"gorm.io/datatypes"
)

type Hospitality_ChargesCreatedEventHandler struct {
}

func NewHospitality_ChargesCreatedEventHandler() *Hospitality_ChargesCreatedEventHandler {
	return &Hospitality_ChargesCreatedEventHandler{}
}

func (h *Hospitality_ChargesCreatedEventHandler) HandleCharges(e sere.Event) error {

	return nil
}

func validateJournal(lines []accounting_models.JournalLine) error {
	var totalDebit, totalCredit float64

	for _, l := range lines {
		totalDebit += l.Debit
		totalCredit += l.Credit
	}

	if totalDebit != totalCredit {
		return errors.New("journal entry is not balanced")
	}

	return nil
}

func buildMeta(e hotelevents.ChargesCreatedInBooking) datatypes.JSON {
	m := map[string]any{
		"booking_id": e.BookingID,
		"hotel_id":   e.HotelID,
		"charge_id":  e.ChargesId,
		"quantity":   e.Quantity,
		"unit_rate":  e.UnitAmount,
	}
	b, _ := json.Marshal(m)
	return datatypes.JSON(b)
}
