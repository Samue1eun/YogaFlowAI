package models

// AIFlowResponse represents the generated yoga flow from Claude

type AIFlowResponse struct {
	FlowName           string   `json:"flow_name"`
	Description        string   `json:"description"`
	Duration           int      `json:"duration"`
	FlowType           string   `json:"flow_type"`
	Difficulty         string   `json:"difficulty"`
	PoseSequence       []AIPose `json:"pose_sequence"`
	WarmupInstructions string   `json:"warmup_instructions"`
	CooldownNotes      string   `json:"cooldown_notes"`
}