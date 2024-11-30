package router

import (
	"github.com/RaihanMalay21/api-service-riors/controller"
	_ "github.com/RaihanMalay21/api-service-riors/docs"
	"github.com/RaihanMalay21/api-service-riors/middlewares"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitRouter(e *echo.Echo, product controller.ProductController, category controller.CategoryController) {

	e.Use(middlewares.CorsMiddlewares)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/product", product.GetAllProduct)
	e.POST("/product/input", product.InputProduct)
	e.GET("/category", category.GetAllCategory)
	e.POST("/category/input", category.InputCategory)

}
