package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/liushuangls/go-anthropic/v2"
	"yogaflow.ai/ai"
	"yogaflow.ai/database"
	"yogaflow.ai/models"
)

// PUT THESE STRUCTS INTO THEIR OWN MODELS (ALL UNDER ONE FILE)

// AIFlowRequest represents the request to generate a yoga flow
type AIFlowRequest struct {
	UserID      int    `json:"user_id"`
	Duration    int    `json:"duration"`    // in minutes (e.g., 30, 45, 60)
	FlowType    string `json:"flow_type"`   // e.g., "vinyasa", "hatha", "restorative", "power"
	FocusArea   string `json:"focus_area"`  // e.g., "flexibility", "strength", "relaxation", "balance"
	Difficulty  string `json:"difficulty"`  // e.g., "beginner", "intermediate", "advanced"
	Description string `json:"description"` // optional free-form description
}

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

// AIPose represents a pose in the AI-generated flow
type AIPose struct {
	Name          string `json:"name"`
	Sanskrit      string `json:"sanskrit"`
	Duration      int    `json:"duration"` // seconds to hold
	Instructions  string `json:"instructions"`
	Modifications string `json:"modifications"` // for injuries or beginners
	Side          string `json:"side"`          // "left", "right", "both", or "center"
}

// GenerateAIFlow generates a personalized yoga flow using Claude AI
func GenerateAIFlow(req AIFlowRequest) (*AIFlowResponse, error) {
	client := ai.GetClient()

	// Get user profile for personalization
	var userProfile models.UserProfile
	var injuries, goals []byte
	err := database.Db.QueryRow(
		"SELECT fitness_level, flexibility_level, strength_level, injuries, goals FROM user_profile WHERE user_id = $1",
		req.UserID,
	).Scan(
		&userProfile.FitnessLevel,
		&userProfile.FlexibilityLevel,
		&userProfile.StrengthLevel,
		&injuries,
		&goals,
	)

	// Parse injuries and goals if user profile exists
	var injuriesList, goalsList []string
	if err == nil {
		json.Unmarshal(injuries, &injuriesList)
		json.Unmarshal(goals, &goalsList)
	}

	// Build the prompt for Claude
	prompt := buildFlowPrompt(req, userProfile, injuriesList, goalsList)

	// Call Claude API
	resp, err := client.CreateMessages(context.Background(), anthropic.MessagesRequest{
		Model: anthropic.ModelClaude3Dot5SonnetLatest,
		Messages: []anthropic.Message{
			anthropic.NewUserTextMessage(prompt),
		},
		MaxTokens: 3999,
	})
	if err != nil {
		return nil, fmt.Errorf("claude API error: %w", err)
	}

	// Parse the response
	if len(resp.Content) == 0 {
		return nil, fmt.Errorf("empty response from Claude")
	}

	responseText := resp.Content[0].GetText()

	// Parse JSON from Claude's response
	var flowResponse AIFlowResponse
	err = json.Unmarshal([]byte(responseText), &flowResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Claude response: %w, response: %s", err, responseText)
	}

	return &flowResponse, nil
}

func buildFlowPrompt(req AIFlowRequest, profile models.UserProfile, injuries, goals []string) string {
	prompt := fmt.Sprintf(`You are an expert yoga instructor. Create a personalized yoga flow based on the following requirements.

REQUIREMENTS:
- Duration: %d minutes
- Flow Type: %s
- Focus Area: %s
- Difficulty Level: %s
`, req.Duration, req.FlowType, req.FocusArea, req.Difficulty)

	if req.Description != "" {
		prompt += fmt.Sprintf("- Additional Notes: %s\n", req.Description)
	}

	// Add user profile info if available
	if profile.FitnessLevel > 0 {
		prompt += fmt.Sprintf(`
USER PROFILE:
- Fitness Level: %d/10
- Flexibility Level: %d/10
- Strength Level: %d/10
`, profile.FitnessLevel, profile.FlexibilityLevel, profile.StrengthLevel)
	}

	if len(injuries) > 0 {
		prompt += fmt.Sprintf("- Injuries/Limitations: %v (IMPORTANT: Avoid poses that aggravate these)\n", injuries)
	}

	if len(goals) > 0 {
		prompt += fmt.Sprintf("- Goals: %v\n", goals)
	}

	prompt += `
RESPONSE FORMAT:
Return ONLY valid JSON (no markdown, no explanation) in this exact format:
{
  "flow_name": "Name of the flow",
  "description": "Brief description of the flow and its benefits",
  "duration": 45,
  "flow_type": "vinyasa",
  "difficulty": "intermediate",
  "warmup_instructions": "Instructions for warming up before the flow",
  "cooldown_notes": "Notes for cooling down after the flow",
  "pose_sequence": [
    {
      "name": "Mountain Pose",
      "sanskrit": "Tadasana",
      "duration": 30,
      "instructions": "Stand tall with feet together, arms at sides...",
      "modifications": "For balance issues, stand with feet hip-width apart",
      "side": "center"
    }
  ]
}

GUIDELINES:
1. Include 10-20 poses depending on duration
2. Start with gentle warm-up poses
3. Build intensity gradually
4. Include both sides for asymmetric poses (specify "left" then "right")
5. End with cooling poses and savasana
6. Respect any injuries mentioned - provide safe alternatives
7. Match the difficulty level appropriately
8. Duration in pose_sequence is in SECONDS

Generate the yoga flow now:`

	return prompt
}
