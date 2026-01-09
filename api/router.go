package api

import (
	v1 "sharpnr-engine/api/v1"
	"sharpnr-engine/app"
	"sharpnr-engine/cookieManager"
	dbtx "sharpnr-engine/core/dbTx"
	org_signup "sharpnr-engine/core/org-signup"
	sharpnr_identity_auth "sharpnr-engine/core/sharpnr-identity"
	"sharpnr-engine/middlewares"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"
	"sharpnr-engine/sere"

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
	sere *sere.SERE,
	cookies cookieManager.Manager,
	jwtmgr jwt_manager.Manager,
	modules app.ModulesProvider,

) *gin.Engine {

	router := gin.Default()

	// ✅ ONE CORS middleware for everything
	router.Use(middlewares.UnifiedCORSMiddleware())

	// ========================
	// 🌐 WEB API (cookies)
	// ========================
	webV1 := router.Group("/api/v1")
	{
		org_signup.OrgSignUpRoute(webV1, cookies, jwtmgr, modules)

		sharpnr_identity_auth.SharpnrIdentityAuthWebRoutes(webV1, cookies, modules)

		v1.RegisterModuleRoutes(webV1, modules)

		v1.UtilsRoutes(webV1)
	}

	// ========================
	// 📱 MOBILE API (tokens)
	// ========================
	mobileV1 := router.Group("/api/v1/mobile")
	{
		sharpnr_identity_auth.SharpnrIdentityAuthMobileRoutes(mobileV1, db, redis, sere)
	}

	return router
}
