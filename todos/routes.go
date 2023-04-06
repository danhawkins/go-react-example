package todos

import "github.com/gofiber/fiber/v2"

func setupRoutes(app *fiber.App) {
	app.Get("/todos", indexHandler)
}
