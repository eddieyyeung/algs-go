/*
Package undirectedgraph 无向图

https://www.geeksforgeeks.org/graph-and-its-representations/
*/
package undirectedgraph

import (
	"container/list"
	"fmt"
)

// Node 邻接列表节点
type Node struct {
	l *list.List
}

// AdjacencyList 邻接列表
type AdjacencyList struct {
	nodes []Node
}

// New 创建邻接列表
func New(n int) *AdjacencyList {
	al := AdjacencyList{
		make([]Node, n),
	}
	for i := range al.nodes {
		al.nodes[i].l = list.New()
	}
	return &al
}

// AddEdge 添加边
func (al *AdjacencyList) AddEdge(u, v int) {
	al.nodes[u].l.PushBack(v)
	al.nodes[v].l.PushBack(u)
}

// DFS A utility function to do DFS of graph
// recursively from a AdjacencyList al.
// https://www.geeksforgeeks.org/graph-implementation-using-stl-for-competitive-programming-set-1-dfs-of-unweighted-and-undirected/
func (al *AdjacencyList) DFS() {
	visited := make([]bool, len(al.nodes))
	for i, node := range al.nodes {
		node.traverse(i, al, &visited)
	}
}

func (node *Node) traverse(idx int, al *AdjacencyList, visited *[]bool) {
	if (*visited)[idx] {
		return
	}
	(*visited)[idx] = true
	fmt.Printf("%v ", idx)
	for e := node.l.Front(); e != nil; e = e.Next() {
		v, _ := e.Value.(int)
		al.nodes[v].traverse(v, al, visited)
	}
}

func (al *AdjacencyList) String() string {
	str := ""
	for i := range al.nodes {
		str += fmt.Sprintf("%v: ", i)
		for e := al.nodes[i].l.Front(); e != nil; e = e.Next() {
			str += fmt.Sprintf("%v ", e.Value)
		}
		str += "\n"
	}
	return str
}
