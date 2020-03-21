/*
Package undirected-weighted-graph 有加权和无向图

https://www.geeksforgeeks.org/graph-implementation-using-stl-for-competitive-programming-set-2-weighted-graph/
*/
package adjacency_list

import (
	"container/list"
	"fmt"
)

type Item struct {
	Value  int
	Weight int
}

// Node 邻接列表节点
type Node struct {
	List *list.List
}

// AdjacencyList 邻接列表
type AdjacencyList struct {
	Nodes []Node
}

// New 创建邻接列表
func New(n int) *AdjacencyList {
	al := AdjacencyList{
		make([]Node, n),
	}
	for i := range al.Nodes {
		al.Nodes[i].List = list.New()
	}
	return &al
}

// AddEdge 添加边
func (al *AdjacencyList) AddEdge(u, v, wt int) {
	al.Nodes[u].List.PushBack(Item{v, wt})
	al.Nodes[v].List.PushBack(Item{u, wt})
}

// DFS A utility function to do DFS of graph
// recursively from a AdjacencyList al.
// https://www.geeksforgeeks.org/graph-implementation-using-stl-for-competitive-programming-set-1-dfs-of-unweighted-and-undirected/
func (al *AdjacencyList) DFS() {
	visited := make([]bool, len(al.Nodes))
	for i, node := range al.Nodes {
		node.traverse(i, al, &visited)
	}
}

func (node *Node) traverse(idx int, al *AdjacencyList, visited *[]bool) {
	if (*visited)[idx] {
		return
	}
	(*visited)[idx] = true
	fmt.Printf("%v ", idx)
	for e := node.List.Front(); e != nil; e = e.Next() {
		item, _ := e.Value.(Item)
		al.Nodes[item.Value].traverse(item.Value, al, visited)
	}
}

func (al *AdjacencyList) String() string {
	str := ""
	for i := range al.Nodes {
		str += fmt.Sprintf("%v: ", i)
		for e := al.Nodes[i].List.Front(); e != nil; e = e.Next() {
			str += fmt.Sprintf("%v ", e.Value)
		}
		str += "\n"
	}
	return str
}
