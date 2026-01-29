package hotel_models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type HotelOperationsTiming struct {
	ID           uuid.UUID
	HotelID      uuid.UUID
	Is24X7       bool
	CheckInTime  string
	CheckOutTime string
}

func (HotelOperationsTiming) TableName() string {
	return "hospitality.hotel_operations_timing"
}

type HotelSetup struct {
	ID          uuid.UUID `gorm:"type:uuid; not null"`
	HotelID     uuid.UUID `gorm:"type:uuid;index;not null"`
	SetupStatus string
	//check:setup_status IN ('PENDING','IN_PROGRESS','COMPLETED')"`
	SetupStep string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (HotelSetup) TableName() string {
	return "hospitality.hotel_setup"
}

type HotelMetaData struct {
	ID        uuid.UUID
	HotelID   uuid.UUID
	Meta      datatypes.JSON
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (HotelMetaData) TableName() string {
	return "hospitality.hotel_meta_data"
}
