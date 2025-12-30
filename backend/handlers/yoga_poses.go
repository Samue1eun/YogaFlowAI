package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"yogaflow.ai/database"
	"yogaflow.ai/models"
	"yogaflow.ai/services"
)

func GetAllYogaPoses(c *gin.Context) {
	var yogaPoses []models.YogaPoses
	rows, err := database.Db.Query("SELECT id, name, sanskrit, category, strength, flexibility, difficulty, level FROM yoga_poses")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var yogaPose models.YogaPoses
		err := rows.Scan(&yogaPose.ID, &yogaPose.Name, &yogaPose.Sanskrit, &yogaPose.Category, &yogaPose.Strength, &yogaPose.Flexibility, &yogaPose.Difficulty, &yogaPose.Level)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		yogaPoses = append(yogaPoses, yogaPose)
	}
	c.IndentedJSON(http.StatusOK, yogaPoses)
}

func GetOneYogaPose(c *gin.Context) {
	id := c.Param("id")
	var yogaPose models.YogaPoses
	err := database.Db.QueryRow(
		"SELECT id, name, sanskrit, category, strength, flexibility, difficulty, level FROM yoga_poses WHERE id = $1",
		id,
	).Scan(&yogaPose.ID, &yogaPose.Name, &yogaPose.Sanskrit, &yogaPose.Category, &yogaPose.Strength, &yogaPose.Flexibility, &yogaPose.Difficulty, &yogaPose.Level)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Yoga pose not found"})
		return
	}
	c.JSON(http.StatusOK, yogaPose)
}

func AddYogaPose(c *gin.Context) {
	var newYogaPose models.YogaPoses
	err := c.ShouldBindJSON(&newYogaPose)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	yogaPose, err := services.CreateYogaPose(newYogaPose)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, yogaPose)
}

func DeleteYogaPose(c *gin.Context) {
	id := c.Param("id")
	var yogaPose models.YogaPoses
	err := database.Db.QueryRow("SELECT id, name, sanskrit, category, strength, flexibility, difficulty, level FROM yoga_poses WHERE id = $1", id).
		Scan(&yogaPose.ID, &yogaPose.Name, &yogaPose.Sanskrit, &yogaPose.Category, &yogaPose.Strength, &yogaPose.Flexibility, &yogaPose.Difficulty, &yogaPose.Level)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Yoga pose not found"})
		return
	}

	_, err = database.Db.Exec("DELETE FROM yoga_poses WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Yoga pose deleted", "Yoga pose": yogaPose})
}

// Add bulk yoga poses

func AddBulkYogaPoses(c *gin.Context) {
	var addBulkYogaPoses []models.YogaPoses
	err := c.ShouldBindBodyWithJSON(&addBulkYogaPoses)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO yoga_poses (id, name, sanskrit, category, strength, flexibility, difficulty, level)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	for _, pose := range addBulkYogaPoses {
		_, err := database.Db.Exec(query,
			pose.ID, pose.Name, pose.Sanskrit, pose.Category,
			pose.Strength, pose.Flexibility, pose.Difficulty, pose.Level)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}