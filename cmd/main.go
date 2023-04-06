package main

import (
	"log"
	"os"

	"github.com/danhawkins/go-react-example/todos"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")+"&application_name=$ go_react_example"), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Fatal("Failed to connect to database \n", err)
		os.Exit(2)
	}
	log.Println("Connected to database")

	app := fiber.New()

	todos.Setup(app, db)

	app.Listen(":3000")
}
