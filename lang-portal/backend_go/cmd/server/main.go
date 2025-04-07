package main

import (
	"github.com/gin-gonic/gin"
	"lang-portal-backend/internal/handlers"
)

func main() {
	r := gin.Default()

	// Setup routes
	handlers.RegisterRoutes(r)

	// Start server
	r.Run(":8080")
}
