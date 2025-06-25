package controllers

import (
	m "IssueTrackerAPI/models"
	"IssueTrackerAPI/services"
	"IssueTrackerAPI/utils"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var body m.Employee
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := services.Register(body.Name, body.Email, body.Password); err != nil { 
		if utils.IsUniqueConstraintError(err) {
			return c.Status(409).JSON(fiber.Map{"error": "Email already exists"})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Registered successfully"})
}

func Login(c *fiber.Ctx) error {
	var body m.Employee	
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	emp, err := services.Login(body.Email, body.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}
	token, _ := utils.GenerateJWT(emp.ID)
	return c.JSON(fiber.Map{"token": token})
}