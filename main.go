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
	ws_engine "sharpnr-engine/ws_engine"

	_ "sharpnr-engine/docs"

	"github.com/gin-gonic/gin"
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

	// Must run before any gin engine is constructed. AppConfig.GinMode was
	// being computed and never applied, so production ran in debug mode —
	// verbose per-request logging and debug warnings included.
	gin.SetMode(cfg.App.GinMode)

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
	defer sseEngine.Close()

	// The socket's origin allow-list is the CORS list: a WebSocket handshake
	// carries the session cookie but is NOT subject to CORS, so this is the
	// only thing stopping another origin from opening a socket as a logged-in
	// user. Sharing the list keeps the two from drifting apart.
	wsEngine := ws_engine.New(ws_engine.Config{AllowedOrigins: cfg.Cors.WebOrgin})
	defer wsEngine.Close()

	sharpnrModules := app.BootstrapModules(db, redisDb, txManager, mailerEng, seroEngine, sseEngine, wsEngine, jwtmgr, cfg.TTL)

	// Mobile refresh sessions are rows without a TTL, so something has to
	// reclaim them. Shares the SERO relay's context, so it stops on shutdown.
	go app.StartExpiredSessionSweeper(ctx, sharpnrModules.AuthenticationService())
	log.Println("🧹 Expired mobile session sweeper started.")

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
