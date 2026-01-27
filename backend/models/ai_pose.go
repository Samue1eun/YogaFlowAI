package models

// AIPose represents a pose in the AI-generated flow

type AIPose struct {
	Name          string `json:"name"`
	Sanskrit      string `json:"sanskrit"`
	Duration      int    `json:"duration"` // seconds to hold
	Instructions  string `json:"instructions"`
	Modifications string `json:"modifications"` // for injuries or beginners
	Side          string `json:"side"`          // "left", "right", "both", or "center"
}