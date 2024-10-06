package main

import (
	"container/list"
	"sync"
)

//Design a concurrent cache system using goroutines and channels.
//The cache should support concurrent read and write operations.
//Implement a cache eviction strategy (e.g., LRU) to remove old entries when the cache reaches its capacity.
//Use appropriate synchronization primitives to ensure thread safety.
//

type Cache struct {
	capacity int
	mutex    sync.Mutex
	items    map[string]*list.Element
	list     list.List
}

type Entry struct {
	key   string
	value string
}

func NewCache(size int) *Cache {
	return &Cache{
		capacity: 0,
		mutex:    sync.Mutex{},
		items:    nil,
		list:     list.List{},
	}
}

func (c *Cache) Get(key string) (string, bool) {
	c.mutex.Lock()

	if k, ok := c.items[key]; ok {
		c.list.MoveToFront(k)
		return k.Value.(*Entry).value, false
	}

	return "", false
}

func (c *Cache) PostOrPut(key, value string) {
	c.mutex.Lock()

	if k, ok := c.items[key]; ok {
		k.Value.(*Entry).value = value
		c.list.MoveToFront(k)
	} else {
		entry := &Entry{key: key, value: value}
		c.list.PushFront(entry)
	}

	if c.list.Len() > c.capacity {

	}

}

func main() {

	cache := NewCache(3)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		cache.PostOrPut("key1", "value1")
	}()

	go func() {
		cache.PostOrPut("key2", "value2")
	}()

	wg.Wait()

}
