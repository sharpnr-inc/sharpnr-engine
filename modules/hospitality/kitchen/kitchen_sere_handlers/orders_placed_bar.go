package kitchen_sere_handlers

import (
	"fmt"
	hotelevents "sharpnr-engine/modules/hospitality/hotel/hotel_events"
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/kitchenRepo"
	"sharpnr-engine/sere"
)

type OrderPlacedFromHotelRoom struct{}

func NewOrderPlacedFromHotelRoom(kitchenRepo kitchenrepo.KitchenRepositoryInterface) *OrderPlacedFromHotelRoom {
	return &OrderPlacedFromHotelRoom{}
}

func (h *OrderPlacedFromHotelRoom) Handle(e sere.Event) error {
	event := e.Payload.(hotelevents.HotelOrderCreated)
	fmt.Println(event)
	return nil

}
