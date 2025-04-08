package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLastStudySession(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Last study session endpoint",
	})
}

func GetStudyProgress(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"total_words_studied": 0,
		"total_available_words": 0,
	})
}

func GetQuickStats(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success_rate": 0,
		"total_study_sessions": 0,
		"total_active_groups": 0,
		"study_streak_days": 0,
	})
}
