package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHelloWorld is the unit test for the /hello endpoint
func TestHelloWorld(t *testing.T) {
	// Set Gin to test mode to disable unnecessary logging
	gin.SetMode(gin.TestMode)

	// Initialize the router
	r := gin.Default()

	// Define the route as you did in main.go
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// Create a test HTTP request
	req, _ := http.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Assert that the HTTP status code is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert that the response body is the expected JSON
	expected := `{"message":"Hello, World!"}`
	assert.JSONEq(t, expected, w.Body.String())
}

// Test an invalid route
func TestInvalidRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()

	// Define some other route
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// Invalid route
	req, _ := http.NewRequest("GET", "/nonexistent", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Assert that the status code is 404 (Not Found)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
