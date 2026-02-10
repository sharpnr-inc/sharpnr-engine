package hotel_sere_handlers

import (
	"log"
	hotelevents "sharpnr-engine/modules/hospitality/hotel/events"
	"sharpnr-engine/sere"
)

type RoomCreatedLogHandler struct{}

func NewRoomCreatedLogHandler() *RoomCreatedLogHandler {
	return &RoomCreatedLogHandler{}
}

func (h *RoomCreatedLogHandler) Handle(e sere.Event) error {

	event := e.Payload.(hotelevents.RoomCreated)

	log.Println("🟢 [TEST] RoomCreated event received")
	log.Println("   RoomID :", event.RoomID)
	log.Println("   HotelID:", event.HotelID)
	log.Println("   Name   :", event.Name)

	// show context propagation
	if rid := e.Ctx.Value("request_id"); rid != nil {
		log.Println("   RequestID:", rid)
	}

	return nil
}
