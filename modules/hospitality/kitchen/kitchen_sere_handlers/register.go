package kitchen_sere_handlers

import (
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/kitchenRepo"
	"sharpnr-engine/sere"
)

func Register(
	sere *sere.SERE,
	kitchenRepo kitchenrepo.KitchenRepositoryInterface,
) {
	handler := NewOrderPlacedFromHotelRoom(kitchenRepo)

	sere.Subscribe("OrderCreated", handler.Handle)
}
