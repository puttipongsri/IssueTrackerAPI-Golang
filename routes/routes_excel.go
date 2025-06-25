package routes

import (
	"IssueTrackerAPI/controllers"
	"IssueTrackerAPI/middleware"

	"github.com/gofiber/fiber/v2"
)

func ExcelRoutes(app *fiber.App) {
	excel := app.Group("/IssueTrackerAPI/excel", middleware.JWTMiddleware)

	excel.Post("/", controllers.ExportProjectExcel)
}