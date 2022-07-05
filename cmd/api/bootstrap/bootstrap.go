package bootstrap

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"simci-go/internal/platform/server"
	infrastructure "simci-go/internal/platform/server/storage/postgresql"
)

func Run() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_HOST := os.Getenv("POSTGRES_HOST")
	POSTGRES_DATABASE := os.Getenv("POSTGRES_DATABASE")
	POSTGRES_PORT := os.Getenv("POSTGRES_PORT")

	dsnTemplate := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"

	dsn := fmt.Sprintf(dsnTemplate, POSTGRES_HOST, POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DATABASE, POSTGRES_PORT)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}

	userRepository := infrastructure.NewUserRepository(db)

	srv := server.New("127.0.0.1", 4000, userRepository)

	return srv.Run()
}
