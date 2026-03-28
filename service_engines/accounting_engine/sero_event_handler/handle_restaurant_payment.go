package accounting_sero_event_handler

import (
	"context"
	"fmt"
	restaurant_events "sharpnr-engine/modules/hospitality/restaurant/events"
	"sharpnr-engine/sero"
	accountingdtos "sharpnr-engine/service_engines/accounting_engine/dto"
	accounting_service "sharpnr-engine/service_engines/accounting_engine/services"
)

func handleRestaurantPaymentSettlment(
	accountingService accounting_service.AccountingServiceInterface,
) sero.Handler {

	return func(ctx context.Context, payload interface{}) error {

		data, ok := payload.(restaurant_events.RestaurantPaymentSettledDto)
		if !ok {
			if err := sero.MapToStruct(payload, &data); err != nil {
				return fmt.Errorf("invalid payload type: %w", err)
			}
		}

		if err := data.ValidateDto(); err != nil {
			return err
		}

		// Step 1 we First we need to create the Journal Entry

		journalEntry := &accountingdtos.JournalEntryRequestDto{
			EntryDate: data.PaidAt,
		}

		if err := accountingService.AddJournalEntry(ctx, data.OrganizationID, data.OrganizationID, *journalEntry); err != nil {
			return err
		}


		return nil
	}
}
