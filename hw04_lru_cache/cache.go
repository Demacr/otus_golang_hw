package hw04_lru_cache //nolint:golint,stylecheck

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool // Добавить значение в кэш по ключу
	Get(key Key) (interface{}, bool)     // Получить значение из кэша по ключу
	Clear()                              // Очистить кэш
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]cacheItem
	mutex    sync.Mutex
}

func (lc *lruCache) Set(key Key, value interface{}) (was bool) {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()

	item, was := lc.items[key]
	if was {
		lc.queue.MoveToFront(item)
		lc.queue.Front().Value = queueItem{value, key}
	} else {
		lc.items[key] = lc.queue.PushFront(queueItem{value, key})
		if lc.queue.Len() >= lc.capacity {
			delete(lc.items, lc.queue.Back().Value.(queueItem).key)
			lc.queue.Remove(lc.queue.Back())
		}
	}
	return
}

func (lc *lruCache) Get(key Key) (value interface{}, ok bool) {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()

	item, ok := lc.items[key]
	if ok {
		value = item.Value.(queueItem).value
		lc.queue.MoveToFront(item)
	}
	return
}

func (lc *lruCache) Clear() {
	lc.queue = NewList()
	lc.items = map[Key]cacheItem{}
}

type cacheItem *listItem
type queueItem struct {
	value interface{}
	key   Key
}

func NewCache(capacity int) Cache {
	return &lruCache{capacity: capacity, queue: NewList(), items: map[Key]cacheItem{}}
}
