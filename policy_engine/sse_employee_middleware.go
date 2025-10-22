package policyengine

import (
	"sharpnr-engine/middlewares"

	"github.com/gin-gonic/gin"
)

func SSEEmployeeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenEmployeeID, exists := c.Get(middlewares.UserIDContextKey)
		if !exists {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden: no user in context"})
			return
		}

		reqEmployeeID := c.Param("employeeId")
		if tokenEmployeeID != reqEmployeeID {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden: cannot access another employee policy"})
			return
		}

		c.Next()
	}
}
