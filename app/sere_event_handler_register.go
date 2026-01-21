package app

import (
	"sharpnr-engine/modules/hospitality/hotel/hotel_sere_handlers"
	"sharpnr-engine/sere"
	accounting_sere_handlers "sharpnr-engine/sharp_services/accounting_engine/sere_handlers"
)

func RegisterSereEventHandler(
	sere *sere.SERE,
) {
	accounting_sere_handlers.Register(sere)
	hotel_sere_handlers.Register(sere)
}
