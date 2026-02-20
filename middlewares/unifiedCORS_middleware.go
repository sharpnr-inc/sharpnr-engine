package middlewares

import (
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func UnifiedCORSMiddleware() gin.HandlerFunc {
	webCORS := cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"http://192.168.1.33:5173",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-CSRF-TOKEN"},
		AllowCredentials: true,
	})

	mobileCORS := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: false,
	})

	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api/v1/mobile") {
			mobileCORS(c)
		} else {
			webCORS(c)
		}

		// Explicitly end preflight
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
