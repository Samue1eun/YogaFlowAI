package router

import (
	"fmt"
	"net/http"

	"yogaflow.ai/handlers"
)

func PageRouter() {
	mainMux := http.NewServeMux()
	mainMux.HandleFunc("/users", handlers.GetAllUser)
	mainMux.HandleFunc("/yoga_poses", handlers.GetAllYogaPoses)

	fmt.Println("Server starting on port :8081")
	err := http.ListenAndServe(":8081", mainMux)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
