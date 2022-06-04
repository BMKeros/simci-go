package core

import (
	"fmt"
	"log"
	"os"
	"simci-go/core/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupModels() (*gorm.DB, error) {
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_HOST := os.Getenv("POSTGRES_HOST")
	POSTGRES_DATABASE := os.Getenv("POSTGRES_DATABASE")
	POSTGRES_PORT := os.Getenv("POSTGRES_PORT")

	dsnTemplate := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"

	dsn := fmt.Sprintf(dsnTemplate, POSTGRES_HOST, POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DATABASE, POSTGRES_PORT)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = db.AutoMigrate(&models.User{}); err != nil {
		log.Println(err)
	}

	return db, err
}
