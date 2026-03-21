package restaurantservices

import (
	"context"
	"github.com/google/uuid"
	org_customer_service "sharpnr-engine/core/customer/service"
	dbtx "sharpnr-engine/core/dbTx"
	restaurant_contracts "sharpnr-engine/modules/hospitality/restaurant/contracts"
	restaurantrepo "sharpnr-engine/modules/hospitality/restaurant/repo"
	payment_engine_service "sharpnr-engine/payment_engine/services"
	"sharpnr-engine/sero"
	invoice_engine_services "sharpnr-engine/service_engines/invoicing_engine/services"
)

type restaurantService struct {
	restaurantRepo restaurantrepo.RestaurantRepositoryInterface
	seroEngine     *sero.SEROEngine
	txManager      dbtx.TxManager
	customerEngine org_customer_service.OrgCustomerServiceInterface
	invoiceEngine  invoice_engine_services.InvoiceEngineServiceInterface
	paymentsEngine payment_engine_service.PaymentServiceInterface
}

// CloseDineInSession implements [RestaurantServiceInterface].
func (r *restaurantService) CloseDineInSession(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, restaurantID uuid.UUID, dineInSessionID uuid.UUID, request restaurant_contracts.CloseSessionRequest) (*restaurant_contracts.CloseSessionResponse, error) {
	panic("unimplemented")
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
	paymentsEngine payment_engine_service.PaymentServiceInterface,
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
