package hotel_sere_handlers

import "sharpnr-engine/sere"

func Register(sere *sere.SERE) {

	logHandler := NewRoomCreatedLogHandler()

	sere.SubscribeAsync(
		"RoomCreated",
		logHandler.Handle,
	)

}
