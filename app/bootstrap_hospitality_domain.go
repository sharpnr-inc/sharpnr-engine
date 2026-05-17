package app

import (
	hospitalitydomain "sharpnr-engine/app/hospitality_domain"
	org_customer_service "sharpnr-engine/core/customer/service"
	barrepo "sharpnr-engine/modules/hospitality/bar/repo"
	barservices "sharpnr-engine/modules/hospitality/bar/services"
	hotelseroeventhandlers "sharpnr-engine/modules/hospitality/hotel/hotel_sero_event_handlers"
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/repo"
	hotelservices "sharpnr-engine/modules/hospitality/hotel/services"
	"sharpnr-engine/modules/hospitality/kitchen/kitchen_sero_event_handlers"
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/repos"
	kitchenservices "sharpnr-engine/modules/hospitality/kitchen/services"
	menurepo "sharpnr-engine/modules/hospitality/menu/repo"
	menuservice "sharpnr-engine/modules/hospitality/menu/service"
	orderEngineSero_event_handlers "sharpnr-engine/modules/hospitality/order_engine/order_engine_event_handlers"
	hospitality_order_engine_repository "sharpnr-engine/modules/hospitality/order_engine/repository"
	hospitality_order_engine_service "sharpnr-engine/modules/hospitality/order_engine/service"
	restaurant_event_handlers "sharpnr-engine/modules/hospitality/restaurant/event_handlers"
	restaurantrepo "sharpnr-engine/modules/hospitality/restaurant/repo"
	restaurantservices "sharpnr-engine/modules/hospitality/restaurant/services"
	payment_engine_service "sharpnr-engine/payment_engine/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
	inventory_engine_services "sharpnr-engine/service_engines/inventory_engine/services"
	invoice_engine_services "sharpnr-engine/service_engines/invoicing_engine/services"
	tax_engine_service "sharpnr-engine/service_engines/tax_engine/service"
	sse_engine "sharpnr-engine/sse"

	"gorm.io/gorm"
)

func bootstrapHospitality(
	db *gorm.DB,
	txManager dbtx.TxManager,
	seroEngine *sero.SEROEngine,
	sseEngine *sse_engine.StreamManager,
	taxEngine tax_engine_service.TaxEngineServiceInterface,
	customerEngineService org_customer_service.OrgCustomerServiceInterface,
	invoiceEngine invoice_engine_services.InvoiceEngineServiceInterface,
	paymentsEngine payment_engine_service.PaymentServiceInterface,
	inventoryEngine inventory_engine_services.InventoryEngineServicesInterface,
) *hospitalitydomain.HospitalityDomain {

	// Kitchen
	kitchenRepo := kitchenrepo.NewKitchenRepository(db)
	kitchenService := kitchenservices.NewKitchenService(kitchenRepo, seroEngine, sseEngine, txManager)
	kitchen_sero_event_handlers.Register(seroEngine, kitchenService)

	// Menu
	menuRepo := menurepo.NewHospitalityMenuRepository(db)
	menuService := menuservice.NewHospitalityMenuService(menuRepo, txManager, taxEngine)

	// order engine
	orderEnginerepo := hospitality_order_engine_repository.NewHospitalityOrderEngine(db)
	orderEngineService := hospitality_order_engine_service.NewHospitalityOrderEngineService(orderEnginerepo, seroEngine, menuService, kitchenService)
	orderEngineSero_event_handlers.Register(seroEngine, orderEngineService)

	// Hotel
	hotelRepo := hotelrepo.NewHotelRepository(db)
	hotelService := hotelservices.NewHotelService(hotelRepo, txManager)
	hotelseroeventhandlers.Register(seroEngine, hotelService)

	// Bar
	barRepo := barrepo.NewBarRepository(db)
	barService := barservices.NewBarServices(barRepo)

	// Restaurant
	restaurantRepo := restaurantrepo.NewRestaurantRepository(db)
	restaurantService := restaurantservices.NewRestaurantService(
		restaurantRepo,
		menuService,
		seroEngine,
		txManager,
		customerEngineService,
		invoiceEngine,
		paymentsEngine,
		inventoryEngine,
	)
	restaurant_event_handlers.Register(seroEngine, restaurantService)

	return hospitalitydomain.NewHospitalityDomain(
		hotelService,
		barService,
		restaurantService,
		kitchenService,
		menuService,
		orderEngineService,
	)
}
