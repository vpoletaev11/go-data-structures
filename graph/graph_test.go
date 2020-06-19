package graph_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/go-data-structures/graph"
)

var datasetFT = []fromTo{
	{0, 6},
	{1, 4},
	{3, 2},
	{3, 9},
	{4, 4},
	{8, 0},
	{3, 3},
	{8, 0},
	{2, 4},
	{5, 1},
}
var datasetFTW = []fromToWeight{
	{0, 6, 1000},
	{1, 4, 38193},
	{3, 2, 14},
	{3, 9, 222222},
	{4, 4, 1309},
	{8, 0, 14981},
	{3, 3, 13992},
	{8, 0, 15},
	{2, 4, 0},
	{5, 1, 0},
}

// Testing AdjacencyMatrixDirectedGraph:
func TestAdjacencyMatrixDirectedGraphSUCCESS(t *testing.T) {
	a := graph.AdjMatrixDirectedGraphInit(10)
	expectedOut := [][]int{{6}, {4}, {4}, {2, 3, 9}, {4}, {1}, []int(nil), []int(nil), {0}, []int(nil)} // indexes of slices == indexes of vertexes

	assert.Equal(t, 10, a.Size())

	for _, value := range datasetFT {
		ok := a.AddEdge(value.from, value.to)
		assert.True(t, ok, "On adding edge with from: %d and to: %d", value.from, value.to)
	}

	for i, out := range expectedOut {
		edges, ok := a.PeekEdges(i)
		assert.Equal(t, out, edges, "On peeking edges with vertex: %d", i)
		assert.True(t, ok, "On peeking edges with vertex: %d", i)
	}

	// remove all edges
	for _, value := range datasetFT {
		ok := a.RemoveEdge(value.from, value.to)
		assert.True(t, ok, "On deleting edge with from: %d and to: %d", value.from, value.to)
	}

	// check that graph haven't edges
	for i := 0; i < a.Size(); i++ {
		edges, ok := a.PeekEdges(i)
		assert.Equal(t, []int(nil), edges, "PeekEdges for node %d %s", i, "should return empty slice")
		assert.True(t, ok, "PeekEdges for node %d %s", i, "should return true")
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
	a := graph.AdjMatrixDirectedWeightedGraphInit(10)
	expectedOut := [][][]int{
		{{6, 1000}},
		{{4, 38193}},
		[][]int(nil),
		{{2, 14}, {3, 13992}, {9, 222222}},
		{{4, 1309}},
		[][]int(nil),
		[][]int(nil),
		[][]int(nil),
		{{0, 15}},
		[][]int(nil),
	}

	assert.Equal(t, 10, a.Size())

	for _, value := range datasetFTW {
		ok := a.AddEdge(value.from, value.to, value.weight)
		assert.True(t, ok, "On adding edge with from: %d, to: %d, weight: %d", value.from, value.to, value.weight)
	}

	for i, out := range expectedOut {
		edges, ok := a.PeekEdges(i)
		assert.Equal(t, out, edges, "On peeking edges with vertex: %d", i)
		assert.True(t, ok, "On peeking edges with vertex: %d", i)
	}
	// remove all edges
	for _, value := range datasetFTW {
		ok := a.RemoveEdge(value.from, value.to)
		assert.True(t, ok, "On deleting edge with from: %d and to: %d", value.from, value.to)
	}

	// check that graph haven't edges
	for i := 0; i < a.Size(); i++ {
		edges, ok := a.PeekEdges(i)
		assert.Equal(t, [][]int(nil), edges, "PeekEdges for node %d %s", i, "should return empty slice")
		assert.True(t, ok, "PeekEdges for node %d %s", i, "should return true")
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
	a := graph.AdjMatrixGraphInit(10)
	expectedOut := [][]int{{6, 8}, {4, 5}, {3, 4}, {2, 3, 9}, {1, 2, 4}, {1}, {0}, []int(nil), {0}, {3}} // indexes of slices == indexes of vertexes

	assert.Equal(t, 10, a.Size())

	for _, value := range datasetFT {
		ok := a.AddEdge(value.from, value.to)
		assert.True(t, ok, "On adding edge with from: %d and to: %d", value.from, value.to)
	}

	for i, out := range expectedOut {
		edges, ok := a.PeekEdges(i)
		assert.Equal(t, out, edges, "On peeking edges with vertex: %d", i)
		assert.True(t, ok, "On peeking edges with vertex: %d", i)
	}

	// remove all edges
	for _, value := range datasetFT {
		ok := a.RemoveEdge(value.from, value.to)
		assert.True(t, ok, "On deleting edge with from: %d and to: %d", value.from, value.to)
	}

	// check that graph haven't edges
	for i := 0; i < a.Size(); i++ {
		edges, ok := a.PeekEdges(i)
		assert.Equal(t, []int(nil), edges, "PeekEdges for node %d %s", i, "should return empty slice")
		assert.True(t, ok, "PeekEdges for node %d %s", i, "should return true")
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
	a := graph.AdjMatrixWeightedGraphInit(10)
	expectedOut := [][][]int{
		{{6, 1000}, {8, 15}},
		{{4, 38193}},
		{{3, 14}},
		{{2, 14}, {3, 13992}, {9, 222222}},
		{{1, 38193}, {4, 1309}},
		[][]int(nil),
		{{0, 1000}},
		[][]int(nil),
		{{0, 15}},
		{{3, 222222}},
	}

	assert.Equal(t, 10, a.Size())

	for _, value := range datasetFTW {
		ok := a.AddEdge(value.from, value.to, value.weight)
		assert.True(t, ok, "On adding edge with from: %d, to: %d, weight: %d", value.from, value.to, value.weight)
	}

	for i, out := range expectedOut {
		edges, ok := a.PeekEdges(i)
		assert.Equal(t, out, edges, "On peeking edges with vertex: %d", i)
		assert.True(t, ok, "On peeking edges with vertex: %d", i)
	}

	// remove all edges
	for _, value := range datasetFTW {
		ok := a.RemoveEdge(value.from, value.to)
		assert.True(t, ok, "On deleting edge with from: %d and to: %d", value.from, value.to)
	}

	// check that graph haven't edges
	for i := 0; i < a.Size(); i++ {
		edges, ok := a.PeekEdges(i)
		assert.Equal(t, [][]int(nil), edges, "PeekEdges for node %d %s", i, "should return empty slice")
		assert.True(t, ok, "PeekEdges for node %d %s", i, "should return true")
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

// BENCHMARKS
// Benchmarking AdjacencyMatrixDirectedGraph:
func BenchmarkAddEdgeAMDirectedGraph(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		g := graph.AdjMatrixDirectedGraphInit(10)
		b.StartTimer()

		for _, ft := range datasetFT {
			g.AddEdge(ft.from, ft.to)
		}
	}
}

func BenchmarkRemoveEdgeAMDirectedGraph(b *testing.B) {
	g := graph.AdjMatrixDirectedGraphInit(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for _, ft := range datasetFT {
			g.AddEdge(ft.from, ft.to)
		}
		b.StartTimer()

		for _, ft := range datasetFT {
			g.RemoveEdge(ft.from, ft.to)
		}
	}
}

func BenchmarkPeekEdgesAMDirectedGraph(b *testing.B) {
	g := graph.AdjMatrixDirectedGraphInit(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for _, ft := range datasetFT {
			g.AddEdge(ft.from, ft.to)
		}
		b.StartTimer()

		for i := 0; i < 10; i++ {
			g.PeekEdges(i)
		}
	}
}

// Benchmarking AdjacencyMatrixDirectedWeightedGraph:
func BenchmarkAddEdgeAMDirectedWeightedGraph(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		g := graph.AdjMatrixDirectedWeightedGraphInit(10)
		b.StartTimer()

		for _, ftw := range datasetFTW {
			g.AddEdge(ftw.from, ftw.to, ftw.weight)
		}
	}
}

func BenchmarkRemoveEdgeAMDirectedWeightedGraph(b *testing.B) {
	g := graph.AdjMatrixDirectedWeightedGraphInit(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for _, ftw := range datasetFTW {
			g.AddEdge(ftw.from, ftw.to, ftw.weight)
		}
		b.StartTimer()

		for _, ft := range datasetFTW {
			g.RemoveEdge(ft.from, ft.to)
		}
	}
}

func BenchmarkPeekEdgesAMDirectedWeightedGraph(b *testing.B) {
	g := graph.AdjMatrixDirectedWeightedGraphInit(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for _, ftw := range datasetFTW {
			g.AddEdge(ftw.from, ftw.to, ftw.weight)
		}
		b.StartTimer()

		for i := 0; i < 10; i++ {
			g.PeekEdges(i)
		}
	}
}

// Benchmarking AdjacencyMatrixGraph:
func BenchmarkAddEdgeAMGraph(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		g := graph.AdjMatrixGraphInit(10)
		b.StartTimer()

		for _, ft := range datasetFT {
			g.AddEdge(ft.from, ft.to)
		}
	}
}

func BenchmarkRemoveEdgeAMGraph(b *testing.B) {
	g := graph.AdjMatrixGraphInit(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for _, ft := range datasetFT {
			g.AddEdge(ft.from, ft.to)
		}
		b.StartTimer()

		for _, ft := range datasetFT {
			g.RemoveEdge(ft.from, ft.to)
		}
	}
}

func BenchmarkPeekEdgesAMGraph(b *testing.B) {
	g := graph.AdjMatrixGraphInit(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for _, ft := range datasetFT {
			g.AddEdge(ft.from, ft.to)
		}
		b.StartTimer()

		for i := 0; i < 10; i++ {
			g.PeekEdges(i)
		}
	}
}

// Benchmarking AdjacencyMatrixDirectedWeightedGraph:
func BenchmarkAddEdgeAMWeightedGraph(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		g := graph.AdjMatrixWeightedGraphInit(10)
		b.StartTimer()

		for _, ftw := range datasetFTW {
			g.AddEdge(ftw.from, ftw.to, ftw.weight)
		}
	}
}

func BenchmarkRemoveEdgeAMWeightedGraph(b *testing.B) {
	g := graph.AdjMatrixWeightedGraphInit(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for _, ftw := range datasetFTW {
			g.AddEdge(ftw.from, ftw.to, ftw.weight)
		}
		b.StartTimer()

		for _, ft := range datasetFTW {
			g.RemoveEdge(ft.from, ft.to)
		}
	}
}

func BenchmarkPeekEdgesAMWeightedGraph(b *testing.B) {
	g := graph.AdjMatrixWeightedGraphInit(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for _, ftw := range datasetFTW {
			g.AddEdge(ftw.from, ftw.to, ftw.weight)
		}
		b.StartTimer()

		for i := 0; i < 10; i++ {
			g.PeekEdges(i)
		}
	}
}

type fromToWeight struct {
	from, to, weight int
}

type fromTo struct {
	from, to int
}
