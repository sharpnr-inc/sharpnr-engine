package app

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	org_signup_repositories "sharpnr-engine/core/org-signup/repositories"
	org_signup_services "sharpnr-engine/core/org-signup/services"
	sharpnr_identity_auth_services "sharpnr-engine/core/sharpnr-identity/services"
	"sharpnr-engine/modules/accounting/accounting_service"
	barservices "sharpnr-engine/modules/hospitality/bar/services"
	hotelservices "sharpnr-engine/modules/hospitality/hotel/services"
	kitchenservices "sharpnr-engine/modules/hospitality/kitchen/kitchenServices"
	menuservice "sharpnr-engine/modules/hospitality/menu/menu_service"
	"sharpnr-engine/modules/hospitality/restaurant/restaurantservices"
	hr_employee_repo "sharpnr-engine/modules/hr/employees/repo"
	hr_employee_services "sharpnr-engine/modules/hr/employees/services"
	"sharpnr-engine/sere"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AppModules struct {
	orgSignupService org_signup_services.OrgSignUpServiceInterface
	sharpAuthService sharpnr_identity_auth_services.SharpnrIdentityAuthServiceInterface

	hospitality_hotelService      hotelservices.HotelService
	hospitality_barService        barservices.BarServices
	hospitality_restaurantService restaurantservices.RestaurantService
	hospitality_kitchenService    kitchenservices.KitchenServiceInterface
	hospitality_menuService       menuservice.MenuService

	hrService         hr_employee_services.HrEmployeeServiceInterface
	accountingService accounting_service.AccountingServiceInterface
}

// AccountingService implements ModulesProvider.
func (m *AppModules) AccountingService() accounting_service.AccountingServiceInterface {
	return m.accountingService
}

// HospitalityBarService implements ModulesProvider.
func (m *AppModules) HospitalityBarService() barservices.BarServices {
	return m.hospitality_barService
}

// HospitalityHotelService implements ModulesProvider.
func (m *AppModules) HospitalityHotelService() hotelservices.HotelService {
	return m.hospitality_hotelService
}

// HospitalityKitchenService implements ModulesProvider.
func (m *AppModules) HospitalityKitchenService() kitchenservices.KitchenServiceInterface {
	return m.hospitality_kitchenService
}

// HospitalityMenuService implements ModulesProvider.
func (m *AppModules) HospitalityMenuService() menuservice.MenuService {
	return m.hospitality_menuService
}

// HospitalityRestaurantService implements ModulesProvider.
func (m *AppModules) HospitalityRestaurantService() restaurantservices.RestaurantService {
	return m.hospitality_restaurantService
}

// SharpAuthService implements ModulesProvider.
func (m *AppModules) SharpAuthService() sharpnr_identity_auth_services.SharpnrIdentityAuthServiceInterface {
	return m.sharpAuthService
}

func (m *AppModules) HrService() hr_employee_services.HrEmployeeServiceInterface {
	return m.hrService
}

func (m *AppModules) OrgSignUpService() org_signup_services.OrgSignUpServiceInterface {
	return m.orgSignupService
}

func BootstrapModules(
	db *gorm.DB,
	redis *redis.Client,
	txManager dbtx.TxManager,
	sereEngine *sere.SERE,
) *AppModules {

	tempOrgSignUpRepo := org_signup_repositories.NewTempOrgSignUpRepository(redis)
	mainOrgSignUpRepo := org_signup_repositories.NewOrgMainSignUpRepository(db)
	orgSignUpService := org_signup_services.NewOrgSignUpService(txManager, sereEngine, tempOrgSignUpRepo, mainOrgSignUpRepo)

	hrEmployeeRepo := hr_employee_repo.NewHrEmployeeRepository(db)
	hrEmployeeService := hr_employee_services.NewHrEmployeeService(hrEmployeeRepo)

	sereEngine.UseTxExecutor(func(ctx context.Context, fn func(ctx context.Context) error) error {
		return txManager.WithTx(ctx, func(tx *gorm.DB) error {
			txCtx := dbtx.InjectTx(ctx, tx)
			return fn(txCtx)
		})
	})

	RegisterAllWorkFlow()
	// RegisterSereEventHandler(sereEngine,)

	return &AppModules{
		orgSignupService: orgSignUpService,
		hrService:        hrEmployeeService,
	}

}
