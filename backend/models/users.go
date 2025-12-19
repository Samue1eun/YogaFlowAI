package models

import "time"

type User struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	PasswordHash string `json:"passwordhash"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Bio string `json:"bio"`
	AvatarURL string `json:"avatarurl"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"updatedat"`
	Role string `json:"role"`
	IsActive bool `json:"isactive"`
}