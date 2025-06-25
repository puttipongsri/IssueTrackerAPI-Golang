package controllers

import (
	"IssueTrackerAPI/services"

	"github.com/gofiber/fiber/v2"
)

func ExportProjectExcel(c *fiber.Ctx) error {
	var body map[string]interface{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	projectIDFloat, ok := body["project_id"].(float64)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "project_id is required and must be a number"})
	}
	projectID := uint(projectIDFloat)

	f, err := services.ExportProjectExcel(projectID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Set("Content-Disposition", `attachment; filename="project.xlsx"`)

	if err := f.Write(c.Context()); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to write excel"})
	}
	return nil
}