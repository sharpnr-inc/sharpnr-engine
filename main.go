package main

// Main entry point of the application
import (
	"context"
	"log"

	"sharpnr-engine/api"
	"sharpnr-engine/app"
	"sharpnr-engine/config"
	"sharpnr-engine/meta"
	"sharpnr-engine/pkg/cookieManager"
	dbtx "sharpnr-engine/pkg/dbTx"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"
	"sharpnr-engine/sero"
	"sharpnr-engine/shared/database"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"
	sse_engine "sharpnr-engine/sse_engine"

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

	log.Printf("🏷️  Sharpnr Engine %s", meta.Get())

	cfg := config.Load()
	redisDb := database.ConnectRedis(cfg.Redis)
	db := database.ConnectDB(cfg.Database)

	err := database.MigrateDB(cfg.Database)
	if err != nil {
		log.Fatalf("DB migration failed: %v", err)
	}

	log.Println("✅ Database migration completed.")

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
	if err := mailerEng.LoadTemplates("./email_templates"); err != nil {
		log.Fatal(err)
	}

	mailerEng.Start()
	defer mailerEng.Shutdown()

	/* ---------------------------------- <***> --------------------------------- */

	cookieMgr := cookieManager.New(cfg.Cookies)

	/* ---------------------------------- <***> --------------------------------- */

	// ToManagerConfig() converts config.JwtManagerConfigs → jwt_manager.Config
	// keeping the two packages independent of each other.
	// New() validates eagerly — missing secret or issuer fails here at startup,
	// not silently during the first token operation.
	jwtmgr, err := jwt_manager.New(cfg.JWT.ToManagerConfig())
	if err != nil {
		log.Fatalf("failed to init jwt manager: %v", err)
	}
	/* ---------------------------------- <***> --------------------------------- */

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
