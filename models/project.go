package models

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model `json:"-"`
	ID uint `json:"id"`
	Name        string `json:"name" validate:"required,min=3"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type ProjectEmployee struct {
	gorm.Model
	ProjectID  uint `json:"project_id"`
	EmployeeID uint `json:"employee_id"`
}