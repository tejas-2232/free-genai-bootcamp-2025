package main

import (
	"log"
	"os"
	"path/filepath"

	"lang-portal-backend/internal/handlers"
	"lang-portal-backend/internal/models"

	"github.com/gin-gonic/gin"

	"lang-portal/internal/handlers"
)

func main() {

	//create database path from environment variable or use  default
	dbPath := os.Getenv("DB_PATH")

	if dbPath == "" {
		dbPath == filepath.Join(".", "words.db")
	}

	dbDir := filepath.Dir(dbPath)

	if err := os.MkdirAll(dbDir, 0755); err != nil {
		log.Fatal("Failed to create database directory", err)
	}

	//initialize database
	db, err := models.NewDB(dbPath)
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	defer db.Close()

	// initialize handlers
	h := handlers.NewHandlers(
		dashboardService,
		wordService,
		groupService,
		studyActivityService,
		studySessionService,
	)

	//create gin router
	r := gin.Default()

	// add middleware
	r.Use(middleware.ErrorHandler())

	// Setup routes
	handlers.RegisterRoutes(r)

	// Start server
	r.Run(":8080")
}
