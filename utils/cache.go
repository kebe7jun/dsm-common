package utils

import (
	"math"
	"sort"
	"sync"
	"time"
)

type Cache interface {
	Get(string) interface{}
	Set(string, interface{})
	Keys() []string
	Clear()
}

type entry struct {
	value    interface{}
	hits     uint64
	lastSeem time.Time
}

type hotCache struct {
	data     map[string]*entry
	capacity uint64
	ttl      time.Duration
	mux      sync.RWMutex
	useTTL   bool
}

func NewHotCache(capacity uint64) Cache {
	if capacity < 1 {
		panic("capacity can not be zero.")
	}
	return &hotCache{
		data:     make(map[string]*entry, 0),
		capacity: capacity,
		mux:      sync.RWMutex{},
		useTTL:   false,
	}
}

func NewHotTTLCache(capacity uint64, ttl time.Duration) Cache {
	if capacity < 1 {
		panic("capacity can not be zero.")
	}
	return &hotCache{
		data:     make(map[string]*entry, 0),
		capacity: capacity,
		mux:      sync.RWMutex{},
		ttl:      ttl,
		useTTL:   true,
	}
}

func (cache *hotCache) Get(key string) interface{} {
	t := time.Now()
	cache.mux.Lock()
	defer cache.mux.Unlock()
	if v, ok := cache.data[key]; ok {
		if cache.useTTL && v.lastSeem.Add(cache.ttl).Before(t) {
			delete(cache.data, key)
			return nil
		}
		v.hits++
		return v.value
	}
	return nil
}

func (cache *hotCache) Set(key string, value interface{}) {
	if _, ok := cache.data[key]; ok {
		// already exists
		return
	}
	t := time.Now()
	cache.mux.Lock()
	defer cache.mux.Unlock()
	if len(cache.data) == int(cache.capacity) {
		var minKey string
		minHits := math.MaxInt64
		for k, v := range cache.data {
			if cache.useTTL && v.lastSeem.Add(cache.ttl).Before(t) {
				// outdated
				minKey = k
				break
			}
			if int(v.hits) < minHits {
				minHits = int(v.hits)
				minKey = k
			}
		}
		delete(cache.data, minKey)
	}
	cache.data[key] = &entry{
		value:    value,
		hits:     0,
		lastSeem: t,
	}
}

func (cache *hotCache) Keys() (keys []string) {
	cache.mux.RLock()
	defer cache.mux.RUnlock()
	for k := range cache.data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return
}

func (cache *hotCache) Clear() {
	cache.mux.Lock()
	defer cache.mux.Unlock()
	cache.data = make(map[string]*entry, 0)
}

type linkEntry struct {
	key   string
	value interface{}
	pre   *linkEntry
	next  *linkEntry
}

type lruCache struct {
	head     *linkEntry
	tail     *linkEntry
	capacity uint64
	len      uint64
	mux      sync.RWMutex
}

func NewLRUCache(capacity uint64) Cache {
	if capacity < 1 {
		panic("capacity can not be zero.")
	}
	return &lruCache{
		capacity: capacity,
	}
}

func (cache *lruCache) Get(key string) interface{} {
	cache.mux.Lock()
	defer cache.mux.Unlock()
	for visit := cache.head; visit != nil; visit = visit.next {
		if visit.key == key {
			if visit.pre == nil {
				// first
				return visit.value
			}
			visit.pre.next = visit.next
			if visit.next != nil {
				visit.next.pre = visit.pre
			} else {
				cache.tail = visit.pre
			}
			visit.pre = nil
			visit.next = cache.head
			cache.head.pre = visit
			cache.head = visit
			return visit.value
		}
	}
	return nil
}

func (cache *lruCache) Set(key string, value interface{}) {
	cache.mux.Lock()
	defer cache.mux.Unlock()
	if cache.len == cache.capacity {
		if cache.len == 1 {
			cache.head = nil
			cache.tail = nil
		} else {
			cache.tail.pre.next = nil
			cache.tail = cache.tail.pre
		}
		cache.len--
	}
	ent := linkEntry{
		key:   key,
		value: value,
		pre:   nil,
		next:  nil,
	}
	if cache.head == nil {
		cache.head = &ent
		cache.tail = &ent
	} else {
		cache.tail.next = &ent
		ent.pre = cache.tail
		cache.tail = &ent
	}
	cache.len++
}

func (cache *lruCache) Keys() (keys []string) {
	cache.mux.RLock()
	defer cache.mux.RUnlock()
	for visit := cache.head; visit != nil; visit = visit.next {
		keys = append(keys, visit.key)
	}
	sort.Strings(keys)
	return
}

func (cache *lruCache) Clear() {
	cache.mux.Lock()
	defer cache.mux.Unlock()
	cache.head = nil
	cache.tail = nil
	cache.len = 0
}
