package models

import "time"

type UserProgress struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	Date time.Time `json:"date"`
	StrengthImprovement float64 `json:"strength_improvement"`
	FlexibilityImprovement float64 `json:"flexibility_improvement"`
	SessionsCompleted int `json:"sessions_completed"`
	TotalTime int `json:"total_time"`
}