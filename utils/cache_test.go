package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewHotCache(t *testing.T) {
	defer func() {
		assert.NotNil(t, recover(), "hot cache should be panic for 0 capacity")
	}()
	NewHotCache(0)
}

func TestNewHotTTLCache(t *testing.T) {
	defer func() {
		assert.NotNil(t, recover(), "hot ttl cache should be panic for 0 capacity")
	}()
	NewHotTTLCache(0, 0)
}

func TestNewLRUCache(t *testing.T) {
	defer func() {
		assert.NotNil(t, recover(), "lru cache should be panic for 0 capacity")
	}()
	NewLRUCache(0)
}

func TestHotCache(t *testing.T) {
	hotCache := NewHotCache(2)
	assert.Nil(t, hotCache.Get(""), "hot cache should be nil")

	hotCache.Set("a", 1)
	assert.Equal(t, hotCache.Get("a"), 1, "hot cache res should be 1")
	hotCache.Set("a", 1)
	hotCache.Set("b", 1)
	hotCache.Set("c", 1)
	assert.Nil(t, hotCache.Get("b"), "hot cache get b should be nil")
	hotCache.Get("c")
	hotCache.Get("c")
	hotCache.Set("b", 1)
	assert.Nil(t, hotCache.Get("a"), "hot cache get a should be nil")
	assert.Equal(t, hotCache.Get("c"), 1, "hot cache get c should be 1")
	assert.Equal(t, len(hotCache.Keys()), 2, "unexpected cache keys")
	hotCache.Clear()
	assert.Nil(t, hotCache.Get("c"), "hot cache get c should be nil")
}

func TestHotTTLCache(t *testing.T) {
	cache := NewHotTTLCache(2, 100*time.Millisecond)
	cache.Set("a", 1)
	assert.Equal(t, cache.Get("a"), 1, "hot ttl cache get a should be 1")
	time.Sleep(100 * time.Millisecond)

	assert.Nil(t, cache.Get("a"), "hot ttl cache get a should be nil")
	cache.Set("a", 1)
	cache.Set("b", 1)
	cache.Set("c", 1)
	assert.Equal(t, cache.Get("c"), 1, "hot ttl cache get c should be 1")
	time.Sleep(100 * time.Millisecond)
	cache.Set("d", 1)
	assert.Nil(t, cache.Get("a"), "hot ttl cache get a should be nil")
	assert.Equal(t, len(cache.Keys()), 2, "unexpected cache keys")
}

func TestLruCache(t *testing.T) {
	cache := NewLRUCache(2)
	assert.Nil(t, cache.Get(""), "lru cache should be nil")
	if cache.Get("") != nil {
		t.Fatal("lru cache should be nil")
	}

	cache.Set("a", 1)
	assert.Equal(t, cache.Get("a"), 1, "lru cache get a should be 1")
	cache.Set("b", 1)
	assert.Equal(t, cache.Get("b"), 1, "lru cache get b should be 1")
	cache.Get("b")
	cache.Set("c", 1)
	assert.Nil(t, cache.Get("a"), "lru cache get a should be nil")
	assert.Equal(t, cache.Get("c"), 1, "lru cache get c should be 1")
	assert.Equal(t, len(cache.Keys()), 2, "unexpected cache keys")
	cache.Clear()
	assert.Nil(t, cache.Get("c"), "lru cache get c should be nil")
}

func TestLruCache3(t *testing.T) {
	cache := NewLRUCache(3)
	cache.Set("a", 1)
	cache.Set("b", 1)
	cache.Set("c", 1)
	assert.Equal(t, cache.Get("a"), 1, "lru cache get a should be 1")
	assert.Equal(t, cache.Get("b"), 1, "lru cache get b should be 1")
	cache.Set("d", 1)
	assert.Nil(t, cache.Get("c"), "lru cache get c should be nil")
	assert.Equal(t, len(cache.Keys()), 3, "unexpected cache keys")
}

func TestLruCache1(t *testing.T) {
	cache := NewLRUCache(1)
	cache.Set("a", 1)
	cache.Set("b", 1)
	assert.Nil(t, cache.Get("a"), "lru cache get a should be nil")
	assert.Equal(t, cache.Get("b"), 1, "lru cache get b should be 1")
	assert.Equal(t, len(cache.Keys()), 1, "unexpected cache keys")
}
