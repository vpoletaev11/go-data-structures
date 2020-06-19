package graph

/*
	ADJACENCY MATRIX GRAPH:

	Size()               AddEdge()            RemoveEdge()         PeekEdges()
	Best:    O(1)        Best:    O(1)        Best:    O(1)        Best:    O(n)
	Average: O(1)        Average: O(1)        Average: O(1)        Average: O(n)
	Worst:   O(1)        Worst:   O(1)        Worst:   O(1)        Worst:   O(n)


	goos: linux
	goarch: amd64
	BenchmarkAddEdgeAMGraph-8                        1000000               103 ns/op               0 B/op          0 allocs/op
	BenchmarkRemoveEdgeAMGraph-8                     1000000               105 ns/op               0 B/op          0 allocs/op
	BenchmarkPeekEdgesAMGraph-8                      1000000              1087 ns/op             208 B/op         17 allocs/op
*/

// AdjacencyMatrixGraph - Graph based on adjacency matrix.
// Quantity of rows or columns of the matrix == quantity of graph vertexes.
// Matrix stores edges of graph vertexes.
// Indexing of vertexes starts from 0.
type AdjacencyMatrixGraph struct {
	adjacencyMatrix [][]bool
}

// Size returns count of vertexes in the graph
func (a *AdjacencyMatrixGraph) Size() int {
	return len(a.adjacencyMatrix)
}

// AdjMatrixGraphInit initializes AdjacencyMatrixDirectedGraph
func AdjMatrixGraphInit(size int) *AdjacencyMatrixGraph {
	if size < 1 {
		return nil
	}
	var a AdjacencyMatrixGraph

	a.adjacencyMatrix = make([][]bool, size)
	for i := 0; i < len(a.adjacencyMatrix); i++ {
		a.adjacencyMatrix[i] = make([]bool, size)
	}

	return &a
}

// AddEdge adds edge between vertexes
func (a *AdjacencyMatrixGraph) AddEdge(from, to int) bool {
	if from > a.Size()-1 || to > a.Size()-1 || from < 0 || to < 0 {
		return false
	}

	a.adjacencyMatrix[from][to] = true
	a.adjacencyMatrix[to][from] = true
	return true
}

// RemoveEdge removes edge between vertexes
func (a *AdjacencyMatrixGraph) RemoveEdge(from, to int) bool {
	if from > a.Size()-1 || to > a.Size()-1 || from < 0 || to < 0 {
		return false
	}

	a.adjacencyMatrix[from][to] = false
	a.adjacencyMatrix[to][from] = false
	return true
}

// PeekEdges returns slice of edges of inputted vertex
func (a *AdjacencyMatrixGraph) PeekEdges(vertex int) (edges []int, ok bool) {
	if vertex > a.Size()-1 || vertex < 0 {
		return []int{}, false
	}

	for i, val := range a.adjacencyMatrix[vertex] {
		if val == true {
			edges = append(edges, i)
		}
	}
	return edges, true
}
