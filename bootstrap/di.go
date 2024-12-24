package boostrap

import (
	"github.com/RaihanMalay21/api-service-riors/config"
	controllerProducts "github.com/RaihanMalay21/api-service-riors/controller/products"
	repositoryProducts "github.com/RaihanMalay21/api-service-riors/repository/products"
	serviceProducts "github.com/RaihanMalay21/api-service-riors/service/products"
	controllerAuth "github.com/RaihanMalay21/api-service-riors/controller/authentication"
	repositoryAuth "github.com/RaihanMalay21/api-service-riors/repository/authentication"
	serviceAuth "github.com/RaihanMalay21/api-service-riors/service/authentication"
	"github.com/RaihanMalay21/api-service-riors/router"
	"github.com/labstack/echo/v4"
)

func SetupDependencies() *echo.Echo{
	e := echo.New()
	db := config.DB
	client := config.Conn

	repoProduct := repositoryProducts.ConstructorProductRepository(db)
	serviceProduct := serviceProducts.ConstructorProductService(repoProduct)
	controllerProduct := controllerProducts.ConstructorProductController(serviceProduct)
	
	repoCategory := repositoryProducts.ConstructorCategoryRepository(db)
	serviceCategory := serviceProducts.ConstructorCategoryService(repoCategory)
	controllerCategory := controllerProducts.ConstructorCategoryController(serviceCategory)

	repoAuth := repositoryAuth.ConstructorAuthenticationRepository(db, client)
	serviceAuth := serviceAuth.ConstructorAuthenticationService(repoAuth)
	controllerAuth := controllerAuth.ConstructorAuthenticationController(serviceAuth)


	router.InitRouter(e, controllerProduct, controllerCategory, controllerAuth)

	return e
}