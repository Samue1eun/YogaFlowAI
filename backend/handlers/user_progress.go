package handlers

import (
	"encoding/json"
	"net/http"
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"

	"yogaflow.ai/models"
	"yogaflow.ai/services"
	"yogaflow.ai/database"
)

func GetAllUserProgress (c *gin.Context) {
	var userProgresses []models.UserProgress
	rows, err := database.Db.Query("SELECT id, user_id, date, strength_improvement, flexibility_improvement, sessions_completed, total_time FROM user_progress")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var userProgress models.UserProgress
		err := rows.Scan(
			&userProgress.ID,
			&userProgress.UserID,
			&userProgress.Date,
			&userProgress.StrengthImprovement,
			&userProgress.FlexibilityImprovement,
			&userProgress.SessionsCompleted,
			&userProgress.TotalTime)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		userProgresses = append(userProgresses, userProgress)
	}
	c.IndentedJSON(http.StatusOK, userProgresses)
}

func GetOneUserProgress (c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var userProgress models.UserProgress
	err = database.Db.QueryRow(
		"SELECT id, user_id, date, strength_improvement, flexibility_improvement, sessions_completed, total_time WHERE id = $1",
		id,
	).Scan(
		&userProgress.ID,
		&userProgress.UserID,
		&userProgress.Date,
		&userProgress.StrengthImprovement,
		&userProgress.FlexibilityImprovement,
		&userProgress.SessionsCompleted,
		&userProgress.TotalTime
	)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "User progress not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "User progress not found"})
		return
	}
	c.JSON(http.StatusOK, userProgress)
}

func CreateUserProgress (c* gin.Context) {
	var newUserProgress
	err := c.ShouldBind(&newUserProgress)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error"})
		return
	}
	newUserProgress, err := services.CreateUserProgress(newUserProgress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newUserProgress)
}

func UpdateUserProgress(c *gin.Context) {
	var updateUserProgress
	err := c.ShouldBindJSON(&updateUserProgress)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updateUserProgress, err := services.UpdateUserProgress(updateUserProgress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, updateUserProgress)
}

func DeleteUserProgress(c *gin.Context) {
	id := c.Param("id")
	var userProgress models.UserProgress
	_, err := database.Db.Exec(`DELETE FROM user_progress WHERE id = $1`, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "User progress deleted", "User Progress": userProgress})
}