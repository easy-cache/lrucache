package lrucache

import (
	"encoding/json"
	"math"
	"testing"
	"time"

	"github.com/easy-cache/cache"
	lru "github.com/hashicorp/golang-lru"
	"github.com/stretchr/testify/assert"
)

var lc, _ = lru.New(math.MaxInt8)
var ac, _ = lru.NewARC(math.MaxInt8)
var q2, _ = lru.New2Q(math.MaxInt8)

var testDataMap = map[string]interface{}{
	"lrucache": []string{"item.1", "item.2"},
}

func TestLruCache(t *testing.T) {
	var lcd = NewLRUDriver(lc)
	var lcc = NewLRUCache(lc)
	testInternal(t, lcd, lcc)
}

func TestArcCache(t *testing.T) {
	var acd = NewARCDriver(ac)
	var acc = NewARCCache(ac)
	testInternal(t, acd, acc)
}

func TestQ2Cache(t *testing.T) {
	var qcd = NewQ2Driver(q2)
	var qcc = NewQ2Cache(q2)
	testInternal(t, qcd, qcc)
}

func testInternal(t *testing.T, driver cache.DriverInterface, cache cache.Interface) {
	var ttl = time.Millisecond * 500
	for key, val := range testDataMap {
		bs, _ := json.Marshal(val)
		assert.Nil(t, driver.Set(key, bs, ttl))

		nbs, ok, err := driver.Get(key)
		assert.True(t, ok)
		assert.Nil(t, err)
		assert.Equal(t, bs, nbs)

		assert.Nil(t, driver.Del(key))

		_, ok, err = driver.Get(key)
		assert.False(t, ok)
		assert.Nil(t, err)
	}

	var tmp []string
	for key, val := range testDataMap {
		assert.True(t, cache.Set(key, val, ttl))
		assert.True(t, cache.Get(key, &tmp))
		assert.EqualValues(t, val, tmp)

		time.Sleep(ttl)

		assert.False(t, cache.Has(key))
	}
}
