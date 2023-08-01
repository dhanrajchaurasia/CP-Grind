package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName  string `validate:"required, min = 3, max = 32"`
	SecondName string `validate:"required, min = 3, max = 32"`
	Username   string `validate:"required, min = 3, max = 32"`
	Email      string `validate:"required, email, min = 6, max = 32"`
	Password   string `validate:"required, password, min = 8, max = 32"`
	AuthToken  string
}
