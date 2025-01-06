package authentication

import (
	"encoding/json"
	"net/http"
	"os"

	// "github.com/RaihanMalay21/api-service-riors/domain"
	"github.com/RaihanMalay21/api-service-riors/config"
	"github.com/RaihanMalay21/api-service-riors/controller/helper"
	"github.com/RaihanMalay21/api-service-riors/dto"
	"github.com/RaihanMalay21/api-service-riors/middlewares"
	"github.com/RaihanMalay21/api-service-riors/service/authentication"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type AuthenticationController struct {
	service *authentication.AuthenticationService
	helper  *helper.HelperController
}

func ConstructorAuthenticationController(service *authentication.AuthenticationService, helper *helper.HelperController) *AuthenticationController {
	return &AuthenticationController{
		service: service,
		helper:  helper,
	}
}

// @summary User login
// @Description This enpoint is used to authenticate users with send cridential data (email dan password).
// @Tags authentication
// @accept  application/x-www-form-urlencoded
// @Produce  application/json
// @Param Login body RegisterUser true "Login Input"
// @Success 200  {object}  ResponseSuccess "Successfuly login, return a token to access enpoint for user"
// @Failure 400  {object}  ResponseErrorBadRequest "Request invalid or the data sent is incorrect"
// @Failure 500  {object}  ResponseErrorInternalServer "Mistake in the server"
// @Router /auth/login/user [post]
func (as *AuthenticationController) LoginUser(c echo.Context) error {
	response := make(map[string]interface{})

	email := c.FormValue("email")
	password := c.FormValue("password")

	cookie, statusCode := as.service.LoginUser(email, password, response)
	if statusCode != 200 {
		return c.JSON(statusCode, response)
	}

	c.SetCookie(cookie)
	return c.JSON(statusCode, response)
}

// @summary Create New User Account
// @Description This endpoint is used to create a new user account by submitting registration data in JSON format. Users are required to provide information email, and password. The data will be validated on the server side before the account is created temporary, client have navigate to endpoint signup verification on the client side.
// @Tags authentication
// @Accept application/json
// @Produce application/json
// @Param Signup body RegisterUser true "User Registration Data (name, email, password)"
// @Success 200 {object} ResponseSuccess "Account successfully created temporery, return a token to verification email"
// @Failure 400 {object} ResponseErrorBadRequest "Invalid request or incomplete data"
// @Failure 500 {object} ResponseErrorInternalServer "Internal server error while processing the request"
// @Router /auth/signup/user [post]
func (as *AuthenticationController) SignupUser(c echo.Context) error {
	response := make(map[string]interface{})
	register := new(dto.RegisterUser)
	if err := c.Bind(&register); err != nil {
		response["error"] = err.Error()
		return c.JSON(http.StatusInternalServerError, response)
	}

	cookie, statusCode := as.service.SignupUser(register, response)
	if statusCode != 200 {
		return c.JSON(statusCode, response)
	}

	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, response)
}

// @summary Verify User Email for Signup
// @Description This endpoint is used to verify the user's email during the registration process. A verification code is sent to the user's email, and the user must input this code to complete the signup process successfully. The request must include the verification code in the form body.
// @Tags authentication
// @Security BearerAuth
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Signup body Verification true "Email Verification Data (code)"
// @Success 200 {object} ResponseSuccess "Email successfully verified, account registration completed"
// @Failure 400 {object} ResponseErrorBadRequest "Invalid or missing verification code"
// @Failure 401 {object} ResponsAuthorization "Unauthorized - Missing or invalid token"
// @Failure 500 {object} ResponseErrorInternalServer "Internal server error while processing the request"
// @Router /auth/signup/user/verification [post]
func (as *AuthenticationController) SignupUserVerification(c echo.Context) error {
	response := make(map[string]interface{})

	email, statusCode := middlewares.VerifyAndExtractTokenClaims(c, response)
	if statusCode != 200 {
		return c.JSON(statusCode, response)
	}

	varificationCode := c.FormValue("code")

	StatusCode := as.service.SignupUserVerification(email, &varificationCode, response)

	return c.JSON(StatusCode, response)
}

