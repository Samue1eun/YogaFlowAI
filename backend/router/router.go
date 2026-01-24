package router

import (
	// "fmt"
	// "net/http"
	"github.com/gin-gonic/gin"

	"yogaflow.ai/handlers"
	"yogaflow.ai/middleware"
)

func PageRouter() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		// AUTH ROUTES (public)
		v1.POST("/auth/register", handlers.Register)
		v1.POST("/auth/login", handlers.Login)

		// user ROUTES (all endpoints satisfied)
		v1.GET("/users", handlers.GetAllUsers)
		v1.GET("/users/adminaccess", middleware.AuthMiddleware(), middleware.AdminMiddleware(), handlers.GetAllUsersAdmin)
		v1.GET("/auth/me", middleware.AuthMiddleware(), handlers.GetMe)
		// v1.GET("/users/:id", handlers.GetOneUser)
		v1.DELETE("/users/:id", handlers.DeleteUser)
		v1.PUT("/users/:id", handlers.UpdateUser)
		v1.POST("/users", handlers.AddUser)

		// yoga_poses ROUTES (all endpoints satisfied)
		v1.GET("/yoga_poses", handlers.GetAllYogaPoses)
		v1.GET("/yoga_poses/:id", handlers.GetOneYogaPose)
		v1.POST("/yoga_poses", handlers.AddYogaPose)
		v1.POST("/yoga_poses/bulkadd", handlers.AddBulkYogaPoses)
		v1.DELETE("/yoga_poses/:id", handlers.DeleteYogaPose)

		// yoga_flow ROUTES (all endpoints satisfied)
		v1.GET("/yoga_flow", handlers.GetAllYogaFlows)
		v1.GET("/yoga_flow/:id", handlers.GetOneYogaFlow)
		v1.POST("/yoga_flow", handlers.CreateYogaFlow)
		v1.POST("/yoga_flow/bulkadd", handlers.AddBulkYogaFlows)
		v1.PUT("/yoga_flow/:id", handlers.UpdateYogaFlow)
		v1.DELETE("/yoga_flow/:id", handlers.DeleteYogaFlow)

		// user_profiles ROUTES (all endpoints satisfied)
		v1.GET("/user_profile", handlers.GetAllUserProfiles)
		v1.GET("/user_profile/:id", handlers.GetOneUserProfile)
		v1.PUT("/user_profile/:id", handlers.UpdateUserProfile)
		v1.POST("/user_profile", handlers.CreateUserProfile)
		v1.DELETE("/user_profile/:id", handlers.DeleteUserProfile)

		// pose_performance ROUTES (all endpoints satisfied)
		// Note: pose performance is based off of how many reps that a user gets on each pose off of the completed workout. The update feature will be the most important part of this feature.
		v1.GET("/pose_performance", handlers.GetAllPosePerformances)
		v1.GET("/pose_performance/:id", handlers.GetOnePosePerformance)
		v1.PUT("/pose_performance/:id", handlers.UpdatePosePerformance)
		v1.POST("/pose_performance", handlers.CreatePosePerformance)
		v1.DELETE("/pose_performance/:id", handlers.DeletePosePerformance)

		// workout_session ROUTES (all endpoints satisfied)
		// Note: workout sessions will be created and updated once the user completes a workout session
		v1.GET("/workout_session", handlers.GetAllWorkoutSession)
		// GET ONE WORKOUT SESSION
		// UPDATE WORKOUT SESSION
		v1.POST("/workout_session", handlers.CreateWorkoutSession)
		v1.DELETE("/workout_session/:id", handlers.DeleteWorkoutSession)

		// user_progress ROUTES (all endpoints satisfied)
		// Note: user progress will be updated based off of averages that the user will need to get from completed flows.
		v1.GET("/user_progress", handlers.GetAllUserProgress)
		v1.GET("/user_progress/:id", handlers.GetOneUserProgress)
		v1.POST("/user_progress", handlers.CreateUserProgress)
		v1.PUT("/user_progress/:id", handlers.UpdateUserProgress)
		v1.DELETE("/user_progress/:id", handlers.DeleteUserProgress)

		// user_favorites ROUTES (all endpoints satisfied)
		v1.GET("/user_favorites", handlers.GetAllUserFavorites)
		v1.PUT("/user_favorites/:id", handlers.UpdateUserFavorite)
		v1.POST("/user_favorites/", handlers.CreateUserFavorite)
		v1.DELETE("/user_favorites/:id", handlers.DeleteUserFavorite)

		// user_flows ROUTES (all endpoints satisfied)
		// User flows do not need an update feature since a user will have a new flow created each time
		v1.GET("/user_flows", handlers.GetAllUserFlows)
		v1.GET("/user_flows/:id", handlers.GetOneUserFlow)
		v1.POST("/user_flows", handlers.CreateUserFlow)
		v1.DELETE("/user_flows/:id", handlers.DeleteUserFlow)

		// AI FLOW GENERATION ROUTES
		// Generate personalized yoga flows using Claude AI
		v1.POST("/ai/generate-flow", handlers.GenerateAIFlow)
		v1.GET("/ai/quick-flow", handlers.QuickGenerateAIFlow)
	}

	r.Run(":8081")
}
