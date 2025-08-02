package main

// This is the Main file and Entry Point to application
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
	config.LoadEnv()           // Load env vars like DB_HOST, DB_USER, etc.
	db := database.ConnectDB() // Connect to DB

	// Run migrations
	err := database.MigrateDB(db)
	if err != nil {
		log.Fatalf("DB migration failed: %v", err)
	}
	log.Println("✅ Database migration completed.")

	// Create and configure the main router
	r := api.SetupRouter(db)

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("App started on :8081...")
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
