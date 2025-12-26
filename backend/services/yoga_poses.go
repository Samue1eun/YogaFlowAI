package services

import (
	"log"

	"yogaflow.ai/database"
	"yogaflow.ai/models"
)

func CreateYogaPose(newYogaPose models.YogaPoses) (models.YogaPoses, error) {
	query := `INSERT INTO yoga_poses (id, name, sanskrit, category, strength, flexibility, difficulty, level) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err := database.Db.QueryRow(
		query,
		newYogaPose.ID,
		newYogaPose.Name,
		newYogaPose.Sanskrit,
		newYogaPose.Category,
		newYogaPose.Strength,
		newYogaPose.Flexibility,
		newYogaPose.Difficulty,
		newYogaPose.Level,
	).Scan(&newYogaPose.ID)
	return newYogaPose, err
}

// Update Yoga Pose
func UpdateYogaPose (updatedYogaPose models.YogaPoses) (models.YogaPoses, error) {
	query := `UPDATE yoga_poses SET name=$1, sandskrit=$2, category=$3, strength=$4, flexibility=$5, difficulty=$6, level=$7`
	_, err := database.Db.Exec(
		query,
		updatedYogaPose.Name,
		updatedYogaPose.Sanskrit,
		updatedYogaPose.Category,
		updatedYogaPose.Strength,
		updatedYogaPose.Flexibility,
		updatedYogaPose.Difficulty,
		updatedYogaPose.Level,
	)
	return updatedYogaPose, err
}

// Delete Yoga Pose
func DeleteYogaPose (deleteYogaPose models.YogaPoses) (models.YogaPoses, error) {
	_, err := database.Db.Exec("DELETE FROM yoga_pose WHERE id =$1", deleteYogaPose.ID)
	if err != nil {
		log.Println(err)
		return deleteYogaPose, err
	}
	return deleteYogaPose, nil
}