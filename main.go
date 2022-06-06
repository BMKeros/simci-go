package main

import (
	"log"
	"simci-go/core"
	"simci-go/docs"
	"simci-go/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// swagger
	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":5000")
}
