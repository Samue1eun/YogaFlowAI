package services

import (
	"time"

	"yogaflow.ai/models"
	"yogaflow.ai/database"
)

func CreateUser(newUser models.User) (models.User, error) {
	now := time.Now()
	newUser.CreatedAt = now
	newUser.UpdatedAt = now

	query := `INSERT INTO users (username, email, passwordhash, firstname, lastname, bio, avatarurl, role, isactive)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING ID`
	err := database.Db.QueryRow(
		query,
		newUser.Username,
		newUser.Email,
		newUser.PasswordHash,
		newUser.FirstName,
		newUser.LastName,
		newUser.Bio,
		newUser.AvatarURL,
		newUser.Role,
		newUser.IsActive,
	).Scan(&newUser.ID)
	return newUser, err
}

// func UpdateUser() {

// }

// func DeleteUser() {

// }