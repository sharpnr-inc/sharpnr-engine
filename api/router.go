package api

import (
	v1 "sharpnr-engine/api/v1"
	"sharpnr-engine/app"
	"sharpnr-engine/config"
	"sharpnr-engine/cookieManager"
	dbtx "sharpnr-engine/core/dbTx"
	legal_entity_routes "sharpnr-engine/core/legal_entity/routes"
	org_signup "sharpnr-engine/core/org_signup"
	sharpnrIdentity_routes "sharpnr-engine/core/sharpnr_identity/routes"
	"sharpnr-engine/middlewares"
	webMiddleware "sharpnr-engine/middlewares/web"
	orchestrator_routes "sharpnr-engine/orchestrator/routes"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"

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

	router := gin.Default()

	// ✅ ONE CORS middleware for everything
	router.Use(middlewares.UnifiedCORSMiddleware(corsConfig))

	// ========================
	// 🌐 WEB API (cookies)
	// ========================
	webV1 := router.Group("/api/v1")
	{
		org_signup.OrgSignUpRoute(webV1, cookies, jwtmgr, modules)

		sharpnrIdentity_routes.RegisterSharpnrIdentityWebRoutes(webV1, cookies, jwtmgr, modules)

		protected := webV1.Group("/")
		protected.Use(
			webMiddleware.WebAuthMiddleware(jwtmgr, cookies),
			webMiddleware.CSRFMiddleware(cookies),
		)

		{
			v1.RegisterModuleRoutes(protected, modules)
			orchestrator_routes.RegisterOrchestratorRoutes(protected, txManager, cookies, jwtmgr, modules)
			legal_entity_routes.RegisterLegalEntityRoutes(protected, modules)
			sharpnrIdentity_routes.RegisterSharpnrIdentityProtectedWebRoutes(protected, cookies, jwtmgr, modules)
		}
	}

	// ========================
	// 📱 MOBILE API (tokens)
	// ========================
	mobileV1 := router.Group("/api/v1/mobile")
	{
		sharpnrIdentity_routes.RegisterSharpnrIdentityMobileRoutes(mobileV1, db, redis)
	}

	return router
}
