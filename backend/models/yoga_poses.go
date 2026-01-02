package models

type YogaPoses struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Sanskrit string `json:"sanskrit"`
	Category string `json:"category"`
	Strength int `json:"strength"`
	Flexibility int `json:"flexibility"`
	Difficulty int `json:"difficulty"`
	Level int `json:"level"`
}

// Update PSQL (Updated January 2nd, 2025)