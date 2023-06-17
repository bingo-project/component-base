package cache

import (
	"encoding/json"
	"errors"
	"log"
	"time"
)

type CacheService struct {
	Store Store
}

func NewService(store Store) *CacheService {
	return &CacheService{
		Store: store,
	}
}

func (cs *CacheService) Set(key string, obj interface{}, expireTime time.Duration) {
	b, err := json.Marshal(&obj)
	if err != nil {
		log.Println(err)
	}

	cs.Store.Set(key, string(b), expireTime)
}

func (cs *CacheService) Get(key string) interface{} {
	stringValue := cs.Store.Get(key)
	if len(stringValue) == 0 {
		return nil
	}

	var wanted interface{}
	err := json.Unmarshal([]byte(stringValue), &wanted)
	if err != nil {
		log.Println(err)
	}

	return wanted
}

func (cs *CacheService) GetObject(key string, wanted interface{}) error {
	val := cs.Store.Get(key)
	if len(val) == 0 {
		return errors.New("not found")
	}

	err := json.Unmarshal([]byte(val), &wanted)
	if err != nil {
		return err
	}

	return nil
}

func (cs *CacheService) Has(key string) bool {
	return cs.Store.Has(key)
}

func (cs *CacheService) Forget(key string) {
	cs.Store.Forget(key)
}

func (cs *CacheService) Forever(key string, value string) {
	cs.Store.Set(key, value, 0)
}

func (cs *CacheService) Remember(key string, expireTime time.Duration, callback func() interface{}) interface{} {
	data := cs.Get(key)
	if data != nil {
		return data
	}

	data = callback()
	cs.Set(key, data, expireTime)

	return data
}

func (cs *CacheService) Flush() {
	cs.Store.Flush()
}
