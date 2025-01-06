package products

import (
	"strconv"

	"github.com/RaihanMalay21/api-service-riors/controller/helper"
	"github.com/RaihanMalay21/api-service-riors/dto"
	"github.com/RaihanMalay21/api-service-riors/service/products"
	"github.com/labstack/echo/v4"
)

// type ProductController interface {
// 	GetAllProduct(e echo.Context) error
// 	InputProduct(e echo.Context) error
// 	GetAllMale(e echo.Context) error
// 	GetAllFemale(e echo.Context) error
// }

type ProductController struct {
	service *products.ProductService
	helper  *helper.HelperController
}

func ConstructorProductController(service *products.ProductService, helper *helper.HelperController) *ProductController {
	return &ProductController{
		service: service,
		helper:  helper,
	}
}

// @summary Retrieve All Product Data
// @Description This endpoint is used to retrieve a list of all products along with detailed information for each product. It also allows you to get the details of a specific product by its ID. The response will include information like product name, description, price, and other relevant details.
// @Tags product
// @Produce application/json
// @Success 200 {object} []ResponseProduct "Successfully retrieved list of products and their details"
// @Failure 404 {object} ResponseErrorNotFound "Products not found"
// @Failure 500 {object} ResponseErrorInternalServer "Internal server error while processing the request"
// @Router /product [get]
func (ps *ProductController) GetAllProduct(e echo.Context) error {
	data, res, statusCode := ps.service.GetAllProduct()
	if statusCode != 200 {
		return e.JSON(statusCode, res)
	}

	return e.JSON(statusCode, data)
}

// @summary Add data Product
// @Description Add a new Product to the system
// @Tags admin
// @Security BearerAuth
// @accept  multipart/form-data
// @produce  application/json
// @Param category body ProductInput true "Product Input"
// @Param categoryGender formData string true "Category Gender (Man, Woman, Unisex)" Enums(Man, Woman, Unisex)
// @Success 200 {object} ResponseSuccess "Product successfully added to the system"
// @Failure 400 {object} ResponseErrorBadRequest "Invalid request or incomplete product data"
// @Failure 500 {object} ResponseErrorInternalServer "Internal server error while processing the request"
// @Router /admin/product [post]
func (ps *ProductController) InputProduct(e echo.Context) error {
	var response = make(map[string]interface{})

	file, fileHeader, Ext, filetype, statusCode := ps.helper.GetFileFromForm(e, response)
	if statusCode != 200 {
		return e.JSON(statusCode, response)
	}

	categoryId, _ := strconv.ParseUint(e.FormValue("categoryId"), 10, 32)
	price, _ := strconv.ParseFloat(e.FormValue("price"), 64)

	data := dto.Product{
		CategoryId:     uint(categoryId),
		ProductName:    e.FormValue("productName"),
		Price:          price,
		DateRelase:     e.FormValue("dateRelase"),
		Desc:           e.FormValue("desc"),
		Type:           e.FormValue("typeProduct"),
		Image:          fileHeader.Filename,
		CategoryGender: e.FormValue("categoryGender"),
		FileSize:       uint(fileHeader.Size),
		Ext:            Ext,
		ImageType:      filetype,
	}

	StatusCode := ps.service.InputProduct(
		file,
		fileHeader,
		&data,
		response,
	)

	return e.JSON(StatusCode, response)
}

// @summary Get All Data Product Male
// @Description Get detailed information of all data Product Male
// @Tags product
// @Produce application/json
// @Success 200 {object} []ResponseProduct "Successfully retrieved list of male products with details"
// @Failure 404 {object} ResponseErrorNotFound "No male products found"
// @Failure 500 {object} ResponseErrorInternalServer "Internal server error while processing the request"
// @Router /product/male [get]
func (ps *ProductController) GetAllMale(e echo.Context) error {
	data, res, statusCode := ps.service.GetAllProductMale()
	if statusCode != 200 {
		return e.JSON(statusCode, res)
	}

	return e.JSON(statusCode, data)
}

// @summary Get All Data Product Female
// @Description Get detailed information of all data Product Female
// @Tags product
// @Produce application/json
// @Success 200 {object} []ResponseProduct "Successfully retrieved list of female products with details"
// @Failure 404 {object} ResponseErrorNotFound "No female products found"
// @Failure 500 {object} ResponseErrorInternalServer "Internal server error while processing the request"
// @Router /product/female [get]
func (ps *ProductController) GetAllFemale(e echo.Context) error {
	data, res, statusCode := ps.service.GetAllProductFemale()
	if statusCode != 200 {
		return e.JSON(statusCode, res)
	}

	return e.JSON(statusCode, data)
}
