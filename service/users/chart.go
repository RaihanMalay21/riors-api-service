package users

import (
	"errors"
	"net/http"

	"github.com/RaihanMalay21/api-service-riors/domain"
	"github.com/RaihanMalay21/api-service-riors/dto"
	"github.com/RaihanMalay21/api-service-riors/mapper"
	repository "github.com/RaihanMalay21/api-service-riors/repository/users"
	"github.com/RaihanMalay21/api-service-riors/service/helper"
	"github.com/RaihanMalay21/api-service-riors/service/validate"
	"gorm.io/gorm"
)

type CartService struct {
	repo     repository.CartRepository
	helper   *helper.HelperService
	validate *validate.ValidateService
}

func ConstructorCartService(repo repository.CartRepository, helper *helper.HelperService, validate *validate.ValidateService) *CartService {
	return &CartService{
		repo:     repo,
		helper:   helper,
		validate: validate,
	}
}

func (cr *CartService) GetAllProduct(userId uint, response map[string]interface{}) ([]dto.Cart, int) {
	cart, err := cr.repo.GetAllProduct(userId)
	if err != nil {
		response["error"] = err.Error()
		return nil, http.StatusInternalServerError
	}

	return mapper.ArrayCartDomainTODTO(cart), http.StatusOK
}

func (cr *CartService) AddProduct(cart *dto.Cart, response map[string]interface{}) int {
	if err := cr.validate.ValidateStructChart(cart, response); err != nil {
		return http.StatusBadRequest
	}

	id, err := cr.repo.GetProductVariantId(cart.ProductId, cart.Color, cart.Size)
	if err != nil {
		response["error"] = err.Error()
		return http.StatusInternalServerError
	}

	cart.ProductVariantId = *id
	data := mapper.CartDTOTODomain(cart)

	if err := cr.repo.AddProduct(&data); err != nil {
		response["error"] = err.Error()
		return http.StatusInternalServerError
	}

	response["success"] = "successfully add product to cart"
	return http.StatusOK
}

func (cr *CartService) UpdateAmountProduct(cart *domain.Cart, response map[string]interface{}) int {
	if err := cr.repo.UpdateAmountProduct(cart); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response["error"] = "Product not found, check product variant id"
			return http.StatusBadRequest
		}
		response["error"] = err.Error()
		return http.StatusInternalServerError
	}

	response["success"] = "successfully update amount product from cart"
	return http.StatusOK
}

func (cr *CartService) DeleteProduct(cart *domain.Cart, response map[string]interface{}) int {
	if err := cr.repo.DeleteProduct(cart); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response["error"] = "product not found, check cart id"
			return http.StatusBadRequest
		}
		response["error"] = err.Error
		return http.StatusInternalServerError
	}

	response["success"] = "successfully delete product from cart"
	return http.StatusOK
}
