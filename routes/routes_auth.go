package routes

import (
	"IssueTrackerAPI/controllers"

	"github.com/gofiber/fiber/v2"

	"IssueTrackerAPI/middleware"
)

func AuthRoutes(app *fiber.App) {
	auth := app.Group("/IssueTrackerAPI/auth")
	auth.Post("/register",  middleware.ValidTest,controllers.Register)
	auth.Post("/login", controllers.Login , middleware.ValidTest)
}
