package services

import (
	"time"
	"log"

	"yogaflow.ai/models"
	"yogaflow.ai/database"
)

func CreateWorkoutSession(newWorkoutSession models.WorkoutSession) (models.WorkoutSession, error) {
	now := time.Now()
	newWorkoutSession.CreatedAt = now
	query := `INSERT INTO workout_session(id, user_id, yoga_flow_id, started_at, completed_at, duration, rating, feedback, created_at)
			VALUES ($1, $2, $3 ,$4, $5, $6, $7, $8, $9) RETURNING id`
	err := database.Db.QueryRow(
		query,
		newWorkoutSession.ID,
		newWorkoutSession.UserID,
		newWorkoutSession.YogaFlowID,
		newWorkoutSession.StartedAt,
		newWorkoutSession.CompletedAt,
		newWorkoutSession.Duration,
		newWorkoutSession.Rating,
		newWorkoutSession.Feedback,
		now,
	).Scan(&newWorkoutSession.ID)
	return newWorkoutSession, err
}

// NEED TO CREATE AN UPDATE FOR THE WORKOUT SESSION? FOR EVERYTIME A USER COMPLETES IT?

func DeleteWorkoutSession (id uint) bool {
	_, err := database.Db.Exec("DELETE FROM workout_session WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}