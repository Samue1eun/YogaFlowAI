package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"yogaflow.ai/database"
	"yogaflow.ai/models"
	"yogaflow.ai/services"
)

func GetAllUserFavorites(c *gin.Context) {
	var userFavorites []models.UserFavorites
	rows, err := database.Db.Query("SELECT id, user_id, favorite_poses, favorite_flows, created_at, updated_at FROM user_favorites")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var userFavorite models.UserFavorites
		var favoritePosesJSON []byte
		var favoriteFlowsJSON []byte
		err := rows.Scan(
			&userFavorite.ID,
			&userFavorite.UserID,
			&favoritePosesJSON,
			&favoriteFlowsJSON,
			&userFavorite.CreatedAt,
			&userFavorite.UpdatedAt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		err = json.Unmarshal(favoritePosesJSON, &userFavorite.FavoritePoses)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		err = json.Unmarshal(favoriteFlowsJSON, &userFavorite.FavoriteFlows)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		userFavorites = append(userFavorites, userFavorite)
	}
	c.IndentedJSON(http.StatusOK, userFavorites)
}

func CreateUserFavorite(c *gin.Context) {
	var newUserFavorite models.UserFavorites
	err := c.ShouldBindJSON(&newUserFavorite)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userFavorite, err := services.CreateUserFavorite(newUserFavorite)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, userFavorite)
}

func UpdateUserFavorite(c *gin.Context) {
	var updateUserFavorite models.UserFavorites
	err := c.ShouldBindJSON(&updateUserFavorite)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userFavorite, err := services.UpdateUserFavorite(updateUserFavorite)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, userFavorite)
}

func DeleteUserFavorite(c *gin.Context) {
	id := c.Param("id")
	_, err := database.Db.Exec("DELETE FROM user_favorites WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User favorites deleted"})
}