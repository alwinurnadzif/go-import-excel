package utils

import (
	"errors"
	"fmt"
	"mime/multipart"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

var ImagesContentTypes map[string]bool = map[string]bool{"image/jpg": true, "image/jpeg": true, "image/png": true}
var ExcelContentTypes map[string]bool = map[string]bool{"application/vnd.openxmlformats-officedocument.spreadsheetml": true, "application/vnd.ms-excel": true}

type ValidationErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Value   string `json:"value"`
}

type ValidationErrorDetailResponse struct {
	Index   uint   `json:"index"`
	Field   string `json:"field"`
	Message string `json:"message"`
	Value   string `json:"value"`
}

var ErrValidationDetail = errors.New("validation error details")
var ErrValidation = errors.New("validation error")

func ValidateStruct(payload interface{}) []*ValidationErrorResponse {

	var errors []*ValidationErrorResponse

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitAfterN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var elem ValidationErrorResponse
			elem.Field = err.Field()
			elem.Message = err.Tag()
			elem.Value = err.Param()

			errors = append(errors, &elem)
		}
	}

	return errors

}

func ValidateUniqueValue(currentValue interface{}, target interface{}, field string, errMessageField string) interface{} {

	errResponse := []ValidationErrorResponse{}

	if currentValue == target {
		errMessage := fmt.Sprintf("%s sudah digunakan", errMessageField)
		err := ValidationErrorResponse{
			Field:   field,
			Message: errMessage,
			Value:   fmt.Sprintf("%v", currentValue),
		}
		errResponse = append(errResponse, err)
		return &errResponse
	}

	return nil
}

func ValidateFileSize(file *multipart.FileHeader, maxFileSizeMb uint, field string) interface{} {

	maxFileSizeByte := maxFileSizeMb * 1024 * 1024

	errResponses := []ValidationErrorResponse{}
	if file.Size > int64(maxFileSizeByte) {
		errResponse := ValidationErrorResponse{
			Field:   field,
			Message: "ukuran file terlalu besar",
		}

		errResponses = append(errResponses, errResponse)
		return errResponses
	}
	return nil
}

func ValidateFileContentType(file *multipart.FileHeader, allowedContentTypes map[string]bool, field string) interface{} {
	contentType := file.Header.Get("Content-Type")

	errResponses := []ValidationErrorResponse{}
	if !allowedContentTypes[contentType] {
		errResponse := ValidationErrorResponse{
			Field:   field,
			Message: "file tidak didukung",
		}

		errResponses = append(errResponses, errResponse)
		return errResponses
	}

	return nil
}

func Validate(payload interface{}, validationError *[]ValidationErrorResponse) error {

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitAfterN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	err := validate.Struct(payload)
	errors := []bool{}
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var elem ValidationErrorResponse
			elem.Field = err.Field()
			elem.Message = err.Tag()
			elem.Value = err.Param()

			(*validationError) = append((*validationError), elem)
			errors = append(errors, true)
		}
	}

	for _, e := range errors {
		if e {
			return ErrValidation
		}
	}

	return nil

}
