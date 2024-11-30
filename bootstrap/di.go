package boostrap

import (
	"github.com/RaihanMalay21/api-service-riors/config"
	"github.com/RaihanMalay21/api-service-riors/controller"
	"github.com/RaihanMalay21/api-service-riors/repository"
	"github.com/RaihanMalay21/api-service-riors/router"
	"github.com/RaihanMalay21/api-service-riors/service"
	"github.com/labstack/echo/v4"
)

func SetupDependencies() *echo.Echo{
	e := echo.New()
	db := config.DB

	repoProduct := repository.ConstructorProductRepository(db)
	repoCategory := repository.ConstructorCategoryRepository(db)

	serviceProduct := service.ConstructorProductService(repoProduct)
	serviceCategory := service.ConstructorCategoryService(repoCategory)

	controllerProduct := controller.ConstructorProductController(serviceProduct)
	controllerCategory := controller.ConstructorCategoryController(serviceCategory)

	router.InitRouter(e, controllerProduct, controllerCategory)

	return e
}