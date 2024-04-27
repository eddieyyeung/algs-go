package cache

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type Pair struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.cache[key]; ok {
		this.list.MoveToFront(elem)
		return elem.Value.(*Pair).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if elem, ok := this.cache[key]; ok {
		this.list.MoveToFront(elem)
		elem.Value.(*Pair).value = value
	} else {
		if len(this.cache) == this.capacity {
			delete(this.cache, this.list.Back().Value.(*Pair).key)
			this.list.Remove(this.list.Back())
		}
		newPair := &Pair{key, value}
		newElem := this.list.PushFront(newPair)
		this.cache[key] = newElem
	}
}

func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1)) // Output: 1
	lru.Put(3, 3)
	fmt.Println(lru.Get(2)) // Output: -1
	lru.Put(4, 4)
	fmt.Println(lru.Get(1)) // Output: -1
	fmt.Println(lru.Get(3)) // Output: 3
	fmt.Println(lru.Get(4)) // Output: 4
}
