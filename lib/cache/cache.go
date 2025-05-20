package cache

import (
	"sync"
)

type Cache struct {
	data sync.Map

	currId int
	mu     sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		currId: -1,
	}
}

func (c *Cache) GetCurrID() int {
	return c.currId
}

func (c *Cache) Set(value interface{}) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.currId++
	c.data.Store(c.currId, value)

	return c.currId
}

func (c *Cache) Get(id int) (interface{}, bool) {
	return c.data.Load(id)
}

func (c *Cache) Delete(id int) {
	c.data.Delete(id)
}
