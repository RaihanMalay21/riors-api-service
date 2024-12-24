package middlewares

import (
	"errors"
	"net/http"

	"github.com/RaihanMalay21/api-service-riors/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func VerifyAndExtractTokenClaims(c echo.Context, response map[string]interface{}) (*string, int) {
	cookie, err := c.Cookie("register_riors_token")
	if err != nil {
		if err.Error() == "http: named cookie not present" {
			response["message"] = "Token is missing"
			return nil, http.StatusUnauthorized
		}
		response["error"] = err.Error()
		return nil, http.StatusInternalServerError
	}

	tokenString := cookie.Value
	claims := &config.JWTClaim{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return config.JWT_KEY, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			response["message"] = "Token signature is invalid"
			return nil, http.StatusUnauthorized
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			response["message"] = "Token signature has Expired"
			return nil, http.StatusUnauthorized
		} else {
			response["error"] = err.Error()
			return nil, http.StatusInternalServerError
		}
	}

	if claim, ok := token.Claims.(*config.JWTClaim); ok && token.Valid {
		email := claim.UserName
		return &email, http.StatusOK
	}

	response["message"] = "Unauthorized"
	return nil, http.StatusUnauthorized
}
