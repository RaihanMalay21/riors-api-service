package service

import (
	"github.com/RaihanMalay21/api-service-riors/mapper"
	"github.com/RaihanMalay21/api-service-riors/repository"

	"mime/multipart"
	"net/http"

	"github.com/RaihanMalay21/api-service-riors/dto"
)

type ProductService interface {
	GetAllProduct() (*[]dto.Product, map[string]string, int)
	InputProduct(file multipart.File, fileHeader *multipart.FileHeader, data *dto.Product, response map[string]interface{}) int
}

type productService struct {
	repository repository.ProductRepository
}

func ConstructorProductService(repository repository.ProductRepository) ProductService {
	return &productService{repository: repository}
}

func (pr *productService) GetAllProduct() (*[]dto.Product, map[string]string, int) {
	data, err := pr.repository.GetAll()
	if err != nil {
		response := map[string]string{"message": "Internal Server encountered an Error"}
		return nil, response, http.StatusInternalServerError
	}

	dataDto := mapper.GetAllProductDomainTODTO(data)

	return dataDto, nil, http.StatusOK
}

func (pr *productService) InputProduct(file multipart.File, fileHeader *multipart.FileHeader, data *dto.Product, response map[string]interface{}) int {
	if err := ValidateStructProduct(data, response); err != nil {
		return http.StatusBadRequest
	}

	dataDomain := mapper.ProductDTOTODomain(data)

	tx := pr.repository.NewTransactionProduct()

	if err := pr.repository.Create(tx, &dataDomain); err != nil {
		tx.Rollback()
		response["message"] = err.Error()
		return http.StatusInternalServerError
	}

	if err := UploadToS3(&dataDomain, file, fileHeader, data.Ext, data.ImageType); err != nil {
		tx.Rollback()
		response["message"] = err.Error()
		return http.StatusInternalServerError
	}

	// update data barang
	if err := pr.repository.UpdateProductImage(tx, &dataDomain); err != nil {
		tx.Rollback()
		response["message"] = err.Error()
		return http.StatusInternalServerError
	}

	tx.Commit()

	response["success"] = "Berhasil Memasukkan Product"
	return http.StatusOK
}
