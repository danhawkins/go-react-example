package todos_web

import (
	"github.com/danhawkins/go-react-example/todos"
	"github.com/gofiber/fiber/v2"
)

type TodosWeb struct {
	app *fiber.App
	svc *todos.TodoService
}

func NewTodosWeb(app *fiber.App, svc *todos.TodoService) *TodosWeb {
	return &TodosWeb{
		app: app,
		svc: svc,
	}
}

func (t *TodosWeb) SetupRoutes() {
	t.app.Get("/todos", t.handleIndex)
}

func (t *TodosWeb) handleIndex(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
