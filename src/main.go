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
	userRoutes.GET("/list", controller.List)
	userRoutes.GET("/get/:id", controller.Get)
	userRoutes.PATCH("/update/:id", controller.Update)
	userRoutes.DELETE("/delete/:id", controller.Delete)

	router.Run(":9000")
	fmt.Print("Server started on port 9000")
}
func main() {
	loadEnv()
	loadDatabase()
	app()
}
