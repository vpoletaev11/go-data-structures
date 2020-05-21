package graph_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/go-data-structures/graph"
)

// Testing AdjacencyMatrixDirectedGraph:
func TestAdjacencyMatrixDirectedGraphSUCCESS(t *testing.T) {
	a := graph.AdjMatrixDirectedGraphInit(5)

	dataset := []fromTo{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}, {1, 3}, {4, 2}, {3, 1}, {2, 0}}

	assert.Equal(t, 5, a.Size())

	for _, value := range dataset {
		a.AddEdge(value.from, value.to)
	}

	edges, ok := a.PeekEdges(0)
	assert.Equal(t, []int{0, 1, 2, 3, 4}, edges)
	assert.True(t, ok)

	edges, ok = a.PeekEdges(1)
	assert.Equal(t, []int{3}, edges)
	assert.True(t, ok)

	edges, ok = a.PeekEdges(2)
	assert.Equal(t, []int{0}, edges)
	assert.True(t, ok)

	edges, ok = a.PeekEdges(3)
	assert.Equal(t, []int{1}, edges)
	assert.True(t, ok)

	edges, ok = a.PeekEdges(4)
	assert.Equal(t, []int{2}, edges)
	assert.True(t, ok)

	// remove all edges
	for _, value := range dataset {
		a.RemoveEdge(value.from, value.to)
	}

	// check that graph haven't edges
	for i := 0; i < a.Size(); i++ {
		edges, ok = a.PeekEdges(i)
		assert.Equal(t, []int(nil), edges)
		assert.True(t, ok)
	}
}

func TestAdjacencyMatrixDirectedGraphOutOfRange(t *testing.T) {
	a := graph.AdjMatrixDirectedGraphInit(5)

	// AdjMatrixDirectedGraphInit out of range
	assert.Nil(t, graph.AdjMatrixDirectedGraphInit(0))

	// AddEdge out of range
	assert.False(t, a.AddEdge(-1, 0))
	assert.False(t, a.AddEdge(0, -1))
	assert.False(t, a.AddEdge(10, 0))
	assert.False(t, a.AddEdge(0, 10))

	// RemoveEdge out of range
	assert.False(t, a.RemoveEdge(-1, 0))
	assert.False(t, a.RemoveEdge(0, -1))
	assert.False(t, a.RemoveEdge(10, 0))
	assert.False(t, a.RemoveEdge(0, 10))

	// PeekEdges out of range
	edges, ok := a.PeekEdges(-1)
	assert.Equal(t, []int{}, edges)
	assert.False(t, ok)

	edges, ok = a.PeekEdges(10)
	assert.Equal(t, []int{}, edges)
	assert.False(t, ok)
}

// Testing AdjacencyMatrixDirectedWeightedGraph:
func TestAdjacencyMatrixDirectedWeightedGraphSUCCESS(t *testing.T) {
	a := graph.AdjMatrixDirectedWeightedGraphInit(5)

	dataset := []fromToWeight{{0, 0, 0}, {0, 1, 100}, {0, 2, 13}, {0, 3, 1000}, {0, 4, 20}, {1, 3, 1}, {4, 2, 8}, {3, 1, 14}, {2, 0, 75}}

	assert.Equal(t, 5, a.Size())

	for _, value := range dataset {
		a.AddEdge(value.from, value.to, value.weight)
	}

	edges, ok := a.PeekEdges(0)
	assert.Equal(t, [][]int{{1, 100}, {2, 13}, {3, 1000}, {4, 20}}, edges)
	assert.True(t, ok)

	edges, ok = a.PeekEdges(1)
	assert.Equal(t, [][]int{{3, 1}}, edges)
	assert.True(t, ok)

	edges, ok = a.PeekEdges(2)
	assert.Equal(t, [][]int{{0, 75}}, edges)
	assert.True(t, ok)

	edges, ok = a.PeekEdges(3)
	assert.Equal(t, [][]int{{1, 14}}, edges)
	assert.True(t, ok)

	edges, ok = a.PeekEdges(4)
	assert.Equal(t, [][]int{{2, 8}}, edges)
	assert.True(t, ok)

	// remove all edges
	for _, value := range dataset {
		a.RemoveEdge(value.from, value.to)
	}

	// check that graph haven't edges
	for i := 0; i < a.Size(); i++ {
		edges, ok = a.PeekEdges(i)
		assert.Equal(t, [][]int(nil), edges)
		assert.True(t, ok)
	}
}

func TestAdjacencyMatrixDirectedWeightedGraphOutOfRange(t *testing.T) {
	a := graph.AdjMatrixDirectedWeightedGraphInit(5)

	// AdjMatrixDirectedWeightedGraphInit out of range
	assert.Nil(t, graph.AdjMatrixDirectedWeightedGraphInit(0))

	// AddEdge out of range
	assert.False(t, a.AddEdge(-1, 0, 0))
	assert.False(t, a.AddEdge(0, -1, 0))
	assert.False(t, a.AddEdge(10, 0, 0))
	assert.False(t, a.AddEdge(0, 10, 0))

	// RemoveEdge out of range
	assert.False(t, a.RemoveEdge(-1, 0))
	assert.False(t, a.RemoveEdge(0, -1))
	assert.False(t, a.RemoveEdge(10, 0))
	assert.False(t, a.RemoveEdge(0, 10))

	// PeekEdges out of range
	edges, ok := a.PeekEdges(-1)
	assert.Equal(t, [][]int{}, edges)
	assert.False(t, ok)

	edges, ok = a.PeekEdges(10)
	assert.Equal(t, [][]int{}, edges)
	assert.False(t, ok)
}

