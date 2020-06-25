package graph

import (
	"github.com/vpoletaev11/go-data-structures/set"
)

// ALGraph - Graph based on Adjacency List.
// Indexing of vertexes starts from 0.
type ALGraph struct {
	adjacencyList []set.Set
}

// Size returns count of vertexes in the graph
func (a *ALGraph) Size() int {
	return len(a.adjacencyList)
}

// ALGraphInit initializes Directed Graph based on Adjacency Matrix
func ALGraphInit(size int) *ALGraph {
	if size < 1 {
		return nil
	}
	var a ALGraph

	a.adjacencyList = make([]set.Set, size)
	for i := 0; i < size; i++ {
		a.adjacencyList[i] = set.Init()
	}
	return &a
}

// AddEdge adds edge between vertexes
func (a *ALGraph) AddEdge(from, to int) bool {
	if from > a.Size()-1 || to > a.Size()-1 || from < 0 || to < 0 {
		return false
	}

	a.adjacencyList[from].Add(to)
	a.adjacencyList[to].Add(from)
	return true
}

// RemoveEdge removes edge between vertexes
func (a *ALGraph) RemoveEdge(from, to int) bool {
	if from > a.Size()-1 || to > a.Size()-1 || from < 0 || to < 0 {
		return false
	}

	a.adjacencyList[from].Delete(to)
	a.adjacencyList[to].Delete(from)
	return true
}

// PeekEdges returns slice of edges of inputted vertex
func (a *ALGraph) PeekEdges(vertex int) ([]int, bool) {
	if vertex > a.Size()-1 || vertex < 0 {
		return []int{}, false
	}

	return a.adjacencyList[vertex].GetSet(), true
}
