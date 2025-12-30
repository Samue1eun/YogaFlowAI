package router

import (
	// "fmt"
	// "net/http"
	"github.com/gin-gonic/gin"

	"yogaflow.ai/handlers"
	// "yogaflow.ai/services"
)

func PageRouter() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/users", handlers.GetAllUsers)
		v1.GET("/users/:id", handlers.GetOneUser)
		v1.DELETE("/users/:id", handlers.DeleteUser)
		v1.PUT("/users/:id", handlers.UpdateUser)
		v1.POST("/users", handlers.AddUser)
		v1.GET("/yoga_poses", handlers.GetAllYogaPoses)
		v1.GET("/yoga_poses/:id", handlers.GetOneYogaPose)
		v1.POST("/yoga_poses", handlers.AddYogaPose)
		v1.DELETE("/yoga_poses/:id", handlers.DeleteYogaPose)
	}
	
	r.Run(":8081")
}
