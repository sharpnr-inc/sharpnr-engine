package main

// Main entry point of the application
import (
	"log"

	"sharpnr-engine/api"
	"sharpnr-engine/config"
	_ "sharpnr-engine/docs"
	"sharpnr-engine/shared/database"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config.LoadEnv() // Load environment variables
	redisDb := database.ConnectRedis()
	db := database.ConnectDB()

	// Run migrations
	err := database.MigrateDB(db)
	if err != nil {
		log.Fatalf("DB migration failed: %v", err)
	}
	log.Println("✅ Database migration completed.")

	// Create and configure the main router
	r := api.SetupRouter(db, redisDb)

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Get local IP
	localIP, err := config.GetLocalIP()
	if err != nil || localIP == "" {
		log.Println("⚠️ Could not determine local IP, using 0.0.0.0")
		localIP = "0.0.0.0"
	}

	port := "8080"
	address := localIP + ":" + port
	log.Printf("App started on http://%s ...", address)

	// Run Gin server on detected local IP
	if err := r.Run(address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
