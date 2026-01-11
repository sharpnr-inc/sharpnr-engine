package hotel_events_test

import (
	"sharpnr-engine/modules/accounting/accounting_repo"
	hotel_events "sharpnr-engine/modules/accounting/accounting_sere_handlers/acc_hospitality/hotel"
	"sharpnr-engine/sere"
	"testing"
)

func TestHospitality_ChargesCreatedEventHandler_HandleCharges(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		accRepo accounting_repo.AccountingRepositoryInterface
		// Named input parameters for target function.
		e       sere.Event
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := hotel_events.NewHospitality_ChargesCreatedEventHandler(tt.accRepo)
			gotErr := h.HandleCharges(tt.e)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("HandleCharges() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("HandleCharges() succeeded unexpectedly")
			}
		})
	}
}
