package main

import (
	"log"
	"simci-go/core"
	"simci-go/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := core.SetupModels()

	if err != nil {
		log.Fatal("Error connect to database")
	}

	r := gin.Default()

	apiRouter := r.Group("/api")

	routers.InitRoutes(apiRouter, db)

	r.Run(":5000")
}
