package graph

// AdjacencyMatrixGraph - Graph based on adjacency matrix.
// Quantity of rows or columns of the matrix == quantity of graph vertexes.
// Matrix stores edges of graph vertexes.
// Indexing of vertexes starts from 0.
type AdjacencyMatrixGraph struct {
	adjacencyMatrix [][]int
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

	a.adjacencyMatrix = make([][]int, size)
	for i := 0; i < len(a.adjacencyMatrix); i++ {
		a.adjacencyMatrix[i] = make([]int, size)
	}

	return &a
}

// AddEdge adds edge between vertexes
func (a *AdjacencyMatrixGraph) AddEdge(from, to int) bool {
	if from > a.Size() || to > a.Size() || from < 0 || to < 0 {
		return false
	}

	a.adjacencyMatrix[from][to] = 1
	a.adjacencyMatrix[to][from] = 1
	return true
}

// RemoveEdge removes edge between vertexes
func (a *AdjacencyMatrixGraph) RemoveEdge(from, to int) bool {
	if from > a.Size() || to > a.Size() || from < 0 || to < 0 {
		return false
	}

	a.adjacencyMatrix[from][to] = 0
	a.adjacencyMatrix[to][from] = 0
	return true
}

// PeekEdges returns slice of edges of inputted vertex
func (a *AdjacencyMatrixGraph) PeekEdges(vertex int) (edges []int, ok bool) {
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
