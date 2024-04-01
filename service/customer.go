package service

import (
	"fmt"
	"go-import-excel/config"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/xuri/excelize/v2"
)

type CustomerService interface {
	ImportCustomer(c *fiber.Ctx, importFile *multipart.FileHeader) error
}

type CustomerServiceImplement struct{}

func (s CustomerServiceImplement) ImportCustomer(c *fiber.Ctx, importFile *multipart.FileHeader) error {

	// save file
	fileName := "customer.xlsx"
	pathSaveFile := filepath.Join(config.ProjectRootPath, "/public/import-excel/", fileName)
	if err := c.SaveFile(importFile, pathSaveFile); err != nil {
		return err
	}

	// read excel file
	spreadsheet, err := excelize.OpenFile(pathSaveFile)
	if err != nil {
		return err
	}

	rows, err := spreadsheet.GetRows("Sheet1")
	if err != nil {
		return err
	}

	customerField := struct {
		Name uint
		Age  uint
	}{
		Name: 0,
		Age:  1,
	}

	for index, row := range rows {
		if index == 0 {
			continue
		}
		log.Infof("name: %s", row[customerField.Name])
		log.Infof("age: %s", row[customerField.Age])
	}

	defer func() {
		log.Info("close and remove file")
		if err := spreadsheet.Close(); err != nil {
			fmt.Println(err)
		}

		if err := os.Remove(pathSaveFile); err != nil {
			log.Info("err: %s", err)
		}
	}()

	return nil
}

func NewCustomerService() *CustomerServiceImplement {
	return &CustomerServiceImplement{}
}

type CustomerExcelField struct {
	Field string
	Index uint
}
