package middlewares

import (
	"sharpnr-engine/config"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func UnifiedCORSMiddleware(cfg config.CORSConfig) gin.HandlerFunc {
	webCORS := cors.New(cors.Config{
		AllowOrigins:     cfg.WebOrgin,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
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

		c.Next()
	}
}
