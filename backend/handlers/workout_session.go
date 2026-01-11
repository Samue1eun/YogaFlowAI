package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"yogaflow.ai/models"
)

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
			&workoutSession.
		)
	}
}