package main

import (
	"aqua-ratsnake/src/controller"
	"aqua-ratsnake/src/database"
	"aqua-ratsnake/src/model"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
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

func app() {
	router := gin.Default()

	userRoutes := router.Group("/user")
	userRoutes.POST("/register", controller.Register)

	router.Run(":9000")
	fmt.Print("Server started on port 9000")
}
func main() {
	loadEnv()
	loadDatabase()
	app()
}
