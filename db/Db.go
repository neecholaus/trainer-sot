package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func CreateDbConnection() bool {
	dsn := "host=postgres user=local password=local dbname=local port=5432 sslmode=disable TimeZone=UTC"
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Could not init db: %s", err)
		return false
	}

	// todo - somehow avoid (var not used)
	// note - figuring out how to use, lol
	type res struct {
		Datname string
	}
	var y res
	Db.Raw("select * from pg_database limit 1").Scan(&y)
	fmt.Println(y)

	return true
}
