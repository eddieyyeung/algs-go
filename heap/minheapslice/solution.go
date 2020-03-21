package minheapslice

import (
	"errors"
	"fmt"
	"math"
)

type MinHeap struct {
	Size int
	Heap []int
}

func New(cap int) *MinHeap {
	return &MinHeap{
		Size: 0,
		Heap: make([]int, cap),
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return i*2 + 1
}

func right(i int) int {
	return i*2 + 2
}

func (mh *MinHeap) Insert(k int) error {
	if len(mh.Heap) == mh.Size {
		return errors.New(fmt.Sprintf("Overflow, Could not insertKey: %v", k))
	}
	mh.Size++
	i := mh.Size - 1
	mh.Heap[i] = k
	for i != 0 && mh.Heap[parent(i)] > mh.Heap[i] {
		mh.Heap[i], mh.Heap[parent(i)] = mh.Heap[parent(i)], mh.Heap[i]
		i = parent(i)
	}
	return nil
}

func (mh *MinHeap) ModifyKey(i, k int) error {
	mh.Heap[i] = k
	for i != 0 && mh.Heap[parent(i)] > mh.Heap[i] {
		mh.Heap[i], mh.Heap[parent(i)] = mh.Heap[parent(i)], mh.Heap[i]
		i = parent(i)
	}
	return nil
}

func (mh *MinHeap) GetMin() int {
	return mh.Heap[0]
}

func (mh *MinHeap) MinHeapify(i int) {
	l := left(i)
	r := right(i)
	smallest := i
	if l < mh.Size && mh.Heap[l] < mh.Heap[smallest] {
		smallest = l
	}
	if r < mh.Size && mh.Heap[r] < mh.Heap[smallest] {
		smallest = r
	}
	if smallest != i {
		mh.Heap[i], mh.Heap[smallest] = mh.Heap[smallest], mh.Heap[i]
		mh.MinHeapify(smallest)
	}
}

func (mh *MinHeap) ExctraMin() (int, error) {
	if mh.Size <= 0 {
		return 0, errors.New("cannot ExtracMin")
	}
	if mh.Size == 1 {
		mh.Size--
		return mh.Heap[0], nil
	}
	root := mh.Heap[0]
	mh.Heap[0] = mh.Heap[mh.Size-1]
	mh.Size--
	mh.MinHeapify(0)
	return root, nil
}

func (mh *MinHeap) DeleteKey(i int) error {
	if err := mh.ModifyKey(i, math.MinInt64); err != nil {
		return err
	}
	if _, err := mh.ExctraMin(); err != nil {
		return err
	}
	return nil
}
