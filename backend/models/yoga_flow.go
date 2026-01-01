package models

type YogaFlow struct {
	ID string `json:"id"`
	Type string `json:"type"`
	TimeLength int `json:"timelength"`
	NumberOfPoses int `json:"numberofposes"`
	PoseList []YogaPoses `json:"poselist"`
	AverageStrength int `json:"averagestrength"`
	AverageFlexibility int `json:"averageflexibility"`
	AverageDifficulty int `json:"averagedifficulty"`
}