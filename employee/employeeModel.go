package employee

import "gorm.io/gorm"

type Employee struct{
	gorm.Model

	Name string
	Email string
	Password string
	OraganizationID uint
	
}