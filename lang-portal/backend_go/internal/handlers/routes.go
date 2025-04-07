package handlers

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	// Dashboard routes
	r.GET("/api/dashboard/last_study_session", GetLastStudySession)
	r.GET("/api/dashboard/study_progress", GetStudyProgress)
	r.GET("/api/dashboard/quick_stats", GetQuickStats)

	// Word routes
	r.GET("/api/words", GetWords)
	r.GET("/api/words/:id", GetWord)
}
