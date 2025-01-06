package products

import (
	"net/http"

	"github.com/RaihanMalay21/api-service-riors/dto"
	"github.com/RaihanMalay21/api-service-riors/mapper"
	repository "github.com/RaihanMalay21/api-service-riors/repository/products"
	"github.com/RaihanMalay21/api-service-riors/service/helper"
	"github.com/RaihanMalay21/api-service-riors/service/validate"
	"github.com/RaihanMalay21/api-service-riors/validation"
	"github.com/go-playground/validator/v10"
)

type CategoryService struct {
	repo     *repository.CategoryRepository
	helper   *helper.HelperService
	validate *validate.ValidateService
}

func ConstructorCategoryService(repo *repository.CategoryRepository, helper *helper.HelperService, validate *validate.ValidateService) *CategoryService {
	return &CategoryService{
		repo:     repo,
		helper:   helper,
		validate: validate,
	}
}

func (rc *CategoryService) GetAllCategory() (*[]dto.Category, map[string]string, int) {
	data, err := rc.repo.GetAll()
	if err != nil {
		response := map[string]string{"error": "Internal Server encountered an Error"}
		return nil, response, http.StatusInternalServerError
	}

	dataDto := mapper.GetAllCategoryDomainTODTO(data)

	return dataDto, nil, http.StatusOK
}

func (rc *CategoryService) InputCategory(data *dto.Category, response map[string]interface{}) int {
	trans := validation.TranslatorIDN()
	validate := validator.New()

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
		return http.StatusBadRequest
	}

	dataDomain := mapper.CategoryDTOTODomain(data)

	if err := rc.repo.Create(&dataDomain); err != nil {
		response["error"] = err.Error()
		return http.StatusInternalServerError
	}

	response["success"] = "Berhasil Input Kategori"
	return http.StatusOK
}
