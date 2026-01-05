package api

import (
	v1 "sharpnr-engine/api/v1"
	"sharpnr-engine/cookieManager"
	sharpnr_identity_auth "sharpnr-engine/core/sharpnrIdentityAuth"
	"sharpnr-engine/middlewares"
	"sharpnr-engine/sere"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// SetupRouter creates and configures the main Gin router.
// It delegates the registration of specific API versions to their respective packages.
func SetupRouter(
	db *gorm.DB,
	redis *redis.Client,
	sere *sere.SERE,
	cookies cookieManager.Manager,
) *gin.Engine {

	router := gin.Default()

	// ✅ ONE CORS middleware for everything
	router.Use(middlewares.UnifiedCORSMiddleware())

	// ========================
	// 🌐 WEB API (cookies)
	// ========================
	webV1 := router.Group("/api/v1")
	{
		// This will be replaced
		v1.SignUpRoutes(webV1, db, redis, sere, cookies)

		sharpnr_identity_auth.SharpnrIdentityAuthWebRoutes(webV1, db, redis, sere, cookies)

		v1.AuthRoutes(webV1, db, redis, sere)

		v1.RegisterModuleRoutes(webV1, db, sere)

		v1.UtilsRoutes(webV1)
	}

	// ========================
	// 📱 MOBILE API (tokens)
	// ========================
	mobileV1 := router.Group("/api/v1/mobile")
	{
		// This will be Replaced
		v1.MobileAuthRoutes(mobileV1, db, redis, sere)

		sharpnr_identity_auth.SharpnrIdentityAuthMobileRoutes(mobileV1, db, redis, sere)
	}

	return router
}
