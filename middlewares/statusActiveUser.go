package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/RaihanMalay21/api-service-riors/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

func DetectionStatusActiveUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		cookie, err := e.Cookie("user_riors_token")
		if err != nil {
			fmt.Println(err)
			return next(e)
		}

		tokenString := cookie.Value
		claim := &config.JWTClaim{}

		token, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})
		if err != nil {
			fmt.Println(err)
			return next(e)
		}

		claims, ok := token.Claims.(*config.JWTClaim)
		if ok {
			// Kirim event ke Redis hanya jika user terautentikasi
			expireAt := time.Now().Add(1 * time.Minute).Unix()
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()
			cmd := config.Conn.XAdd(ctx, &redis.XAddArgs{
				Stream: "active_stream_user",
				Values: map[string]interface{}{
					"user_id":   claims.Id,
					"expire_at": expireAt,
				},
			})

			if cmd.Err() != nil {
				response := map[string]interface{}{"error": fmt.Sprintf("Failed to add user activity to Redis: %v", cmd.Err())}
				return e.JSON(http.StatusInternalServerError, response)
			}
		}

		return next(e)
	}
}
