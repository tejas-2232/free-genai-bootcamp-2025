package main

import (
	"lang-portal-backend/internal/handlers"

	"github.com/gin-gonic/gin"

	"lang-portal/internal/handlers"
)

func main() {
	r := gin.Default()

	// Setup routes
	handlers.RegisterRoutes(r)

	// Start server
	r.Run(":8080")
}
