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

func main() {
	config.LoadEnv()           // Load env vars like DB_HOST, DB_USER, etc.
	db := database.ConnectDB() // Connect to DB

	// Run migrations
	err := database.MigrateDB(db)
	if err != nil {
		log.Fatalf("DB migration failed: %v", err)
	}
	log.Println("✅ Database migration completed.")

	// Create Gin router
	r := gin.Default()

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register your API routes and inject DB
	v1.RegisterV1Routes(r, db)

	log.Println("App started...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
