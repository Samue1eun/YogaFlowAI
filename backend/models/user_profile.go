package models

import "time"

type UserProfile struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	FitnessLevel int `json:"fitness_level"`
	FlexibilityLevel int `json:"flexibility_level"`
	StrengthLevel int `json:"strength_level"`
	Injuries []string `json:"injuries"`
	Goals []string `json:"goals"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Update PSQL (Updated on January 2nd, 2025)