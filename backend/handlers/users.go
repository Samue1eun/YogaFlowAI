package handlers

import (
	// "encoding/json"
	"net/http"
	// "net/http"
	"github.com/gin-gonic/gin"

	"yogaflow.ai/database"
	"yogaflow.ai/models"
	"yogaflow.ai/services"
)

// Update User handler

//Get All Users (Gin)

func GetAllUsers(c *gin.Context) {
	var users []models.User
	rows, err := database.Db.Query("SELECT id, username, email, firstname, lastname, bio, avatarurl, role, usertype, tier, isactive FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.FirstName, &user.LastName, &user.Bio, &user.AvatarURL, &user.Role, &user.Tier, &user.IsActive)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, user)
	}

	c.IndentedJSON(http.StatusOK, users)
}

// func GetOneUser(c *gin.Context) {
// 	id := c.Param("id")
// 	var user models.User
// 	rows, err := database.Db.Query("SELECT id, username, email, firstname, lastname, bio, avatarurl, role, isactive FROM users WHERE id = $1", id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	defer rows.Close()

// 	if rows != nil {
// 		for rows.Next() {
// 			var (
// 				id        int
// 				username  string
// 				email     string
// 				firstname string
// 				lastname  string
// 				bio       string
// 				avatarurl string
// 				role      string
// 				isactive  string
// 			)
// 			err := rows.Scan(&id, &username, &email, &firstname, &lastname, &bio, &avatarurl, &role, &isactive)
// 			if err != nil {
// 				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 				return
// 			}
// 			var isActiveBool bool
// 			if isactive == "true" || isactive == "1" {
// 				isActiveBool = true
// 			} else {
// 				isActiveBool = false
// 			}
// 			user = models.User{ID: id, Username: username, Email: email, FirstName: firstname, LastName: lastname, Bio: bio, AvatarURL: avatarurl, Role: role, IsActive: isActiveBool}
// 		}
// 	}
// 	c.JSON(http.StatusOK, user)
// }

// // Add User
// func AddUser(c *gin.Context) {
// 	var newUser models.User
// 	err := c.ShouldBindJSON(&newUser)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	user, err := services.CreateUser(newUser)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, user)
// }

// // Update User
// func UpdateUser(c *gin.Context) {
// 	var updateUser models.User
// 	err := c.ShouldBindJSON(&updateUser)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	user, err := services.UpdateUser(updateUser)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, user)
// }

// // Delete User
// func DeleteUser(c *gin.Context) {
// 	id := c.Param("id")
// 	var user models.User
// 	err := database.Db.QueryRow("SELECT id, username, email, firstname, lastname, bio, avatarurl, role, isactive FROM users WHERE id = $1", id).
// 		Scan(&user.ID, &user.Username, &user.Email, &user.FirstName, &user.LastName, &user.Bio, &user.AvatarURL, &user.Role, &user.IsActive)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}

// 	_, err = database.Db.Exec("DELETE FROM users WHERE id = $1", id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "User deleted", "user": user})
// }


