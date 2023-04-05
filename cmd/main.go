package main

import (
	"github.com/danhawkins/go-react-example/database"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	database.ConnectDB()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Worlds 👋!")
	})

	app.Listen(":3000")
}
