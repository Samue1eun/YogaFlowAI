package models

import "time"

type PosePerformance struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	PoseID int `json:"pose_id"`
	Attempts int `json:"attempts"`
	SuccessRate float64 `json:"success_rate"`
	DifficultyRating int `json:"difficulty_rating"`
	LastAttempted time.Time `json:"last_attempted"`
}