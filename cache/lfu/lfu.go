package lfu

import (
	"container/heap"
	"fmt"
)

type LFUCache struct {
	capacity int
	cache    map[int]*entry
	freq     map[int]*freqList
	minFreq  int
}

type entry struct {
	key   int
	value int
	freq  int
	index int
}

type freqList struct {
	freq    int
	entries []*entry
}

type freqListHeap []*freqList

func (h freqListHeap) Len() int {
	return len(h)
}

func (h freqListHeap) Less(i, j int) bool {
	return h[i].freq < h[j].freq
}

func (h freqListHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].freq = i
	h[j].freq = j
}

func (h *freqListHeap) Push(x interface{}) {
	*h = append(*h, x.(*freqList))
}

func (h *freqListHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		cache:    make(map[int]*entry),
		freq:     make(map[int]*freqList),
		minFreq:  0,
	}
}

func (lfu *LFUCache) Get(key int) int {
	if entry, ok := lfu.cache[key]; ok {
		lfu.updateEntry(entry)
		return entry.value
	}
	return -1
}

func (lfu *LFUCache) Put(key int, value int) {
	if lfu.capacity == 0 {
		return
	}

	if entry, ok := lfu.cache[key]; ok {
		entry.value = value
		lfu.updateEntry(entry)
	} else {
		if len(lfu.cache) >= lfu.capacity {
			removedEntry := lfu.removeEntry()
			delete(lfu.cache, removedEntry.key)
		}

		newEntry := &entry{
			key:   key,
			value: value,
			freq:  1,
		}
		lfu.cache[key] = newEntry
		lfu.insertEntry(newEntry)
	}
}

func (lfu *LFUCache) updateEntry(e *entry) {
	f := lfu.freq[e.freq]
	f.removeEntry(e)

	if len(f.entries) == 0 && lfu.minFreq == e.freq {
		lfu.minFreq++
	}

	e.freq++
	newF, ok := lfu.freq[e.freq]
	if !ok {
		newF = &freqList{freq: e.freq}
		lfu.freq[e.freq] = newF
	}
	newF.insertEntry(e)
}

func (lfu *LFUCache) removeEntry() *entry {
	minFreqList := lfu.freq[lfu.minFreq]
	removedEntry := minFreqList.removeOldestEntry()

	if len(minFreqList.entries) == 0 {
		delete(lfu.freq, lfu.minFreq)
		lfu.minFreq++
	}

	return removedEntry
}

func (fl *freqList) insertEntry(e *entry) {
	e.index = len(fl.entries)
	fl.entries = append(fl.entries, e)
	heap.Fix(fl, e.index)
}

func (fl *freqList) removeEntry(e *entry) {
	heap.Remove(fl, e.index)
}

func (fl *freqList) removeOldestEntry() *entry {
	n := len(fl.entries)
	oldestEntry := fl.entries[n-1]
	fl.entries = fl.entries[0 : n-1]
	return oldestEntry
}

func main() {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // Output: 1

	cache.Put(3, 3)
	fmt.Println(cache.Get(2)) // Output: -1

	cache.Put(4, 4)
	fmt.Println(cache.Get(1)) // Output: -1
	fmt.Println(cache.Get(3)) // Output: 3
	fmt.Println(cache.Get(4)) // Output: 4
}
