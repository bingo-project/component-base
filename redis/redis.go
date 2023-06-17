package redis

import (
	"context"
	"time"

	redis "github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client  *redis.Client
	Context context.Context
}

func NewClient(addr string, password string, db int) (rds *RedisClient, err error) {
	rds = &RedisClient{}
	rds.Context = context.Background()

	// New client
	rds.Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// Ping
	err = rds.Ping()
	if err != nil {
		return
	}

	return
}

func (rds RedisClient) Ping() (err error) {
	_, err = rds.Client.Ping(rds.Context).Result()

	return
}

func (rds RedisClient) Set(key string, value interface{}, expiration time.Duration) bool {
	if err := rds.Client.Set(rds.Context, key, value, expiration).Err(); err != nil {
		return false
	}

	return true
}

func (rds RedisClient) Get(key string) string {
	result, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		return ""
	}

	return result
}

func (rds RedisClient) Has(key string) bool {
	_, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		return false
	}

	return true
}

func (rds RedisClient) Del(keys ...string) bool {
	if err := rds.Client.Del(rds.Context, keys...).Err(); err != nil {
		return false
	}

	return true
}

func (rds RedisClient) FlushDB() bool {
	if err := rds.Client.FlushDB(rds.Context).Err(); err != nil {
		return false
	}

	return true
}

func (rds RedisClient) Incr(key string) bool {
	if err := rds.Client.Incr(rds.Context, key).Err(); err != nil {
		return false
	}

	return true
}

func (rds RedisClient) IncrBy(key string, value int64) bool {
	if err := rds.Client.IncrBy(rds.Context, key, value).Err(); err != nil {
		return false
	}

	return true
}

func (rds RedisClient) Decr(key string) bool {
	if err := rds.Client.Decr(rds.Context, key).Err(); err != nil {
		return false
	}

	return true
}

func (rds RedisClient) DecrBy(key string, value int64) bool {
	if err := rds.Client.DecrBy(rds.Context, key, value).Err(); err != nil {
		return false
	}

	return true
}
