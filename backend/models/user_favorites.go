package models

import "time"

type UserFavorites struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	FavoritePoses []string `json:"favorite_poses"`
	FavoriteFlows []string `json:"favorite_flows"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}