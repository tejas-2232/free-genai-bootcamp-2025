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

	//initialize services
	dashboardService := handlers.NewDashboardService(db)
	wordService := handlers.NewWordService(db)
	groupService := handlers.NewGroupService(db)
	studyActivityService := handlers.NewStudyActivityService(db)
	studySessionService := handlers.NewStudySessionService(db)

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

	// enable CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Setup routes
	handlers.RegisterRoutes(r)

	log.Printf("server starting on http://localhost:8081")

	if err := r.Run(":8081"); err != nil {
		log.Fatal("Failed to start server", err)
	}
}