// @summary Admin login
// @Description This endpoint is used to login by email and password.
// @Tags authentication
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Login body RegisterUser true "login data"
// @Success 200 {object} ResponseSuccess "user successfully login"
// @Failure 400 {object} ResponseErrorBadRequest "Invalid or missing"
// @Failure 401 {object} ResponsAuthorization "Unauthorized - Missing or invalid token"
// @Failure 500 {object} ResponseErrorInternalServer "Internal server error while processing the request"
// @Router /auth/login/admin [post]
func (as *AuthenticationController) LoginAdmin(c echo.Context) error {
	response := make(map[string]interface{})

	email := c.FormValue("email")
	password := c.FormValue("password")

	cookieAdmin, cookieOwner, statusCode := as.service.LoginAdmin(email, password, response)
	if statusCode != 200 {
		return c.JSON(statusCode, response)
	}

	if cookieOwner == nil {
		c.SetCookie(cookieAdmin)
	} else {
		c.SetCookie(cookieAdmin)
		c.SetCookie(cookieOwner)
	}

	return c.JSON(statusCode, response)
}

// @summary signup for admin side
// @Description This endpoint is used to register new employee to access admin side.
// @Tags admin
// @Security BearerAuth
// @Accept multipart/form-data
// @Produce application/json
// @Param Signup body SignupEmploye true "Signup data"
// @Param employementType formData string true "Employment Type (Tetap, Kontrak, Freelance)" Enums(Tetap, Kontrak, Freelance)
// @Param gender formData string true "Gender (Man, Woman)" Enums(Man, Woman)
// @Param position formData string true "Position (Staff, Owner)" Enums(Staff, Owner)
// @Success 200 {object} ResponseSuccess "successfuly register employee"
// @Failure 400 {object} ResponseErrorBadRequest "Invalid or missing data"
// @Failure 401 {object} ResponsAuthorization "Unauthorized - Missing or invalid token"
// @Failure 500 {object} ResponseErrorInternalServer "Internal server error while processing the request"
// @Router /admin/owner/register/employe [post]
func (as *AuthenticationController) SignupEmploye(c echo.Context) error {
	response := make(map[string]interface{})

	file, fileHeader, Ext, filetype, statusCode := as.helper.GetFileFromForm(c, response)
	if statusCode != 200 {
		return c.JSON(statusCode, response)
	}

	data := dto.Employee{
		Name:            c.FormValue("name"),
		Email:           c.FormValue("email"),
		Whatsapp:        c.FormValue("whatsapp"),
		Password:        os.Getenv("AUTO_PASSWORD_SIGNUP"),
		Position:        c.FormValue("position"),
		EmployementType: c.FormValue("employementType"),
		DateOfBirth:     c.FormValue("dateOfBirth"),
		Gender:          c.FormValue("gender"),
		Address:         c.FormValue("address"),
		Image:           fileHeader.Filename,
		FileSize:        uint(fileHeader.Size),
		Ext:             Ext,
		ImageType:       filetype,
	}

	StatusCode := as.service.SignupEmploye(file, fileHeader, &data, response)

	return c.JSON(StatusCode, response)
}

// @summary change password admin
// @Description This endpoint is used to change password in admin side.
// @Tags admin
// @Security BearerAuth
// @Accept x-www-form-urlencoded
// @Produce application/json
// @Param ChangePassword body ChangePassword true "Change password data"
// @Success 200 {object} ResponseSuccess "successfuly change password
// @Failure 400 {object} ResponseErrorBadRequest "Invalid or missing data"
// @Failure 401 {object} ResponsAuthorization "Unauthorized - Missing or invalid token"
// @Failure 500 {object} ResponseErrorInternalServer "Internal server error while processing the request"
// @Router /admin/change/password [patch]
func (as *AuthenticationController) ChangePasswordAdmin(c echo.Context) error {
	response := make(map[string]interface{})

	claims, ok := c.Get("admin_claims").(*config.JWTClaim)
	if !ok {
		response["error"] = "Unable to retrieve claims"
		return c.JSON(http.StatusUnauthorized, response)
	}

	data := &dto.ChangePassword{
		Id:             claims.Id,
		PasswordBefore: c.FormValue("passwordBefore"),
		Password:       c.FormValue("password"),
	}

	statusCode := as.service.ChangePasswordAdmin(data, response)

	return c.JSON(statusCode, response)
}

