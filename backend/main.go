package main

import (
	"yogaflow.ai/ai"
	"yogaflow.ai/database"
	"yogaflow.ai/router"
)

func main() {
	database.ConnectDatabase()
	ai.InitClient()
	router.PageRouter()
}
