package utlitiesApi_repo

import "gorm.io/gorm"

type utlitiesApiRepo struct {
	db *gorm.DB
}

func NewUtilitiesApiRepo(db *gorm.DB) *utlitiesApiRepo {
	return &utlitiesApiRepo{
		db: db,
	}
}
