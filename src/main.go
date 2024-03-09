package main

import (
	"aqua-ratsnake/src/database"
	"aqua-ratsnake/src/model"
	"log"

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	// Connect to the database
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
}
func main() {
	loadEnv()
	loadDatabase()
}
