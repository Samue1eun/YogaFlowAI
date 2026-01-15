package handlers

import (
	"encoding/json"
	"net/http"
	"database/sql"
	"strconv"

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
			return
		}
		userProfiles = append(userProfiles, userProfile)
	}
	c.IndentedJSON(http.StatusOK, userProfiles)
}

func GetOneUserProfile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid"})
	}
	var userProfile models.UserProfile
	var injuries []byte
	var goals []byte
	err = database.Db.QueryRow(
		"SELECT id, user_id, fitness_level, strength_level, injuries, goals, created_at, updated_at FROM user_profile WHERE id = $1",
		id,
	).Scan(
		&userProfile.ID,
		&userProfile.UserID,
		&userProfile.FitnessLevel,
		&userProfile.StrengthLevel,
		&userProfile.Injuries,
		&userProfile.Goals,
		&userProfile.CreatedAt,
		&userProfile.UpdatedAt,
	)
		if err == sql.ErrNoRows {
        c.JSON(http.StatusNotFound, gin.H{"error": "User profile not found"})
        return
    }
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "User profile not found"})
		return
	}
		if err := json.Unmarshal(injuries, &userProfile.Injuries); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
	}
			if err := json.Unmarshal(goals, &userProfile.Goals); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, userProfile)
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

func UpdateUserProfile(c *gin.Context) {
	var updateUserProfile models.UserProfile
	err := c.ShouldBindJSON(&updateUserProfile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userProfile, err := services.UpdateUserProfile(updateUserProfile)
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