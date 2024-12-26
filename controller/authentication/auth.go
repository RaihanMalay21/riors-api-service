package authentication

import (
	"net/http"
	"os"

	// "github.com/RaihanMalay21/api-service-riors/domain"
	"github.com/RaihanMalay21/api-service-riors/config"
	"github.com/RaihanMalay21/api-service-riors/controller"
	"github.com/RaihanMalay21/api-service-riors/dto"
	"github.com/RaihanMalay21/api-service-riors/middlewares"
	service "github.com/RaihanMalay21/api-service-riors/service/authentication"
	"github.com/labstack/echo/v4"
)

type AuthenticationController interface {
	LoginUser(c echo.Context) error
	SignupUser(c echo.Context) error
	SignupUserVerification(c echo.Context) error
	LoginAdmin(c echo.Context) error
	SignupEmploye(c echo.Context) error
	ChangePasswordAdmin(c echo.Context) error
	ChangePasswordUser(c echo.Context) error
}

type authenticationController struct {
	service service.AuthenticationService
}

func ConstructorAuthenticationController(service service.AuthenticationService) AuthenticationController {
	return &authenticationController{service: service}
}

// @summary User Login
// @Description This enpoint is used to authenticate users with send cridential data (email dan password).
// @Tags authentication
// @accept  application/x-www-form-urlencoded
// @Produce  application/json
// @Param Login body RegisterUser true "Login Input"
// @Success 200  {object}  ResponseSuccess "Successfuly login, return a token to access enpoint for user"
// @Failure 400  {object}  ResponseErrorBadRequest "Request invalid or the data sent is incorrect"
// @Failure 500  {object}  ResponseErrorInternalServer "Mistake in the server"
// @Router /login/user [post]
func (as *authenticationController) LoginUser(c echo.Context) error {
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
// @Router /signup/user [post]
func (as *authenticationController) SignupUser(c echo.Context) error {
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
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Param Signup body Verification true "Email Verification Data (code)"
// @Success 200 {object} ResponseSuccess "Email successfully verified, account registration completed"
// @Failure 400 {object} ResponseErrorBadRequest "Invalid or missing verification code"
// @Failure 401 {object} ResponsAuthorization "Unauthorized - Missing or invalid token"
// @Failure 500 {object} ResponseErrorInternalServer "Internal server error while processing the request"
// @Router /signup/user/verification [post]
func (as *authenticationController) SignupUserVerification(c echo.Context) error {
	response := make(map[string]interface{})

	email, statusCode := middlewares.VerifyAndExtractTokenClaims(c, response)
	if statusCode != 200 {
		return c.JSON(statusCode, response)
	}

	varificationCode := c.FormValue("code")

	StatusCode := as.service.SignupUserVerification(email, &varificationCode, response)

	return c.JSON(StatusCode, response)
}

func (as *authenticationController) LoginAdmin(c echo.Context) error {
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

func (as *authenticationController) SignupEmploye(c echo.Context) error {
	response := make(map[string]interface{})

	file, fileHeader, Ext, filetype, statusCode := controller.GetFileFromForm(c, response)
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

func (as *authenticationController) ChangePasswordAdmin(c echo.Context) error {
	response := make(map[string]interface{})

	claims, ok := c.Get("admin_claims").(*config.JWTClaim)
	if !ok {
		response["error"] = "Unable to retrieve claims"
		return c.JSON(http.StatusUnauthorized, response)
	}

	data := &dto.ChangePassword{
		Id : claims.Id,
		PasswordBefore: c.FormValue("passwordBefore"),
		Password : c.FormValue("password"),
	}

	statusCode := as.service.ChangePasswordAdmin(data, response)

	return c.JSON(statusCode, response)
}

func (as *authenticationController) ChangePasswordUser(c echo.Context) error {
	response := make(map[string]interface{})

	claims, ok := c.Get("user_claims").(*config.JWTClaim)
	if !ok {
		response["error"] = "Unable to retrieve claims"
		return c.JSON(http.StatusUnauthorized, response)
	}

	data := &dto.ChangePassword{
		Id: claims.Id,
		PasswordBefore: c.FormValue("passwordBefore"),
		Password: c.FormValue("password"),
	}

	statusCode := as.service.ChangePasswordUser(data, response)

	return c.JSON(statusCode, response)
}
 