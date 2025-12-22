package handlers

import (
	"encoding/json"
	"net/http"

	"yogaflow.ai/models"
	"yogaflow.ai/database"
)

// Get All Yoga Poses
func GetAllYogaPoses(w http.ResponseWriter, r *http.Request) {
	var yogaPoses []models.YogaPoses
	rows, err := database.Db.Query("SELECT id, name, sanskrit, category, strength, flexibility, difficulty, level FROM yoga_poses")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var yogaPose models.YogaPoses
		err := rows.Scan(&yogaPose.ID, &yogaPose.Name, &yogaPose.Sanskrit, &yogaPose.Category, &yogaPose.Strength, &yogaPose.Flexibility, &yogaPose.Difficulty, &yogaPose.Level)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		yogaPoses = append(yogaPoses, yogaPose)
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(yogaPoses)
}