package detect_circle_in_a_directed_graph

type Graph struct {
	V   int     // Number of vertices
	Adj [][]int // Adjacent List
}

func NewGraph(v int) *Graph {
	return &Graph{
		V:   v,
		Adj: make([][]int, v),
	}
}

func (graph *Graph) AddEdge(source, dest int) {
	graph.Adj[source] = append(graph.Adj[source], dest)
}

func (graph *Graph) isCyclicUtil(i int, visited []bool, recStack []bool) bool {
	if recStack[i] {
		return true
	}
	if visited[i] {
		return false
	}
	visited[i] = true
	recStack[i] = true
	for _, dest := range graph.Adj[i] {
		if graph.isCyclicUtil(dest, visited, recStack) {
			return true
		}
	}
	recStack[i] = false
	return false
}

func (graph *Graph) IsCyclic() bool {
	visited := make([]bool, graph.V)
	recStack := make([]bool, graph.V)
	for i := 0; i < graph.V; i++ {
		if graph.isCyclicUtil(i, visited, recStack) {
			return true
		}
	}
	return false
}
