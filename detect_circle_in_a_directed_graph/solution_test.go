package detect_circle_in_a_directed_graph

import (
	"testing"
)

func TestGraph(t *testing.T) {
	graph := NewGraph(4)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 0)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 3)
	if got := graph.IsCyclic(); got != true {
		t.Errorf("not correct")
	}
}
