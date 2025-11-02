package hospitalityrepo

import (
	"sharpnr-engine/shared/database"

	"gorm.io/gorm"
)

type hotelRepository struct {
	db *gorm.DB
}

func NewHotelRepository(db *gorm.DB) *hotelRepository {
	return &hotelRepository{
		db: database.DB,
	}
}
