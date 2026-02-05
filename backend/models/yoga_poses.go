package models

// Add in side and as a string
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