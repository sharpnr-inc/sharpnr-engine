package organization

import (
	"github.com/sharpnr-inc/sharpnr-engine/employee"
	"gorm.io/gorm"
)

type Organization struct {
	gorm.Model

	Name string `gorm:"unique: not null"`
	Organization_id string 
	Address string
	Phone int
	Email string
	Employees []employee.Employee

}