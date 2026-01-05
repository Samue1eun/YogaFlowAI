package services

import (
	"time"
	"log"

	"yogaflow.ai/database"
	"yogaflow.ai/models"
)

// Update service based off of New User Model

func CreateUser(newUser models.User) (models.User, error) {
	now := time.Now()
	newUser.CreatedAt = now
	newUser.UpdatedAt = now

	query := `INSERT INTO users (id, username, email, password_hash, first_name, last_name, bio, avatar_url, created_at, updated_at, role, user_type, tier, is_active)
        	VALUES (DEFAULT, $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id`
	err := database.Db.QueryRow(
		query,
		newUser.Username,
		newUser.Email,
		newUser.PasswordHash,
		newUser.FirstName,
		newUser.LastName,
		newUser.Bio,
		newUser.AvatarURL,
		now,
		now,
		newUser.Role,
		newUser.UserType,
		newUser.Tier,
		newUser.IsActive,
	).Scan(&newUser.ID)
	return newUser, err
}

func UpdateUser(updatedUser models.User) (models.User, error) {
	now := time.Now()
	updatedUser.UpdatedAt = now
	
	query := `UPDATE users SET username=$1, email=$2, password_hash=$3, first_name=$4, last_name=$5, bio=$6, avatar_url=$7, updated_at=Now(), role=$8, user_type=$9, tier=$10, is_active=$11 WHERE id=$12`
	_, err := database.Db.Exec(
		query,
		updatedUser.Username,
		updatedUser.Email,
		updatedUser.PasswordHash,
		updatedUser.FirstName,
		updatedUser.LastName,
		updatedUser.Bio,
		updatedUser.AvatarURL,
		updatedUser.Role,
		updatedUser.UserType,
		updatedUser.Tier,
		updatedUser.IsActive,
		updatedUser.ID,
	)
	return updatedUser, err
}

func DeleteUser(id uint) bool {
	_, err := database.Db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}