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

// AMDirectedGraph - Directed Graph based on Adjacency Matrix.
// Quantity of rows and columns of the matrix == quantity of graph vertexes.
// Matrix stores edges of graph vertexes.
// Indexing of vertexes starts from 0.
type AMDirectedGraph struct {
	adjacencyMatrix [][]bool
}

// Size returns count of vertexes in the graph
func (a *AMDirectedGraph) Size() int {
	return len(a.adjacencyMatrix)
}

// AMDirectedGraphInit initializes Directed Graph based on Adjacency Matrix
func AMDirectedGraphInit(size int) *AMDirectedGraph {
	if size < 1 {
		return nil
	}
	var a AMDirectedGraph

	a.adjacencyMatrix = make([][]bool, size)
	for i := 0; i < len(a.adjacencyMatrix); i++ {
		a.adjacencyMatrix[i] = make([]bool, size)
	}

	return &a
}

// AddEdge adds edge between vertexes
func (a *AMDirectedGraph) AddEdge(from, to int) bool {
	if from > a.Size()-1 || to > a.Size()-1 || from < 0 || to < 0 {
		return false
	}

	a.adjacencyMatrix[from][to] = true
	return true
}

// RemoveEdge removes edge between vertexes
func (a *AMDirectedGraph) RemoveEdge(from, to int) bool {
	if from > a.Size()-1 || to > a.Size()-1 || from < 0 || to < 0 {
		return false
	}

	a.adjacencyMatrix[from][to] = false
	return true
}

// PeekEdges returns slice of edges of inputted vertex
func (a *AMDirectedGraph) PeekEdges(vertex int) (edges []int, ok bool) {
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
