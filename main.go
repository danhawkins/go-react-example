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
	// Create a file server for the embedded assets
	assetsServer := http.FileServer(http.FS(content))

	// Create a new gin router
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/health", healthHandler)
	}

	r.GET("/assets/*filepath", func(c *gin.Context) { staticHandler(c, assetsServer) })

	r.GET("/", indexHandler)

	// Start the server
	r.Run(":3000")
}

// Return the index.html from
func indexHandler(c *gin.Context) {
	// Serve the index.html file from the embedded assets
	data, err := content.ReadFile("client/dist/index.html")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", data)
}

// Health check handler
func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// Static file handler
func staticHandler(c *gin.Context, assetServer http.Handler) {
	c.Request.URL.Path = "client/dist/assets/" + c.Param("filepath")
	assetServer.ServeHTTP(c.Writer, c.Request)
}
