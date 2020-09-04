package main

import (
	"github.com/joho/godotenv"
	"mostlikelyto/src/infrastructure/persistence/mysql/database"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	mysql := database.New(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SCHEMA"),
		)

	db := mysql.Connect()
	defer db.Close()

	database.MigrateDatabase(db)

}
