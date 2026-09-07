package app

import (
	hospitalitydomain "sharpnr-engine/app/hospitality_domain"
	org_customer_service "sharpnr-engine/core/customer/service"
	payment_engine_service "sharpnr-engine/payment_engine/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	barrepo "sharpnr-engine/property_business_ops/hospitality/bar/repo"
	barservices "sharpnr-engine/property_business_ops/hospitality/bar/services"
	hotelseroeventhandlers "sharpnr-engine/property_business_ops/hospitality/hotel/hotel_sero_event_handlers"
	hotelrepo "sharpnr-engine/property_business_ops/hospitality/hotel/repo"
	hotelservices "sharpnr-engine/property_business_ops/hospitality/hotel/services"
	orderEngineSero_event_handlers "sharpnr-engine/property_business_ops/hospitality/order_engine/order_engine_event_handlers"
	hospitality_order_engine_repository "sharpnr-engine/property_business_ops/hospitality/order_engine/repository"
	hospitality_order_engine_service "sharpnr-engine/property_business_ops/hospitality/order_engine/service"
	restaurant_event_handlers "sharpnr-engine/property_business_ops/hospitality/restaurant/event_handlers"
	restaurantrepo "sharpnr-engine/property_business_ops/hospitality/restaurant/repo"
	restaurantservices "sharpnr-engine/property_business_ops/hospitality/restaurant/services"
	"sharpnr-engine/property_infra/hospitality_infra/kitchen/kitchen_sero_event_handlers"
	kitchenrepo "sharpnr-engine/property_infra/hospitality_infra/kitchen/repos"
	kitchenservices "sharpnr-engine/property_infra/hospitality_infra/kitchen/services"
	menurepo "sharpnr-engine/property_resources/hospitality_resources/menu/repo"
	menuservice "sharpnr-engine/property_resources/hospitality_resources/menu/service"
	"sharpnr-engine/sero"
	stockwaveEngineServices "sharpnr-engine/stockwave_engine/services"

	invoice_engine_services "sharpnr-engine/sharp_accounting_engine/invoicing_engine/services"
	sse_engine "sharpnr-engine/sse_engine"

	"gorm.io/gorm"
)

func bootstrapHospitality(
	db *gorm.DB,
	txManager dbtx.TxManager,
	seroEngine *sero.SEROEngine,
	sseEngine *sse_engine.StreamManager,
	customerEngineService org_customer_service.OrgCustomerServiceInterface,
	invoiceEngine invoice_engine_services.InvoiceEngineServiceInterface,
	paymentsEngine payment_engine_service.PaymentServiceInterface,
	inventoryEngine stockwaveEngineServices.StockwaveEngineServicesInterface,
) *hospitalitydomain.HospitalityDomain {

	// Kitchen
	kitchenRepo := kitchenrepo.NewKitchenRepository(db)
	kitchenService := kitchenservices.NewKitchenService(kitchenRepo, seroEngine, sseEngine, txManager)
	kitchen_sero_event_handlers.Register(seroEngine, kitchenService)

	// Menu
	menuRepo := menurepo.NewHospitalityMenuRepository(db)
	menuService := menuservice.NewHospitalityMenuService(menuRepo, txManager)

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
