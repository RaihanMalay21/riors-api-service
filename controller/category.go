package controller

import (
	_ "github.com/RaihanMalay21/api-service-riors/docs"
	"github.com/RaihanMalay21/api-service-riors/dto"
	"github.com/RaihanMalay21/api-service-riors/service"

	"github.com/labstack/echo/v4"
)

type CategoryController interface {
	GetAllCategory(e echo.Context) error
	InputCategory(e echo.Context) error
}

type categoryController struct {
	service service.CategoryService
}

func ConstructorCategoryController(service service.CategoryService) CategoryController {
	return &categoryController{service: service}
}

// @summary Get All Data Category
// @Description Get detailed information of all data category and product based on category
// @Tags Category
// @Produce  application/json
// @Success 200  {object}  []dto.Category
// @Failure 404  {object}  ResponseErrorNotFound
// @Failure 500  {object}  ResponseErrorInternalServer
// @Router /category [get]
func (cs *categoryController) GetAllCategory(e echo.Context) error {
	data, res, statusCode := cs.service.GetAllCategory()
	if statusCode != 200 {
		return e.JSON(statusCode, res)
	}

	return e.JSON(statusCode, data)
}

// @summary post data Category
// @Description Add a new category to the system
// @Tags Category
// @accept  application/x-www-form-urlencoded
// @produce  application/json
// @Param category body CategoryInput true "Category Input"
// @Success 200  {object}  ResponseSuccess
// @Failure 400  {object}  ResponseErrorBadRequest
// @Failure 500  {object}  ResponseErrorInternalServer
// @Router /category/input [post]
func (cs *categoryController) InputCategory(e echo.Context) error {
	var response = make(map[string]interface{})

	data := dto.Category{
		CategoryName: e.QueryParam("category"),
	}

	statusCode := cs.service.InputCategory(&data, response)

	return e.JSON(statusCode, response)
}
