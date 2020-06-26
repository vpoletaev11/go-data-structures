package graph_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/go-data-structures/graph"
)

type graphI interface {
	AddEdge(int, int) bool
	RemoveEdge(int, int) bool
	PeekEdges(int) ([]int, bool)
	Size() int
}

type weightedGraphI interface {
	AddEdge(int, int, int) bool
	RemoveEdge(int, int) bool
	PeekEdges(int) ([][]int, bool)
	Size() int
}

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

var expectedOut = [][]int{
	{6, 8},
	{4, 5},
	{3, 4},
	{2, 3, 9},
	{1, 2, 4},
	{1},
	{0},
	{},
	{0},
	{3},
}

var expectedOutDirected = [][]int{
	{6},
	{4},
	{4},
	{2, 3, 9},
	{4},
	{1},
	{},
	{},
	{0},
	{},
}

// Testing NOT WEIGHTED GRAPHS (AMDirectedGraph ALDirectedGraph and AMGraph ALGraph):
func TestAMGraphSuccess(t *testing.T) {
	graphSuccess(t, graph.AMGraphInit(10), expectedOut)
}

func TestAMDirectedGraphSuccess(t *testing.T) {
	graphSuccess(t, graph.AMDirectedGraphInit(10), expectedOutDirected)
}

func TestALGraphSuccess(t *testing.T) {
	graphSuccess(t, graph.ALGraphInit(10), expectedOut)
}

func TestALDirectedGraphSuccess(t *testing.T) {
	graphSuccess(t, graph.ALDirectedGraphInit(10), expectedOutDirected)
}

func graphSuccess(t *testing.T, graph graphI, expectedOut [][]int) {
	assert.Equal(t, 10, graph.Size())

	for _, value := range datasetFT {
		ok := graph.AddEdge(value.from, value.to)
		assert.True(t, ok, "On adding edge with from: %d and to: %d", value.from, value.to)
	}

	for i, out := range expectedOut {
		edges, ok := graph.PeekEdges(i)
		assert.ElementsMatch(t, out, edges, "On peeking edges with vertex: %d", i)
		assert.True(t, ok, "On peeking edges with vertex: %d", i)
	}

	// remove all edges
	for _, value := range datasetFT {
		ok := graph.RemoveEdge(value.from, value.to)
		assert.True(t, ok, "On deleting edge with from: %d and to: %d", value.from, value.to)
	}

	// check that graph haven't edges
	for i := 0; i < graph.Size(); i++ {
		edges, ok := graph.PeekEdges(i)
		assert.Equal(t, []int{}, edges, "PeekEdges for node %d", i)
		assert.True(t, ok, "PeekEdges for node %d", i)
	}
}

func TestAMGraphOutOfRange(t *testing.T) {
	assert.Nil(t, graph.AMGraphInit(0)) // out of range

	graphOutOfRange(t, graph.AMGraphInit(5))
}

func TestAMDirectedGraphOutOfRange(t *testing.T) {
	assert.Nil(t, graph.AMDirectedGraphInit(0)) // out of range

	graphOutOfRange(t, graph.AMDirectedGraphInit(5))
}

func TestALGraphOutOfRange(t *testing.T) {
	assert.Nil(t, graph.ALGraphInit(0)) // out of range

	graphOutOfRange(t, graph.ALGraphInit(5))
}

func TestALDirectedGraphOutOfRange(t *testing.T) {
	assert.Nil(t, graph.ALDirectedGraphInit(0)) // out of range

	graphOutOfRange(t, graph.ALDirectedGraphInit(5))
}

func graphOutOfRange(t *testing.T, graph graphI) {
	// AddEdge out of range
	assert.False(t, graph.AddEdge(-1, 0))
	assert.False(t, graph.AddEdge(0, -1))
	assert.False(t, graph.AddEdge(10, 0))
	assert.False(t, graph.AddEdge(0, 10))

	// RemoveEdge out of range
	assert.False(t, graph.RemoveEdge(-1, 0))
	assert.False(t, graph.RemoveEdge(0, -1))
	assert.False(t, graph.RemoveEdge(10, 0))
	assert.False(t, graph.RemoveEdge(0, 10))

	// PeekEdges out of range
	edges, ok := graph.PeekEdges(-1)
	assert.Equal(t, []int{}, edges)
	assert.False(t, ok)

	edges, ok = graph.PeekEdges(10)
	assert.Equal(t, []int{}, edges)
	assert.False(t, ok)
}

// Testing AMDirectedWeightedGraph and AMWeightedGraph:
func TestAMDirectedWeightedGraphSuccess(t *testing.T) {
	expectedOut := [][][]int{
		{{6, 1000}},
		{{4, 38193}},
		{},
		{{2, 14}, {3, 13992}, {9, 222222}},
		{{4, 1309}},
		{},
		{},
		{},
		{{0, 15}},
		{},
	}
	weightedGraphSuccess(t, graph.AMDirectedWeightedGraphInit(10), expectedOut)
}

func TestAMWeightedGraphSuccess(t *testing.T) {
	expectedOut := [][][]int{
		{{6, 1000}, {8, 15}},
		{{4, 38193}},
		{{3, 14}},
		{{2, 14}, {3, 13992}, {9, 222222}},
		{{1, 38193}, {4, 1309}},
		{},
		{{0, 1000}},
		{},
		{{0, 15}},
		{{3, 222222}},
	}
	weightedGraphSuccess(t, graph.AMWeightedGraphInit(10), expectedOut)
}

