package cache

import (
	"log"
	"time"
)

type MultiStore struct {
	LocalStore *LocalStore
	RedisStore *RedisStore
}

func (s *MultiStore) Set(key string, value string, expireTime time.Duration) {
	s.LocalStore.Set(key, value, 0)
	s.RedisStore.Set(key, value, expireTime)
}

func (s *MultiStore) Get(key string) string {
	value := s.LocalStore.Get(key)
	if len(value) > 0 {
		log.Println("get from local store")

		return value
	}

	log.Println("get from redis store")

	return s.RedisStore.Get(key)
}

func (s *MultiStore) Has(key string) bool {
	has := s.LocalStore.Has(key)
	if has {
		return has
	}

	return s.RedisStore.Has(key)
}

func (s *MultiStore) Forget(key string) {
	s.LocalStore.Forget(key)
	s.RedisStore.Forget(key)
}

func (s *MultiStore) Forever(key string, value string) {
	s.LocalStore.Forever(key, value)
	s.RedisStore.Forever(key, value)
}

func (s *MultiStore) Flush() {
	s.LocalStore.Flush()
	s.RedisStore.Flush()
}

func (s *MultiStore) Incr(key string) {
	s.LocalStore.Incr(key)
	s.RedisStore.Incr(key)
}

func (s *MultiStore) IncrBy(key string, value int64) {
	s.LocalStore.IncrBy(key, value)
	s.RedisStore.IncrBy(key, value)
}

func (s *MultiStore) Decr(key string) {
	s.LocalStore.Decr(key)
	s.RedisStore.Decr(key)
}

func (s *MultiStore) DecrBy(key string, value int64) {
	s.LocalStore.DecrBy(key, value)
	s.RedisStore.DecrBy(key, value)
}
