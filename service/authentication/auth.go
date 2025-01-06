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
	"github.com/RaihanMalay21/api-service-riors/service/helper"
	"github.com/RaihanMalay21/api-service-riors/service/validate"
	"github.com/RaihanMalay21/api-service-riors/validation"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthenticationService struct {
	repo     *repository.AuthenticationRepository
	helper   *helper.HelperService
	validate *validate.ValidateService
}

func ConstructorAuthenticationService(repo *repository.AuthenticationRepository, helper *helper.HelperService, validate *validate.ValidateService) *AuthenticationService {
	return &AuthenticationService{
		repo:     repo,
		helper:   helper,
		validate: validate,
	}
}

func (ar *AuthenticationService) LoginUser(email string, password string, response map[string]interface{}) (*http.Cookie, int) {
	var res = make(map[string]string)
	if err := ar.validate.ValidateLogin(email, password, &res); err != nil {
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
	cookie, err := middlewares.CreateJWT(user.Name, user.Id, expToken, "user_riors_token", expCookie, response)
	if err != nil {
		return nil, http.StatusInternalServerError
	}

	response["success"] = "Login Successfully"
	return cookie, http.StatusOK
}

func (ar *AuthenticationService) SignupUser(register *dto.RegisterUser, response map[string]interface{}) (*http.Cookie, int) {
	if err := ar.validate.ValidateStructRegister(register, response); err != nil {
		return nil, http.StatusBadRequest
	}

	register.Code = ar.helper.GenerateRandomNumber()
	expr := 5 * time.Minute

	if err := ar.repo.PushRedisRegister(register, expr); err != nil {
		response["error"] = err.Error()
		return nil, http.StatusInternalServerError
	}

	if err := ar.helper.SendEmailVerificationCode(&register.Email, &register.Code); err != nil {
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

func (ar *AuthenticationService) SignupUserVerification(email *string, code *string, response map[string]interface{}) int {
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

func (ar *AuthenticationService) LoginAdmin(email string, password string, response map[string]interface{}) (*http.Cookie, *http.Cookie, int) {
	var res = make(map[string]string)
	if err := ar.validate.ValidateLogin(email, password, &res); err != nil {
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

func (ar *AuthenticationService) SignupEmploye(file multipart.File, fileHeader *multipart.FileHeader, data *dto.Employee, response map[string]interface{}) int {
	if err := ar.validate.ValidateStructEmployee(data, response); err != nil {
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

	if err := ar.helper.UploadToS3Admin(&dataDomain, file, fileHeader, data.Ext, data.ImageType); err != nil {
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

func (ar *AuthenticationService) ChangePasswordAdmin(data *dto.ChangePassword, response map[string]interface{}) int {
	if err := ar.validate.ValidateStructChangePassword(data, response); err != nil {
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

func (ar *AuthenticationService) ChangePasswordUser(data *dto.ChangePassword, response map[string]interface{}) int {
	if err := ar.validate.ValidateStructChangePassword(data, response); err != nil {
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

func (ar *AuthenticationService) HandleGoogleCallback(email string, response map[string]interface{}) (*http.Cookie, int) {
	boolUser, User := validation.IsUniqueEmailUser(email)
	if boolUser {
		if err := ar.repo.CreateUser(&domain.User{Email: email}); err != nil {
			response["error"] = err.Error()
			return nil, http.StatusInternalServerError
		}
		response["success"] = "Signup berhasil"
		return nil, http.StatusOK
	}

	expJwt := time.Now().Add(7 * 24 * 60 * 60)
	expCookie := 7 * 24 * 60 * 60
	cookieUser, err := middlewares.CreateJWT(User.Email, User.Id, expJwt, "user_riors_token", expCookie, response)
	if err != nil {
		return nil, http.StatusInternalServerError
	}

	return cookieUser, http.StatusOK
}

func (ar *AuthenticationService) ForgotPasswordUser(email string, response map[string]interface{}) int {
	if email == "" {
		response["ErrorField"] = map[string]string{"email": "Email field cannot be empty"}
		return http.StatusBadRequest
	}

	boolUser, user := validation.IsUniqueEmailUser(email)
	if boolUser {
		response["ErrorField"] = map[string]string{"email": "Your email not found"}
		return http.StatusBadRequest
	}

	tokenStr, err := middlewares.GenerateResetPasswordToken(email, user.Id, time.Now().Add(5*time.Minute), response)
	if err != nil {
		return http.StatusInternalServerError
	}

	endpointResetPassword := fmt.Sprintf("http://localhost:8080/auth/reset/password?token=%s", tokenStr)

	if err := ar.helper.SendEmailForgotPassword(email, user.Name, endpointResetPassword); err != nil {
		response["error"] = err.Error()
		return http.StatusInternalServerError
	}

	response["success"] = "Wait until we send an email and check your email"
	return http.StatusOK
}

func (ar *AuthenticationService) ResetPasswordUser(resetPassword *dto.ResetPassword, response map[string]interface{}) int {
	_, userId, statusCode := middlewares.VerifyResetPasswordToken(resetPassword.Token, response)
	if statusCode != 200 {
		fmt.Println("errprvnlfnfl")
		return statusCode
	}

	if err := ar.validate.ValidateStructResetPassword(resetPassword, response); err != nil {
		return http.StatusBadRequest
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(resetPassword.Password), bcrypt.DefaultCost)

	if err := ar.repo.UpdatePasswordById(string(hashPassword), userId, "user"); err != nil {
		response["error"] = err.Error()
		return http.StatusInternalServerError
	}

	response["success"] = "Successfully reset password, please login again"
	return http.StatusOK
}
