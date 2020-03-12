package undirected_weighted_graph

import (
	"fmt"
	"testing"
)

func TestAdjacencyList(t *testing.T) {
	al := New(5)
	al.AddEdge(0, 1, 10)
	al.AddEdge(0, 4, 20)
	al.AddEdge(1, 2, 30)
	al.AddEdge(1, 3, 40)
	al.AddEdge(1, 4, 50)
	al.AddEdge(2, 3, 60)
	al.AddEdge(3, 4, 70)
	fmt.Print(al)
	al.DFS()
}
