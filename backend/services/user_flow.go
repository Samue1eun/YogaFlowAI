package services

// import (
// 	"log"

// 	"yogaflow.ai/database"
// 	"yogaflow.ai/models"
// )

// Need to work on the Schema Table prior to implementing user_flow_services

// func CreateUserFlow(newFlow models.UserFlows) (models.UserFlows, error) {
// 	query := `INSERT INTO user_flows (id, type, timelength, numberofposes, poselist, user_id)
// 			VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
// 	err := database.Db.QueryRow(
// 		query,
// 		newFlow.ID,
// 		newFlow.Type,
// 		newFlow.TimeLength,
// 		newFlow.NumberOfPoses,
// 		newFlow.PoseList,
// 		newFlow.UserID,
// 	).Scan(&newFlow.ID)
// 	return newFlow, err
// }

// func DeleteUserFlow(id uint) bool {
// 	_, err := database.Db.Exec("DELETE FROM user_flows WHERE id = $1", id)
// 	if err != nil {
// 		log.Println(err)
// 		return false
// 	}
// 	return true
// }