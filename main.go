package main

import (
	"Assessment/internal/db"
	"Assessment/internal/routes"

	"github.com/gofiber/fiber"
)

func main() {

	db.Init(&db.Config{
		URL:       "postgresql://postgres:postgres@localhost:5432/employee?sslmode=disable",
		MaxDBConn: 5,
	})

	app := fiber.New()

	routes.SetupRoutes(app)
	// fmt.Printf("Server is running on port %s\n", cnf.Port)

	app.Listen(":3007")
}
