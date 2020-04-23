package graph

// AdjacencyMatrixDirectedGraph - Directed Graph based on adjacency matrix.
// Quantity of rows or columns of the matrix == quantity of graph vertexes.
// Matrix stores edges of graph vertexes.
// Indexing of vertexes starts from 0.
type AdjacencyMatrixDirectedGraph struct {
	adjacencyMatrix [][]int
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

	a.adjacencyMatrix = make([][]int, size)
	for i := 0; i < len(a.adjacencyMatrix); i++ {
		a.adjacencyMatrix[i] = make([]int, size)
	}

	return &a
}

// AddEdge adds edge between vertexes
func (a *AdjacencyMatrixDirectedGraph) AddEdge(from, to int) bool {
	if from > a.Size() || to > a.Size() || from < 0 || to < 0 {
		return false
	}

	a.adjacencyMatrix[from][to] = 1
	return true
}

// RemoveEdge removes edge between vertexes
func (a *AdjacencyMatrixDirectedGraph) RemoveEdge(from, to int) bool {
	if from > a.Size() || to > a.Size() || from < 0 || to < 0 {
		return false
	}

	a.adjacencyMatrix[from][to] = 0
	return true
}

// PeekEdges returns slice of edges of inputted vertex
func (a *AdjacencyMatrixDirectedGraph) PeekEdges(vertex int) (edges []int, ok bool) {
	if vertex > a.Size() || vertex < 0 {
		return []int{}, false
	}

	for i, val := range a.adjacencyMatrix[vertex] {
		if val != 0 {
			edges = append(edges, i)
		}
	}
	return edges, true
}