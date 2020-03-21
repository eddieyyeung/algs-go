package adjacency_list

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	al := New(5)
	al.AddEdge(0, 1)
	al.AddEdge(0, 4)
	al.AddEdge(1, 2)
	al.AddEdge(1, 3)
	al.AddEdge(1, 4)
	al.AddEdge(2, 3)
	al.AddEdge(3, 4)
	fmt.Print(al)
	al.DFS()
}
