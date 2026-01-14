package services

import (
	"time"
	"log"

	"yogaflow.ai/models"
	"yogaflow.ai/database"
)

func CreateUserProgress(newUserProgress models.UserProgress) (models.UserProgress, error) {
	now := time.Now()
	newUserProgress.Date = now
	query := `INSERT INTO user_progress (id, user_id, date, strength_improvement, flexibility_improvement, sessions_completed, total_time)
			VALUES (DEFAULT, $1, $2, $3, $4, $5, $6)`
	err := database.Db.QueryRow(
		query,
		newUserProgress.ID,
		newUserProgress.UserID,
		now,
		newUserProgress.StrengthImprovement,
		newUserProgress.FlexibilityImprovement,
		newUserProgress.SessionsCompleted,
		newUserProgress.TotalTime,
	).Scan(&newUserProgress.ID)
	return newUserProgress, err
}

func UpdateUserProgress(updatedUserProgress models.UserProgress) (models.UserProgress, error) {
	now := time.Now()
	query := `UPDATE user_progress SET date=$1, strength_improvement=$2, flexibility_improvement=$3, sessions_completed=$4, total_time=$5`
	_, err := database.Db.Exec(
		query,
		now,
		updatedUserProgress.StrengthImprovement,
		updatedUserProgress.FlexibilityImprovement,
		updatedUserProgress.SessionsCompleted,
		updatedUserProgress.TotalTime,
	)
	return updatedUserProgress, err
}

func DeleteUserProgress(id uint) bool {
	_, err := database.Db.Exec("DELETE FROM user_progress WHERE id = $1, id")
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}