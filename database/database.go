package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/danhawkins/go-react-example/todos"
)

type DBInstance struct {
	DB *gorm.DB
}

var DB DBInstance

func ConnectDB() {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")+"&application_name=$ go_react_example"), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Fatal("Failed to connect to database \n", err)
		os.Exit(2)
	}
	log.Println("Connected to database")

	log.Println("Running migrations")
	db.AutoMigrate(&todos.Todo{})

	DB = DBInstance{DB: db}
}
