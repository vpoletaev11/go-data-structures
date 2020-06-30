package graph

/*
	ADJACENCY LIST DIRECTED WEIGHTED GRAPH:

	Size()               AddEge()             RemoveEdge()         PeekEdges()
	Best:    O(1)        Best:    O(1)        Best:    O(1)        Best:    O(e) e - count of edges
	Average: O(1)        Average: O(1)        Average: O(1)        Average: O(e)
	Worst:   O(1)        Worst:   O(1)        Worst:   O(1)        Worst:   O(e)
*/

// ALDirectedWeightedGraph - Directed Weighted Graph based on Adjacency List.
// Indexing of vertexes starts from 0.
type ALDirectedWeightedGraph struct {
	adjacencyList []map[int]int
}

// Size returns count of vertexes in the graph
func (a *ALDirectedWeightedGraph) Size() int {
	return len(a.adjacencyList)
}

// ALDirectedWeightedGraphInit initializes Directed Weighted Graph based on Adjacency List
func ALDirectedWeightedGraphInit(size int) *ALDirectedWeightedGraph {
	if size < 1 {
		return nil
	}
	var a ALDirectedWeightedGraph

	a.adjacencyList = make([]map[int]int, size)
	for i := 0; i < len(a.adjacencyList); i++ {
		a.adjacencyList[i] = make(map[int]int)
	}

	return &a
}

// AddEdge adds edge between vertexes
func (a *ALDirectedWeightedGraph) AddEdge(from, to, weight int) bool {
	if from > a.Size()-1 || to > a.Size()-1 || from < 0 || to < 0 {
		return false
	}

	a.adjacencyList[from][to] = weight
	return true
}

// RemoveEdge removes edge between vertexes
func (a *ALDirectedWeightedGraph) RemoveEdge(from, to int) bool {
	if from > a.Size()-1 || to > a.Size()-1 || from < 0 || to < 0 {
		return false
	}

	delete(a.adjacencyList[from], to)
	return true
}

// PeekEdges returns slice of pairs edges and edges weight of inputted vertex
func (a *ALDirectedWeightedGraph) PeekEdges(vertex int) ([][]int, bool) {
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
