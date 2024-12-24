package router

import (
	"html/template"

	controllerAuth "github.com/RaihanMalay21/api-service-riors/controller/authentication"
	controllerProducts "github.com/RaihanMalay21/api-service-riors/controller/products"
	_ "github.com/RaihanMalay21/api-service-riors/docs"
	"github.com/RaihanMalay21/api-service-riors/middlewares"
	templateModule "github.com/RaihanMalay21/api-service-riors/template"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitRouter(e *echo.Echo, product controllerProducts.ProductController, category controllerProducts.CategoryController, auth controllerAuth.AuthenticationController) {
	renderer := templateModule.ConstructorTemplate(template.Must(template.ParseGlob("././template/*.html")))
	e.Renderer = renderer

	e.Use(middlewares.CorsMiddlewares)
	e.Use(middlewares.SetLimiterMiddleware)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/login/user", auth.LoginUser)
	e.POST("/signup/user", auth.SignupUser)
	e.POST("/signup/verification", auth.SignupUserVerification)
	e.GET("/category", category.GetAllCategory)
	e.GET("/product", product.GetAllProduct)

	protectedUser := e.Group("/user")
	protectedUser.Use(middlewares.ArmorUser)
	protectedUser.GET("/product/male", product.GetAllMale)
	protectedUser.GET("/product/female", product.GetAllFemale)

	protectedAdmin := e.Group("/admin")
	protectedAdmin.Use(middlewares.ArmorAdmin)
	protectedAdmin.POST("/category", category.InputCategory)
	protectedAdmin.POST("/product", product.InputProduct)

}
