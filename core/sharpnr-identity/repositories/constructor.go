package sharpnr_identity_auth_repositories

import "gorm.io/gorm"

type sharpnrIdentityAuthRepo struct {
	db *gorm.DB
}

func NewSharpnrIdentityAuthRepository(db *gorm.DB) *sharpnrIdentityAuthRepo {
	return &sharpnrIdentityAuthRepo{
		db: db,
	}
}
