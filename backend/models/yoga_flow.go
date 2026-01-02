package models

type YogaFlow struct {
	ID string `json:"id"`
	Type string `json:"type"`
	TimeLength int `json:"time_length"`
	NumberOfPoses int `json:"number_of_poses"`
	PoseList []YogaPoses `json:"pose_list"`
	AverageStrength int `json:"average_strength"`
	AverageFlexibility int `json:"average_flexibility"`
	AverageDifficulty int `json:"average_difficulty"`
}

// Update PSQL (January 2nd, 2026)