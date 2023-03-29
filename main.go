package main

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

// content holds our static web server content.
//
//go:embed client/dist/assets/*
//go:embed client/dist/index.html
var content embed.FS

func main() {

	// Create a new gin router
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/health", func(c *gin.Context) {
			// Write a 200 response with JSON data
			data := map[string]string{
				"message": "Success",
			}
			c.JSON(http.StatusOK, data)
		})
	}

	assetsServer := http.FileServer(http.FS(content))
	r.GET("/assets/*filepath", func(c *gin.Context) {
		c.Request.URL.Path = "client/dist/assets/" + c.Param("filepath")
		assetsServer.ServeHTTP(c.Writer, c.Request)
	})

	r.GET("/", func(c *gin.Context) {
		// Serve the index.html file from the embedded assets
		data, err := content.ReadFile("client/dist/index.html")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})

	// Start the server
	r.Run(":3000")
}
