package payment_engine_service

import (
	"fmt"
	"time"
)

func generateOrderNumber() string {
	return fmt.Sprintf("ORD-%d", time.Now().UnixMilli())
}

func generatePaymentNumber() string {
	return fmt.Sprintf("PAY-%d", time.Now().UnixMilli())
}

func generateRefundNumber() string {
	return fmt.Sprintf("RFN-%d", time.Now().UnixMilli())
}
