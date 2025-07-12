package api

import "github.com/gin-gonic/gin"

// SetupRouter creates and configures the main Gin router.
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Add middleware here

	// Register v1 routes
	// v1.RegisterRoutes(router)

	return router
}
