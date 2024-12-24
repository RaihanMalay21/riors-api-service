package products

import (
	"github.com/RaihanMalay21/api-service-riors/domain"
	"github.com/RaihanMalay21/api-service-riors/mapper"
	repository "github.com/RaihanMalay21/api-service-riors/repository/products"
	service "github.com/RaihanMalay21/api-service-riors/service"

	"mime/multipart"
	"net/http"

	"github.com/RaihanMalay21/api-service-riors/dto"
)

type ProductService interface {
	GetAllProduct() (*[]domain.Product, map[string]string, int)
	InputProduct(file multipart.File, fileHeader *multipart.FileHeader, data *dto.Product, response map[string]interface{}) int
	GetAllProductMale() (*[]domain.Product, map[string]string, int)
	GetAllProductFemale() (*[]domain.Product, map[string]string, int)
}

type productService struct {
	repository repository.ProductRepository
}

func ConstructorProductService(repository repository.ProductRepository) ProductService {
	return &productService{repository: repository}
}

func (pr *productService) GetAllProduct() (*[]domain.Product, map[string]string, int) {
	data, err := pr.repository.GetAll()
	if err != nil {
		response := map[string]string{"error": "Internal Server encountered an Error"}
		return nil, response, http.StatusInternalServerError
	}

	return data, nil, http.StatusOK
}

func (pr *productService) InputProduct(file multipart.File, fileHeader *multipart.FileHeader, data *dto.Product, response map[string]interface{}) int {
	if err := service.ValidateStructProduct(data, response); err != nil {
		return http.StatusBadRequest
	}

	dataDomain := mapper.ProductDTOTODomain(data)

	tx := pr.repository.NewTransactionProduct()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
		if tx != nil {
			tx.Rollback()
		}
	}()

	if err := pr.repository.Create(tx, &dataDomain); err != nil {
		tx.Rollback()
		response["error"] = err.Error()
		return http.StatusInternalServerError
	}

	if err := service.UploadToS3(&dataDomain, file, fileHeader, data.Ext, data.ImageType); err != nil {
		tx.Rollback()
		response["error"] = err.Error()
		return http.StatusInternalServerError
	}

	// update data barang
	if err := pr.repository.UpdateProductImage(tx, &dataDomain); err != nil {
		tx.Rollback()
		response["error"] = err.Error()
		return http.StatusInternalServerError
	}

	tx.Commit()

	response["success"] = "Berhasil Memasukkan Product"
	return http.StatusOK
}

func (pr *productService) GetAllProductMale() (*[]domain.Product, map[string]string, int) {
	data, err := pr.repository.GetAllMale()
	if err != nil {
		response := map[string]string{"error": "Internal Server encountered an Error"}
		return nil, response, http.StatusInternalServerError
	}

	return data, nil, http.StatusOK
}

func (pr *productService) GetAllProductFemale() (*[]domain.Product, map[string]string, int) {
	data, err := pr.repository.GetAllFemale()
	if err != nil {
		response := map[string]string{"error": "Internal Server encountered an Error"}
		return nil, response, http.StatusInternalServerError
	}

	return data, nil, http.StatusOK
}