// @summary change password user
// @Description This endpoint is used to change password in client side or user.
// @Tags user
// @Accept x-www-form-urlencoded
// @Security BearerAuth
// @Produce application/json
// @Param ChangePassword body ChangePassword true "Change password data"
// @Success 200 {object} ResponseSuccess "successfuly change password
// @Failure 400 {object} ResponseErrorBadRequest "Invalid or missing data"
// @Failure 401 {object} ResponsAuthorization "Unauthorized - Missing or invalid token"
// @Failure 500 {object} ResponseErrorInternalServer "Internal server error while processing the request"
// @Router /user/change/password [patch]
func (as *AuthenticationController) ChangePasswordUser(c echo.Context) error {
	response := make(map[string]interface{})

	claims, ok := c.Get("user_claims").(*config.JWTClaim)
	if !ok {
		response["error"] = "Unable to retrieve claims"
		return c.JSON(http.StatusUnauthorized, response)
	}

	data := &dto.ChangePassword{
		Id:             claims.Id,
		PasswordBefore: c.FormValue("passwordBefore"),
		Password:       c.FormValue("password"),
	}

	statusCode := as.service.ChangePasswordUser(data, response)

	return c.JSON(statusCode, response)
}

func (as *AuthenticationController) HandleGoogleLogin(c echo.Context) error {
	url := config.GoogleOauth2Config.AuthCodeURL(config.RandomState, oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusFound, url)
}

func (as *AuthenticationController) HandleGoogleCallback(c echo.Context) error {
	var response = make(map[string]interface{})
	if c.QueryParam("state") != config.RandomState {
		response["message"] = "State did not match"
		return c.JSON(http.StatusBadRequest, response)
	}

	code := c.QueryParam("code")
	token, err := config.GoogleOauth2Config.Exchange(c.Request().Context(), code)
	if err != nil {
		response["error"] = "Failed to exchange token"
		return c.JSON(http.StatusInternalServerError, response)
	}

	// Menggunakan token untuk mendapatkan informasi pengguna
	client := config.GoogleOauth2Config.Client(c.Request().Context(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		response["error"] = "Failed to get user info"
		return c.JSON(http.StatusInternalServerError, response)
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		response["error"] = "Failed to parse user info"
		return c.JSON(http.StatusInternalServerError, response)
	}

	cookie, statusCode := as.service.HandleGoogleCallback(userInfo["email"].(string), response)
	if cookie == nil {
		// redirect ke halamn login
		return c.JSON(statusCode, response)
	}

	// redirect ke halaman user
	c.SetCookie(cookie)
	response["message"] = "login berhasil"
	return c.JSON(http.StatusOK, response)
}

func (as *AuthenticationController) ForgotPasswordUser(e echo.Context) error {
	response := make(map[string]interface{})

	statusCode := as.service.ForgotPasswordUser(e.QueryParam("email"), response)

	return e.JSON(statusCode, response)
}

func (as *AuthenticationController) ResetPasswordUser(e echo.Context) error {
	response := make(map[string]interface{})

	password := dto.ResetPassword{
		Token: e.FormValue("token"),
		Password: e.FormValue("password"),
		PasswordRepeat: e.FormValue("passwordRepeat"),
	}

	statusCode := as.service.ResetPasswordUser(&password, response)
	
	return e.JSON(statusCode, response)
}
