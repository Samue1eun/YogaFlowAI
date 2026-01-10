package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"yogaflow.ai/database"
	"yogaflow.ai/models"
	"yogaflow.ai/services"
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
			&userProfile.FlexibilityLevel,
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
		err = json.Unmarshal(goalsListJSON, &userProfile.Goals)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		userProfiles = append(userProfiles, userProfile)
	}
	c.IndentedJSON(http.StatusOK, userProfiles)
}

func CreateUserProfile(c *gin.Context) {
	var newUserProfile models.UserProfile
	err := c.ShouldBindJSON(&newUserProfile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userProfile, err := services.CreateUserProfile(newUserProfile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, userProfile)
}

func DeleteUserProfile(c *gin.Context) {
	id := c.Param("id")
	var userProfile models.UserProfile
	_, err := database.Db.Exec("DELETE FROM user_profile WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Profile deleted", "user profile": userProfile})
}