package models

var Strength = [5]int{1, 2, 3, 4, 5}
var Flexibility = [5]int{1, 2, 3, 4, 5}
var Dificulty = [5]int{1, 2, 3, 4, 5}
var Level = [5]int{1, 2, 3, 4, 5}

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