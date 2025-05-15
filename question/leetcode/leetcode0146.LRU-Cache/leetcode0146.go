package leetcode

import "container/list"

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}
type Pair struct {
	key   int
	value int
}

func Consturctor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

// O(1)
func (lru *LRUCache) Get(key int) int {
	if elem, ok := lru.cache[key]; ok {
		lru.list.MoveToFront(elem)
		return elem.Value.(Pair).value
	}
	return -1
}

// O(1)
func (lru *LRUCache) Put(key int, value int) {
	if elem, ok := lru.cache[key]; ok {
		//如果找到了，就把这个元素移动到最前面
		lru.list.MoveToFront(elem)
		elem.Value = Pair{key: key, value: value}
	} else {
		//如果没有找到
		if lru.list.Len() >= lru.capacity {
			//如果超出了容量，就删除最后一个元素
			//先从cache中删除
			//再从list中删除
			delete(lru.cache, lru.list.Back().Value.(Pair).key)
			lru.list.Remove(lru.list.Back())
		}
		//然后把新元素插入到最前面
		//更新cache
		lru.list.PushFront(Pair{key: key, value: value})
		lru.cache[key] = lru.list.Front()
	}
}

// Remove removes a key from the cache if it exists.
// Returns true if the key was found and removed, false otherwise.
// O(1)
func (lru *LRUCache) Remove(key int) bool {
	if elem, ok := lru.cache[key]; ok {
		// Remove from the linked list
		lru.list.Remove(elem)
		// Remove from the cache map
		delete(lru.cache, key)
		return true
	}
	return false
}
