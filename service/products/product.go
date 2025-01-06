package products

import (
	"github.com/RaihanMalay21/api-service-riors/domain"
	"github.com/RaihanMalay21/api-service-riors/dto"
	"github.com/RaihanMalay21/api-service-riors/mapper"
	repository "github.com/RaihanMalay21/api-service-riors/repository/products"
	"github.com/RaihanMalay21/api-service-riors/service/helper"
	"github.com/RaihanMalay21/api-service-riors/service/validate"

	"mime/multipart"
	"net/http"
)

type ProductService struct {
	repository *repository.ProductRepository
	helper     *helper.HelperService
	validate   *validate.ValidateService
}

func ConstructorProductService(repository *repository.ProductRepository, helper *helper.HelperService, validate *validate.ValidateService) *ProductService {
	return &ProductService{
		repository: repository,
		helper:     helper,
		validate:   validate,
	}
}

func (pr *ProductService) GetAllProduct() (*[]domain.Product, map[string]string, int) {
	data, err := pr.repository.GetAll()
	if err != nil {
		response := map[string]string{"error": "Internal Server encountered an Error"}
		return nil, response, http.StatusInternalServerError
	}

	return data, nil, http.StatusOK
}

func (pr *ProductService) InputProduct(file multipart.File, fileHeader *multipart.FileHeader, data *dto.Product, response map[string]interface{}) int {
	if err := pr.validate.ValidateStructProduct(data, response); err != nil {
		return http.StatusBadRequest
	}

	dateRelase := pr.helper.ConvertDateStringToTime(data.DateRelase, response)
	if dateRelase.IsZero() {
		return http.StatusInternalServerError
	}

	dataDomain := mapper.ProductDTOTODomain(data)
	dataDomain.DateRelase = dateRelase

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

	if err := pr.helper.UploadToS3(&dataDomain, file, fileHeader, data.Ext, data.ImageType); err != nil {
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

func (pr *ProductService) GetAllProductMale() (*[]domain.Product, map[string]string, int) {
	data, err := pr.repository.GetAllMale()
	if err != nil {
		response := map[string]string{"error": "Internal Server encountered an Error"}
		return nil, response, http.StatusInternalServerError
	}

	return data, nil, http.StatusOK
}

func (pr *ProductService) GetAllProductFemale() (*[]domain.Product, map[string]string, int) {
	data, err := pr.repository.GetAllFemale()
	if err != nil {
		response := map[string]string{"error": "Internal Server encountered an Error"}
		return nil, response, http.StatusInternalServerError
	}

	return data, nil, http.StatusOK
}
