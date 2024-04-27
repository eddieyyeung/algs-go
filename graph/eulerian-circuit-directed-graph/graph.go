// Euler Circuit in a Directed Graph
// https://www.geeksforgeeks.org/euler-circuit-directed-graph/
package euleriancircuitdirectedgraph

// How to check if a directed graph is eulerian?
// A directed graph has an eulerian cycle if following conditions are true
// - All vertices with nonzero degree belong to a single strongly connected component.
//   - A strongly connected component is the component of a directed graph that has a path from every vertex to every
//     other vertex in that component. It can only be used in a directed graph.
// - In degree is equal to the out degree for every vertex.

// Time complexity: O(n * (n + m)), because for each pair of vertices we are checking whether a path exists between them.
// Auxiliary Space: O(N)

type Graph struct {
	Size  int
	Edges [][]int
}

func NewGraph(size int) *Graph {
	return &Graph{
		Size:  size,
		Edges: make([][]int, size),
	}
}

func (g *Graph) AddEdge(src, dest int) *Graph {
	g.Edges[src] = append(g.Edges[src], dest)
	return g
}

func (g *Graph) IsPath(src, dest int) bool {
	visited := make([]int, g.Size)
	var dfs func(cur, dest int, visited []int) bool
	dfs = func(cur, dest int, visited []int) bool {
		if cur == dest {
			return true
		}
		visited[cur] = 1
		for _, next := range g.Edges[cur] {
			if visited[next] == 0 {
				if dfs(next, dest, visited) {
					return true
				}
			}
		}
		return false
	}
	return dfs(src, dest, visited)
}

// FindScc return all the strongly connected components of a graph
func (g *Graph) FindSccs() [][]int {
	var ans [][]int
	is_scc := make([]int, g.Size+1)
	for i := 0; i < g.Size; i++ {
		scc := []int{i}
		if is_scc[i] == 0 {
			for j := i + 1; j < g.Size; j++ {
				if is_scc[j] == 0 && g.IsPath(i, j) && g.IsPath(j, i) {
					is_scc[j] = 1
					scc = append(scc, j)
				}
			}

			ans = append(ans, scc)
		}
	}
	return ans
}

func (g *Graph) GetTranspose() *Graph {
	newG := NewGraph(g.Size)
	for u, vs := range g.Edges {
		for _, v := range vs {
			newG.AddEdge(v, u)
		}
	}
	return newG
}

func (g *Graph) DFSUtil(v int, visited []bool) {
	visited[v] = true
	for _, i := range g.Edges[v] {
		if !visited[i] {
			g.DFSUtil(i, visited)
		}
	}
}

func (g *Graph) IsSc() bool {
	visited := make([]bool, g.Size)

	g.DFSUtil(0, visited)

	for _, v := range visited {
		if !v {
			return false
		}
	}

	gr := g.GetTranspose()
	visited = make([]bool, gr.Size)
	gr.DFSUtil(0, visited)
	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}
