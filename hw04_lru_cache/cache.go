package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (lru *lruCache) Set(key Key, value interface{}) bool {
	cache := cacheItem{key, value}
	if item, ok := lru.items[key]; ok {
		item.Value = cache
		lru.queue.MoveToFront(item)

		return true
	}

	item := lru.queue.PushFront(cache)
	lru.items[key] = item

	if len(lru.items) > lru.capacity {
		back := lru.queue.Back()
		lru.queue.Remove(back)
		item := back.Value.(cacheItem)
		delete(lru.items, item.key)
	}

	return false
}

func (lru *lruCache) Get(key Key) (interface{}, bool) {
	item, ok := lru.items[key]
	if !ok {
		return item, ok
	}

	lru.queue.MoveToFront(item)

	return item.Value.(cacheItem).value, ok
}

func (lru *lruCache) Clear() {
	lru.queue = NewList()
	lru.items = make(map[Key]*ListItem, lru.capacity)
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
