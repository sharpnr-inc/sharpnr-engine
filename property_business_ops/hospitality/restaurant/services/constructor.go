package restaurantservices

import (
	"context"
	org_customer_service "sharpnr-engine/core/customer/service"
	payment_engine_service "sharpnr-engine/payment_engine/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	restaurantrepo "sharpnr-engine/property_business_ops/hospitality/restaurant/repo"
	menuservice "sharpnr-engine/property_resources/hospitality_resources/menu/service"
	"sharpnr-engine/sero"
	stockwaveEngineServices "sharpnr-engine/stockwave_engine/services"

	invoice_engine_services "sharpnr-engine/sharp_accounting_engine/invoicing_engine/services"
)

type restaurantService struct {
	restaurantRepo         restaurantrepo.RestaurantRepositoryInterface
	menuService            menuservice.MenuService
	seroEngine             *sero.SEROEngine
	txManager              dbtx.TxManager
	customerEngine         org_customer_service.OrgCustomerServiceInterface
	invoiceEngine          invoice_engine_services.InvoiceEngineServiceInterface
	paymentsEngine         payment_engine_service.PaymentServiceInterface
	stockwaveEngineService stockwaveEngineServices.StockwaveEngineServicesInterface
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
	menuService menuservice.MenuService,
	seroEngine *sero.SEROEngine,
	txManager dbtx.TxManager,
	customerEngine org_customer_service.OrgCustomerServiceInterface,
	invoiceEngine invoice_engine_services.InvoiceEngineServiceInterface,
	paymentsEngine payment_engine_service.PaymentServiceInterface,
	stockwaveEngineService stockwaveEngineServices.StockwaveEngineServicesInterface,
) *restaurantService {
	return &restaurantService{
		restaurantRepo:         restaurantRepo,
		menuService:            menuService,
		seroEngine:             seroEngine,
		txManager:              txManager,
		customerEngine:         customerEngine,
		invoiceEngine:          invoiceEngine,
		paymentsEngine:         paymentsEngine,
		stockwaveEngineService: stockwaveEngineService,
	}

}
