package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	TrainerId string
	Email     string
	Phone     string
	FirstName string
	LastName  string
}
