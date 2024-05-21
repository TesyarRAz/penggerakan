package config

import (
	"fmt"

	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	"github.com/redis/go-redis/v9"
)

func NewRedis(env model.DotEnvConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprint(env.RedisHost(), ":", env.RedisPort()),
		Password: env.RedisPassword(),
		DB:       0,
	})
}
