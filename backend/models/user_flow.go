package models

type UserFlows struct {
	ID string `json:"id"`
	NumberOfFlows int `json:"numberofflows"`
	UserFlowList []YogaFlow `json:"flowlist"`
	UserID int `json:"user_id"`
}

func (uf *UserFlows) UpdateNumberOfPoses() {
	uf.NumberOfFlows = len(uf.UserFlowList)
}


