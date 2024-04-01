package main

import (
	"go-import-excel/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.SetupRoutes(app)

	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("failed run the app: %v", err)
	}
}
