package sharphrrepository

import "gorm.io/gorm"

type sharphrRepository struct {
	db *gorm.DB
}

func NewSharpHREmployeeRepo(datab *gorm.DB) *sharphrRepository {
	return &sharphrRepository{
		db: datab,
	}
}
