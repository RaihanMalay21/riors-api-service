package boostrap

import (
	"github.com/RaihanMalay21/api-service-riors/config"
	controllerAuth "github.com/RaihanMalay21/api-service-riors/controller/authentication"
	controllerHelper "github.com/RaihanMalay21/api-service-riors/controller/helper"
	controllerProducts "github.com/RaihanMalay21/api-service-riors/controller/products"
	repositoryAuth "github.com/RaihanMalay21/api-service-riors/repository/authentication"
	repositoryProducts "github.com/RaihanMalay21/api-service-riors/repository/products"
	"github.com/RaihanMalay21/api-service-riors/router"
	serviceAuth "github.com/RaihanMalay21/api-service-riors/service/authentication"
	serviceHelper "github.com/RaihanMalay21/api-service-riors/service/helper"
	serviceProducts "github.com/RaihanMalay21/api-service-riors/service/products"
	serviceValidate "github.com/RaihanMalay21/api-service-riors/service/validate"
	"github.com/labstack/echo/v4"
)

func SetupDependencies() *echo.Echo {
	e := echo.New()
	db := config.DB
	client := config.Conn

	helperController := controllerHelper.NewHelperController()

	helperService := serviceHelper.NewHelperService()
	validateService := serviceValidate.NewValidateService()

	repoCategory := repositoryProducts.ConstructorCategoryRepository(db)
	repoProduct := repositoryProducts.ConstructorProductRepository(db)
	repoAuth := repositoryAuth.ConstructorAuthenticationRepository(db, client)

	serviceCategory := serviceProducts.ConstructorCategoryService(repoCategory, helperService, validateService)
	serviceProduct := serviceProducts.ConstructorProductService(repoProduct, helperService, validateService)
	serviceAuth := serviceAuth.ConstructorAuthenticationService(repoAuth, helperService, validateService)

	controllerCategory := controllerProducts.ConstructorCategoryController(serviceCategory, helperController)
	controllerProduct := controllerProducts.ConstructorProductController(serviceProduct, helperController)
	controllerAuth := controllerAuth.ConstructorAuthenticationController(serviceAuth, helperController)

	router.InitRouter(e, controllerProduct, controllerCategory, controllerAuth)

	return e
}
