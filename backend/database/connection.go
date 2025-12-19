package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
	"strconv"
)

var Db *sql.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error: Database was not successfully loaded")
	}

	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("post"))
	user := os.Getenv("USER")
	db_name := os.Getenv("DB_NAME")
	pass := os.Getenv("PASSWORD")

	psqlSetup := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
	host, port, user, db_name, pass)
	db, errSql := sql.Open("postgres", psqlSetup)
	if errSql != nil {
		fmt.Println("There was an error while trying to connect your database")
	} else {
		Db = db
		fmt.Println("Successfully connected to database")
	}
}