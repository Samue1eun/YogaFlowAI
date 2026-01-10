package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"yogaflow.ai/database"
	"yogaflow.ai/models"
)

func GetAllUserProfiles (c *gin.Context) {
	var userProfiles []models.UserProfile
	// Add in model
	rows, err := database.Db.Query("SELECT id, user_id, fitness_level, flexibility_level, strength_level, injuries, goals FROM user_profile")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var userProfiles models.YogaFlow
		var injuriesListJSON []byte
		var goalsListJSON []byte
		// Take a look at the model
		err := rows.Scan(
			&userProfiles.ID,
			&userProfiles.UserID,
			&userProfiles.FitnessLevel,
			&userProfiles.StrengthLevel,
			&injuriesListJSON,
			&goalsListJSON,
		)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Unmarshal JSON bytes into []Injuries and []Goal
	

	}
}