package main

import (
	"gin_api_rest/models"
	"gin_api_rest/routes"
	"log"

	"github.com/gin-gonic/gin"

	db "gin_api_rest/config"
)

func main() {
	r := gin.Default()

	db, err := db.GetDb()

	if err != nil {
		log.Fatal("Error connecting to database")
	}
	db.AutoMigrate(&models.User{})

	routes.SetupRoutes(r)

	r.Run(":8080")
}
