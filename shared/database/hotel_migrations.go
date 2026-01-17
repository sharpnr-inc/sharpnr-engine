package database

import (
	hotel_models "sharpnr-engine/modules/hospitality/hotel/models"

	"gorm.io/gorm"
)

func migrateHotel(db *gorm.DB) error {
	return db.AutoMigrate(
		// ────────────── CORE HOTEL STRUCTURE ──────────────
		&hotel_models.Hotel{},
		&hotel_models.HotelFloor{},
		&hotel_models.Room{},
		&hotel_models.HotelRoomPrice{},

		// ────────────── GUESTS ──────────────
		&hotel_models.HotelGuest{},

		// ────────────── BOOKINGS (PARENT) ──────────────
		&hotel_models.HotelRoomBooking{},

		// ────────────── BOOKING CHILD TABLES ──────────────
		&hotel_models.BookingGuest{},
		&hotel_models.HotelRoomBookingPayment{},
		&hotel_models.HotelBookingCharge{},

		// ────────────── SERVICES ──────────────
		&hotel_models.HotelService{},
		&hotel_models.HotelServiceItem{},

		// ────────────── TAX ──────────────
		&hotel_models.HotelTaxRule{},
		&hotel_models.HotelTaxComponent{},

		// ────────────── BILLING ──────────────
		&hotel_models.HotelBill{},
		&hotel_models.HotelBillTax{},
	)
}
