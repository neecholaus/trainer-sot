package db

import "gorm.io/gorm"

type Trainer struct {
	gorm.Model
	Id       int
	Email    string
	Password string
}
