package restaurantservices

import (
	"context"
	org_customer_service "sharpnr-engine/core/customer/service"
	dbtx "sharpnr-engine/core/dbTx"
	restaurantrepo "sharpnr-engine/modules/hospitality/restaurant/repo"
	payment_service "sharpnr-engine/payment_engine/services"
	"sharpnr-engine/sero"
	invoice_engine_services "sharpnr-engine/service_engines/invoicing_engine/services"
)

type restaurantService struct {
	restaurantRepo restaurantrepo.RestaurantRepositoryInterface
	seroEngine     *sero.SEROEngine
	txManager      dbtx.TxManager
	customerEngine org_customer_service.OrgCustomerServiceInterface
	invoiceEngine  invoice_engine_services.InvoiceEngineServiceInterface
	paymentsEngine payment_service.PaymentServiceInterface
}

// CreateOrder implements RestaurantServiceInterface.
func (r *restaurantService) CreateOrder(ctx context.Context) {
	panic("unimplemented")
}

// AddRestaurantGeneralDetails implements RestaurantServiceInterface.
func (r *restaurantService) AddRestaurantGeneralDetails(ctx context.Context) {
	panic("unimplemented")
}

func NewRestaurantService(
	restaurantRepo restaurantrepo.RestaurantRepositoryInterface,
	seroEngine *sero.SEROEngine,
	txManager dbtx.TxManager,
	customerEngine org_customer_service.OrgCustomerServiceInterface,
	invoiceEngine invoice_engine_services.InvoiceEngineServiceInterface,
	paymentsEngine payment_service.PaymentServiceInterface,
) *restaurantService {
	return &restaurantService{
		restaurantRepo: restaurantRepo,
		seroEngine:     seroEngine,
		txManager:      txManager,
		customerEngine: customerEngine,
		invoiceEngine:  invoiceEngine,
		paymentsEngine: paymentsEngine,
	}

}
