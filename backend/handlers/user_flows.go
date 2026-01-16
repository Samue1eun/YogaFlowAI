package handlers

import (
	"net/http"
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	"yogaflow.ai/database"
	"yogaflow.ai/models"
	"yogaflow.ai/services"
)

func GetAllUserFlows(c *gin.Context) {
	var userFlows []models.UserFlows
	rows, err := database.Db.Query("SELECT id, user_id, yoga_flow_id, created_at FROM user_flows")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	
	for rows.Next() {
		var userFlow models.UserFlows
		err := rows.Scan(
			&userFlow.ID,
			&userFlow.UserID,
			&userFlow.YogaFlowID,
			&userFlow.CreatedAt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		userFlows = append(userFlows, userFlow)
	}
	c.IndentedJSON(http.StatusOK, userFlows)
}

func GetOneUserFlow(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid"})
	}
	var userFlow models.UserFlows
	err = database.Db.QueryRow(
		"SELECT id, user_id, yoga_flow_id, created_at FROM user_profile WHERE id=$1",
		id,
	).Scan(
		&userFlow.ID,
		&userFlow.UserID,
		&userFlow.YogaFlowID,
		&userFlow.CreatedAt,
	)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "User flow now found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User flow now found"})
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func CreateUserFlow(c *gin.Context) {
	var newUserFlow models.UserFlows
	err := c.ShouldBindJSON(&newUserFlow)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userFlow, err := services.CreateUserFlow(newUserFlow)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, userFlow)
}

func DeleteUserFlow(c *gin.Context) {
	id := c.Param("id")
	var userFlows models.UserFlows
	_, err := database.Db.Exec("DELETE FROM user_flows WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User flow deleted", "user flow": userFlows})
}