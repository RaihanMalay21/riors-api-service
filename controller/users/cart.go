package users

import (
	"net/http"
	"strconv"

	"github.com/RaihanMalay21/api-service-riors/config"
	"github.com/RaihanMalay21/api-service-riors/controller/helper"
	"github.com/RaihanMalay21/api-service-riors/domain"
	"github.com/RaihanMalay21/api-service-riors/dto"
	"github.com/RaihanMalay21/api-service-riors/service/users"
	"github.com/labstack/echo/v4"
)

type CartController struct {
	service *users.CartService
	helper  *helper.HelperController
}

func ConstructorCartController(service *users.CartService, helper *helper.HelperController) *CartController {
	return &CartController{
		service: service,
		helper:  helper,
	}
}

func (cs *CartController) GetUserIdFromClaims(e echo.Context, response map[string]interface{}) uint {
	claims, ok := e.Get("user_claims").(*config.JWTClaim)
	if !ok {
		response["error"] = "Id user not found"
		return 0
	}

	return claims.Id
}

func (cs *CartController) GetAllProduct(e echo.Context) error {
	response := make(map[string]interface{})

	userId := cs.GetUserIdFromClaims(e, response)
	if userId == 0 {
		return e.JSON(http.StatusInternalServerError, response)
	}

	charts, statusCode := cs.service.GetAllProduct(userId, response)
	if statusCode != 200 {
		return e.JSON(statusCode, response)
	}

	return e.JSON(statusCode, charts)
}

func (cs *CartController) AddProductToCart(e echo.Context) error {
	var response = make(map[string]interface{})

	userId := cs.GetUserIdFromClaims(e, response)
	if userId == 0 {
		return e.JSON(http.StatusInternalServerError, response)
	}

	cart := new(dto.Cart)
	if err := e.Bind(&cart); err != nil {
		response["error"] = err.Error()
		return e.JSON(http.StatusInternalServerError, response)
	}
	cart.UserId = userId

	statusCode := cs.service.AddProduct(cart, response)

	return e.JSON(statusCode, response)
}

func (cs *CartController) UpdateAmountProduct(e echo.Context) error {
	response := make(map[string]interface{})

	userId := cs.GetUserIdFromClaims(e, response)
	if userId == 0 {
		return e.JSON(http.StatusInternalServerError, response)
	}

	chartId, err := strconv.Atoi(e.FormValue("cartId"))
	if err != nil {
		response["error"] = err.Error()
		return e.JSON(http.StatusInternalServerError, response)
	}

	amountPrice, err := strconv.Atoi(e.FormValue("amountPrice"))
	if err != nil {
		response["error"] = err.Error()
		return e.JSON(http.StatusInternalServerError, response)
	}

	amountItems, err := strconv.Atoi(e.FormValue("amountItem"))
	if err != nil {
		response["error"] = err.Error()
		return e.JSON(http.StatusInternalServerError, response)
	}

	cart := domain.Cart{
		Id:          uint(chartId),
		UserId:      userId,
		AmountPrice: float64(amountPrice),
		AmountItem:  float64(amountItems),
	}

	statusCode := cs.service.UpdateAmountProduct(&cart, response)

	return e.JSON(statusCode, response)
}

func (cs *CartController) DeleteProduct(e echo.Context) error {
	response := make(map[string]interface{})

	userId := cs.GetUserIdFromClaims(e, response)
	if userId == 0 {
		return e.JSON(http.StatusInternalServerError, response)
	}

	cartId, err := strconv.Atoi(e.FormValue("cartId"))
	if err != nil {
		response["error"] = err.Error()
		return e.JSON(http.StatusInternalServerError, response)
	}

	productVariantId, err := strconv.Atoi(e.FormValue("productVariantId"))
	if err != nil {
		response["error"] = err.Error()
		return e.JSON(http.StatusInternalServerError, response)
	}

	cart := domain.Cart{
		Id:               uint(cartId),
		ProductVariantId: uint(productVariantId),
	}

	statusCode := cs.service.DeleteProduct(&cart, response)

	return e.JSON(statusCode, response)
}
