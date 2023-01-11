package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	TrainerId string
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
