package models

import "time"

type User struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	PasswordHash string `json:"password_hash"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Bio string `json:"bio"`
	AvatarURL string `json:"avatar_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Role string `json:"role"`
	UserType string `json:"user_type"`
	Tier string `json:"tier"`
	IsActive bool `json:"is_active"`
}

// Update PSQL (Table Created January 2nd, 2026)

// Role is for admin!
// Add in what type of user it is (teacher, student)
// Add in what type of paid user it is (free, premium, teacher)


// User favorites
// Add in favorite yoga poses for user
// Add in favorite yoga poses for user
// Add in favorite yoga flows for user
