package database

import (
	hotel_models "sharpnr-engine/modules/hospitality/hotel/models"

	"gorm.io/gorm"
)

func migrateHotel(db *gorm.DB) error {
	return db.AutoMigrate(
		&hotel_models.Hotel{},
		&hotel_models.HotelFloor{},

		//--------Rooms--------------//
		&hotel_models.HotelRoom{},
		&hotel_models.HotelRoomRate{},
		&hotel_models.HotelRoomDetail{},

		&hotel_models.HotelStay{},
		&hotel_models.HotelCharge{},
		&hotel_models.HotelStayPayment{},

		// ────────────── GUESTS ──────────────
		&hotel_models.HotelGuest{},
		&hotel_models.HotelStayGuest{},
	)
}
