package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"yogaflow.ai/database"
	"yogaflow.ai/models"
	"yogaflow.ai/services"
)

type UpdatePoseAttemptRequest struct {
	PoseID int `json:"pose_id" binding:"required"`
	WasSuccessful bool `json:"was_successful"`
}

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

func GetOnePosePerformance(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var posePerformance models.PosePerformance
	err = database.Db.QueryRow(
		"SELECT id, user_id, pose_id, attempts, success_rate, difficulty_rating, last_attempted FROM pose_performance WHERE id = $1",
		id,
	).Scan(
		&posePerformance.ID,
		&posePerformance.UserID,
		&posePerformance.PoseID,
		&posePerformance.Attempts,
		&posePerformance.SuccessRate,
		&posePerformance.DifficultyRating,
		&posePerformance.LastAttempted,
	)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pose performance not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pose performance not found"})
		return
	}
	c.JSON(http.StatusOK, posePerformance)
}

func UpdatePosePerformance(c *gin.Context) {
	var updatePosePerformance models.PosePerformance
	err := c.ShouldBindJSON(&updatePosePerformance)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	posePerformance, err := services.UpdatePosePerformance(updatePosePerformance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, posePerformance)
}


// Update every time a flow calls the pose (NOT COMPLETE)
func UpdateUserPosePerformance(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var req UpdatePoseAttemptRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var updateUserPosePerformance models.PosePerformance
	err = database.Db.QueryRow(
		"SELECT id, user_id, pose_id, attempts, success_rate FROM pose_perforance WHERE id=$1",
		id,
	).Scan(
		&updateUserPosePerformance.ID,
		&updateUserPosePerformance.UserID,
		&updateUserPosePerformance.PoseID,
		&updateUserPosePerformance.Attempts,
		&updateUserPosePerformance.SuccessRate,
	)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pose performance not found"})
		return
	}
	// performance, err := services.GetPosePerformanceByUserAndPose(userID, poseID)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	// performance.Attempts += 1

	// updatedPerformance, err := services.UpdatePosePerformance(performance)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	// c.JSON(http.StatusCreated, updatedPerformance)
}

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