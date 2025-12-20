package main

import (
	"yogaflow.ai/database"
	"yogaflow.ai/router"
)

func main() {
	database.ConnectDatabase()
	router.PageRouter()
}