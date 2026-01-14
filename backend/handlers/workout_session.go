package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"yogaflow.ai/database"
	"yogaflow.ai/models"
	"yogaflow.ai/services"
)

// NOT FINISHED WITH THIS YET.

func GetAllWorkoutSession(c* gin.Context){
	var workoutSessions []models.WorkoutSession
	rows, err := database.Db.Query("SELECT id,")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
	}
	defer rows.Close()

	for rows.Next() {
		var workoutSession models.WorkoutSession
		err := rows.Scan(
			&workoutSession.ID,
			&workoutSession.UserID,
			&workoutSession.YogaFlowID,
			&workoutSession.StartedAt,
			&workoutSession.CompletedAt,
			&workoutSession.Duration,
			&workoutSession.Rating,
			&workoutSession.Feedback,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		workoutSessions = append(workoutSessions, workoutSession)
	}
	c.IndentedJSON(http.StatusOK, workoutSessions)
}

func CreateWorkoutSession(c *gin.Context) {
	var newWorkoutSession models.WorkoutSession
	err := c.ShouldBindJSON(&newWorkoutSession)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	workoutSession, err := services.CreateWorkoutSession(newWorkoutSession)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, workoutSession)
}

func DeleteWorkoutSession(c *gin.Context) {
	id := c.Param("id")
	var workoutSession models.WorkoutSession
	_, err := database.Db.Exec("DELETE FROM workout_session WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Workout session deleted", "Workout session": workoutSession})
}