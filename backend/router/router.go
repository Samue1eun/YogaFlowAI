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
		// /user ROUTES
		v1.GET("/users", handlers.GetAllUsers)
		v1.GET("/users/adminaccess", handlers.GetAllUsersAdmin)
		// v1.GET("/users/:id", handlers.GetOneUser)
		v1.DELETE("/users/:id", handlers.DeleteUser)
		v1.PUT("/users/:id", handlers.UpdateUser)
		v1.POST("/users", handlers.AddUser)

		// /yoga_poses ROUTES
		v1.GET("/yoga_poses", handlers.GetAllYogaPoses)
		v1.GET("/yoga_poses/:id", handlers.GetOneYogaPose)
		v1.POST("/yoga_poses", handlers.AddYogaPose)
		v1.POST("/yoga_poses/bulkadd", handlers.AddBulkYogaPoses)
		v1.DELETE("/yoga_poses/:id", handlers.DeleteYogaPose)

		// /yoga_flow ROUTES
		v1.GET("/yoga_flow", handlers.GetAllYogaFlows)
		v1.GET("/yoga_flow/:id", handlers.GetOneYogaFlow)
		v1.POST("/yoga_flow", handlers.CreateYogaFlow)
		v1.POST("/yoga_flow/bulkadd", handlers.AddBulkYogaFlows)
		v1.PUT("/yoga_flow/:id", handlers.UpdateYogaFlow)
		v1.DELETE("/yoga_flow/:id", handlers.DeleteYogaFlow)

		// user_profiles ROUTES
		v1.GET("/user_profile", handlers.GetAllUserProfiles)
		v1.
		
	}

	r.Run(":8081")
}
