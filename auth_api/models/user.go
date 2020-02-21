package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	Email string `json:"e-mail" validate:"required"`
	Password string `json:"password" validate:"required"`
}