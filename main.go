package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	var a App
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Yellow")
		log.Println(err.Error())
	}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)

	a.Run(":8010")
}
