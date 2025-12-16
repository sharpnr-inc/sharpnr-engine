package hospitalityinvoicingservices

import (
	barrepo "sharpnr-engine/modules/hospitality/bar/repo"
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/repo"
	"sharpnr-engine/modules/hospitality/restaurant/restaurantrepo"
	hospitalityinvoicingrepo "sharpnr-engine/modules/hospitalityInvoicing/hospitalityInvoicingRepo"
)

// hospitalityInvoicingService implements the business logic for generating,
// managing, and updating invoices across the hospitality domain.
//
// This service unifies invoicing for Restaurant, Bar, and Hotel modules,
// ensuring consistent tax calculation, subtotal/total generation, and invoice
// linkage across all hospitality services.
//
// Dependencies injected into this service include:
//   - hospInvoiceRepo:  Repository for invoice persistence (CRUD operations)
//   - hotelRepo:        Provides access to hotel bookings used for invoice generation
//   - barRepo:          Provides access to bar sessions for billing
//   - restaurantRepo:   Provides access to restaurant table sessions for billing
//
// By combining all three domain repositories, the invoicing service becomes
// the central authority for all hospitality bill generation, while keeping
// domain modules fully decoupled and independent.
type hospitalityInvoicingService struct {
	hospInvoiceRepo hospitalityinvoicingrepo.HospitalityInvoicingRepositoryInterfaces

	hotelRepo      hotelrepo.HotelRepositoryInterface
	barRepo        barrepo.BarRepository
	restaurantRepo restaurantrepo.RestaurantRepository
}

// NewHospitalityInvoicingService initializes and returns a new instance of
// hospitalityInvoicingService with all required dependencies.
//
// This constructor follows dependency injection principles, allowing the
// service to operate independently of specific database or storage layers.
// Each domain repository is injected to allow the invoicing service to fetch
// the necessary session/booking data for billing computations.
//
// Parameters:
//   - hospInvoiceRepo: Storage interface for invoice operations
//   - hotelRepo:       Hotel domain repository
//   - barRepo:         Bar domain repository
//   - restaurantRepo:  Restaurant domain repository
//
// Returns:
//
//	*hospitalityInvoicingService — a fully initialized invoicing service.
func NewHospitalityInvoicingService(
	hospInvoiceRepo hospitalityinvoicingrepo.HospitalityInvoicingRepositoryInterfaces,
	hotelRepo hotelrepo.HotelRepositoryInterface,
	barRepo barrepo.BarRepository,
	restaurantRepo restaurantrepo.RestaurantRepository) *hospitalityInvoicingService {

	return &hospitalityInvoicingService{
		hospInvoiceRepo: hospInvoiceRepo,
		hotelRepo:       hotelRepo,
		barRepo:         barRepo,
		restaurantRepo:  restaurantRepo,
	}
}
