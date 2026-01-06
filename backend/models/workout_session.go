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

// Will need to make a replacement for this model to make an array for the yoga flows to cater to the workout session demands.
// The minimum for a yoga session should be 45 minutes

// Update PSQL (Table Created January 2nd, 2026)