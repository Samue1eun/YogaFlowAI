package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"yogaflow.ai/database"
	"yogaflow.ai/models"
)

func GetAllUserProfiles(c *gin.Context) {
	var userProfiles []models.UserProfile
	rows, err := database.Db.Query("SELECT id, user_id, fitness_level, flexibility_level, strength_level, injuries, goals FROM user_profile")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var userProfile models.UserProfile
		var injuriesListJSON []byte
		var goalsListJSON []byte
		err := rows.Scan(
			&userProfile.ID,
			&userProfile.UserID,
			&userProfile.FitnessLevel,
			&userProfile.StrengthLevel,
			&injuriesListJSON,
			&goalsListJSON,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// Unmarshal JSON bytes into []Injuries and []Goal
		err = json.Unmarshal(injuriesListJSON, &userProfile.Injuries)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		err = json.Unmarshal(goalsListJSON, &userProfile.)

	}
}
