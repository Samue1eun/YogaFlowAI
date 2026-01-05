package services

import (
	"encoding/json"

	"yogaflow.ai/models"
	"yogaflow.ai/database"
)

// Need to work on the Schema Table prior to implementing user_flow_services

func CreateYogaFlow (newYogaFlow models.YogaFlow) (models.YogaFlow, error) {
	poseListJSON, err := json.Marshal(newYogaFlow.PoseList)
    if err != nil {
        return newYogaFlow, err
    }
	query := `INSERT INTO yoga_flows (id, type, time_length, number_of_poses, pose_list, average_strength, average_flexibility, average_difficulty)
			VALUES (DEFAULT, $1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err = database.Db.QueryRow(
		query,
		newYogaFlow.Type,
		newYogaFlow.TimeLength,
		newYogaFlow.NumberOfPoses,
		poseListJSON,
		newYogaFlow.AverageStrength,
		newYogaFlow.AverageFlexibility,
		newYogaFlow.AverageDifficulty,
	).Scan(&newYogaFlow.ID)
	return newYogaFlow, err
}