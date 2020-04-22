package graph

type AdjacencyMatrixGraph struct {
	adjacencyMatrix [][]int
}

func (a AdjacencyMatrixGraph) Size() int {
	return len(a.adjacencyMatrix)
}

func AMGraphInit(size int) AdjacencyMatrixGraph {
	var a AdjacencyMatrixGraph

	a.adjacencyMatrix = make([][]int, size)
	for i := 0; i < len(a.adjacencyMatrix); i++ {
		a.adjacencyMatrix[i] = make([]int, size)
	}

	return a
}
