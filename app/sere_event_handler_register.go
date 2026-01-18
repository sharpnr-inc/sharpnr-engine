package app

import (
	"sharpnr-engine/modules/accounting/accounting_sere_handlers"
	"sharpnr-engine/modules/hospitality/hotel/hotel_sere_handlers"
	"sharpnr-engine/sere"
)

func RegisterSereEventHandler(
	sere *sere.SERE,
) {
	accounting_sere_handlers.Register(sere)
	hotel_sere_handlers.Register(sere)
}
