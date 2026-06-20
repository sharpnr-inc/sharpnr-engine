package cables_repo

import "gorm.io/gorm"

type cablesRepo struct {
	db *gorm.DB
}
