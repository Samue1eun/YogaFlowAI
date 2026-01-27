package models

// AIFlowRequest represents the request to generate a yoga flow

type AIFlowRequest struct {
	UserID      int    `json:"user_id"`
	Duration    int    `json:"duration"`    // in minutes (e.g., 30, 45, 60)
	FlowType    string `json:"flow_type"`   // e.g., "vinyasa", "hatha", "restorative", "power"
	FocusArea   string `json:"focus_area"`  // e.g., "flexibility", "strength", "relaxation", "balance"
	Difficulty  string `json:"difficulty"`  // e.g., "beginner", "intermediate", "advanced"
	Description string `json:"description"` // optional free-form description
}