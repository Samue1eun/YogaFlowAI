package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"yogaflow.ai/models"
	"yogaflow.ai/services"
)

// GenerateAIFlow handles the AI yoga flow generation endpoint
func GenerateAIFlow(c *gin.Context) {
	var req models.AIFlowRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate required fields
	if req.Duration <= 0 {
		req.Duration = 45 // default to 45 minutes
	}
	if req.FlowType == "" {
		req.FlowType = "vinyasa"
	}
	if req.Difficulty == "" {
		req.Difficulty = "beginner"
	}
	if req.FocusArea == "" {
		req.FocusArea = "general wellness"
	}

	// Generate the flow using Claude AI
	flowResponse, err := services.GenerateAIFlow(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Yoga flow generated successfully",
		"flow":    flowResponse,
	})
}

// QuickGenerateAIFlow generates a flow with minimal input
func QuickGenerateAIFlow(c *gin.Context) {
	duration := c.DefaultQuery("duration", "30")
	flowType := c.DefaultQuery("type", "vinyasa")
	difficulty := c.DefaultQuery("difficulty", "beginner")

	// Convert duration to int
	var durationInt int
	if _, err := c.GetQuery("duration"); err {
		durationInt = 30
	} else {
		fmt.Sscanf(duration, "%d", &durationInt)
	}

	req := models.AIFlowRequest{
		Duration:   durationInt,
		FlowType:   flowType,
		Difficulty: difficulty,
		FocusArea:  "general wellness",
	}

	// Get user_id from auth context if available
	if userID, exists := c.Get("user_id"); exists {
		req.UserID = userID.(int)
	}

	flowResponse, err := services.GenerateAIFlow(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Yoga flow generated successfully",
		"flow":    flowResponse,
	})
}
