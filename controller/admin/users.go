package admin

import (
	"github.com/RaihanMalay21/api-service-riors/controller/helper"
	"github.com/RaihanMalay21/api-service-riors/service/admin"
)

type AdminUserController struct {
	service *admin.AdminUserService
	helper  *helper.HelperController
}

func ConstructorAdminUserController(service *admin.AdminUserService, helper *helper.HelperController) *AdminUserController {
	return &AdminUserController{
		service: service,
		helper:  helper,
	}
}

// func (auc *AdminUserController) GetUserActive(e echo.Context) error {
// 	response := make(map[string]interface{})

// 	latesUserIdRedis := e.QueryParam("lastUserId")

// 	latestId, users, statusCode := auc.service.GetUserActive(latesUserIdRedis, response)
// 	if statusCode != 200 {
// 		return e.JSON(statusCode, response)
// 	}

// 	mapUsers := map[string]interface{}{
// 		"latestId": latestId,
// 		"users":    users,
// 	}

// 	return e.JSON(statusCode, mapUsers)
// }
