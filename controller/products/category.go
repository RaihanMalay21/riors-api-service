package products

import (
	_ "github.com/RaihanMalay21/api-service-riors/docs"
	"github.com/RaihanMalay21/api-service-riors/dto"
	service "github.com/RaihanMalay21/api-service-riors/service/products"

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
// @Tags category
// @Produce  application/json
// @Success 200 {object} []dto.Category "Successfully retrieved list of categories with details"
// @Failure 404 {object} ResponseErrorNotFound "No categories found"
// @Failure 500 {object} ResponseErrorInternalServer "Internal server error while processing the request"
// @Router /category [get]
func (cs *categoryController) GetAllCategory(e echo.Context) error {
	data, res, statusCode := cs.service.GetAllCategory()
	if statusCode != 200 {
		return e.JSON(statusCode, res)
	}

	return e.JSON(statusCode, data)
}

// @summary Add data Category
// @Description Add a new category to the system
// @Tags admin
// @Security BearerAuth
// @accept  application/x-www-form-urlencoded
// @produce  application/json
// @Param category body CategoryInput true "Category Input"
// @Success 200 {object} ResponseSuccess "Category successfully added to the system"
// @Failure 400 {object} ResponseErrorBadRequest "Invalid request or incomplete category data"
// @Failure 500 {object} ResponseErrorInternalServer "Internal server error while processing the request"
// @Router /admin/category [post]
func (cs *categoryController) InputCategory(e echo.Context) error {
	var response = make(map[string]interface{})

	data := dto.Category{
		CategoryName: e.QueryParam("category"),
	}

	statusCode := cs.service.InputCategory(&data, response)

	return e.JSON(statusCode, response)
}
