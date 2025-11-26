package main

import (
	"log"

	"mini-ecommerce/config"
	database "mini-ecommerce/db"
	"mini-ecommerce/internal/router"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to database
	db := database.Connect(cfg)

	// Run migrations
	if err := database.Migrate(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	// Setup router with database
	r := router.SetupRouter(db)

	log.Println("Server running on port", cfg.Port)
	r.Run(":" + cfg.Port)
}
