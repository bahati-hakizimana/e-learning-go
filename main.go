package main

import (
	"log"
	"os"

	"github.com/bahati-hakizimana/e-learning-backend/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	database.Connect()

	err:=godotenv.Load()

	if err != nil {

		log.Fatal("Could not connect to database")
		
	}

	port:=os.Getenv("PORT")
    app:=fiber.New()
	app.Listen(":"+port)
}