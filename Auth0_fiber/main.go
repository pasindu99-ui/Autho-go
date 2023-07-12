package main

import (
	"AUTH0_FIBER/platform/authenticator"
	"AUTH0_FIBER/platform/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	app := fiber.New()

	router.New(app, auth)

	log.Print("Server listening on http://localhost:3000/")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
