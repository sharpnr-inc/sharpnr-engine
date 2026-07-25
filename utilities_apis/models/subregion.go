package utilitiesApi_models

import "time"

// Subregion is a sub-division of a Region (e.g. Northern Africa) from the
// countries-states-cities dataset. IDs are upstream canonical IDs, not UUIDs.
type Subregion struct {
	ID           int64      `gorm:"column:id;primaryKey"`
	Name         string     `gorm:"column:name;type:varchar(100);not null"`
	Translations *string    `gorm:"column:translations;type:text"`
	RegionID     int64      `gorm:"column:region_id;not null"`
	CreatedAt    *time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at;not null"`
	Flag         int16      `gorm:"column:flag;not null;default:1"`
	WikiDataID   *string    `gorm:"column:wikiDataId;type:varchar(255)"`

	Region *Region `gorm:"foreignKey:RegionID"`
}

func (Subregion) TableName() string { return "geo.subregions" }
