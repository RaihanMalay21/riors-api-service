package authentication

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/RaihanMalay21/api-service-riors/domain"
	"github.com/RaihanMalay21/api-service-riors/dto"
	"github.com/RaihanMalay21/api-service-riors/middlewares"
	repository "github.com/RaihanMalay21/api-service-riors/repository/authentication"
	"github.com/RaihanMalay21/api-service-riors/service"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthenticationService interface {
	LoginUser(email string, password string, response map[string]interface{}) (*http.Cookie, int)
	SignupUser(register *dto.RegisterUser, response map[string]interface{}) (*http.Cookie, int)
	SignupUserVerification(email *string, code *string, response map[string]interface{}) int
}

type authenticationService struct {
	repo repository.AuthenticationRepository
}

func ConstructorAuthenticationService(repo repository.AuthenticationRepository) AuthenticationService {
	return &authenticationService{repo: repo}
}

func (ar *authenticationService) LoginUser(email string, password string, response map[string]interface{}) (*http.Cookie, int) {
	var res = make(map[string]string)
	if err := service.ValidateLoginUser(email, password, &res); err != nil {
		response["ErrorFields"] = res
		return nil, http.StatusBadRequest
	}

	user, err := ar.repo.GetUserByEmail(email)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			res["email"] = fmt.Sprintf("Email %s tidak ditemukan", email)
			response["ErrorFields"] = res
			return nil, http.StatusBadRequest
		default:
			response["error"] = err.Error()
			return nil, http.StatusInternalServerError
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			res["password"] = "Password anda salah"
			response["ErrorFields"] = res
			return nil, http.StatusBadRequest
		default:
			response["error"] = err.Error()
			return nil, http.StatusInternalServerError
		}
	}

	expToken := time.Now().Add(7 * 24 * time.Hour)
	expCookie := 7 * 24 * 60 * 60
	cookie, err := middlewares.CreateJWT(user.UserName, user.Id, expToken, "user_riors_token", expCookie, response)
	if err != nil {
		return nil, http.StatusInternalServerError
	}

	response["success"] = "Login Successfully"
	return cookie, http.StatusOK
}

func (ar *authenticationService) SignupUser(register *dto.RegisterUser, response map[string]interface{}) (*http.Cookie, int) {
	if err := service.ValidateStructRegister(register, response); err != nil {
		return nil, http.StatusBadRequest
	}

	register.Code = service.GenerateRandomNumber()
	expr := 5 * time.Minute

	if err := ar.repo.PushRedisRegister(register, expr); err != nil {
		response["error"] = err.Error()
		return nil, http.StatusInternalServerError
	}

	if err := service.EmailVerificationCode(&register.Email, &register.Code); err != nil {
		if err := ar.repo.DeleteRedisRegister(register.Email); err != nil {
			response["error"] = err.Error()
			return nil, http.StatusInternalServerError
		}
		response["error"] = err.Error()
		return nil, http.StatusInternalServerError
	}

	expToken := time.Now().Add(5 * time.Minute)
	expCookie := 5 * 60
	cookie, err := middlewares.CreateJWT(register.Email, 0, expToken, "register_riors_token", expCookie, response)
	if err != nil {
		return nil, http.StatusInternalServerError
	}

	response["success"] = "Signup berhasil, silahkan verifikasi account anda"
	return cookie, http.StatusOK
}

func (ar *authenticationService) SignupUserVerification(email *string, code *string, response map[string]interface{}) int {
	value, err := ar.repo.GetRediRegistrationByEmail(*email)
	if err != nil {
		if err == redis.Nil {
			response["message"] = fmt.Sprintf("Email %s not Found", *email)
			return http.StatusBadRequest
		}
		response["error"] = err.Error()
		return http.StatusInternalServerError
	}

	varificationCode, err := strconv.Atoi(*code)
	if err != nil {
		response["error"] = err.Error()
		return http.StatusInternalServerError
	}

	if value.Try > 3 {
		response["message"] = "terlalu banyak percobaan, silahkan signup kembali"
		return http.StatusTooManyRequests
	}

	if varificationCode != value.Code {
		response["ErrorField"] = "kode verifikasi anda salah"
		value.Try = value.Try + 1
		expr := 2 * time.Minute
		if err := ar.repo.PushRedisRegister(value, expr); err != nil {
			response["error"] = err.Error()
			fmt.Println("redis: ", err)
			return http.StatusInternalServerError
		}
		return http.StatusBadRequest
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte((value.Password)), bcrypt.DefaultCost)
	if err != nil {
		response["error"] = err.Error()
		fmt.Println(err)
		return http.StatusInternalServerError
	}

	user := &domain.User{
		Email:    value.Email,
		Password: string(hashPassword),
	}

	if err := ar.repo.CreateRegistration(user); err != nil {
		response["error"] = err.Error()
		fmt.Println("postgresql: ", err)
		return http.StatusInternalServerError
	}

	response["success"] = "verifikasi berhasil, silahkan login"
	return http.StatusOK
}
