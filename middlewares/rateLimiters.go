package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/RaihanMalay21/api-service-riors/config"
	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
)

const (
	rateLimitTTL = time.Second // Waktu TTL untuk rate limiter (60 detik)
	maxRequests  = 5
)

func getLimiter(ip string) *rate.Limiter {
	ctx := context.Background()
	client := config.Conn

	count, err := client.Get(ctx, ip).Int()
	if err != nil && err.Error() != "redis: nil" {
		fmt.Println("Error Redis:", err)
		return rate.NewLimiter(rate.Limit(2), 5)
	}

	if count > maxRequests {
		return rate.NewLimiter(0, 0)
	}

	limiter := rate.NewLimiter(rate.Limit(2), 5)

	if _, err := client.Set(ctx, ip, count+1, rateLimitTTL).Result(); err != nil {
		fmt.Println("Error Redis Set:", err)
	}

	return limiter
}

func SetLimiterMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ip := c.RealIP()

		limiter := getLimiter(ip)

		if !limiter.Allow() {
			return c.JSON(http.StatusTooManyRequests, map[string]string{
				"message": "Rate limit exceeded",
			})
		}

		return next(c)
	}
}
