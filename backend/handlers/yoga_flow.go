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

func GetOneYogaFlow (c *gin.Context) {
	id := c.Param("id")
	var yogaFlow models.YogaFlow
	err := database.Db.QueryRow(
		"SELECT id, type, time_length, number_of_poses, pose_list, average_strength, average_flexibiltiy, average_difficulty FROM yoga_flows WHERE id = $1",
		id,
	).Scan(&yogaFlow.ID, &yogaFlow.Type, &yogaFlow.TimeLength, &yogaFlow.NumberOfPoses, &yogaFlow.PoseList, &yogaFlow.AverageStrength, &yogaFlow.AverageFlexibility, &yogaFlow.AverageDifficulty)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Yoga flow not found"})
		return
	}
	c.JSON(http.StatusOK, yogaFlow)
	// declare a variable for the id
	// Query for the id
	// if statement if the id matches
	// return the data in the id
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

// Update Yoga Flow

func UpdateYogaFlow(c *gin.Context) {
	var updateYogaFlow models.YogaFlow
	err := c.ShouldBindJSON(&updateYogaFlow)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	yogaFlow, err := services.UpdateYogaFlow(updateYogaFlow)
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
	_, err := database.Db.Exec(`DELETE FROM yoga_flows WHERE id = $1`, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Yoga Flow deleted", "Yoga Flow": yogaFlow})
}