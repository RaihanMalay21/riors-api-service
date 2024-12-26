package service

import (
	"errors"
	"mime/multipart"
	"path/filepath"

	"github.com/RaihanMalay21/api-service-riors/dto"
	"github.com/RaihanMalay21/api-service-riors/validation"
	"github.com/go-playground/validator/v10"
)

func ValidateStructProduct(data *dto.Product, response map[string]interface{}) error {
	trans := validation.TranslatorIDN()
	validate := validator.New()
	validation.RegisterCustomValidationsProduct(validate, trans)

	if err := validate.Struct(data); err != nil {
		var errsMessage []map[string]string
		for _, err := range err.(validator.ValidationErrors) {
			errField := err.StructField()
			errTranslate := err.Translate(trans)
			if errField == "Ext" || errField == "FileSize" {
				errField = "Image"
			}
			errs := map[string]string{
				errField: errTranslate,
			}
			errsMessage = append(errsMessage, errs)
		}
		response["ErrorField"] = errsMessage
		return err
	}

	return nil
}

func ValidateStructRegister(data *dto.RegisterUser, response map[string]interface{}) error {
	trans := validation.TranslatorIDN()
	validate := validator.New()
	validation.RegisterCustomValidationsProduct(validate, trans)

	if err := validate.Struct(data); err != nil {
		var errsMessage []map[string]string
		for _, err := range err.(validator.ValidationErrors) {
			errField := err.StructField()
			errTranslate := err.Translate(trans)
			errs := map[string]string{
				errField: errTranslate,
			}
			errsMessage = append(errsMessage, errs)
		}
		response["ErrorField"] = errsMessage
		return err
	}

	return nil
}

func ValidateStructEmployee(data *dto.Employee, response map[string]interface{}) error {
	trans := validation.TranslatorIDN()
	validate := validator.New()
	validation.RegisterCustomValidationsProduct(validate, trans)

	if err := validate.Struct(data); err != nil {
		var errsMessage []map[string]string
		for _, err := range err.(validator.ValidationErrors) {
			errField := err.StructField()
			errTranslate := err.Translate(trans)
			errs := map[string]string{
				errField: errTranslate,
			}
			errsMessage = append(errsMessage, errs)
		}
		response["ErrorField"] = errsMessage
		return err
	}

	return nil
}

func ValidateStructChangePassword(data *dto.ChangePassword, response map[string]interface{}) error {
	trans := validation.TranslatorIDN()
	validate := validator.New()
	validation.RegisterCustomValidationsProduct(validate, trans)

	if err := validate.Struct(data); err != nil {
		var errsMessage []map[string]string
		for _, err := range err.(validator.ValidationErrors) {
			errField := err.StructField()
			errTranslate := err.Translate(trans)
			errs := map[string]string{
				errField: errTranslate,
			}
			errsMessage = append(errsMessage, errs)
		}
		response["ErrorField"] = errsMessage
		return err
	}

	return nil
}

func ValidateFileExtention(fileHeader *multipart.FileHeader) (string, error) {
	ext := filepath.Ext(fileHeader.Filename) // extention file image
	if ext != ".jpg" && ext != ".png" && ext != ".gift" {
		return "", errors.New("Invalid file extention, .jpg, .png, and .gift only are allowed")
	}
	return ext, nil
}

func ValidateFileSize(fileHeader *multipart.FileHeader) error {
	if fileHeader.Size > 2000000 {
		return errors.New("file to large, max size file 2mb")
	}
	return nil
}

func ValidateLogin(email string, password string, response *map[string]string) error {
	if email == "" {
		(*response)["email"] = "email tidak boleh kosong"
	} else if password == "" {
		(*response)["password"] = "password tidak boleh kosong"
	}

	if len(*response) > 0 {
		return errors.New("Validation Login User Error")
	}

	return nil
}
