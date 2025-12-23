package services

import (
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

// Delete Yoga Pose

