package graph

/*
	ADJACENCY MATRIX WEIGHTED GRAPH:

	Size()               AddEdge()            RemoveEdge()         PeekEdges()
	Best:    O(1)        Best:    O(1)        Best:    O(1)        Best:    O(n)
	Average: O(1)        Average: O(1)        Average: O(1)        Average: O(n)
	Worst:   O(1)        Worst:   O(1)        Worst:   O(1)        Worst:   O(n)
*/

// AdjacencyMatrixWeightedGraph - Weighted Graph based on adjacency matrix.
// Quantity of rows or columns of the matrix == quantity of graph vertexes.
// Matrix stores edges of graph vertexes.
// Indexing of vertexes starts from 0.
type AdjacencyMatrixWeightedGraph struct {
	adjacencyMatrix [][]int
}

// Size returns count of vertexes in the graph
func (a *AdjacencyMatrixWeightedGraph) Size() int {
	return len(a.adjacencyMatrix)
}

// AdjMatrixWeightedGraphInit initializes AdjacencyMatrixDirectedGraph
func AdjMatrixWeightedGraphInit(size int) *AdjacencyMatrixWeightedGraph {
	if size < 1 {
		return nil
	}
	var a AdjacencyMatrixWeightedGraph

	a.adjacencyMatrix = make([][]int, size)
	for i := 0; i < len(a.adjacencyMatrix); i++ {
		a.adjacencyMatrix[i] = make([]int, size)
	}

	return &a
}

// AddEdge adds edge between vertexes
func (a *AdjacencyMatrixWeightedGraph) AddEdge(from, to, weight int) bool {
	if from > a.Size() || to > a.Size() || from < 0 || to < 0 {
		return false
	}

	a.adjacencyMatrix[from][to] = weight
	a.adjacencyMatrix[to][from] = weight
	return true
}

// RemoveEdge removes edge between vertexes
func (a *AdjacencyMatrixWeightedGraph) RemoveEdge(from, to int) bool {
	if from > a.Size() || to > a.Size() || from < 0 || to < 0 {
		return false
	}

	a.adjacencyMatrix[from][to] = 0
	a.adjacencyMatrix[to][from] = 0
	return true
}

// PeekEdges returns slice of pairs edges and edges weight of inputted vertex
func (a *AdjacencyMatrixWeightedGraph) PeekEdges(vertex int) (edges [][]int, ok bool) {
	if vertex > a.Size() || vertex < 0 {
		return [][]int{}, false
	}

	for i, val := range a.adjacencyMatrix[vertex] {
		if val != 0 {
			edges = append(edges, []int{i, val})
		}
	}
	return edges, true
}