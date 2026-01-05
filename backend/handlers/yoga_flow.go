package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"yogaflow.ai/models"
	"yogaflow.ai/services"
	"yogaflow.ai/database"
)

// Need to work on the Scheme before working on yoga_flow handlers

// Get All Yoga Flows

func GetAllYogaFlows (c *gin.Context) {
	var yogaFlows []models.YogaFlow
	rows, err := database.Db.Query("SELECT id, type, time_length, number_of_poses, pose_list, average_strength, average_flexibility, average_difficulty")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	for rows.Next() {
		var yogaFlow models.YogaFlow
		err := rows.Scan(&yogaFlow.ID, &yogaFlow.Type, &yogaFlow.TimeLength, &yogaFlow.NumberOfPoses, &yogaFlow.PoseList, &yogaFlow.AverageStrength, &yogaFlow.AverageFlexibility, &yogaFlow.AverageDifficulty)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		yogaFlows = append(yogaFlows, yogaFlow)
	}
	defer rows.Close()
	c.IndentedJSON(http.StatusOK, yogaFlows)
}


// Add Yoga Flow
func CreateYogaFlow (c *gin.Context) {
	var newYogaFlows models.YogaFlow
	err := c.ShouldBindJSON(&newYogaFlows)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	yogaFlow, err := services.CreateYogaFlow(newYogaFlows)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, yogaFlow)
}