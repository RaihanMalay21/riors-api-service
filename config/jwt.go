package config

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var JWT_KEY = []byte(os.Getenv("KEY_JWT"))

type JWTClaim struct {
	UserName string
	Id uint
	jwt.RegisteredClaims
}
