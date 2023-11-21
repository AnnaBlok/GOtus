package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*cacheItem
	sync.Mutex
}

type cacheItem struct {
	listItem *ListItem
	Value    interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*cacheItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.Lock()
	defer c.Unlock()

	cItem, ok := c.items[key]
	var lItem *ListItem

	if ok {
		lItem = cItem.listItem
		c.queue.MoveToFront(lItem)
	} else {
		if c.queue.Len() == c.capacity {
			delete(c.items, c.queue.Back().Value.(Key))
			c.queue.Remove(c.queue.Back())
		}
		lItem = c.queue.PushFront(key)
	}
	c.items[key] = &cacheItem{Value: value, listItem: lItem}

	return ok
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.Lock()
	defer c.Unlock()

	v, ok := c.items[key]
	var value interface{}
	if ok {
		value = v.Value
		c.queue.MoveToFront(c.items[key].listItem)
	}
	return value, ok
}

func (c *lruCache) Clear() {
	c.Lock()
	defer c.Unlock()

	c.items = make(map[Key]*cacheItem, c.capacity)
	c.queue = NewList()
}
