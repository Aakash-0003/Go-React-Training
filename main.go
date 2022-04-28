package main

import (
	"github.com/Aakash-0003/Go-React-Training/database"
	"github.com/Aakash-0003/Go-React-Training/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.CreateConnection()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.RootSetup(app)
	app.Listen(":8000")
}
