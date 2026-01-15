package services

import (
	"encoding/json"
	"log"
	"time"

	"yogaflow.ai/database"
	"yogaflow.ai/models"
)

func CreateUserProfile (newUserProfile models.UserProfile) (models.UserProfile, error) {
	now := time.Now()
	newUserProfile.CreatedAt = now
	newUserProfile.UpdatedAt = now
	injuriesListJSON, err := json.Marshal(newUserProfile.Injuries)
	if err != nil {
		return newUserProfile, err
	}
	goalsListJSON, err := json.Marshal(newUserProfile.Goals)
	if err != nil {
		return newUserProfile, err
	}
	query := `INSERT INTO user_profile(id, user_id, fitness_level, strength_level, injuries, goals, created_at, updated_at)
			VALUES (DEFAULT, $1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err = database.Db.QueryRow(
		query,
		newUserProfile.UserID,
		newUserProfile.FitnessLevel,
		newUserProfile.StrengthLevel,
		injuriesListJSON,
		goalsListJSON,
		now,
		now,
	).Scan(&newUserProfile.ID)
	return newUserProfile, err
}

// Update User Profile

func UpdateUserProfile () {
	
}

func DeleteUserProfile (id uint) bool {
	_, err := database.Db.Exec("DELETE FROM user_profile WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}