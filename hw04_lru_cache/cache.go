package hw04_lru_cache //nolint:golint,stylecheck
import "fmt"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool // Добавить значение в кэш по ключу.
	Get(key Key) (interface{}, bool)     // Получить значение из кэша по ключу.
	Clear()                              // Очистить кэш.
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

/*
iteams[key] -> *ListItem
queue.Front() -> *ListItem
c.queue.PushFront(interface) -> *ListItem

*/

func (c *lruCache) Set(key Key, value interface{}) bool {
	newCacheItem := new(cacheItem)
	newCacheItem.key = string(key)
	newCacheItem.value = value

	if _, ok := c.items[key]; ok {
		fmt.Println(c.items[key].Next, c.items[key].Prev)
		c.queue.MoveToFront(c.items[key])
		c.queue.Front().Value = value
		c.items[key] = c.queue.Front()
		fmt.Println(c.items[key].Next, c.items[key].Prev)
		return true
	} else {
		if c.capacity <= c.queue.Len() {
			if cacheItem, ok := c.queue.Back().Value.(*cacheItem); ok {
				delete(c.items, Key(cacheItem.key))
				c.queue.Remove(c.queue.Back())
			} else {
				c.Clear()
			}
		}
		//iCacheItem := new(ListItem)
		//iCacheItem.Value = newCacheItem
		c.items[key] = c.queue.PushFront(newCacheItem)
		/*
			fmt.Println("And thats in map", c.items[key])

			for i, _ := range c.items {

				fmt.Println(c.items[i])
			}
			fmt.Println(len(c.items))

		*/
		return false
	}
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if _, ok := c.items[key]; ok {
		c.queue.MoveToFront(c.items[key])
		c.items[key] = c.queue.Front()
		if cacheItem, ok := c.queue.Front().Value.(*cacheItem); ok {
			return cacheItem.value, true
		} else {
			fmt.Println("bad get")
			return c.items[key].Value, true
		}
	} else {
		return nil, false
	}
}

func (c *lruCache) Clear() {
	if c.queue.Back() != nil {
		//cleanMap := make(map[Key]*ListItem, c.capacity)
		//var newCacheItem *cacheItem
		//c.items = cleanMap
		/*if cacheItem, ok := c.queue.Back().Value.(cacheItem); ok {
			delete(c.items, Key(cacheItem.key))
		} else {
			return item.Value, true
		}
		*/
		for i, _ := range c.items {
			c.queue.Remove(c.items[i].Value.(*ListItem))
			delete(c.items, i)
		}
	}
	c.queue.Remove(c.queue.Back())
}

type cacheItem struct {
	key   string
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
