package v1

import (
	authcontrollers "sharpnr-engine/core/controllers/auth_controllers"
	signupcontrollers "sharpnr-engine/core/controllers/signup_controllers"
	organizationrepositories "sharpnr-engine/core/repositories/organization_repo"
	authservice "sharpnr-engine/core/services/auth_service"
	signupservice "sharpnr-engine/core/services/signup_service"
	"sharpnr-engine/middlewares"
	barrepo "sharpnr-engine/modules/hospitality/bar/repo"
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/repo"
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/kitchenRepo"
	menurepo "sharpnr-engine/modules/hospitality/menu/menu_repo"
	"sharpnr-engine/modules/hospitality/restaurant/restaurantrepo"
	hr_employee_repo "sharpnr-engine/modules/hr/employees/repo"

	"sharpnr-engine/pkg/jwtpack"
	policyengine "sharpnr-engine/policy_engine"
	sharpnruirenderer "sharpnr-engine/sharpnr_ui_renderer"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// These Are the Auth routes for the authentication registration
func AuthRoutes(v1 *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	newPolicyEngineRepo := policyengine.NewOrgPolicyRepo(db)

	newHotelRepo := hotelrepo.NewHotelRepository(db)
	newBarRepo := barrepo.NewBarRepository(db)
	newRestaurantRepo := restaurantrepo.NewRestaurantRepository(db)
	newKitchenRepo := kitchenrepo.NewKitchenRepository(db)
	newMenuRepo := menurepo.NewHospitalityMenuRepository(db)

	newUIRendererRepo := sharpnruirenderer.NewOrgUIRendererRepo(db)
	newHrEmployeeRepo := hr_employee_repo.NewHrEmployeeRepository(db)

	newOrgRepo := organizationrepositories.NewOrganizationRepo(db, redis)
	newSignupService := signupservice.NewSignupService(newOrgRepo, newHrEmployeeRepo, newPolicyEngineRepo, newUIRendererRepo, newHotelRepo, newBarRepo, newRestaurantRepo, newKitchenRepo, newMenuRepo)
	newSignupController := signupcontrollers.NewSignUpController(newSignupService)

	newAuthService := authservice.NewAuthService(newOrgRepo, newHrEmployeeRepo)
	newAuthController := authcontrollers.NewAuthController(newAuthService)

	//Register
	v1.POST("/register", newSignupController.RegisterOrg)
	//Need Middleware

	v1.POST("/otp-verify", middlewares.RequireTokenType(jwtpack.TokenTypeVerifyOTP), newSignupController.VerifyRegistrationOTP)
	v1.POST("/complete-register", middlewares.RequireTokenType(jwtpack.TokenTypeCompleteRegistration), newSignupController.CompleteRegistration)

	// Here we will gonna implement the login
	v1.POST("/login", newAuthController.Login)
	v1.GET("/check-cookies", newAuthController.CheckCookies)

	v1.POST("/reset-password", newAuthController.ResetPassword)

	v1.GET("/signup/status", newSignupController.SignupStatus)

	v1.GET("/get-org-render-tree", newAuthController.GetOrgUIRenderer)

	v1.GET("/sse/policy/:employeeId", middlewares.RequireTokenType(jwtpack.TokenTypeAuth),
		policyengine.SSEEmployeeMiddleware(),
		policyengine.PolicySSEHandler(newPolicyEngineRepo))

}
