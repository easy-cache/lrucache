package lrucache

import (
	"time"

	"github.com/easy-cache/cache"
	lru "github.com/hashicorp/golang-lru"
)

type lruCacheDriver struct {
	lrucache *lru.Cache
}

func (lcd lruCacheDriver) Get(key string) ([]byte, bool, error) {
	item, hit := lcd.lrucache.Get(key)
	if hit == false {
		return nil, false, nil
	}
	bs, ok := item.(*cache.Item).GetValue()
	if ok == false {
		_ = lcd.Del(key)
	}
	return bs, ok, nil
}

func (lcd lruCacheDriver) Set(key string, val []byte, ttl time.Duration) error {
	lcd.lrucache.Add(key, cache.NewItem(val, ttl))
	return nil
}

func (lcd lruCacheDriver) Del(key string) error {
	lcd.lrucache.Remove(key)
	return nil
}

func NewLRUDriver(lrucache *lru.Cache) cache.DriverInterface {
	return lruCacheDriver{lrucache: lrucache}
}

func NewLRUCache(lrucache *lru.Cache, args ...interface{}) cache.Interface {
	return cache.New(append(args, NewLRUDriver(lrucache))...)
}

func NewLRUCache2(size int, args ...interface{}) cache.Interface {
	lruCache, _ := lru.New(size)
	return NewLRUCache(lruCache, args...)
}
