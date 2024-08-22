package engineering

import (
	"GolangPractice/utils/logger"
	"context"
	"github.com/allegro/bigcache/v3"
	"github.com/coocood/freecache"
	lru "github.com/hashicorp/golang-lru/v2"
	"testing"
	"time"
)

// https://github.com/coocood/freecache
// thread-safe
func TestFreeCache(t *testing.T) {
	cache := freecache.NewCache(1 * 1024 * 1024) // 1MB

	_ = cache.Set([]byte("key"), []byte("value"), -1)

	val, err := cache.Get([]byte("key"))
	logger.Infoln(string(val), err)
}

// https://github.com/allegro/bigcache
// thread-safe.
// based on the benchmarking results, bigcache is faster than freecache.
func TestBigCache(t *testing.T) {
	cache, _ := bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))

	_ = cache.Set("my-unique-key", []byte("value"))

	entry, err := cache.Get("my-unique-key")
	logger.Infoln(string(entry), err)
}

// https://github.com/hashicorp/golang-lru
// thread-safe
func TestLRU(t *testing.T) {
	l, _ := lru.New[int, any](128)
	
	l.Add(10, nil)
	l.Get(10)
}