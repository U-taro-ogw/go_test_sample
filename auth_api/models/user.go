package models

import (
	//"gopkg.in/go-playground/validator.v9"
)

type User struct {
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}