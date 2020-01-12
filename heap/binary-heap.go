package heap

import (
	"fmt"
)

type MinHeap struct {
	Harr     []int
	Capacity int
	HeapSize int
}

func (mh MinHeap) New(cap int) {
	mh.HeapSize = 0
	mh.Capacity = cap
	mh.Harr = make([]int, cap)
}

func (mh MinHeap) parent(i int) int {
	return (i - 1) / 2
}

func (mh MinHeap) InsertKey(k int) error {
	if mh.HeapSize == mh.Capacity {
		return fmt.Errorf("Overflow: Cound not insertKey %d", k)
	}
	mh.HeapSize++
	i := mh.HeapSize - 1
	mh.Harr[i] = k
	for i != 0 && mh.Harr[mh.parent(i)] > mh.Harr[i] {
		mh.Harr[i], mh.Harr[mh.parent(i)] = mh.Harr[mh.parent(i)], mh.Harr[i]
		i = mh.parent(i)
	}
	return nil
}
