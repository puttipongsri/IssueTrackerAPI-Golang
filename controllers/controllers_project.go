package controllers

import (
	"IssueTrackerAPI/models"
	"IssueTrackerAPI/services"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)


func CreateProject(c *fiber.Ctx) error {
	name := c.FormValue("name")
	description := c.FormValue("description")
	employeeIDsRaw := c.FormValue("employee_ids")
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Image is required"})
	}

	timestamp := time.Now().Format("02012006150405") 
	ext := filepath.Ext(file.Filename)       
	newFilename := timestamp + ext
	saveDir := "./fileimg/"
	os.MkdirAll(saveDir, os.ModePerm)
	savePath := "./fileimg/" + newFilename
	if err := c.SaveFile(file, savePath); err != nil {
		return c.Status(500).SendString("Failed to save image")
	}

	ids := strings.Split(employeeIDsRaw, ",")
	var employeeIDs []uint
	for _, idStr := range ids {
		id, _ := strconv.Atoi(strings.TrimSpace(idStr))
		employeeIDs = append(employeeIDs, uint(id))
	}

	if err := services.CreateProject(name, description, employeeIDs, savePath); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Project created successfully"})
}
func GetProjects(c *fiber.Ctx) error {
	projects, err := services.GetProjects()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(projects)
}
func GetProjectByID(c *fiber.Ctx) error {
	var body models.Project
		if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	project, err := services.GetProjectByID(body.ID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Project not found"})
	}
	return c.JSON(project)
}
func UpdateProject(c *fiber.Ctx) error {
	idStr := c.FormValue("id")
	name := c.FormValue("name")
	description := c.FormValue("description")
	image, _ := c.FormFile("image")

	idUint, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid project ID"})
	}

	var imagePath string
	if image != nil {
		timestamp := time.Now().Format("02012006150405")
		ext := filepath.Ext(image.Filename)
		newFilename := timestamp + ext
		saveDir := "./fileimg/"
		os.MkdirAll(saveDir, os.ModePerm)
		savePath := "./fileimg/" + newFilename
		if err := c.SaveFile(image, savePath); err != nil {
			return c.Status(500).SendString("Failed to save image")
		}
		imagePath = savePath
	}

	if err := services.UpdateProject(uint(idUint), name, description, imagePath); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Project updated successfully"})
}
func DeleteProject(c *fiber.Ctx) error {
	var body models.Project
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := services.DeleteProject(body.ID); err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Project not found"})
	}
	return c.JSON(fiber.Map{"message": "Project deleted successfully"})
}
