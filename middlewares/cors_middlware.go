package middlewares

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleWare() gin.HandlerFunc {
	env := os.Getenv("APP_ENV")

	var allowedOrigins []string
	if env == "production" {
		allowedOrigins = []string{
			"http://localhost:5173",
			"http://10.60.106.1:5173",
		}
	} else {
		allowedOrigins = []string{"http://localhost:5173", "http://10.60.106.1:5173"}
	}

	corsConfig := cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	return cors.New(corsConfig)

}
