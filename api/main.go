package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Pradumnasaraf/url-short/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:shortPath", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Unable to load .env file. Lodding from ENV if available")
	}
	app := fiber.New()

	app.Use(logger.New())

	setupRoutes(app)
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
