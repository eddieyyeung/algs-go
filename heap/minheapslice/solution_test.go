package minheapslice

import (
	"fmt"
	"testing"
)

func TestMinHeap(t *testing.T) {
	mh := New(11)
	mh.Insert(3)
	mh.Insert(2)
	mh.DeleteKey(1)
	mh.Insert(15)
	mh.Insert(5)
	mh.Insert(4)
	mh.Insert(45)
	extramin, _ := mh.ExctraMin()
	fmt.Println("extramin", extramin)
	min := mh.GetMin()
	fmt.Println("min", min)
	mh.ModifyKey(2, 1)
	min = mh.GetMin()
	fmt.Println("min", min)
}
