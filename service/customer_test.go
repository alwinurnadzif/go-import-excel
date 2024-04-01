package service

import (
	"testing"

	"github.com/gofiber/fiber/v2/log"
)

func TestH(t *testing.T) {

	customerField := struct {
		Name uint
		Age  uint
	}{
		Name: 0,
		Age:  1,
	}

	log.Info(customerField)
}
