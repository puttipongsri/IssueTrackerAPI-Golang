package routes

import (
	"IssueTrackerAPI/controllers"
	"IssueTrackerAPI/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProjectRoutes(app *fiber.App){
	project := app.Group("/IssueTrackerAPI/project", middleware.JWTMiddleware)

	project.Post("/create", controllers.CreateProject)
	project.Get("/", controllers.GetProjects)
	project.Get("/id", controllers.GetProjectByID)
	project.Put("/update", controllers.UpdateProject)
	project.Delete("/delete", controllers.DeleteProject)
}