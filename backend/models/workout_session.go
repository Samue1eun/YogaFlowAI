package models

import "time"

type WorkoutSession struct {
	ID string `json:"id"`
	UserID int `json:"user_id"`
	YogaFlowID string `json:"yoga_flow_id"`
	StartedAt time.Time `json:"started_at"`
	CompletedAt time.Time `json:"completed_at"`
	Duration int `json:"duration"`
	Rating int `json:"rating"`
	Feedback string `json:"feedback"`
	CreatedAt time.Time `json:"created_at"`
}

// Update PSQL (Created at January 2nd, 2026)