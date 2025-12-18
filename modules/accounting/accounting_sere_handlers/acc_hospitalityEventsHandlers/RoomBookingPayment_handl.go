package acc_hospitality_events_handlers

import (
	"sharpnr-engine/modules/accounting/accounting_repo"
	"sharpnr-engine/sere"
)

type HospitalityRoomBookingPaymentEventHandler struct {
	accRepo accounting_repo.AccountingRepositoryInterface
}

func NewHospitalityRoomBookingPaymentEventHandler(
	accRepo accounting_repo.AccountingRepositoryInterface,
) *HospitalityRoomBookingPaymentEventHandler {
	return &HospitalityRoomBookingPaymentEventHandler{
		accRepo: accRepo,
	}
}

func (h *HospitalityRoomBookingPaymentEventHandler) Handle(e sere.SERE) error {
	panic("unimplemented")
}
