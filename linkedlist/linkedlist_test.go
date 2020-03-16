package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertHead(t *testing.T) {
	var l LinkedList

	l.InsertHead("one")
	l.InsertHead("two")
	l.InsertHead("three")

	val, ok := l.GetHead()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = l.GetHead()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = l.GetHead()
	assert.Equal(t, "one", val)
	assert.True(t, ok)
}

func TestInsertTail(t *testing.T) {
	var l LinkedList

	l.InsertTail("one")
	l.InsertTail("two")
	l.InsertTail("three")

	val, ok := l.GetHead()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = l.GetHead()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = l.GetHead()
	assert.Equal(t, "three", val)
	assert.True(t, ok)
}

func TestGetHeadSuccess(t *testing.T) {
	var l LinkedList

	l.InsertTail("one")
	l.InsertTail("two")
	l.InsertTail("three")

	val, ok := l.GetHead()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = l.GetHead()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = l.GetHead()
	assert.Equal(t, "three", val)
	assert.True(t, ok)
}

func TestGetHeadEmptyList(t *testing.T) {
	var l LinkedList

	val, ok := l.GetHead()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

func TestGetTailSuccess(t *testing.T) {
	var l LinkedList

	l.InsertTail("one")
	l.InsertTail("two")
	l.InsertTail("three")

	val, ok := l.GetTail()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = l.GetTail()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = l.GetTail()
	assert.Equal(t, "one", val)
	assert.True(t, ok)
}

func TestGetTailEmptyList(t *testing.T) {
	var l LinkedList

	val, ok := l.GetTail()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

func TestPeekHeadSuccess(t *testing.T) {
	var l LinkedList

	l.InsertHead("one")
	l.InsertHead("two")
	l.InsertHead("three")

	val, ok := l.PeekHead()
	assert.Equal(t, "three", val)
	assert.True(t, ok)
}

func TestPeekHeadEmptyList(t *testing.T) {
	var l LinkedList

	val, ok := l.PeekHead()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

func TestPeekTailSuccess(t *testing.T) {
	var l LinkedList

	l.InsertHead("one")
	l.InsertHead("two")
	l.InsertHead("three")

	val, ok := l.PeekTail()
	assert.Equal(t, "one", val)
	assert.True(t, ok)
}

func TestPeekTailEmptyList(t *testing.T) {
	var l LinkedList

	val, ok := l.PeekTail()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}
