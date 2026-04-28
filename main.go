package main

// Main entry point of the application
import (
	"context"
	"log"

	"sharpnr-engine/api"
	"sharpnr-engine/app"
	"sharpnr-engine/config"
	"sharpnr-engine/cookieManager"
	dbtx "sharpnr-engine/core/dbTx"
	"sharpnr-engine/migrations"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"
	"sharpnr-engine/sero"
	"sharpnr-engine/shared/database"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"
	sse_engine "sharpnr-engine/sse"

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

// @host api.sharpnr.test
// @schemes http https
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

	err = migrations.RunDataMigrations(db)
	if err != nil {
		log.Fatalf("Data migration failed: %v", err)
	}

	/* ---------------------------------- <***> --------------------------------- */

	// --- Initialize SERO Engine ---
	seroRepo := sero.NewGormOutboxRepo(db)
	seroEngine := sero.NewSEROEngine(seroRepo)

	// Start background relay
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go seroEngine.StartRelay(ctx)
	log.Println("🚀 SERO Event Relay started.")

	/* ---------------------------------- <***> --------------------------------- */

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
	sseEngine := sse_engine.NewStreamManager()
	sharpnrModules := app.BootstrapModules(db, redisDb, txManager, mailerEng, seroEngine, sseEngine)

	// Create and configure the main router
	r := api.SetupRouter(
		db,
		txManager,
		redisDb,
		cookieMgr,
		jwtmgr,
		sharpnrModules,
		cfg.Cors,
	)
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := cfg.App.Port

	log.Printf("App started on :%s ...", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
