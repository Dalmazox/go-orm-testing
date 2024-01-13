package main

import (
	"fmt"
	"os"

	"dalmazox/go/cmd/server/controllers"
	"dalmazox/go/domain/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}

	port := os.Getenv("PORT")
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)

	r.Run(fmt.Sprintf(":%s", port))
}
