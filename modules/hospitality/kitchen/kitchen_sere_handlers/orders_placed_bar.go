package kitchen_sere_handlers

import (
	"fmt"
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/kitchenRepo"
	"sharpnr-engine/sere"
	"sharpnr-engine/sere/events/hospitality_events/hotel_events"
)

type OrderPlacedFromHotelRoom struct{}

func NewOrderPlacedFromHotelRoom(kitchenRepo kitchenrepo.KitchenRepositoryInterface) *OrderPlacedFromHotelRoom {
	return &OrderPlacedFromHotelRoom{}
}

func (h *OrderPlacedFromHotelRoom) Handle(e sere.Event) error {
	event := e.Payload.(hotel_events.HotelOrderCreated)
	fmt.Println(event)
	return nil

}
