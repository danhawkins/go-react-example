package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/danhawkins/go-react-example/todos"
	"github.com/danhawkins/go-react-example/todos_web"
)

func main() {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")+"&application_name=$ go_react_example"), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Fatal("Failed to connect to database \n", err)
		os.Exit(2)
	}
	log.Println("Connected to database")

	app := fiber.New()

	// Setup Todos Service
	todosRepo := todos.NewTodoRepository(db)
	if err := todosRepo.Migrate(); err != nil {
		log.Fatal("Failed to migrate database \n", err)
		os.Exit(2)
	}

	todosWeb := todos_web.NewTodosWeb(app, todos.NewTodoService(todosRepo))
	todosWeb.SetupRoutes()

	app.Listen(":3000")
}
