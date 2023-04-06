package main

import (
	"embed"
	"log"
	"os"

	"github.com/danhawkins/go-react-example/todos"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// content holds our static web server content.
//
//go:embed ./../client/dist/assets/*
//go:embed ./../client/dist/index.html
var content embed.FS

func main() {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")+"&application_name=$ go_react_example"), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Fatal("Failed to connect to database \n", err)
		os.Exit(2)
	}
	log.Println("Connected to database")

	app := fiber.New()

	todos.Setup(app, db)
}

// func main() {
// db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")+"&application_name=$ go_react_example"), &gorm.Config{})
// if err != nil {
// 	log.Fatalf("Failed to connect to database %s", err)
// }

// err = todos.Setup(db)
// if err != nil {
// 	log.Fatalf("Failed to migrate database %s", err)
// }

// // Create a file server for the embedded assets
// assetsServer := http.FileServer(http.FS(content))

// // Create a new gin router
// r := gin.Default()

// api := r.Group("/api")
// {
// 	api.GET("/health", healthHandler)
// }

// r.GET("/assets/*filepath", func(c *gin.Context) { staticHandler(c, assetsServer) })
// r.GET("/", indexHandler)

// // Start the server
// r.Run(":3000")
// }

// Return the index.html from
// func indexHandler(c *gin.Context) {
// 	// Serve the index.html file from the embedded assets
// 	data, err := content.ReadFile("client/dist/index.html")
// 	if err != nil {
// 		c.AbortWithError(http.StatusInternalServerError, err)
// 		return
// 	}
// 	c.Data(http.StatusOK, "text/html; charset=utf-8", data)
// }

// // Health check handler
// func healthHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"status": "ok",
// 	})
// }

// // Static file handler
// func staticHandler(c *gin.Context, assetServer http.Handler) {
// 	c.Request.URL.Path = "client/dist/assets/" + c.Param("filepath")
// 	assetServer.ServeHTTP(c.Writer, c.Request)
// }
