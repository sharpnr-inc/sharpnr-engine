package api

import (
	"fmt"
	v1 "sharpnr-engine/api/v1"
	"sharpnr-engine/api/v1mobile"
	"sharpnr-engine/app"
	authenticationEngineRoutes "sharpnr-engine/authentication/routes"
	"sharpnr-engine/config"
	org_signup "sharpnr-engine/core/org_signup"
	iamEngineRoutes "sharpnr-engine/iam_engine/routes"
	"sharpnr-engine/meta"
	"sharpnr-engine/middlewares"
	mobile_middleware "sharpnr-engine/middlewares/mobile"
	webMiddleware "sharpnr-engine/middlewares/web"
	orchestrator_routes "sharpnr-engine/orchestrator/routes"
	"sharpnr-engine/pkg/cookieManager"
	dbtx "sharpnr-engine/pkg/dbTx"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"
	utilitiesApi_routes "sharpnr-engine/utilities_apis/routes"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// SetupRouter creates and configures the main Gin router.
// It delegates the registration of specific API versions to their respective packages.
func SetupRouter(
	db *gorm.DB,
	txManager dbtx.TxManager,
	redis *redis.Client,
	cookies cookieManager.Manager,
	jwtmgr jwt_manager.Manager,
	modules app.ModulesProvider,
	corsConfig config.CORSConfig,

) *gin.Engine {
	gin.ForceConsoleColor()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		methodColors := map[string]string{
			"GET":    "\033[1;32m", // bold green
			"POST":   "\033[1;34m", // bold blue
			"PUT":    "\033[1;33m", // bold yellow
			"PATCH":  "\033[1;35m", // bold magenta
			"DELETE": "\033[1;31m", // bold red
		}

		reset := "\033[0m"
		pathColor := "\033[1;36m"    // bold cyan  — the endpoint path
		handlerColor := "\033[2;37m" // dim white  — handler name (less important)
		arrowColor := "\033[1;90m"   // dark gray  — the --> arrow

		color, ok := methodColors[httpMethod]
		if !ok {
			color = "\033[1;37m"
		}

		fmt.Printf("[GIN-debug] %s%-7s%s %s%-60s%s %s-->%s %s%s%s \033[90m(%d handlers)%s\n",
			color, httpMethod, reset,
			pathColor, absolutePath, reset,
			arrowColor, reset,
			handlerColor, handlerName, reset,
			nuHandlers, reset,
		)
	}

	router := gin.Default()

	// ✅ ONE CORS middleware for everything
	router.Use(middlewares.UnifiedCORSMiddleware(corsConfig))

	// ========================
	// 🩺 Health & version (public, no auth)
	// ========================
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "version": meta.Version})
	})
	router.GET("/version", func(c *gin.Context) {
		c.JSON(200, meta.Get())
	})

	// ========================
	// 🌐 WEB API (cookies)
	// ========================
	webV1 := router.Group("/api/v1")
	{
		org_signup.OrgSignUpRoute(webV1, cookies, jwtmgr, modules)

		authenticationEngineRoutes.RegisterAuthenticationWebRoutes(webV1, cookies, jwtmgr, modules)

		protected := webV1.Group("/")
		protected.Use(
			webMiddleware.WebAuthMiddleware(jwtmgr, cookies),
			webMiddleware.CSRFMiddleware(cookies),
		)

		{
			v1.RegisterModuleRoutes(protected, modules)
			iamEngineRoutes.RegisterIamEngineRoutes(protected, cookies, modules)
			orchestrator_routes.RegisterOrchestratorRoutes(protected, txManager, cookies, jwtmgr, modules)
			authenticationEngineRoutes.RegisterAuthenticationProtectedWebRoutes(protected, cookies, jwtmgr, modules)
			utilitiesApi_routes.UtilitesApiGeoRoutes(protected, modules)
		}
	}

	// ========================
	// 📱 MOBILE API (tokens)
	// ========================
	mobileV1 := router.Group("/api/v1/mobile")
	{
		// Unprotected routes for the Login
		authenticationEngineRoutes.RegisterAuthenticationMobileRoutes(mobileV1, jwtmgr, modules, cookies)
		protected_mobile_routes := mobileV1.Group("/")

		// these are mobile routes for the other user
		protected_mobile_routes.Use(
			mobile_middleware.MobileMiddleware(jwtmgr),
		)

		orchestrator_routes.RegisterOrchestratorRoutes(protected_mobile_routes, txManager, cookies, jwtmgr, modules)
		v1mobile.RegisterMobilesModulesRoutes(protected_mobile_routes, modules)
		authenticationEngineRoutes.RegisterAuthenticationProtectedMobileRoutes(protected_mobile_routes, modules, cookies, jwtmgr)

	}

	return router
}
