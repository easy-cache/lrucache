package lrucache

import (
	"time"

	"github.com/easy-cache/cache"
	"github.com/hashicorp/golang-lru"
)

type q2CacheDriver struct {
	q2cache *lru.TwoQueueCache
}

func (qcd q2CacheDriver) Get(key string) ([]byte, bool, error) {
	item, hit := qcd.q2cache.Get(key)
	if hit == false {
		return nil, false, nil
	}
	bs, ok := item.(*cache.Item).GetValue()
	return bs, ok, nil
}

func (qcd q2CacheDriver) Set(key string, val []byte, ttl time.Duration) error {
	qcd.q2cache.Add(key, cache.NewItem(val, ttl))
	return nil
}

func (qcd q2CacheDriver) Del(key string) error {
	qcd.q2cache.Remove(key)
	return nil
}

func (qcd q2CacheDriver) Has(key string) (bool, error) {
	_, ok, err := qcd.Get(key)
	return ok, err
}

func NewQ2Driver(q2cache *lru.TwoQueueCache) cache.DriverInterface {
	return q2CacheDriver{q2cache: q2cache}
}

func NewQ2Cache(q2cache *lru.TwoQueueCache, args ...interface{}) *cache.Cache {
	return cache.New(append(args, NewQ2Driver(q2cache))...)
}
