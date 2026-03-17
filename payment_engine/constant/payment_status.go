package paymentconstant

type PaymentStatus string

const (
	SUCCESS PaymentStatus = ""
)

type PaymentGatewayProvider string

const (
	Paytm     PaymentGatewayProvider = "paytm"
	PhonePe   PaymentGatewayProvider = "phonepe"
	AmazonPay PaymentGatewayProvider = "amazon_pay"
)
