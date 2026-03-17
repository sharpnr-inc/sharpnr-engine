package payment_engine_repo

import "gorm.io/gorm"

type paymentRepo struct {
	db *gorm.DB
}

// payment Constructor
func NewPaymentRepository(db *gorm.DB) *paymentRepo {
	return &paymentRepo{
		db: db,
	}
}
