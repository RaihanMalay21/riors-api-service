package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CorsMiddlewares(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Response().Header().Set("Access-Control-Allow-Methods", "POST, PUT, PATCH, GET, DELETE, OPTIONS")
		c.Response().Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Authorization")
		c.Response().Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request().Method == http.MethodOptions {
			return c.NoContent(http.StatusOK)
		}

		return next(c)
	}
}
