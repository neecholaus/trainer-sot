package models

import "gorm.io/gorm"

type Trainer struct {
	gorm.Model
	Email string
}
