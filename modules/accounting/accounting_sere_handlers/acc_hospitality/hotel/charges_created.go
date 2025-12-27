package hotel_events

import (
	"encoding/json"
	"errors"
	"log"
	accountingconstants "sharpnr-engine/modules/accounting/accounting_constants"
	"sharpnr-engine/modules/accounting/accounting_models"
	"sharpnr-engine/modules/accounting/accounting_repo"
	hotelevents "sharpnr-engine/modules/hospitality/hotel/hotel_events"
	"sharpnr-engine/sere"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
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
	ctx := e.Ctx

	log.Printf(
		"[Accounting][START] Charge event received | Org=%s Booking=%s Amount=%.2f Type=%s",
		event.OrgID,
		event.BookingID,
		event.TotalAmount,
		event.ChargesType,
	)

	return h.accRepo.WithTx(ctx, func(txRepo accounting_repo.AccountingRepositoryInterface) error {

		// 1️⃣ Find Accounting Book
		book, err := txRepo.GetAccountingBookByOrgAndSubModule(ctx, event.OrgID, "HOTEL")
		if err != nil {
			log.Printf("[Accounting][ERROR] Failed to resolve accounting book | Org=%s err=%v", event.OrgID, err)
			return err
		}
		log.Printf("[Accounting] Book resolved | BookID=%s", book.ID)

		// 2️⃣ Find Fiscal Year
		fy, err := txRepo.GetOpenFiscalYearByBookAndDate(ctx, book.ID, event.OccurredAt)
		if err != nil {
			log.Printf("[Accounting][ERROR] Failed to resolve fiscal year | BookID=%s Date=%s err=%v",
				book.ID, event.OccurredAt, err)
			return err
		}
		log.Printf("[Accounting] Fiscal year resolved | FYID=%s", fy.ID)

		// 3️⃣ Resolve Receivable Account
		receivableAccount, err := txRepo.ResolveAccountMapping(
			ctx,
			event.OrgID,
			"HOTEL",
			"SYSTEM",
			"RECEIVABLE",
		)
		if err != nil {
			log.Printf("[Accounting][ERROR] Receivable account mapping failed | Org=%s err=%v", event.OrgID, err)
			return err
		}
		log.Printf("[Accounting] Receivable account resolved | AccountID=%s", receivableAccount.ID)

		// 4️⃣ Resolve Revenue Account
		revenueAccount, err := txRepo.ResolveAccountMapping(
			ctx,
			event.OrgID,
			"HOTEL",
			"CHARGE",
			string(event.ChargesType),
		)
		if err != nil {
			log.Printf("[Accounting][ERROR] Revenue account mapping failed | Org=%s ChargeType=%s err=%v",
				event.OrgID, event.ChargesType, err)
			return err
		}
		log.Printf("[Accounting] Revenue account resolved | AccountID=%s", revenueAccount.ID)

		// 5️⃣ Create Journal Entry
		je := &accounting_models.JournalEntry{
			ID:               uuid.New(),
			AccountingBookID: book.ID,
			AccountingBook:   *book,
			FiscalYearID:     fy.ID,
			FiscalYear:       *fy,
			EntryDate:        event.OccurredAt,
			Reference:        event.BookingID.String(),
			Narration:        "Hotel Charge recorded",
			EntryType:        accountingconstants.JournalNormal,
			CreatedAt:        time.Now(),
		}

		log.Printf("[Accounting] Creating journal entry | JEID=%s Ref=%s", je.ID, je.Reference)

		if err := txRepo.CreateJournalEntry(ctx, je); err != nil {
			log.Printf("[Accounting][ERROR] Failed to create journal entry | JEID=%s err=%v", je.ID, err)
			return err
		}

		// 6️⃣ Build Journal Lines
		lines := []accounting_models.JournalLine{
			{
				ID:             uuid.New(),
				JournalEntryID: je.ID,
				AccountID:      receivableAccount.ID,
				Debit:          event.TotalAmount,
				Credit:         0,
				PartyID:        &event.BookingID,
				Meta:           buildMeta(event),
				CreatedAt:      time.Now(),
			},
			{
				ID:             uuid.New(),
				JournalEntryID: je.ID,
				AccountID:      revenueAccount.ID,
				Debit:          0,
				Credit:         event.TotalAmount,
				PartyID:        &event.BookingID,
				Meta:           buildMeta(event),
				CreatedAt:      time.Now(),
			},
		}

		log.Printf(
			"[Accounting] Journal lines prepared | JEID=%s Lines=%d Debit=%.2f Credit=%.2f",
			je.ID,
			len(lines),
			event.TotalAmount,
			event.TotalAmount,
		)

		// 7️⃣ Validate Journal
		if err := validateJournal(lines); err != nil {
			log.Printf("[Accounting][ERROR] Journal validation failed | JEID=%s err=%v", je.ID, err)
			return err
		}

		// 8️⃣ Persist Journal Lines
		log.Printf("[Accounting] Persisting journal lines | JEID=%s", je.ID)

		if err := txRepo.CreateJournalLines(ctx, lines); err != nil {
			log.Printf("[Accounting][ERROR] Failed to persist journal lines | JEID=%s err=%v", je.ID, err)
			return err
		}

		log.Printf("[Accounting][SUCCESS] Journal posted successfully | JEID=%s Booking=%s",
			je.ID, event.BookingID)

		return nil
	})
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
