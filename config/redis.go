package config

import (
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	Conn *redis.Client
)

func ConnectionRedis() {
	addrRedis := os.Getenv("REDIS_URL")

	opt, err := redis.ParseURL(addrRedis)
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)

	Conn = client
}
