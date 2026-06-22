package utilitiesApi_models

import "time"

// City is a city record from the countries-states-cities dataset.
// IDs are upstream canonical IDs, not UUIDs.
type City struct {
	ID          int64     `gorm:"column:id;primaryKey"`
	Name        string    `gorm:"column:name;type:varchar(255);not null"`
	StateID     int64     `gorm:"column:state_id;not null"`
	StateCode   string    `gorm:"column:state_code;type:varchar(255);not null"`
	CountryID   int64     `gorm:"column:country_id;not null"`
	CountryCode string    `gorm:"column:country_code;type:char(2);not null"`
	Type        *string   `gorm:"column:type;type:varchar(191)"`
	Level       *int      `gorm:"column:level"`
	ParentID    *int64    `gorm:"column:parent_id"`
	Latitude    float64   `gorm:"column:latitude;type:numeric(10,8);not null"`
	Longitude   float64   `gorm:"column:longitude;type:numeric(11,8);not null"`
	Native      *string   `gorm:"column:native;type:varchar(255)"`
	Population  *int64    `gorm:"column:population"`
	Timezone    *string   `gorm:"column:timezone;type:varchar(255)"`
	CreatedAt   time.Time `gorm:"column:created_at;not null"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null"`
	Flag        int16     `gorm:"column:flag;not null;default:1"`
	WikiDataID  *string   `gorm:"column:wikiDataId;type:varchar(255)"`

	State   *State   `gorm:"foreignKey:StateID"`
	Country *Country `gorm:"foreignKey:CountryID"`
}

func (City) TableName() string { return "geo.cities" }
