package todos

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TodoService struct {
	app *fiber.App
	db  *gorm.DB
}

func NewTodoService(app *fiber.App, db *gorm.DB) *TodoService {
	return &TodoService{
		app: app,
		db:  db,
	}
}

func Setup(app *fiber.App, db *gorm.DB) {
	setupRoutes(app)
	setupDb(db)
}

func setupDb(db *gorm.DB) {
	db.AutoMigrate(&Todo{})
}
