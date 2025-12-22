package handlers

import (
	"encoding/json"
	"net/http"

	"yogaflow.ai/database"
	"yogaflow.ai/models"
)

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	rows, err := database.Db.Query("SELECT id, username, email, firstname, lastname, bio, avatarurl, role, isactive FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.FirstName, &user.LastName, &user.Bio, &user.AvatarURL, &user.Role, &user.IsActive)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func GetAllUserAdmin(w http.ResponseWriter, r*http.Request) {
	var users[[]models.User]
	rows, err := database.Db.Query("SELECT"
	)
}

func GetOneUser() {

}
