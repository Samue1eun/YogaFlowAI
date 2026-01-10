package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"yogaflow.ai/models"
	"yogaflow.ai/services"
)

func CreatePosePerformance(c *gin.Context) {
	var newPosePerformance models.PosePerformance
	err := c.ShouldBindJSON(&newPosePerformance)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userPosePerformance, err := services.CreatePosePerformance(newPosePerformance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, userPosePerformance)

}