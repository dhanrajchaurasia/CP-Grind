package initializers

import (
	"log"
	"os"

	"github.com/dhanrajchaurasia/CP-GRIND/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectionToDB() {
	dsn := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database!")
	}
}

func SyncDB() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Error auto migrating the database!")
	}
}