// Testing AdjacencyMatrixGraph:
func TestAdjacencyMatrixGraphSUCCESS(t *testing.T) {
	a := graph.AdjMatrixGraphInit(5)

	dataset := []fromTo{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}, {1, 3}, {4, 2}, {3, 1}, {2, 0}}

	assert.Equal(t, 5, a.Size())

	for _, value := range dataset {
		a.AddEdge(value.from, value.to)
	}

	edges, ok := a.PeekEdges(0)
	assert.Equal(t, []int{0, 1, 2, 3, 4}, edges)
	assert.True(t, ok)

	edges, ok = a.PeekEdges(1)
	assert.Equal(t, []int{0, 3}, edges)
	assert.True(t, ok)

	edges, ok = a.PeekEdges(2)
	assert.Equal(t, []int{0, 4}, edges)
	assert.True(t, ok)

	edges, ok = a.PeekEdges(3)
	assert.Equal(t, []int{0, 1}, edges)
	assert.True(t, ok)

	edges, ok = a.PeekEdges(4)
	assert.Equal(t, []int{0, 2}, edges)
	assert.True(t, ok)

	// remove all edges
	for _, value := range dataset {
		a.RemoveEdge(value.from, value.to)
	}

	// check that graph haven't edges
	for i := 0; i < a.Size(); i++ {
		edges, ok = a.PeekEdges(i)
		assert.Equal(t, []int(nil), edges)
		assert.True(t, ok)
	}
}

func TestAdjacencyMatrixGraphOutOfRange(t *testing.T) {
	a := graph.AdjMatrixGraphInit(5)

	// AdjMatrixDirectedGraphInit out of range
	assert.Nil(t, graph.AdjMatrixGraphInit(0))

	// AddEdge out of range
	assert.False(t, a.AddEdge(-1, 0))
	assert.False(t, a.AddEdge(0, -1))
	assert.False(t, a.AddEdge(10, 0))
	assert.False(t, a.AddEdge(0, 10))

	// RemoveEdge out of range
	assert.False(t, a.RemoveEdge(-1, 0))
	assert.False(t, a.RemoveEdge(0, -1))
	assert.False(t, a.RemoveEdge(10, 0))
	assert.False(t, a.RemoveEdge(0, 10))

	// PeekEdges out of range
	edges, ok := a.PeekEdges(-1)
	assert.Equal(t, []int{}, edges)
	assert.False(t, ok)

	edges, ok = a.PeekEdges(10)
	assert.Equal(t, []int{}, edges)
	assert.False(t, ok)
}

// Testing AdjacencyMatrixWeightedGraph:
func TestAdjacencyMatrixWeightedGraphSUCCESS(t *testing.T) {
	a := graph.AdjMatrixWeightedGraphInit(5)

	dataset := []fromToWeight{{0, 0, 0}, {0, 1, 100}, {0, 2, 13}, {0, 3, 1000}, {0, 4, 20}, {1, 3, 1}, {4, 2, 8}, {3, 1, 14}, {2, 0, 75}}

	assert.Equal(t, 5, a.Size())

	for _, value := range dataset {
		a.AddEdge(value.from, value.to, value.weight)
	}

	edges, ok := a.PeekEdges(0)
	assert.Equal(t, [][]int{{1, 100}, {2, 75}, {3, 1000}, {4, 20}}, edges)
	assert.True(t, ok)

	edges, ok = a.PeekEdges(1)
	assert.Equal(t, [][]int{{0, 100}, {3, 14}}, edges)
	assert.True(t, ok)

	edges, ok = a.PeekEdges(2)
	assert.Equal(t, [][]int{{0, 75}, {4, 8}}, edges)
	assert.True(t, ok)

	edges, ok = a.PeekEdges(3)
	assert.Equal(t, [][]int{{0, 1000}, {1, 14}}, edges)
	assert.True(t, ok)

	edges, ok = a.PeekEdges(4)
	assert.Equal(t, [][]int{{0, 20}, {2, 8}}, edges)
	assert.True(t, ok)

	// remove all edges
	for _, value := range dataset {
		a.RemoveEdge(value.from, value.to)
	}

	// check that graph haven't edges
	for i := 0; i < a.Size(); i++ {
		edges, ok = a.PeekEdges(i)
		assert.Equal(t, [][]int(nil), edges)
		assert.True(t, ok)
	}
}

func TestAdjacencyMatrixWeightedGraphOutOfRange(t *testing.T) {
	a := graph.AdjMatrixWeightedGraphInit(5)

	// AdjMatrixDirectedWeightedGraphInit out of range
	assert.Nil(t, graph.AdjMatrixWeightedGraphInit(0))

	// AddEdge out of range
	assert.False(t, a.AddEdge(-1, 0, 0))
	assert.False(t, a.AddEdge(0, -1, 0))
	assert.False(t, a.AddEdge(10, 0, 0))
	assert.False(t, a.AddEdge(0, 10, 0))

	// RemoveEdge out of range
	assert.False(t, a.RemoveEdge(-1, 0))
	assert.False(t, a.RemoveEdge(0, -1))
	assert.False(t, a.RemoveEdge(10, 0))
	assert.False(t, a.RemoveEdge(0, 10))

	// PeekEdges out of range
	edges, ok := a.PeekEdges(-1)
	assert.Equal(t, [][]int{}, edges)
	assert.False(t, ok)

	edges, ok = a.PeekEdges(10)
	assert.Equal(t, [][]int{}, edges)
	assert.False(t, ok)
}

type fromToWeight struct {
	from, to, weight int
}

type fromTo struct {
	from, to int
}
