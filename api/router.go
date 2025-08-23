package api

import (
	v1 "sharpnr-engine/api/v1"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// SetupRouter creates and configures the main Gin router.
// It delegates the registration of specific API versions to their respective packages.
func SetupRouter(db *gorm.DB, redis *redis.Client) *gin.Engine {
	router := gin.Default()

	// Create a router group for the v1 API.
	v1Group := router.Group("/api/v1")

	// Register all V1 routes.
	v1.CoreV1Routes(v1Group, db, redis)
	v1.DepartmentsRoutes(v1Group, db, redis)
	v1.RegisterModuleRoutes(v1Group, db)

	// If you had a v2, you would register it here:
	// v2.RegisterV2Routes(router, db)

	v1.UtilsRoutes(v1Group)

	return router
}
