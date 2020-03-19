package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test InsertHead
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

	val, ok = l.GetHead()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// test Insert
func TestInsertSuccess(t *testing.T) {
	var l LinkedList

	ok := l.Insert(0, "two")
	assert.True(t, ok)

	ok = l.Insert(0, "one")
	assert.True(t, ok)

	ok = l.Insert(1, "three")
	assert.True(t, ok)

	ok = l.Insert(3, "four")
	assert.True(t, ok)

	val, ok := l.GetHead()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = l.GetHead()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = l.GetHead()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = l.GetHead()
	assert.Equal(t, "four", val)
	assert.True(t, ok)

	val, ok = l.GetHead()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

func TestInsertOutOfRange(t *testing.T) {
	var l LinkedList

	ok := l.Insert(5, "one")
	assert.False(t, ok)

	ok = l.Insert(-1, "two")
	assert.False(t, ok)
}

// test InsertTail
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

	val, ok = l.GetHead()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// test GetHead
func TestGetHeadSuccess(t *testing.T) {
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

	val, ok = l.GetHead()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

func TestGetHeadEmptyList(t *testing.T) {
	var l LinkedList

	val, ok := l.GetHead()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// test Get
func TestGetSuccess(t *testing.T) {
	var l LinkedList

	l.InsertHead("one")
	l.InsertHead("two")
	l.InsertHead("three")

	val, ok := l.Get(2)
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = l.Get(0)
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = l.Get(0)
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = l.Get(0)
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

func TestGetEmptyList(t *testing.T) {
	var l LinkedList

	val, ok := l.Get(0)
	assert.False(t, ok)
	assert.Equal(t, "", val)
}

func TestGetOutOfRange(t *testing.T) {
	var l LinkedList

	l.InsertHead("one")

	val, ok := l.Get(5)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = l.Get(-1)
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// test GetTail
func TestGetTailSuccess(t *testing.T) {
	var l LinkedList

	l.InsertHead("one")
	l.InsertHead("two")
	l.InsertHead("three")

	val, ok := l.GetTail()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = l.GetTail()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = l.GetTail()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = l.GetTail()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

func TestGetTailEmptyList(t *testing.T) {
	var l LinkedList

	val, ok := l.GetTail()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// test PeekHead
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

// test Peek
func TestPeekSuccess(t *testing.T) {
	var l LinkedList

	l.InsertHead("one")
	l.InsertHead("two")
	l.InsertHead("three")

	val, ok := l.Peek(2)
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = l.Peek(1)
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = l.Peek(0)
	assert.Equal(t, "three", val)
	assert.True(t, ok)
}

func TestPeekEmptyList(t *testing.T) {
	var l LinkedList

	val, ok := l.Peek(0)
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

func TestPeekOutOfRange(t *testing.T) {
	var l LinkedList

	l.InsertHead("one")

	val, ok := l.Peek(2)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = l.Peek(-1)
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// test PeekTail
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
