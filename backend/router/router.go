package router

import (
	"net/http"
	"fmt"
)

func PageRouter() {
	mainMux := http.NewServeMux()

	v1 := http.NewServeMux()
	



	fmt.Println("Server starting on port :8081")
	err := http.ListenAndServe(":8081", mainMux)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}