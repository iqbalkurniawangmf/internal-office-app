package config

import (
	"fmt"
	"log"
	"os"

	"internal-office-backend/internal/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	DB = database
	log.Println("Database Connected Successfully")

	// AUTO MIGRATION
	err = DB.AutoMigrate(
		&model.Department{},
		&model.User{},
		&model.Leave{},
		&model.Task{},
		&model.Announcement{},
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration completed successfully")
}