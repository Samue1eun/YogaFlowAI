package router

import (
	// "fmt"
	// "net/http"
	"github.com/gin-gonic/gin"

	"yogaflow.ai/handlers"
)

func PageRouter() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		// user ROUTES
		v1.GET("/users", handlers.GetAllUsers)
		v1.GET("/users/adminaccess", handlers.GetAllUsersAdmin)
		// v1.GET("/users/:id", handlers.GetOneUser)
		v1.DELETE("/users/:id", handlers.DeleteUser)
		v1.PUT("/users/:id", handlers.UpdateUser)
		v1.POST("/users", handlers.AddUser)

		// yoga_poses ROUTES
		v1.GET("/yoga_poses", handlers.GetAllYogaPoses)
		v1.GET("/yoga_poses/:id", handlers.GetOneYogaPose)
		v1.POST("/yoga_poses", handlers.AddYogaPose)
		v1.POST("/yoga_poses/bulkadd", handlers.AddBulkYogaPoses)
		v1.DELETE("/yoga_poses/:id", handlers.DeleteYogaPose)

		// yoga_flow ROUTES
		v1.GET("/yoga_flow", handlers.GetAllYogaFlows)
		v1.GET("/yoga_flow/:id", handlers.GetOneYogaFlow)
		v1.POST("/yoga_flow", handlers.CreateYogaFlow)
		v1.POST("/yoga_flow/bulkadd", handlers.AddBulkYogaFlows)
		v1.PUT("/yoga_flow/:id", handlers.UpdateYogaFlow)
		v1.DELETE("/yoga_flow/:id", handlers.DeleteYogaFlow)

		// user_profiles ROUTES
		v1.GET("/user_profile", handlers.GetAllUserProfiles)
		// GET ONE User Profile
		// UPDATE User Profile
		v1.POST("/user_profile", handlers.CreateUserProfile)
		v1.DELETE("/user_profile/:id", handlers.DeleteUserProfile)

		// pose_performance ROUTES
		// Note: pose performance is based off of how many reps that a user gets on each pose off of the completed workout. The update feature will be the most important part of this feature.
		// GET ONE Pose Performance Data
		// UPDATE Pose Performance

		// NEED TO TEST THIS END POINT
		v1.GET("/pose_performance", handlers.GetAllPosePerformances)
		v1.POST("/pose_performance", handlers.CreatePosePerformance)
		v1.DELETE("/pose_performance", handlers.DeletePosePerformance)

		// NEED TO TEST THIS END POINT
		// workout_session ROUTES
		// Note: workout sessions will be created and updated once the user completes a workout session
		v1.GET("/workout_session", handlers.GetAllWorkoutSession)
		v1.POST("/workout_session", handlers.CreateWorkoutSession)
		v1.DELETE("/workout_session", handlers.DeleteWorkoutSession)
		
		// user_progress ROUTES
		// Note: user progress will be updated based off of averages that the user will need to get from completed flows.

		// user_favorites ROUTES
		
		// user_flow ROUTES
	}

	r.Run(":8081")
}
