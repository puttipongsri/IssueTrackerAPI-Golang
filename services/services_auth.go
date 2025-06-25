package services

import (
	"IssueTrackerAPI/database"
	m "IssueTrackerAPI/models"
	"IssueTrackerAPI/utils"

	"github.com/gofiber/fiber/v2"
)

func Register(name, email, password string ,) error {
	db := database.DBConn
	hashedPwd, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	employee := m.Employee{
		Name:     name,
		Email:    email,
		Password: hashedPwd,
		Role:     "employee",
	}
	return db.Create(&employee).Error
}

func Login(email, password string) (m.Employee, error) {
	db := database.DBConn
	var emp m.Employee
	err := db.Where("email = ?", email).First(&emp).Error
	if err != nil || !utils.CheckPasswordHash(password, emp.Password) {
		return emp, fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}
	return emp, nil
}