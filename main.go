package main

import (
	"go-rest-api-mock/app"
	"go-rest-api-mock/config"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error .env file at %s", err)
		panic(err)
	}

	err = config.InitGorm()
	if err != nil {
		panic(err)
	}
}

func main() {
	app.StartApp()
}
