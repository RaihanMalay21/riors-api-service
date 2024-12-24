package middlewares

import (
	"errors"
	"net/http"

	"github.com/RaihanMalay21/api-service-riors/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		response := make(map[string]interface{})

// 		cookie, err := c.Cookie("user_riors_token")
// 		if err != nil {
// 			response["message"] = "Token is missing"
// 			return c.JSON(http.StatusUnauthorized, response)
// 		}

// 		tokenString := cookie.Value
// 		claims := &config.JWTClaim{}

// 		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
// 			return config.JWT_KEY, nil
// 		})
// 		if err != nil {
// 			if errors.Is(err, jwt.ErrTokenSignatureInvalid){
// 				response["message"] = "Token signature is invalid"
// 				c.JSON(http.StatusUnauthorized, response)
// 			} else if errors.Is(err, jwt.ErrTokenExpired) {
// 				response["message"] = "Token signature has Expired"
// 				c.JSON(http.StatusUnauthorized, response)
// 			} else {
// 				response["error"] = err.Error()
// 				c.JSON(http.StatusInternalServerError, response)
// 			}
// 		}

// 		if claims, ok := token.Claims.(*config.JWTClaim); ok && token.Valid{
// 			role := claims.Role
// 			endPoint := c.Request().URL.Path
// 			method := c.Request().Method

// 			if err := endPointCanAccess(&role, &endPoint, &method); err != nil {
// 				response["message"] = err.Error()
// 				return c.JSON(http.StatusUnauthorized, response)
// 			}
// 		} else {
// 			response["message"] = "Unauthorized"
// 			c.JSON(http.StatusUnauthorized, response)
// 		}

// 		return next(c)
// 	}
// }

// func endPointCanAccess(role, endPoint, method *string) error {
// 	var endPoinAdmin = map[string][]string{
// 		"POST" : {"/category", "/product"},
// 	}

// 	var endPoinUser = map[string][]string{
// 	}

// 	switch *role {
// 	case "user":
// 		for _, en := range endPoinUser[*method] {
// 			if strings.HasPrefix(*endPoint, en) {
// 				return nil
// 			}
// 		}
// 	case "admin":
// 		for _, en := range endPoinAdmin[*method] {
// 			if strings.HasPrefix(*endPoint, en) {
// 				return nil
// 			}
// 		}
// 	}

// 	return fmt.Errorf("access denied to endpoint: %s", endPoint)
// }

func ArmorUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		response := make(map[string]interface{})

		cookie, err := c.Cookie("user_riors_token")
		if err != nil {
			response["message"] = "Token is missing"
			return c.JSON(http.StatusUnauthorized, response)
		}

		tokenString := cookie.Value
		claims := &config.JWTClaim{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})
		if err != nil {
			if errors.Is(err, jwt.ErrTokenSignatureInvalid){
				response["message"] = "Token signature is invalid"
				return c.JSON(http.StatusUnauthorized, response)
			} else if errors.Is(err, jwt.ErrTokenExpired) {
				response["message"] = "Token signature has expired"
				return c.JSON(http.StatusUnauthorized, response)
			} else {
				response["error"] = err.Error()
				return c.JSON(http.StatusInternalServerError, response)
			}
		}

		if claims, ok := token.Claims.(*config.JWTClaim); ok && token.Valid{
			c.Set("user_claims", claims)
		} else {
			response["message"] = "Unauthorized"
			return c.JSON(http.StatusUnauthorized, response)
		}

		return next(c)
	}
}

func ArmorAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		response := make(map[string]interface{})

		cookie, err := c.Cookie("admin_riors_token")
		if err != nil {
			response["message"] = "Token is missing"
			return c.JSON(http.StatusUnauthorized, response)
		}

		tokenString := cookie.Value
		claims := &config.JWTClaim{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})
		if err != nil {
			if errors.Is(err, jwt.ErrTokenSignatureInvalid){
				response["message"] = "Token signature is invalid"
				return c.JSON(http.StatusUnauthorized, response)
			} else if errors.Is(err, jwt.ErrTokenExpired) {
				response["message"] = "Token signature has Expired"
				return c.JSON(http.StatusUnauthorized, response)
			} else {
				response["error"] = err.Error()
				return c.JSON(http.StatusInternalServerError, response)
			}
		}

		if claims, ok := token.Claims.(*config.JWTClaim); ok && token.Valid{
			c.Set("admin_claims", claims)
		} else {
			response["message"] = "Unauthorized"
			return c.JSON(http.StatusUnauthorized, response)
		}

		return next(c)
	}
}