package handlers

import (
	"encoding/json"
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
	rows, err := database.Db.Query("SELECT id, type, time_length, number_of_poses, pose_list, average_strength, average_flexibility, average_difficulty FROM yoga_flows")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	
	for rows.Next() {
		var yogaFlow models.YogaFlow
		var poseListJSON []byte
		err := rows.Scan(
			&yogaFlow.ID, 
			&yogaFlow.Type, 
			&yogaFlow.TimeLength, 
			&yogaFlow.NumberOfPoses, 
			&poseListJSON, 
			&yogaFlow.AverageStrength, 
			&yogaFlow.AverageFlexibility, 
			&yogaFlow.AverageDifficulty)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// Unmarshal JSON bytes into []YogaPoses
        if err := json.Unmarshal(poseListJSON, &yogaFlow.PoseList); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
		yogaFlows = append(yogaFlows, yogaFlow)
	}

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

// Delete Yoga Flow

func DeleteYogaFlow(c *gin.Context) {
	id := c.Param("id")
	var yogaFlow models.YogaFlow
	_, err := database.Db.Exec(`DELETE FROM yoga_flow WHERE id = $1`, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Yoga Flow deleted", "Yoga Flow": yogaFlow})
}