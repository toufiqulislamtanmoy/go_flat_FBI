package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	dsn := os.Getenv("DB_DSN")
	log.Println("dsn string:", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	return db
}
