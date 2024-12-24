package middlewares

import (
	"net/http"
	"time"

	"github.com/RaihanMalay21/api-service-riors/config"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(username string, id uint, expToken time.Time, jwtName string, expCookie int, response map[string]interface{}) (*http.Cookie, error) {
	
	claims := &config.JWTClaim{
		UserName: username,
		Id: id, 
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expToken),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response["error"] = err.Error()
		return nil, err
	}

	cookie := &http.Cookie{
		Name:     jwtName,
		Value:    token,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		MaxAge:   expCookie,
		SameSite: http.SameSiteNoneMode,
	}

	return cookie, nil
}