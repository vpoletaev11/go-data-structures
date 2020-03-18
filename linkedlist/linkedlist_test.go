package linkedlist

import (
	"fmt"
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
}

// test Insert
func TestInsertSuccess(t *testing.T) {
	var l LinkedList

	err := l.Insert(0, "two")
	assert.Nil(t, err)

	err = l.Insert(0, "one")
	assert.Nil(t, err)

	err = l.Insert(1, "three")
	assert.Nil(t, err)

	err = l.Insert(3, "four")
	assert.Nil(t, err)

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
}

func TestInsertOutOfRange(t *testing.T) {
	var l LinkedList

	l.InsertHead("one")

	err := l.Insert(5, "two")
	assert.Equal(t, fmt.Errorf("Index out of range"), err)
}

func TestInsertOutOfRangeEmptyList(t *testing.T) {
	var l LinkedList

	err := l.Insert(5, "one")
	assert.Equal(t, fmt.Errorf("Index out of range"), err)
}

func TestInsertNegativeIndex(t *testing.T) {
	var l LinkedList

	err := l.Insert(-1, "value")

	assert.Equal(t, fmt.Errorf("Cannot use negative index: -1"), err)
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

	val, err := l.Get(2)
	assert.Nil(t, err)
	assert.Equal(t, "one", val)

	val, err = l.Get(0)
	assert.Nil(t, err)
	assert.Equal(t, "three", val)

	val, err = l.Get(0)
	assert.Nil(t, err)
	assert.Equal(t, "two", val)
}

func TestGetEmptyList(t *testing.T) {
	var l LinkedList

	val, err := l.Get(0)
	assert.Equal(t, fmt.Errorf("Empty list"), err)
	assert.Equal(t, "", val)
}

func TestGetNegativeIndex(t *testing.T) {
	var l LinkedList

	l.InsertHead("one")

	val, err := l.Get(-1)
	assert.Equal(t, fmt.Errorf("Cannot use negative index: -1"), err)
	assert.Equal(t, "", val)
}

func TestGetOutOfRange(t *testing.T) {
	var l LinkedList

	l.InsertHead("one")

	val, err := l.Get(5)
	assert.Equal(t, fmt.Errorf("Index out of range"), err)
	assert.Equal(t, "", val)
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

	val, err := l.Peek(2)
	assert.Equal(t, "one", val)
	assert.Nil(t, err)
}

func TestPeekEmptyList(t *testing.T) {
	var l LinkedList

	val, err := l.Peek(0)
	assert.Equal(t, "", val)
	assert.Equal(t, fmt.Errorf("Empty list"), err)
}

func TestPeekOutOfRange(t *testing.T) {
	var l LinkedList

	l.InsertHead("one")

	val, err := l.Peek(2)
	assert.Equal(t, "", val)
	assert.Equal(t, fmt.Errorf("Index out of range"), err)
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
