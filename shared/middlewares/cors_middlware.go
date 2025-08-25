package middlewares

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleWare() gin.HandlerFunc {
	env := os.Getenv("APP_ENV")

	var allowedOrigins []string
	if env == "prodcution" {
		allowedOrigins = []string{
			"some",
			"thing",
		}
	} else {
		allowedOrigins = []string{"*"}
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
