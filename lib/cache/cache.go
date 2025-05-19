package cache

import (
	"sync"
	"sync/atomic"
)

type Cache struct {
	data   sync.Map
	currId uint64
}

func NewCache() *Cache {
	return &Cache{}
}

func (c *Cache) GetCurrID() uint64 {
	return c.currId
}

func (c *Cache) Set(value interface{}) uint64 {
	id := atomic.AddUint64(&c.currId, 1) - 1
	c.data.Store(id, value)
	return id
}

func (c *Cache) Get(id uint64) (interface{}, bool) {
	return c.data.Load(id)
}

func (c *Cache) Delete(id uint64) {
	c.data.Delete(id)
}
