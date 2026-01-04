package services

import (
	"time"
	// "log"

	"yogaflow.ai/database"
	"yogaflow.ai/models"
)

// Update service based off of New User Model

func CreateUser(newUser models.User) (models.User, error) {
	now := time.Now()
	newUser.CreatedAt = now
	newUser.UpdatedAt = now

	query := `INSERT INTO users (username, email, passwordhash, first_name, last_name, bio, avatar_url, role, is_active)
        	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
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
		newUser.UserType,
		newUser.Tier,
		newUser.IsActive,
	).Scan(&newUser.ID)
	return newUser, err
}

// func UpdateUser(updatedUser models.User) (models.User, error) {
// 	query := `UPDATE users SET username=$1, email=$2, passwordhash=$3, firstname=$4, lastname=$5, bio=$6, avatarurl=$7, role=$8, isactive=$9, updatedat=Now() WHERE id=$10`
// 	_, err := database.Db.Exec(
// 		query,
// 		updatedUser.Username,
// 		updatedUser.Email,
// 		updatedUser.PasswordHash,
// 		updatedUser.FirstName,
// 		updatedUser.LastName,
// 		updatedUser.Bio,
// 		updatedUser.AvatarURL,
// 		updatedUser.Role,
// 		updatedUser.IsActive,
// 		updatedUser.ID,
// 	)
// 	return updatedUser, err
// }

// func DeleteUser(id uint) bool {
// 	_, err := database.Db.Exec("DELETE FROM users WHERE id = $1", id)
// 	if err != nil {
// 		log.Println(err)
// 		return false
// 	}
// 	return true
// }