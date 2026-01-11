package services

import (
	"log"
	// "net/http"

	"yogaflow.ai/database"
	"yogaflow.ai/models"
)



func CreatePosePerformance(newPosePerformance models.PosePerformance) (models.PosePerformance, error) {
	
	query := `INSERT INTO pose_performance(id, user_id, pose_id, attempts, success_rate, difficulty_rating, last_attempted)
			VALUES (DEFAULT, $1, $2, $3, $4, $5, $6, $7) RETURNING id`
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

func UpdatePosePerformance () {

}

func DeletePosePerformance (id uint) bool {
	_, err := database.Db.Exec("DELETE FROM pose_performance WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}