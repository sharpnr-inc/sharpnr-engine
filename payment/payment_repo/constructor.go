package payment_repo

import "gorm.io/gorm"

type PaymentRepo struct {
	db *gorm.DB
}

// payment Constructor
func NewPaymentRepository(db *gorm.DB) *PaymentRepo {
	return &PaymentRepo{
		db: db,
	}
}
