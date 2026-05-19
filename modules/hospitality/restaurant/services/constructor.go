package restaurantservices

import (
	"context"
	org_customer_service "sharpnr-engine/core/customer/service"
	menuservice "sharpnr-engine/modules/hospitality/menu/service"
	restaurantrepo "sharpnr-engine/modules/hospitality/restaurant/repo"
	payment_engine_service "sharpnr-engine/payment_engine/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
	inventory_engine_services "sharpnr-engine/service_engines/inventory_engine/services"
	invoice_engine_services "sharpnr-engine/service_engines/invoicing_engine/services"
)

type restaurantService struct {
	restaurantRepo  restaurantrepo.RestaurantRepositoryInterface
	menuService     menuservice.MenuService
	seroEngine      *sero.SEROEngine
	txManager       dbtx.TxManager
	customerEngine  org_customer_service.OrgCustomerServiceInterface
	invoiceEngine   invoice_engine_services.InvoiceEngineServiceInterface
	paymentsEngine  payment_engine_service.PaymentServiceInterface
	inventoryEngine inventory_engine_services.InventoryEngineServicesInterface
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
	inventoryEngine inventory_engine_services.InventoryEngineServicesInterface,
) *restaurantService {
	return &restaurantService{
		restaurantRepo:  restaurantRepo,
		menuService:     menuService,
		seroEngine:      seroEngine,
		txManager:       txManager,
		customerEngine:  customerEngine,
		invoiceEngine:   invoiceEngine,
		paymentsEngine:  paymentsEngine,
		inventoryEngine: inventoryEngine,
	}

}
