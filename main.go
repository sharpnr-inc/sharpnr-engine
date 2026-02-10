package main

// Main entry point of the application
import (
	"log"

	"sharpnr-engine/api"
	"sharpnr-engine/app"
	"sharpnr-engine/config"
	"sharpnr-engine/cookieManager"
	dbtx "sharpnr-engine/core/dbTx"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"
	"sharpnr-engine/sere"
	"sharpnr-engine/shared/database"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"

	_ "sharpnr-engine/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Sharpnr Engine APIs
// @version 1.0
// @description Sharpnr Engine API Documentation
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.sharpnr.com/support
// @contact.email support@sharpnr.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host 192.168.1.35:8080
// @BasePath /api/v1
func main() {
	cfg := config.Load()

	redisDb := database.ConnectRedis(cfg.Redis)
	db := database.ConnectDB(cfg.Database)

	// Run migrations
	err := database.MigrateDB(db)
	if err != nil {
		log.Fatalf("DB migration failed: %v", err)
	}
	log.Println("✅ Database migration completed.")

	sereEngine := sere.NewSERE()
	sereEngine.Start()
	sereEngine.WaitUntilStarted()

	log.Println("✅ SERE is up and running")

	mailerEng := sharp_mailer_engine.New(cfg)
	if err := mailerEng.LoadTemplates("./sharp_mailer_engine/templates"); err != nil {
		log.Fatal(err)
	}
	mailerEng.Start()
	defer mailerEng.Shutdown()

	defs := config.CookieRegistry(cfg.App, cfg.TTL)
	cookieManager.Init(defs)
	cookieMgr := cookieManager.New(cookieManager.Config{
		Secure: cfg.Cookie.Secure,
		Domain: cfg.Cookie.Domain,
	})

	jwtmgr := jwt_manager.New(jwt_manager.Config{
		Secret:            cfg.JWT.Secret,
		Issuer:            cfg.JWT.Issuer,
		Audience:          cfg.JWT.Audience,
		AccessTTL:         cfg.JWT.AccessTTL,
		RefreshTTL:        cfg.JWT.RefreshTTL,
		SignupTTL:         cfg.JWT.SignupTTL,
		WebCookieTokenTTL: cfg.JWT.WebCookieTokenTTL,
	})

	txManager := dbtx.NewGormTxManager(db)

	sharpnrModules := app.BootstrapModules(db, redisDb, txManager, sereEngine, mailerEng)

	// Create and configure the main router
	r := api.SetupRouter(
		db,
		txManager,
		redisDb,
		sereEngine,
		cookieMgr,
		jwtmgr,
		sharpnrModules,
	)

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Get local IP
	localIP, err := config.GetLocalIP()
	if err != nil || localIP == "" {
		log.Println("⚠️ Could not determine local IP, using 0.0.0.0")
		localIP = "0.0.0.0"
	}

	port := "8080"
	address := localIP + ":" + port
	log.Printf("App started on http://%s ...", address)

	// Run Gin server on detected local IP
	if err := r.Run(address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
