package models

import "gorm.io/gorm"

type Trainer struct {
	gorm.Model
	Email     string `json:"email"`
	Password  string
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`

	Clients []Client `gorm:"foreignKey:TrainerId"`
}
