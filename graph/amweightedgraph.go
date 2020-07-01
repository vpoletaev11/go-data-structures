package graph

/*
	ADJACENCY MATRIX WEIGHTED GRAPH:

	Size()               AddEdge()            RemoveEdge()         PeekEdges()
	Best:    O(1)        Best:    O(1)        Best:    O(1)        Best:    O(n)
	Average: O(1)        Average: O(1)        Average: O(1)        Average: O(n)
	Worst:   O(1)        Worst:   O(1)        Worst:   O(1)        Worst:   O(n)


	goos: linux
	goarch: amd64
	BenchmarkAddEdgeAMWeightedGraph-8              	12128056	        99.4 ns/op	       0 B/op	       0 allocs/op
	BenchmarkRemoveEdgeAMWeightedGraph-8           	13610829	        88.2 ns/op	       0 B/op	       0 allocs/op
	BenchmarkPeekEdgesAMWeightedGraph-8            	  652176	      2053 ns/op	     688 B/op	      24 allocs/op
*/

// AMWeightedGraph - Weighted Graph based on Adjacency Matrix.
// Quantity of rows or columns of the matrix == quantity of graph vertexes.
// Matrix stores edges of graph vertexes.
// Indexing of vertexes starts from 0.
type AMWeightedGraph struct {
	adjacencyMatrix [][]int
}

// Size returns count of vertexes in the graph
func (a *AMWeightedGraph) Size() int {
	return len(a.adjacencyMatrix)
}

// AMWeightedGraphInit initializes Weighted Graph based on Adjacency Matrix
func AMWeightedGraphInit(size int) *AMWeightedGraph {
	if size < 1 {
		return nil
	}
	var a AMWeightedGraph

	a.adjacencyMatrix = make([][]int, size)
	for i := 0; i < len(a.adjacencyMatrix); i++ {
		a.adjacencyMatrix[i] = make([]int, size)
	}

	return &a
}

// AddEdge adds edge between vertexes
func (a *AMWeightedGraph) AddEdge(from, to, weight int) bool {
	if from > a.Size()-1 || to > a.Size()-1 || from < 0 || to < 0 {
		return false
	}

	a.adjacencyMatrix[from][to] = weight
	a.adjacencyMatrix[to][from] = weight
	return true
}

// RemoveEdge removes edge between vertexes
func (a *AMWeightedGraph) RemoveEdge(from, to int) bool {
	if from > a.Size()-1 || to > a.Size()-1 || from < 0 || to < 0 {
		return false
	}

	a.adjacencyMatrix[from][to] = 0
	a.adjacencyMatrix[to][from] = 0
	return true
}

// PeekEdges returns slice of pairs edges and edges weight of inputted vertex
func (a *AMWeightedGraph) PeekEdges(vertex int) ([][]int, bool) {
	if vertex > a.Size()-1 || vertex < 0 {
		return [][]int{}, false
	}

	edges := [][]int{}
	for i, val := range a.adjacencyMatrix[vertex] {
		if val != 0 {
			edges = append(edges, []int{i, val})
		}
	}
	return edges, true
}