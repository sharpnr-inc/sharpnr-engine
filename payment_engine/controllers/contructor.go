package payment_engine_controllers

import payment_engine_service "sharpnr-engine/payment_engine/services"

type paymentEngineController struct {
	paymentEngineService payment_engine_service.PaymentServiceInterface
}

func NewPaymentEngineController(
	paymentEngineService payment_engine_service.PaymentServiceInterface,
) *paymentEngineController {
	return &paymentEngineController{
		paymentEngineService: paymentEngineService,
	}
}
