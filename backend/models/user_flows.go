package models

import (
	"time"
)

type UserFlows struct {
	ID string `json:"id"`
	UserID int `json:"user_id"`
	YogaFlowID string `json:"yoga_flow_id"`
	CreatedAt time.Time `json:"created_at"`
}

// Update PSQL (Table Created January 2nd, 2026)
