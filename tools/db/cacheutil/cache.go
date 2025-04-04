package cahceutil

import "sync"

type Cache[K comparable, V any] struct {
	m sync.Map
}

func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{}
}

func (c *Cache[K, V]) Load(key K) (value V, ok bool) {
	v, ok := c.m.Load(key)
	if !ok {
		// default value is (nil,false)
		return
	} else {
		return v.(V), ok
	}
}

func (c *Cache[K, V]) Store(key K, value V) {
	c.m.Store(key, value)
}

// f means the convert function which generate key from value
func (c *Cache[K, V]) StoreAll(f func(value V) K, values []V) {
	for _, v := range values {
		c.m.Store(f(v), v)
	}
}

func (c *Cache[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	v, loaded := c.m.LoadOrStore(key, value)
	return v.(V), loaded
}

func (c *Cache[K, V]) Delete(key K) {
	c.m.Delete(key)
}

func (c *Cache[K, V]) DeleteAll() {
	c.m.Range(func(key, value interface{}) bool {
		c.m.Delete(key)
		return true
	})
}

// traverse
func (c *Cache[K, V]) RangeAll() (values []V) {
	c.m.Range(func(key, value interface{}) bool {
		values = append(values, value.(V))
		return true
	})
	return values
}

// traverse by specify condition f
func (c *Cache[K, V]) RangeCon(f func(key K, value V) bool) (values []V) {
	c.m.Range(func(k, v interface{}) bool {
		if f(k.(K), v.(V)) {
			values = append(values, v.(V))
		}
		return true
	})
	return values
}
