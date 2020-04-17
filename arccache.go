package lrucache

import (
	"time"

	"github.com/easy-cache/cache"
	"github.com/hashicorp/golang-lru"
)

type arcCacheDriver struct {
	arccache *lru.ARCCache
}

func (acd arcCacheDriver) Get(key string) ([]byte, bool, error) {
	item, hit := acd.arccache.Get(key)
	if hit == false {
		return nil, false, nil
	}
	bs, ok := item.(*cache.Item).GetValue()
	return bs, ok, nil
}

func (acd arcCacheDriver) Set(key string, val []byte, ttl time.Duration) error {
	acd.arccache.Add(key, cache.NewItem(val, ttl))
	return nil
}

func (acd arcCacheDriver) Del(key string) error {
	acd.arccache.Remove(key)
	return nil
}

func (acd arcCacheDriver) Has(key string) (bool, error) {
	_, ok, err := acd.Get(key)
	return ok, err
}

func NewARCDriver(arccache *lru.ARCCache) cache.DriverInterface {
	return arcCacheDriver{arccache: arccache}
}

func NewARCCache(arccache *lru.ARCCache, args ...interface{}) cache.Interface {
	return cache.New(append(args, NewARCDriver(arccache))...)
}
