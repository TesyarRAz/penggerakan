package repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	Redis *redis.Client
}

func NewRedisRepository(redis *redis.Client) *RedisRepository {
	return &RedisRepository{Redis: redis}
}

func (r *RedisRepository) SetMapByTags(ctx context.Context, key string, values map[string]interface{}, expire time.Duration, tags ...string) error {
	pipe := r.Redis.TxPipeline()
	for _, tag := range tags {
		pipe.SAdd(ctx, tag, key)
		pipe.Expire(ctx, tag, expire)
	}

	pipe.HSet(ctx, key, values)

	_, err := pipe.Exec(ctx)
	return err
}

func (r *RedisRepository) GetMapByTags(ctx context.Context, key string) (map[string]string, error) {
	return r.Redis.HGetAll(ctx, key).Result()
}

func (r *RedisRepository) Delete(ctx context.Context, key string) error {
	return r.Redis.Del(ctx, key).Err()
}

func (r *RedisRepository) SetMap(ctx context.Context, key string, values map[string]interface{}) error {
	return r.Redis.HSet(ctx, key, values).Err()
}

func (r *RedisRepository) Exists(ctx context.Context, key string) (bool, error) {
	count, err := r.Redis.Exists(ctx, key).Result()

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *RedisRepository) GetMap(ctx context.Context, key string) (map[string]string, error) {
	return r.Redis.HGetAll(ctx, key).Result()
}

func (r *RedisRepository) Set(ctx context.Context, key string, value interface{}, expire time.Duration, tags ...string) error {
	pipe := r.Redis.TxPipeline()
	for _, tag := range tags {
		pipe.SAdd(ctx, tag, key)
	}
	pipe.Set(ctx, key, value, expire)

	_, err := pipe.Exec(ctx)

	return err
}

func (r *RedisRepository) Get(ctx context.Context, key string) (string, error) {
	return r.Redis.Get(ctx, key).Result()
}
