package controller

import (
	"strconv"

	"github.com/RaihanMalay21/api-service-riors/dto"
	"github.com/RaihanMalay21/api-service-riors/service"
	"github.com/labstack/echo/v4"
)

type ProductController interface {
	GetAllProduct(e echo.Context) error
	InputProduct(e echo.Context) error
	GetAllMale(e echo.Context) error
	GetAllFemale(e echo.Context) error
}

type productController struct {
	service service.ProductService
}

func ConstructorProductController(service service.ProductService) ProductController {
	return &productController{service: service}
}

// @summary Get All Data Product
// @Description Get detailed information of all data Product and detailProduct by Id Product
// @Tags Product
// @Produce application/json
// @Success 200  {object}  []ResponseProduct
// @Failure 404  {object}  ResponseErrorNotFound
// @Failure 500  {object}  ResponseErrorInternalServer
// @Router /product [get]
func (ps *productController) GetAllProduct(e echo.Context) error {
	data, res, statusCode := ps.service.GetAllProduct()
	if statusCode != 200 {
		return e.JSON(statusCode, res)
	}

	return e.JSON(statusCode, data)
}

// @summary post data Product
// @Description Add a new Product to the system
// @Tags Product
// @accept  multipart/form-data
// @produce  application/json
// @Param category body ProductInput true "Product Input"
// @Success 200  {object}  ResponseSuccess
// @Failure 400  {object}  ResponseErrorBadRequest
// @Failure 500  {object}  ResponseErrorInternalServer
// @Router /product [post]
func (ps *productController) InputProduct(e echo.Context) error {
	var response = make(map[string]interface{})

	file, fileHeader, Ext, filetype, statusCode := GetFileFromForm(e, response)
	if statusCode != 200 {
		return e.JSON(statusCode, response)
	}

	categoryId, _ := strconv.ParseUint(e.FormValue("categoryId"), 10, 32)
	hargaBarang, _ := strconv.ParseFloat(e.FormValue("hargaProduct"), 64)

	data := dto.Product{
		CategoryId:     uint(categoryId),
		ProductName:    e.FormValue("namaProduct"),
		HargaBarang:    hargaBarang,
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
// @Tags Product
// @Produce application/json
// @Success 200  {object}  []ResponseProduct
// @Failure 404  {object}  ResponseErrorNotFound
// @Failure 500  {object}  ResponseErrorInternalServer
// @Router /product/male [get]
func (ps *productController) GetAllMale(e echo.Context) error {
	data, res, statusCode := ps.service.GetAllProductMale()
	if statusCode != 200 {
		return e.JSON(statusCode, res)
	}

	return e.JSON(statusCode, data)
}

// @summary Get All Data Product Female
// @Description Get detailed information of all data Product Female
// @Tags Product
// @Produce application/json
// @Success 200  {object}  []ResponseProduct
// @Failure 404  {object}  ResponseErrorNotFound
// @Failure 500  {object}  ResponseErrorInternalServer
// @Router /product/female [get]
func (ps *productController) GetAllFemale(e echo.Context) error {
	data, res, statusCode := ps.service.GetAllProductFemale()
	if statusCode != 200 {
		return e.JSON(statusCode, res)
	}

	return e.JSON(statusCode, data)
}
