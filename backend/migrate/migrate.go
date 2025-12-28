package main

import (
	"go-jwt-auth/initializers"
	"go-jwt-auth/models"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
}

func main() {
	
}