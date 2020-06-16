package graph

/*
	ADJACENCY MATRIX DIRECTED GRAPH:

	Size()               AddEdge()            RemoveEdge()         PeekEdges()
	Best:    O(1)        Best:    O(1)        Best:    O(1)        Best:    O(n)
	Average: O(1)        Average: O(1)        Average: O(1)        Average: O(n)
	Worst:   O(1)        Worst:   O(1)        Worst:   O(1)        Worst:   O(n)


	goos: linux
	goarch: amd64
	BenchmarkAddEdgeAMDirectedGraph-8                1000000               109 ns/op               0 B/op          0 allocs/op
	BenchmarkRemoveEdgeAMDirectedGraph-8             1000000                89.3 ns/op             0 B/op          0 allocs/op
	BenchmarkPeekEdgesAMDirectedGraph-8              1000000               625 ns/op              80 B/op          9 allocs/op
*/

// AdjacencyMatrixDirectedGraph - Directed Graph based on adjacency matrix.
// Quantity of rows or columns of the matrix == quantity of graph vertexes.
// Matrix stores edges of graph vertexes.
// Indexing of vertexes starts from 0.
type AdjacencyMatrixDirectedGraph struct {
	adjacencyMatrix [][]bool
}

// Size returns count of vertexes in the graph
func (a *AdjacencyMatrixDirectedGraph) Size() int {
	return len(a.adjacencyMatrix)
}

// AdjMatrixDirectedGraphInit initializes AdjacencyMatrixDirectedGraph
func AdjMatrixDirectedGraphInit(size int) *AdjacencyMatrixDirectedGraph {
	if size < 1 {
		return nil
	}
	var a AdjacencyMatrixDirectedGraph

	a.adjacencyMatrix = make([][]bool, size)
	for i := 0; i < len(a.adjacencyMatrix); i++ {
		a.adjacencyMatrix[i] = make([]bool, size)
	}

	return &a
}

// AddEdge adds edge between vertexes
func (a *AdjacencyMatrixDirectedGraph) AddEdge(from, to int) bool {
	if from > a.Size() || to > a.Size() || from < 0 || to < 0 {
		return false
	}

	a.adjacencyMatrix[from][to] = true
	return true
}

// RemoveEdge removes edge between vertexes
func (a *AdjacencyMatrixDirectedGraph) RemoveEdge(from, to int) bool {
	if from > a.Size() || to > a.Size() || from < 0 || to < 0 {
		return false
	}

	a.adjacencyMatrix[from][to] = false
	return true
}

// PeekEdges returns slice of edges of inputted vertex
func (a *AdjacencyMatrixDirectedGraph) PeekEdges(vertex int) (edges []int, ok bool) {
	if vertex > a.Size() || vertex < 0 {
		return []int{}, false
	}

	for i, val := range a.adjacencyMatrix[vertex] {
		if val == true {
			edges = append(edges, i)
		}
	}
	return edges, true
}
