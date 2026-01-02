package api

import (
	v1 "sharpnr-engine/api/v1"
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
) *gin.Engine {

	router := gin.Default()

	// ✅ ONE CORS middleware for everything
	router.Use(middlewares.UnifiedCORSMiddleware())

	// ========================
	// 🌐 WEB API (cookies)
	// ========================
	webV1 := router.Group("/api/v1")
	{
		v1.SignUpRoutes(webV1, db, redis, sere)

		v1.AuthRoutes(webV1, db, redis, sere)

		v1.RegisterModuleRoutes(webV1, db, sere)

		v1.UtilsRoutes(webV1)
	}

	// ========================
	// 📱 MOBILE API (tokens)
	// ========================
	mobileV1 := router.Group("/api/v1/mobile")
	{
		v1.MobileAuthRoutes(mobileV1, db, redis, sere)
	}

	return router
}
