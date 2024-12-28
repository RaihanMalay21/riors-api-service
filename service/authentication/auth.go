package authentication

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/RaihanMalay21/api-service-riors/domain"
	"github.com/RaihanMalay21/api-service-riors/dto"
	"github.com/RaihanMalay21/api-service-riors/mapper"
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
	LoginAdmin(email string, password string, response map[string]interface{}) (*http.Cookie, *http.Cookie, int)
	SignupEmploye(file multipart.File, fileHeader *multipart.FileHeader, data *dto.Employee, response map[string]interface{}) int
	ChangePasswordAdmin(data *dto.ChangePassword, response map[string]interface{}) int
	ChangePasswordUser(data *dto.ChangePassword, response map[string]interface{}) int
}

type authenticationService struct {
	repo repository.AuthenticationRepository
}

func ConstructorAuthenticationService(repo repository.AuthenticationRepository) AuthenticationService {
	return &authenticationService{repo: repo}
}

func (ar *authenticationService) LoginUser(email string, password string, response map[string]interface{}) (*http.Cookie, int) {
	var res = make(map[string]string)
	if err := service.ValidateLogin(email, password, &res); err != nil {
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
			return http.StatusInternalServerError
		}
		return http.StatusBadRequest
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(value.Password), bcrypt.DefaultCost)
	if err != nil {
		response["error"] = err.Error()
		return http.StatusInternalServerError
	}

	user := &domain.User{
		Email:    value.Email,
		Password: string(hashPassword),
	}

	if err := ar.repo.CreateUser(user); err != nil {
		response["error"] = err.Error()
		return http.StatusInternalServerError
	}

	response["success"] = "verifikasi berhasil, silahkan login"
	return http.StatusOK
}

func (ar *authenticationService) LoginAdmin(email string, password string, response map[string]interface{}) (*http.Cookie, *http.Cookie, int) {
	var res = make(map[string]string)
	if err := service.ValidateLogin(email, password, &res); err != nil {
		response["ErrorFields"] = res
		return nil, nil, http.StatusBadRequest
	}

	employee, err := ar.repo.GetEmployeeByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res["email"] = fmt.Sprintf("Email %s not found", email)
			response["ErrorField"] = res
			return nil, nil, http.StatusBadRequest
		}
		response["error"] = err.Error()
		return nil, nil, http.StatusInternalServerError
	}

	if err := bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(password)); err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			res["password"] = "Password anda salah"
			response["ErrorFields"] = res
			return nil, nil, http.StatusBadRequest
		default:
			response["error"] = err.Error()
			return nil, nil, http.StatusInternalServerError
		}
	}

	expToken := time.Now().Add(2 * time.Hour)
	expCookie := 2 * 60 * 60
	var cookieAdmin, cookieOwner *http.Cookie
	switch employee.Position {
	case "Staff":
		cookieAdmin, err = middlewares.CreateJWT(employee.Email, employee.Id, expToken, "admin_riors_token", expCookie, response)
		if err != nil {
			return nil, nil, http.StatusInternalServerError
		}
		response["success"] = "Anda berhasil login"
		return cookieAdmin, nil, http.StatusOK
	case "Owner":
		var err error
		cookieAdmin, err = middlewares.CreateJWT(employee.Email, employee.Id, expToken, "admin_riors_token", expCookie, response)
		if err != nil {
			return nil, nil, http.StatusInternalServerError
		}
		cookieOwner, err = middlewares.CreateJWT(employee.Email, employee.Id, expToken, "owner_riors_token", expCookie, response)
		if err != nil {
			return nil, nil, http.StatusInternalServerError
		}
		response["success"] = "Anda berhasil login"
		return cookieAdmin, cookieOwner, http.StatusOK
	}

	response["error"] = "Internal server error"
	return nil, nil, http.StatusInternalServerError
}

func (ar *authenticationService) SignupEmploye(file multipart.File, fileHeader *multipart.FileHeader, data *dto.Employee, response map[string]interface{}) int {
	if err := service.ValidateStructEmployee(data, response); err != nil {
		return http.StatusBadRequest
	}

	dataDomain := mapper.EmployeeDTOToEmployeeDomain(data)

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(dataDomain.Password), bcrypt.DefaultCost)
	if err != nil {
		response["error"] = err.Error()
		return http.StatusInternalServerError
	}
	dataDomain.Password = string(hashPassword)

	tx := ar.repo.NewTransactionAuth()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
		if tx != nil {
			tx.Rollback()
		}
	}()

	if err := ar.repo.CreateEmployee(tx, &dataDomain); err != nil {
		tx.Rollback()
		response["error"] = err.Error()
		return http.StatusInternalServerError
	}

	if err := service.UploadToS3Admin(&dataDomain, file, fileHeader, data.Ext, data.ImageType); err != nil {
		tx.Rollback()
		response["error"] = err.Error()
		return http.StatusInternalServerError
	}

	if err := ar.repo.UpdateEmployeImage(tx, &dataDomain); err != nil {
		tx.Rollback()
		response["error"] = err.Error()
		return http.StatusInternalServerError
	}

	tx.Commit()

	response["success"] = "Berhasil registrasi employee"
	return http.StatusOK
}

func (ar *authenticationService) ChangePasswordAdmin(data *dto.ChangePassword, response map[string]interface{}) int {
	if err := service.ValidateStructChangePassword(data, response); err != nil {
		return http.StatusBadRequest
	}

	employee, err := ar.repo.GetPasswordById(data.Id, "employee")
	if err != nil {
		response["error"] = err.Error()
		return http.StatusInternalServerError
	}

	Employee, ok := employee.(*domain.Employee)
	if !ok {
		response["error"] = "Invalid type assertion for Employee"
		return http.StatusInternalServerError
	}

	if err := bcrypt.CompareHashAndPassword([]byte(Employee.Password), []byte(data.PasswordBefore)); err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			response["ErrorFields"] = map[string]string{"password": "Password anda salah"}
			return http.StatusBadRequest
		default:
			response["error"] = err.Error()
			return http.StatusInternalServerError
		}
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		response["error"] = err.Error()
		return http.StatusInternalServerError
	}

	Employees := domain.Employee{
		Id:       data.Id,
		Password: string(hashPassword),
	}

	if err := ar.repo.UpdatePasswordById(Employees.Password, Employees.Id, "employee"); err != nil {
		response["error"] = err.Error()
		return http.StatusInternalServerError
	}

	response["success"] = "succesfully change password"
	return http.StatusOK
}

func (ar *authenticationService) ChangePasswordUser(data *dto.ChangePassword, response map[string]interface{}) int {
	if err := service.ValidateStructChangePassword(data, response); err != nil {
		return http.StatusBadRequest
	}

	user, err := ar.repo.GetPasswordById(data.Id, "user")
	if err != nil {
		response["error"] = err.Error()
		return http.StatusInternalServerError
	}

	User, ok := user.(*domain.User)
	if !ok {
		response["error"] = "Invalid type assertion for user"
	}

	if err := bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(data.PasswordBefore)); err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			response["ErrorFields"] = map[string]string{"password": "Password anda salah"}
			return http.StatusBadRequest
		default:
			response["error"] = err.Error()
			return http.StatusInternalServerError
		}
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		response["error"] = err.Error()
		return http.StatusInternalServerError
	}

	Users := domain.User{
		Id:       data.Id,
		Password: string(hashPassword),
	}

	if err := ar.repo.UpdatePasswordById(Users.Password, Users.Id, "user"); err != nil {
		response["error"] = err.Error()
		return http.StatusInternalServerError
	}

	response["success"] = "successfully Change Password"
	return http.StatusOK
}
