package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"mini-ecommerce/config"
	"mini-ecommerce/internal/admin"
	"mini-ecommerce/internal/order"
	"mini-ecommerce/internal/product"
	"mini-ecommerce/internal/user"
)

func Connect(cfg config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBName,
		cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to database successfully")
	return db
}

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&product.Product{}, &admin.Admin{}, &user.User{}, &order.Order{}); err != nil {
		log.Fatalf("Auto migration failed: %v", err)
		return err
	}
	log.Println("Database migration completed successfully")
	return nil
}