func TestALDirectedWeightedGraphSuccess(t *testing.T) {
	expectedOut := [][][]int{
		{{6, 1000}},
		{{4, 38193}},
		{},
		{{2, 14}, {3, 13992}, {9, 222222}},
		{{4, 1309}},
		{},
		{},
		{},
		{{0, 15}},
		{},
	}
	weightedGraphSuccess(t, graph.ALDirectedWeightedGraphInit(10), expectedOut)
}

func weightedGraphSuccess(t *testing.T, graph weightedGraphI, expectedOut [][][]int) {
	assert.Equal(t, 10, graph.Size())

	for _, value := range datasetFTW {
		ok := graph.AddEdge(value.from, value.to, value.weight)
		assert.True(t, ok, "On adding edge with from: %d, to: %d, weight: %d", value.from, value.to, value.weight)
	}

	for i, out := range expectedOut {
		edges, ok := graph.PeekEdges(i)
		assert.ElementsMatch(t, out, edges, "On peeking edges with vertex: %d", i)
		assert.True(t, ok, "On peeking edges with vertex: %d", i)
	}
	// remove all edges
	for _, value := range datasetFTW {
		ok := graph.RemoveEdge(value.from, value.to)
		assert.True(t, ok, "On deleting edge with from: %d and to: %d", value.from, value.to)
	}

	// check that graph haven't edges
	for i := 0; i < graph.Size(); i++ {
		edges, ok := graph.PeekEdges(i)
		assert.Equal(t, [][]int{}, edges, "PeekEdges for node %d", i)
		assert.True(t, ok, "PeekEdges for node %d", i)
	}
}

func TestAMDirectedWeightedGraphOutOfRange(t *testing.T) {
	assert.Nil(t, graph.AMDirectedWeightedGraphInit(0)) // out of range

	weightedGraphOutOfRange(t, graph.AMDirectedWeightedGraphInit(5))
}

func TestAMWeightedGraphOutOfRange(t *testing.T) {
	assert.Nil(t, graph.AMWeightedGraphInit(0)) // out of range

	weightedGraphOutOfRange(t, graph.AMWeightedGraphInit(5))
}

func TestALDirectedWeightedGraphOutOfRange(t *testing.T) {
	assert.Nil(t, graph.ALDirectedWeightedGraphInit(0)) // out of range

	weightedGraphOutOfRange(t, graph.ALDirectedWeightedGraphInit(5))
}

func weightedGraphOutOfRange(t *testing.T, graph weightedGraphI) {
	// AddEdge out of range
	assert.False(t, graph.AddEdge(-1, 0, 0))
	assert.False(t, graph.AddEdge(0, -1, 0))
	assert.False(t, graph.AddEdge(10, 0, 0))
	assert.False(t, graph.AddEdge(0, 10, 0))

	// RemoveEdge out of range
	assert.False(t, graph.RemoveEdge(-1, 0))
	assert.False(t, graph.RemoveEdge(0, -1))
	assert.False(t, graph.RemoveEdge(10, 0))
	assert.False(t, graph.RemoveEdge(0, 10))

	// PeekEdges out of range
	edges, ok := graph.PeekEdges(-1)
	assert.Equal(t, [][]int{}, edges)
	assert.False(t, ok)

	edges, ok = graph.PeekEdges(10)
	assert.Equal(t, [][]int{}, edges)
	assert.False(t, ok)
}

// BENCHMARKS
// Benchmarking AMDirectedGraph:
func BenchmarkAddEdgeAMDirectedGraph(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		g := graph.AMDirectedGraphInit(10)
		b.StartTimer()

		for _, ft := range datasetFT {
			g.AddEdge(ft.from, ft.to)
		}
	}
}

func BenchmarkRemoveEdgeAMDirectedGraph(b *testing.B) {
	g := graph.AMDirectedGraphInit(10)
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
	g := graph.AMDirectedGraphInit(10)
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

// Benchmarking AMDirectedWeightedGraph:
func BenchmarkAddEdgeAMDirectedWeightedGraph(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		g := graph.AMDirectedWeightedGraphInit(10)
		b.StartTimer()

		for _, ftw := range datasetFTW {
			g.AddEdge(ftw.from, ftw.to, ftw.weight)
		}
	}
}

func BenchmarkRemoveEdgeAMDirectedWeightedGraph(b *testing.B) {
	g := graph.AMDirectedWeightedGraphInit(10)
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
	g := graph.AMDirectedWeightedGraphInit(10)
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

// Benchmarking AMGraph:
func BenchmarkAddEdgeAMGraph(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		g := graph.AMGraphInit(10)
		b.StartTimer()

		for _, ft := range datasetFT {
			g.AddEdge(ft.from, ft.to)
		}
	}
}

func BenchmarkRemoveEdgeAMGraph(b *testing.B) {
	g := graph.AMGraphInit(10)
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
	g := graph.AMGraphInit(10)
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

// Benchmarking AMDirectedWeightedGraph:
func BenchmarkAddEdgeAMWeightedGraph(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		g := graph.AMWeightedGraphInit(10)
		b.StartTimer()

		for _, ftw := range datasetFTW {
			g.AddEdge(ftw.from, ftw.to, ftw.weight)
		}
	}
}

func BenchmarkRemoveEdgeAMWeightedGraph(b *testing.B) {
	g := graph.AMWeightedGraphInit(10)
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
	g := graph.AMWeightedGraphInit(10)
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
