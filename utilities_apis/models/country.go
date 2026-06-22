package utilitiesApi_models

import "time"

// Country is a country record from the countries-states-cities dataset
// (countrystatecity.in). IDs are upstream canonical IDs, not UUIDs.
type Country struct {
	ID               int64      `gorm:"column:id;primaryKey"`
	Name             string     `gorm:"column:name;type:varchar(100);not null"`
	ISO3             *string    `gorm:"column:iso3;type:char(3)"`
	NumericCode      *string    `gorm:"column:numeric_code;type:char(3)"`
	ISO2             *string    `gorm:"column:iso2;type:char(2)"`
	PhoneCode        *string    `gorm:"column:phonecode;type:varchar(255)"`
	Capital          *string    `gorm:"column:capital;type:varchar(255)"`
	Currency         *string    `gorm:"column:currency;type:varchar(255)"`
	CurrencyName     *string    `gorm:"column:currency_name;type:varchar(255)"`
	CurrencySymbol   *string    `gorm:"column:currency_symbol;type:varchar(255)"`
	TLD              *string    `gorm:"column:tld;type:varchar(255)"`
	Native           *string    `gorm:"column:native;type:varchar(255)"`
	Population       *int64     `gorm:"column:population"`
	GDP              *int64     `gorm:"column:gdp"`
	Region           *string    `gorm:"column:region;type:varchar(255)"`
	RegionID         *int64     `gorm:"column:region_id"`
	Subregion        *string    `gorm:"column:subregion;type:varchar(255)"`
	SubregionID      *int64     `gorm:"column:subregion_id"`
	Nationality      *string    `gorm:"column:nationality;type:varchar(255)"`
	AreaSqKm         *float64   `gorm:"column:area_sq_km"`
	PostalCodeFormat *string    `gorm:"column:postal_code_format;type:varchar(255)"`
	PostalCodeRegex  *string    `gorm:"column:postal_code_regex;type:varchar(255)"`
	Timezones        *string    `gorm:"column:timezones;type:text"`
	Translations     *string    `gorm:"column:translations;type:text"`
	Latitude         *float64   `gorm:"column:latitude;type:numeric(10,8)"`
	Longitude        *float64   `gorm:"column:longitude;type:numeric(11,8)"`
	Emoji            *string    `gorm:"column:emoji;type:varchar(191)"`
	EmojiU           *string    `gorm:"column:emojiU;type:varchar(191)"`
	CreatedAt        *time.Time `gorm:"column:created_at"`
	UpdatedAt        time.Time  `gorm:"column:updated_at;not null"`
	Flag             int16      `gorm:"column:flag;not null;default:1"`
	WikiDataID       *string    `gorm:"column:wikiDataId;type:varchar(255)"`
}

func (Country) TableName() string { return "geo.countries" }
