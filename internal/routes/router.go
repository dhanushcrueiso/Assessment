package routes

import (
	"Assessment/internal/handlers"

	"github.com/gofiber/fiber"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/v1")

	api.Post("/employee/saveAttendance", handlers.SaveAttendance)
	api.Get("/employee/:eid/getattendance", handlers.GetAttendance)
	api.Get("/employee/get", handlers.GetAllWithLess)
}
