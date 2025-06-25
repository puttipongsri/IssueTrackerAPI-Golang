package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name     string `json:"name" validate:"required,min=3,max=32"`
	Email    string `gorm:"unique" json:"email" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=3,max=32"`
	Role     string `json:"role"`
}