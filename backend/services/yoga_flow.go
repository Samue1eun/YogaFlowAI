package services

import (
	"encoding/json"
	"log"
	"net/http"

	"yogaflow.ai/models"
	"yogaflow.ai/database"
)

// Need to work on the Schema Table prior to implementing user_flow_services


// Create a YogaFlow

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

// Bulk Yoga Flow (Need to work on)

// func AddBulkYogaFlows (c *gin.Context) {
// 	var addBulkYogaFlows []models.YogaPoses
// 	err := c.ShouldBindBodyWithJSON(&addBulkYogaFlows)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	query := `INSERT INTO yoga_flows (id, type, time_length, number_of_poses, pose_list, average_strength, average_flexibility, average_difficulty)
// 				VALUES ($1, $2, $3, $4, $5, $6, $7, %8)`
// }

// Update Yoga Flow

func UpdateYogaFlow (updatedYogaFlow models.YogaFlow) (models.YogaFlow, error) {
	poseListJSON, err := json.Marshal(updatedYogaFlow.PoseList)
    if err != nil {
        return updatedYogaFlow, err
    }
	query := `UPDATE yoga_flows SET type=$1, time_length=$2, number_of_poses=$3, pose_list=$4, average_strength=$5, average_flexibility=$6, average_difficulty=$7 WHERE id=$8`
	_, err = database.Db.Exec(
		query,
		updatedYogaFlow.Type,
		updatedYogaFlow.TimeLength,
		updatedYogaFlow.NumberOfPoses,
		poseListJSON,
		updatedYogaFlow.AverageStrength,
		updatedYogaFlow.AverageFlexibility,
		updatedYogaFlow.AverageDifficulty,
		updatedYogaFlow.ID,
	)
	return updatedYogaFlow, err
}

// Delete Yoga Flow

func DeleteYogaFlow(id uint) bool {
	_, err := database.Db.Exec("DELETE FROM yoga_flows WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}