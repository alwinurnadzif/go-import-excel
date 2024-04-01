package router

import (
	"go-import-excel/handler"
	"go-import-excel/service"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello")
	})

	// customer
	customerService := service.NewCustomerService()
	customerHandler := handler.NewCustomerHandler(customerService)
	customerRoutes := app.Group("/customers")
	customerRoutes.Post("/import", customerHandler.Import)
}
