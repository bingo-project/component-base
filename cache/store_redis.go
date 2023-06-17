package cache

import (
	"time"

	"github.com/bingo-project/component-base/redis"
)

type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

func (s *RedisStore) Set(key string, value string, expireTime time.Duration) {
	s.RedisClient.Set(s.KeyPrefix+key, value, expireTime)
}

func (s *RedisStore) Get(key string) string {
	return s.RedisClient.Get(s.KeyPrefix + key)
}

func (s *RedisStore) Has(key string) bool {
	return s.RedisClient.Has(s.KeyPrefix + key)
}

func (s *RedisStore) Forget(key string) {
	s.RedisClient.Del(s.KeyPrefix + key)
}

func (s *RedisStore) Forever(key string, value string) {
	s.RedisClient.Set(s.KeyPrefix+key, value, 0)
}

func (s *RedisStore) Flush() {
	s.RedisClient.FlushDB()
}

func (s *RedisStore) Incr(key string) {
	s.RedisClient.Incr(s.KeyPrefix + key)
}

func (s *RedisStore) IncrBy(key string, value int64) {
	s.RedisClient.IncrBy(s.KeyPrefix+key, value)
}

func (s *RedisStore) Decr(key string) {
	s.RedisClient.Decr(s.KeyPrefix + key)
}

func (s *RedisStore) DecrBy(key string, value int64) {
	s.RedisClient.DecrBy(s.KeyPrefix+key, value)
}
