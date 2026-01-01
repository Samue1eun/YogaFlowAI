package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"yogaflow.ai/database"
	"yogaflow.ai/models"
)

// Need to work on the Scheme before working on user_flow handlers

// func GetAllUserFlows(c *gin.Context) {
// 	var userFlows []models.UserFlows
// 	rows, err := database.Db.Query("SELECT id, type, timelength, numberofposes, poselist, FROM user_flows")
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// }