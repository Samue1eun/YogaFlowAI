package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"yogaflow.ai/database"
	"yogaflow.ai/models"
	"yogaflow.ai/services"
)

func GetAllPosePerformances(c *gin.Context) {
	var posePerformances []models.PosePerformance
	rows, err := database.Db.Query("SELECT id, user_id, pose_id, attempts, success_rate, difficulty_rating, last_attempted FROM pose_performance")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var posePerformance models.PosePerformance
		err := rows.Scan(
			&posePerformance.ID,
			&posePerformance.UserID,
			&posePerformance.PoseID,
			&posePerformance.Attempts,
			&posePerformance.SuccessRate,
			&posePerformance.DifficultyRating,
			&posePerformance.LastAttempted,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		posePerformances = append(posePerformances, posePerformance)
	}
	c.IndentedJSON(http.StatusOK, posePerformances)
}

func GetOnePosePerformance(c *gin.Context)
	idStr :=c.Param

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

func DeletePosePerformance(c *gin.Context) {
	id := c.Param("id")
	var posePerformance models.UserProfile
	_, err := database.Db.Exec("DELETE FROM pose_performance WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Pose performance deleted", "Pose performance": posePerformance})
}