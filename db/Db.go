package db

import (
	"fmt"
	"nicholas/trainer-sot/db/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func CreateDbConnection() (*gorm.DB, error) {
	if Db != nil {
		fmt.Println("db connection already created")
		return Db, nil
	}

	// todo - replace with env values
	dsn := "host=postgres user=local password=local dbname=local port=5432 sslmode=disable TimeZone=UTC"

	// Declare so that Db uses global
	var err error

	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Could not init db: %s", err)
		return nil, err
	}

	return Db, nil
}

func Migrate() {
	_ = Db.AutoMigrate(&models.Trainer{})
}
