// routes/v1/api.go
package v1

import "github.com/gin-gonic/gin"

func RegisterV1Routes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	// Register each module's routes under /api/v1
	RegisterOrganizationRoutes(v1)
}
