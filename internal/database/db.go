package database

import (
	"ecommerce-app/internal/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectAndMigrate() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		log.Fatal("DB_PATH not set in .env")
	}

	database, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	// "Open a database connection using SQLite as the backend (called a dialect) and configure GORM with the settings I pass inside gorm.Config{}."
	// A dialect in GORM is the database engine you want to use.
	// analogy : car := NewCar(engine, &CarOptions{}) // engine = petrol/diesel/electric, car = final object, car options = ac / sunroof etc etc

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	DB = database

	// AutoMigrate all models here
	err = DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.Order{},
	)
	if err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}

	fmt.Println("✅ Database connected and migrated")
}
