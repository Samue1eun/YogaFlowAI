package services

import (
	"encoding/json"
	"time"
	"log"

	"yogaflow.ai/models"
	"yogaflow.ai/database"
)

func CreateUserFavorite(newUserFavorite models.UserFavorites) (models.UserFavorites, error) {
	now := time.Now()
	newUserFavorite.CreatedAt = now
	newUserFavorite.UpdatedAt = now
	favoritePosesListJSON, err := json.Marshal(newUserFavorite.FavoritePoses)
	if err != nil {
		return newUserFavorite, err
	}
	favoriteFlowsListJSON, err := json.Marshal(newUserFavorite.FavoriteFlows)
	if err != nil{
		return newUserFavorite, err
	}
	query := `INSERT INTO user_favorites (id, user_id, favorite_poses, favorite_flows, created_at, updated_at)
		VALUES (DEFAULT, $1, $2, $3, $4, $5) RETURNING id`
	err = database.Db.QueryRow(
		query,
		newUserFavorite.UserID,
		favoritePosesListJSON,
		favoriteFlowsListJSON,
		now,
		now,
	).Scan(&newUserFavorite.ID)
	return newUserFavorite, err
}

func UpdateUserFavorite (updatedUserFavorite models.UserFavorites) (models.UserFavorites, error) {
	now := time.Now()
	updatedUserFavorite.UpdatedAt = now
	favoritePosesListJSON, err := json.Marshal(updatedUserFavorite.FavoritePoses)
	if err != nil {
		return updatedUserFavorite, err
	}
	favoriteFlowsListJSON, err := json.Marshal(updatedUserFavorite.FavoriteFlows)
	if err != nil {
		return updatedUserFavorite, err
	}
	query := `UPDATE user_favorites SET favorite_poses=$1, favorite_flows=$2, created_at=$3 WHERE id=$4`
	_, err = database.Db.Exec(
		query,
		favoritePosesListJSON,
		favoriteFlowsListJSON,
		now,
		updatedUserFavorite.ID,
	)
	return updatedUserFavorite, err
}

func DeleteUserFavorite(id uint) bool {
	_, err := database.Db.Exec("DELETE FROM user_id WHERE id=$1", id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}