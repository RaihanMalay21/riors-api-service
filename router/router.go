package router

import (
	controllerAuth "github.com/RaihanMalay21/api-service-riors/controller/authentication"
	controllerProducts "github.com/RaihanMalay21/api-service-riors/controller/products"
	_ "github.com/RaihanMalay21/api-service-riors/docs"
	"github.com/RaihanMalay21/api-service-riors/middlewares"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitRouter(
	e *echo.Echo,
	product *controllerProducts.ProductController,
	category *controllerProducts.CategoryController,
	auth *controllerAuth.AuthenticationController,
) {

	e.Use(middlewares.CorsMiddlewares)
	e.Use(middlewares.SetLimiterMiddleware)

	// general access without token access
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/category", category.GetAllCategory)
	e.GET("/product", product.GetAllProduct)
	e.GET("/product/female", product.GetAllFemale)
	e.GET("/product/male", product.GetAllMale)

	// authentication
	protectedAuthentication := e.Group("auth")
	protectedAuthentication.POST("/login/user", auth.LoginUser)
	protectedAuthentication.POST("/login/admin", auth.LoginAdmin)
	protectedAuthentication.POST("/signup/user", auth.SignupUser)
	protectedAuthentication.POST("/signup/user/verification", auth.SignupUserVerification) // access with token register_riors_token
	protectedAuthentication.GET("/google", auth.HandleGoogleLogin)
	protectedAuthentication.GET("/google/callback", auth.HandleGoogleCallback)
	protectedAuthentication.POST("/forgot/password/user", auth.ForgotPasswordUser)
	protectedAuthentication.POST("/reset/password", auth.ResetPasswordUser)

	// user access with token user_riors_token
	protectedUser := e.Group("/user")
	protectedUser.Use(middlewares.ArmorUser)
	protectedUser.PATCH("/change/password", auth.ChangePasswordUser)

	// admin access with token admin_riors_token
	protectedAdmin := e.Group("/admin")
	protectedAdmin.Use(middlewares.ArmorAdmin)
	protectedAdmin.PATCH("/change/password", auth.ChangePasswordAdmin)
	protectedAdmin.POST("/category", category.InputCategory)
	protectedAdmin.POST("/product", product.InputProduct)

	// owner access with token owner_riors_token based on position employee
	protectedOwner := protectedAdmin.Group("/owner")
	protectedOwner.Use(middlewares.ArmorOwner)
	protectedOwner.POST("/register/employe", auth.SignupEmploye)

}
