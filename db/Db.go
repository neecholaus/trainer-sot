package db

import (
	"fmt"
	"nicholas/trainer-sot/db/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func EnsureDbConnection() (*gorm.DB, error) {
	if Db != nil {
		fmt.Println("db connection already created")
		return Db, nil
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASS"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	// Declare so that Db uses global
	var err error

	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Could not init db: %s", err)
		Db = nil
		return nil, err
	}

	return Db, nil
}

func Migrate() {
	_ = Db.AutoMigrate(&models.Trainer{}, &models.Client{})
}
