package payment_service

import "sharpnr-engine/payment_gateway/payment_repo"

type PaymentService struct {
	paymentRepo payment_repo.PaymentRepositoryInterface
}

func NewPaymentService(paymentRepo payment_repo.PaymentRepositoryInterface) *PaymentService {
	return &PaymentService{
		paymentRepo: paymentRepo,
	}
}
