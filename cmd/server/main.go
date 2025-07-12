package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sharpnr-inc/sharpnr-engine/config"
	_ "github.com/sharpnr-inc/sharpnr-engine/docs"

	"github.com/sharpnr-inc/sharpnr-engine/core/repositories"
	"github.com/sharpnr-inc/sharpnr-engine/core/services"
	orgservices "github.com/sharpnr-inc/sharpnr-engine/core/services/orgnisation"
	"github.com/sharpnr-inc/sharpnr-engine/modules"
	base_modules "github.com/sharpnr-inc/sharpnr-engine/modules/base"
	hotel_modules "github.com/sharpnr-inc/sharpnr-engine/modules/hotel"

	"github.com/sharpnr-inc/sharpnr-engine/shared/database"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Import controllers to register routes
	"github.com/sharpnr-inc/sharpnr-engine/api/v1"
	"github.com/sharpnr-inc/sharpnr-engine/controllers/organisationControllers"
)

func main() {
	config.LoadEnv()           // Load env vars like DB_HOST, DB_USER, etc.
	db := database.ConnectDB() // Connect to DB

	// Run migrations
	err := database.MigrateDB(db)
	if err != nil {
		log.Fatalf("DB migration failed: %v", err)
	}
	log.Println("✅ Database migration completed.")

	// Initialize Repositories
	orgRepo := repositories.NewOrganizationRepository(db)
	deptRepo := repositories.NewDepartmentRepository(db)
	termRepo := repositories.NewTerminalRepository(db)

	// Initialize Services
	deptService := services.NewDepartmentService(deptRepo)
	termService := services.NewTerminalService(termRepo)

	// Initialize Module Manager
	moduleManager := modules.NewModuleManager()

	// Register Base Module Initializer
	baseModuleInitializer := base_modules.NewBaseModule(deptService, termService)
	moduleManager.RegisterBaseInitializer(baseModuleInitializer)

	// Register Hotel Module Initializer
	hotelModuleInitializer := hotel_modules.NewHotelModule()
	moduleManager.RegisterModuleInitializer("HOTEL", hotelModuleInitializer) // Assuming "HOTEL" is the constant for hotel type

	// Initialize Organization Registration Service with Module Manager
	orgRegisterService := orgservices.NewOrganizationRegistrationService(orgRepo, moduleManager)

	// Create Gin router
	r := gin.Default()

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register API routes
	v1.RegisterV1Routes(r, orgRegisterService) // Pass the service to the route registration

	log.Println("App started...")
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}