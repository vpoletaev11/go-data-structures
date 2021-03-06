package graph

import (
	"github.com/vpoletaev11/go-data-structures/set"
)

/*
	ADJACENCY LIST DIRECTED GRAPH:

	Size()               AddEge()             RemoveEdge()         PeekEdges()
	Best:    O(1)        Best:    O(1)        Best:    O(1)        Best:    O(e) e - count of edges
	Average: O(1)        Average: O(1)        Average: O(1)        Average: O(e)
	Worst:   O(1)        Worst:   O(1)        Worst:   O(1)        Worst:   O(e)


	goos: linux
	goarch: amd64
	BenchmarkAddEdgeALDirectedGraph-8              	 2065872	       582 ns/op	     560 B/op	       7 allocs/op
	BenchmarkRemoveEdgeALDirectedGraph-8           	 4847061	       254 ns/op	       0 B/op	       0 allocs/op
	BenchmarkPeekEdgesALDirectedGraph-8            	 1292356	       924 ns/op	      80 B/op	       7 allocs/op
*/

// ALDirectedGraph - Directed Graph based on Adjacency List.
// Indexing of vertexes starts from 0.
type ALDirectedGraph struct {
	adjacencyList []set.Set
}

// Size returns count of vertexes in the graph
func (a *ALDirectedGraph) Size() int {
	return len(a.adjacencyList)
}

// ALDirectedGraphInit initializes Directed Graph based on Adjacency Matrix
func ALDirectedGraphInit(size int) *ALDirectedGraph {
	if size < 1 {
		return nil
	}
	var a ALDirectedGraph

	a.adjacencyList = make([]set.Set, size)
	for i := 0; i < size; i++ {
		a.adjacencyList[i] = set.Init()
	}
	return &a
}

// AddEdge adds edge between vertexes
func (a *ALDirectedGraph) AddEdge(from, to int) bool {
	if from > a.Size()-1 || to > a.Size()-1 || from < 0 || to < 0 {
		return false
	}

	a.adjacencyList[from].Add(to)
	return true
}

// RemoveEdge removes edge between vertexes
func (a *ALDirectedGraph) RemoveEdge(from, to int) bool {
	if from > a.Size()-1 || to > a.Size()-1 || from < 0 || to < 0 {
		return false
	}

	a.adjacencyList[from].Delete(to)
	return true
}

// PeekEdges returns slice of edges of inputted vertex
func (a *ALDirectedGraph) PeekEdges(vertex int) ([]int, bool) {
	if vertex > a.Size()-1 || vertex < 0 {
		return []int{}, false
	}

	return a.adjacencyList[vertex].GetSet(), true
}
