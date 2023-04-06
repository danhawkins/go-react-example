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
	svc := &TodoService{
		app: app,
		db:  db,
	}

	svc.Setup()

	return svc
}

func (s *TodoService) Setup(app *fiber.App, db *gorm.DB) {
	setupRoutes(app)
	setupDb(db)
}

func setupDb(db *gorm.DB) {
	db.AutoMigrate(&Todo{})
}
