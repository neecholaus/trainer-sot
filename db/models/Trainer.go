package models

import "gorm.io/gorm"

type Trainer struct {
	gorm.Model
	Email     string
	Password  string
	FirstName string
	LastName  string

	Clients []Client `gorm:"foreignKey:TrainerId"`
}
