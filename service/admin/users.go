package admin

import (
	//"context"
	//"net/http"

	//"github.com/RaihanMalay21/api-service-riors/domain"
	"github.com/RaihanMalay21/api-service-riors/repository/admin"
	"github.com/RaihanMalay21/api-service-riors/service/helper"
	"github.com/RaihanMalay21/api-service-riors/service/validate"
)

type AdminUserService struct {
	repo     *admin.AdminUsersRepository
	helper   *helper.HelperService
	validate *validate.ValidateService
}

func ConstructorAdminUserController(repo *admin.AdminUsersRepository, helper *helper.HelperService, validate *validate.ValidateService) *AdminUserService {
	return &AdminUserService{
		repo:     repo,
		helper:   helper,
		validate: validate,
	}
}

// func (aus *AdminUserService) GetUserActive(latestUserId string, response map[string]interface{}) (*string, *[]domain.User, int) {
// 	aus.repo.CleanupExpiredMessages()

// 	ctx := context.Background()

// 	latesId, usersId, err := aus.repo.RedisGetDataUserActive(&latestUserId, ctx)
// 	if err != nil {
// 		response["error"] = err.Error()
// 		return nil, nil, http.StatusInternalServerError
// 	}

// 	if latesId == nil && usersId == nil {
// 		response["message"] = "No new user activity found"
// 		return nil, nil, http.StatusNoContent
// 	}

// 	users, err := aus.repo.GetUserById(usersId)
// 	if err != nil {
// 		response["error"] = err.Error()
// 		return nil, nil, http.StatusInternalServerError
// 	}

// 	return latesId, users, http.StatusOK
// }
