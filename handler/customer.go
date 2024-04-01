package handler

import (
	"errors"
	"go-import-excel/service"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type CustomerHandler interface {
	Import(c *fiber.Ctx) error
}

type CustomerHandlerImplement struct {
	Service service.CustomerService
}

func (handler CustomerHandlerImplement) Import(c *fiber.Ctx) error {
	file, err := c.FormFile("excel")
	if err != nil {
		if !errors.Is(err, fasthttp.ErrMissingFile) {
			return c.Status(fiber.StatusBadRequest).JSON(Response{Message: err.Error()})
		}
	}

	err = handler.Service.ImportCustomer(c, file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{Message: err.Error()})
	}

	return c.JSON(Response{Message: "oke", Data: nil})
}

func NewCustomerHandler(service service.CustomerService) *CustomerHandlerImplement {
	return &CustomerHandlerImplement{Service: service}
}
