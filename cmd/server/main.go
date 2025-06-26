package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sharpnr-inc/sharpnr-engine/config"
	"github.com/sharpnr-inc/sharpnr-engine/database"
	_ "github.com/sharpnr-inc/sharpnr-engine/docs"
	v1 "github.com/sharpnr-inc/sharpnr-engine/routes/v1"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title SharpNR API
// @version 1.0
// @description Backend APIs for SharpNR system.
// @termsOfService http://sharpnr.com/terms/

// @contact.name API Support
// @contact.email support@sharpnr.com

// @host localhost:8080
// @BasePath /api/v1
func main() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	config.LoadEnv()

	database.ConnectDB()
	log.Println("App started...")
	v1.RegisterV1Routes(r)
	r.Run(":8080")
}
