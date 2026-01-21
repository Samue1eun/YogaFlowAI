package services

import (
	"log"

	"yogaflow.ai/database"
	"yogaflow.ai/models"
)



func CreatePosePerformance(newPosePerformance models.PosePerformance) (models.PosePerformance, error) {
	
	query := `INSERT INTO pose_performance(id, user_id, pose_id, attempts, success_rate, difficulty_rating, last_attempted)
			VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := database.Db.QueryRow(
		query,
		newPosePerformance.ID,
		newPosePerformance.UserID,
		newPosePerformance.PoseID,
		newPosePerformance.Attempts,
		newPosePerformance.SuccessRate,
		newPosePerformance.DifficultyRating,
		newPosePerformance.LastAttempted,
	).Scan(&newPosePerformance.ID)
	return newPosePerformance, err
}

// Update Pose Performance

func UpdatePosePerformance (updatePosePerformance models.PosePerformance) (models.PosePerformance, error) {
	query := `UPDATE pose_performance SET user_id=$1, pose_id=$2, attempts=$3, success_rate=$4, difficulty_rating=$5, last_attempted=$6 WHERE id=$7`
	_, err := database.Db.Exec(
		query,
		updatePosePerformance.UserID,
		updatePosePerformance.PoseID,
		updatePosePerformance.Attempts,
		updatePosePerformance.SuccessRate,
		updatePosePerformance.DifficultyRating,
		updatePosePerformance.LastAttempted,
		updatePosePerformance.ID,
	)
	return updatePosePerformance, err
}

func DeletePosePerformance (id uint) bool {
	_, err := database.Db.Exec("DELETE FROM pose_performance WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}