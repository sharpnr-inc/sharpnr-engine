package utilitiesApi_models

import "time"

// Region is a top-level geographic region (e.g. Africa, Americas) from the
// countries-states-cities dataset. IDs are upstream canonical IDs, not UUIDs.
type Region struct {
	ID           int64      `gorm:"column:id;primaryKey"`
	Name         string     `gorm:"column:name;type:varchar(100);not null"`
	Translations *string    `gorm:"column:translations;type:text"`
	CreatedAt    *time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at;not null"`
	Flag         int16      `gorm:"column:flag;not null;default:1"`
	WikiDataID   *string    `gorm:"column:wikiDataId;type:varchar(255)"`
}

func (Region) TableName() string { return "geo.regions" }
