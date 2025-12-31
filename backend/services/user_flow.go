package services

import (
	"yogaflow.ai/database"
	"yogaflow.ai/models"
)

func CreateUserFlow(newFlow models.UserFlows) (models.UserFlows, error) {
	query := `INSERT INTO user_flows (id, type, timelength, numberofposes, poselist, user_id)
			VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := database.Db.QueryRow(
		query,
		newFlow.ID,
		newFlow.Type,
		newFlow.TimeLength,
		newFlow.NumberOfPoses,
		newFlow.PoseList,
		newFlow.UserID,
	).Scan(&newFlow.ID)
	return newFlow, err
}