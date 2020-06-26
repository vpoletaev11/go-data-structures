package graph

// ALWeightedGraph - Directed Weighted Graph based on Adjacency List.
// Indexing of vertexes starts from 0.
type ALWeightedGraph struct {
	adjacencyList []map[int]int
}

// Size returns count of vertexes in the graph
func (a *ALWeightedGraph) Size() int {
	return len(a.adjacencyList)
}

// ALWeightedGraphInit initializes Weighted Graph based on Adjacency List
func ALWeightedGraphInit(size int) *ALWeightedGraph {
	if size < 1 {
		return nil
	}
	var a ALWeightedGraph

	a.adjacencyList = make([]map[int]int, size)
	for i := 0; i < len(a.adjacencyList); i++ {
		a.adjacencyList[i] = make(map[int]int)
	}

	return &a
}

// AddEdge adds edge between vertexes
func (a *ALWeightedGraph) AddEdge(from, to, weight int) bool {
	if from > a.Size()-1 || to > a.Size()-1 || from < 0 || to < 0 {
		return false
	}

	a.adjacencyList[from][to] = weight
	a.adjacencyList[to][from] = weight
	return true
}

// RemoveEdge removes edge between vertexes
func (a *ALWeightedGraph) RemoveEdge(from, to int) bool {
	if from > a.Size()-1 || to > a.Size()-1 || from < 0 || to < 0 {
		return false
	}

	delete(a.adjacencyList[from], to)
	delete(a.adjacencyList[to], from)
	return true
}

// PeekEdges returns slice of pairs edges and edges weight of inputted vertex
func (a *ALWeightedGraph) PeekEdges(vertex int) ([][]int, bool) {
	if vertex > a.Size()-1 || vertex < 0 {
		return [][]int{}, false
	}

	edges := [][]int{}
	for key, val := range a.adjacencyList[vertex] {
		if val != 0 {
			edges = append(edges, []int{key, val})
		}
	}
	return edges, true
}
