package app

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	hospitality_org_factory_repository "sharpnr-engine/core/hospitality-org-factory/repository"
	org_signup_repositories "sharpnr-engine/core/org-signup/repositories"
	org_signup_services "sharpnr-engine/core/org-signup/services"
	sharpnr_identity_auth_repositories "sharpnr-engine/core/sharpnr-identity/repositories"
	sharpnr_identity_auth_services "sharpnr-engine/core/sharpnr-identity/services"
	"sharpnr-engine/modules/accounting/accounting_repo"
	"sharpnr-engine/modules/accounting/accounting_service"
	barrepo "sharpnr-engine/modules/hospitality/bar/repo"
	barservices "sharpnr-engine/modules/hospitality/bar/services"
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/hotel_repo"
	hotelservices "sharpnr-engine/modules/hospitality/hotel/services"
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/kitchenRepo"
	kitchenservices "sharpnr-engine/modules/hospitality/kitchen/kitchenServices"
	menurepo "sharpnr-engine/modules/hospitality/menu/menu_repo"
	menuservice "sharpnr-engine/modules/hospitality/menu/menu_service"
	"sharpnr-engine/modules/hospitality/restaurant/restaurantrepo"
	"sharpnr-engine/modules/hospitality/restaurant/restaurantservices"
	hr_employee_repo "sharpnr-engine/modules/hr/employees/repo"
	hr_employee_services "sharpnr-engine/modules/hr/employees/services"
	"sharpnr-engine/sere"
	sharp_mailer_engine "sharpnr-engine/sharp-mailer-engine"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func BootstrapModules(
	db *gorm.DB,
	redis *redis.Client,
	txManager dbtx.TxManager,
	sereEngine *sere.SERE,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
) *AppModules {

	tempOrgSignUpRepo := org_signup_repositories.NewTempOrgSignUpRepository(redis)
	mainOrgSignUpRepo := org_signup_repositories.NewOrgMainSignUpRepository(db)
	orgSignUpService := org_signup_services.NewOrgSignUpService(txManager, sereEngine, tempOrgSignUpRepo, mainOrgSignUpRepo, mailerEngine)

	hospitalityOrgFactoryRepository := hospitality_org_factory_repository.NewHospitalityOrgFactoryRepository(db)

	// Hospitality kitchen
	hospitality_kitchenRepo := kitchenrepo.NewKitchenRepository(db)
	hospitality_kitchenService := kitchenservices.NewKitchenService(hospitality_kitchenRepo)

	// hospitality Menu
	newMeneRepo := menurepo.NewHospitalityMenuRepository(db)
	newMenuService := menuservice.NewHospitalityMenuService(newMeneRepo)

	// hospitality Hotel
	hospitality_hotelRepo := hotelrepo.NewHotelRepository(db)
	hospitality_hotelService := hotelservices.NewHotelService(sereEngine, hospitality_hotelRepo)

	// hospitality Bar
	hospitality_barRepo := barrepo.NewBarRepository(db)
	hospitality_barService := barservices.NewBarServices(hospitality_barRepo)

	// hospitality Restaurant
	hospitality_restaurantRepo := restaurantrepo.NewRestaurantRepository(db)
	hospitality_restaurantService := restaurantservices.NewRestaurantService(hospitality_restaurantRepo)

	accountingRepo := accounting_repo.NewAccountingRepository(db)
	accountingService := accounting_service.NewAccountingService(accountingRepo)

	empIdentityRepo := hr_employee_repo.NewEmployeeIdentityRepository(db)
	empAuthRepo := hr_employee_repo.NewEmployeeAuthRepository(db)
	empOnboardingRepo := hr_employee_repo.NewEmployeeOnboardingRepository(db)
	empVerificationRepo := hr_employee_repo.NewEmployeeVerificationRepo(db)
	empProfileRepo := hr_employee_repo.NewEmployeeProfileRepo(db)
	empEmploymentRepo := hr_employee_repo.NewEmployeeEmploymentRepo(db)

	policyRepo := hr_employee_repo.NewEmployeePolicyRepo(db)
	policyRuleRepo := hr_employee_repo.NewEmployeePolicyRuleRepository(db)
	privilegePolicyRepo := hr_employee_repo.NewEmployeePrivilegePolicyRepository(db)

	hrEmployeeService := hr_employee_services.NewHrEmployeeService(
		empIdentityRepo,
		empAuthRepo,
		empOnboardingRepo,
		empVerificationRepo,
		empProfileRepo,
		empEmploymentRepo,

		// Policy
		policyRepo,
		policyRuleRepo,
		privilegePolicyRepo,
	)

	sharpnrIdentityRepo := sharpnr_identity_auth_repositories.NewSharpnrIdentityAuthRepository(db)
	sharpnrIdentityService := sharpnr_identity_auth_services.NewSharpnrIdentityAuthService(
		empIdentityRepo,
		empAuthRepo,
		sharpnrIdentityRepo,
	)

	sereEngine.UseTxExecutor(func(ctx context.Context, fn func(ctx context.Context) error) error {
		return txManager.WithTx(ctx, func(tx *gorm.DB) error {
			txCtx := dbtx.InjectTx(ctx, tx)
			return fn(txCtx)
		})
	})

	RegisterAllWorkFlow(
		sereEngine,
		hrEmployeeService,
		sharpnrIdentityService,
		hospitality_kitchenRepo,
		hospitality_hotelRepo,
		hospitality_hotelService,

		hospitality_barRepo,
		hospitality_barService,

		hospitality_restaurantRepo,
		hospitality_restaurantService,

		newMeneRepo,
		newMenuService,

		accountingRepo,
		accountingService,
		hospitalityOrgFactoryRepository,
	)

	RegisterSereEventHandler(sereEngine, accountingRepo)

	return &AppModules{
		orgSignupService: orgSignUpService,

		hrService: hrEmployeeService,

		hospitality_hotelService:      hospitality_hotelService,
		hospitality_barService:        hospitality_barService,
		hospitality_restaurantService: hospitality_restaurantService,
		hospitality_kitchenService:    hospitality_kitchenService,
		hospitality_menuService:       newMenuService,
		accountingService:             accountingService,
		sharpAuthService:              sharpnrIdentityService,
	}

}
