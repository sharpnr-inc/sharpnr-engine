package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	// beacuse of gorm.Model we already Have ID
	Name            string
	Email           string
	Password        string
	OraganizationID uint
}
