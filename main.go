package main

import (
	"Assessment/internal/db"

	"github.com/gofiber/fiber"
)

func main() {

	db.Init(&db.Config{
		URL:       "postgresql://postgres:postgres@localhost:5432/order_management?sslmode=disable",
		MaxDBConn: 5,
	})

	app := fiber.New()

	// routes.SetupRoutes(app)
	// fmt.Printf("Server is running on port %s\n", cnf.Port)

	// app.Listen(cnf.Port)
}
