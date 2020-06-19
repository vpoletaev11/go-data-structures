package graph

/*
	ADJACENCY MATRIX DIRECTED WEIGHTED GRAPH:

	Size()               AddEdge()            RemoveEdge()         PeekEdges()
	Best:    O(1)        Best:    O(1)        Best:    O(1)        Best:    O(n)
	Average: O(1)        Average: O(1)        Average: O(1)        Average: O(n)
	Worst:   O(1)        Worst:   O(1)        Worst:   O(1)        Worst:   O(n)


	goos: linux
	goarch: amd64
	BenchmarkAddEdgeAMDirectedWeightedGraph-8        1000000                92.6 ns/op             0 B/op          0 allocs/op
	BenchmarkRemoveEdgeAMDirectedWeightedGraph-8     1000000                87.3 ns/op             0 B/op          0 allocs/op
	BenchmarkPeekEdgesAMDirectedWeightedGraph-8      1000000              1527 ns/op             448 B/op         18 allocs/op
*/

// AdjacencyMatrixDirectedWeightedGraph - Directed Weighted Graph based on adjacency matrix.
// Quantity of rows or columns of the matrix == quantity of graph vertexes.
// Matrix stores weighted edges of graph vertexes.
// Indexing of vertexes starts from 0.
type AdjacencyMatrixDirectedWeightedGraph struct {
	adjacencyMatrix [][]int
}

// Size returns count of vertexes in the graph
func (a *AdjacencyMatrixDirectedWeightedGraph) Size() int {
	return len(a.adjacencyMatrix)
}

// AdjMatrixDirectedWeightedGraphInit initializes AdjacencyMatrixDirectedWeightedGraph
func AdjMatrixDirectedWeightedGraphInit(size int) *AdjacencyMatrixDirectedWeightedGraph {
	if size < 1 {
		return nil
	}
	var a AdjacencyMatrixDirectedWeightedGraph

	a.adjacencyMatrix = make([][]int, size)
	for i := 0; i < len(a.adjacencyMatrix); i++ {
		a.adjacencyMatrix[i] = make([]int, size)
	}

	return &a
}

// AddEdge adds edge between vertexes
func (a *AdjacencyMatrixDirectedWeightedGraph) AddEdge(from, to, weight int) bool {
	if from > a.Size()-1 || to > a.Size()-1 || from < 0 || to < 0 {
		return false
	}

	a.adjacencyMatrix[from][to] = weight
	return true
}

// RemoveEdge removes edge between vertexes
func (a *AdjacencyMatrixDirectedWeightedGraph) RemoveEdge(from, to int) bool {
	if from > a.Size()-1 || to > a.Size()-1 || from < 0 || to < 0 {
		return false
	}

	a.adjacencyMatrix[from][to] = 0
	return true
}

// PeekEdges returns slice of pairs edges and edges weight of inputted vertex
func (a *AdjacencyMatrixDirectedWeightedGraph) PeekEdges(vertex int) (edges [][]int, ok bool) {
	if vertex > a.Size()-1 || vertex < 0 {
		return [][]int{}, false
	}

	for i, val := range a.adjacencyMatrix[vertex] {
		if val != 0 {
			edges = append(edges, []int{i, val})
		}
	}
	return edges, true
}
