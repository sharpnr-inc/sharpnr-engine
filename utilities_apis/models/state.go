package utilitiesApi_models

import "time"

// State is a state/province record from the countries-states-cities dataset.
// IDs are upstream canonical IDs, not UUIDs.
type State struct {
	ID           int64      `gorm:"column:id;primaryKey"`
	Name         string     `gorm:"column:name;type:varchar(255);not null"`
	CountryID    int64      `gorm:"column:country_id;not null"`
	CountryCode  string     `gorm:"column:country_code;type:char(2);not null"`
	FipsCode     *string    `gorm:"column:fips_code;type:varchar(255)"`
	ISO2         *string    `gorm:"column:iso2;type:varchar(255)"`
	ISO31662     *string    `gorm:"column:iso3166_2;type:varchar(10)"`
	Type         *string    `gorm:"column:type;type:varchar(191)"`
	Level        *int       `gorm:"column:level"`
	ParentID     *int64     `gorm:"column:parent_id"`
	Native       *string    `gorm:"column:native;type:varchar(255)"`
	Latitude     *float64   `gorm:"column:latitude;type:numeric(10,8)"`
	Longitude    *float64   `gorm:"column:longitude;type:numeric(11,8)"`
	Timezone     *string    `gorm:"column:timezone;type:varchar(255)"`
	Translations *string    `gorm:"column:translations;type:text"`
	CreatedAt    *time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at;not null"`
	Flag         int16      `gorm:"column:flag;not null;default:1"`
	WikiDataID   *string    `gorm:"column:wikiDataId;type:varchar(255)"`
	Population   *string    `gorm:"column:population;type:varchar(255)"`

	Country *Country `gorm:"foreignKey:CountryID"`
}

func (State) TableName() string { return "geo.states" }
