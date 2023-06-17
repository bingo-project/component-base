package cache

import "time"

type Store interface {
	Set(key string, value string, expireTime time.Duration)
	Get(key string) string
	Has(key string) bool
	Forget(key string)
	Forever(key string, value string)
	Flush()
	Incr(key string)
	IncrBy(key string, value int64)
	Decr(key string)
	DecrBy(key string, value int64)
}
