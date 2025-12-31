package models

type UserFlows struct {
	ID string `json:"id"`
	Type string `json:"type"`
	TimeLength int `json:"timelength"`
	NumberOfPoses int `json:"numberofposes"`
	PoseList []YogaPoses `json:"poselist"`
	UserID int `json:"user_id"`
}

func (uf *UserFlows) UpdateNumberOfPoses() {
	uf.NumberOfPoses = len(uf.PoseList)
}


