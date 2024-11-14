package main

import (
	_ "github.com/sbriggs-wtc/go_swagger/docs" // Import the generated docs
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Hello World API
// @version 1.0
// @description This is a simple Hello World API with Swagger documentation
// @contact.name API Support
// @contact.url https://www.example.com/support
// @contact.email support@example.com
// @host localhost:8080
// @BasePath /
// @Summary Get Hello message
// @Description Get a Hello World message
// @Tags greetings
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /hello [get]
func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// Swagger UI route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run the server
	// r.Run(":8080")
	r.Run("localhost:8080")
}
