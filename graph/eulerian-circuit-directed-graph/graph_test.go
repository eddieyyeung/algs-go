// Euler Circuit in a Directed Graph
// https://www.geeksforgeeks.org/euler-circuit-directed-graph/
package euleriancircuitdirectedgraph

import (
	"fmt"
	"testing"
)

func TestNewGraph(t *testing.T) {
	g1 := NewGraph(5)
	g1.AddEdge(0, 2)
	g1.AddEdge(0, 3)
	g1.AddEdge(1, 0)
	g1.AddEdge(2, 1)
	g1.AddEdge(3, 4)
	g1.AddEdge(4, 3)
	ans := g1.FindSccs()
	for _, scc := range ans {
		fmt.Println(scc)
	}
}

func TestGraph_IsSc(t *testing.T) {
	g1 := NewGraph(5)
	g1.AddEdge(0, 1)
	g1.AddEdge(1, 2)
	g1.AddEdge(2, 3)
	g1.AddEdge(3, 0)
	g1.AddEdge(2, 4)
	g1.AddEdge(4, 2)
	fmt.Println("g1 isSc", g1.IsSc())

	g2 := NewGraph(4)
	g2.AddEdge(0, 1)
	g2.AddEdge(1, 2)
	g2.AddEdge(2, 3)
	fmt.Println("g2 isSc", g2.IsSc())
}
