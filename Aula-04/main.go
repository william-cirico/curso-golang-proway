package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/william-cirico/rest-api-golang/infrastructure"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gormDB := infrastructure.NewGormDBConnection()
	srv := infrastructure.NewMuxRouter(gormDB)
	srv.Start()
}
