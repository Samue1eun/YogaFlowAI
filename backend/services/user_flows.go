package services

import (
	"log"
	"time"

	"yogaflow.ai/database"
	"yogaflow.ai/models"
)


func CreateUserFlow(newFlow models.UserFlows) (models.UserFlows, error) {
	now := time.Now()
	newFlow.CreatedAt = now
	query := `INSERT INTO user_flows (user_id, yoga_flow_id, created_at)
			VALUES ($1, $2, $3) RETURNING id`
	err := database.Db.QueryRow(
		query,
		newFlow.ID,
		newFlow.UserID,
		newFlow.YogaFlowID,
		newFlow.CreatedAt,
	).Scan(&newFlow.ID)
	return newFlow, err
}

func DeleteUserFlow(id uint) bool {
	_, err := database.Db.Exec("DELETE FROM user_flows WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}