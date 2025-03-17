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

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.cache[key]; ok {
		this.list.MoveToFront(elem)
		return elem.Value.(Pair).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if elem, ok := this.cache[key]; ok {
		//如果找到了，就把这个元素移动到最前面
		this.list.MoveToFront(elem)
		elem.Value = Pair{key: key, value: value}
	} else {
		//如果没有找到
		if this.list.Len() >= this.capacity {
			//如果超出了容量，就删除最后一个元素
			//先从cache中删除
			//再从list中删除
			delete(this.cache, this.list.Back().Value.(Pair).key)
			this.list.Remove(this.list.Back())
		}
		//然后把新元素插入到最前面
		//更新cache
		this.list.PushFront(Pair{key: key, value: value})
		this.cache[key] = this.list.Front()
	}
}