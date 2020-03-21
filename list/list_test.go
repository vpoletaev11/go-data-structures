package list_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/go-data-structures/list"
)

// Testing singly linked list:
//
// test InsertHead
func TestInsertHeadSinglyList(t *testing.T) {
	var l list.SinglyLinkedList

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
func TestInsertSinglyListSuccess(t *testing.T) {
	var l list.SinglyLinkedList

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

func TestInsertOutOfRangeSinglyList(t *testing.T) {
	var l list.SinglyLinkedList

	ok := l.Insert(5, "one")
	assert.False(t, ok)

	ok = l.Insert(-1, "two")
	assert.False(t, ok)
}

// test InsertTail
func TestInsertTailSinglyList(t *testing.T) {
	var l list.SinglyLinkedList

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
func TestGetHeadSinglyListSuccess(t *testing.T) {
	var l list.SinglyLinkedList

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

func TestGetHeadEmptyListSinglyList(t *testing.T) {
	var l list.SinglyLinkedList

	val, ok := l.GetHead()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// test Get
func TestGetSinglyListSuccess(t *testing.T) {
	var l list.SinglyLinkedList

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

func TestGetEmptyListSinglyList(t *testing.T) {
	var l list.SinglyLinkedList

	val, ok := l.Get(0)
	assert.False(t, ok)
	assert.Equal(t, "", val)
}

func TestGetOutOfRangeSinglyList(t *testing.T) {
	var l list.SinglyLinkedList

	l.InsertHead("one")

	val, ok := l.Get(5)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = l.Get(-1)
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// test GetTail
func TestGetTailSinglyListSuccess(t *testing.T) {
	var l list.SinglyLinkedList

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

func TestGetTailEmptyListSinglyList(t *testing.T) {
	var l list.SinglyLinkedList

	val, ok := l.GetTail()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// test PeekHead
func TestPeekHeadSinglyListSuccess(t *testing.T) {
	var l list.SinglyLinkedList

	l.InsertHead("one")
	l.InsertHead("two")
	l.InsertHead("three")

	val, ok := l.PeekHead()
	assert.Equal(t, "three", val)
	assert.True(t, ok)
}

func TestPeekHeadEmptyListSinglyList(t *testing.T) {
	var l list.SinglyLinkedList

	val, ok := l.PeekHead()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// test Peek
func TestPeekSinglyListSuccess(t *testing.T) {
	var l list.SinglyLinkedList

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

func TestPeekEmptyListSinglyList(t *testing.T) {
	var l list.SinglyLinkedList

	val, ok := l.Peek(0)
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

func TestPeekOutOfRangeSinglyList(t *testing.T) {
	var l list.SinglyLinkedList

	l.InsertHead("one")

	val, ok := l.Peek(2)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = l.Peek(-1)
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// test PeekTail
func TestPeekTailSinglyListSuccess(t *testing.T) {
	var l list.SinglyLinkedList

	l.InsertHead("one")
	l.InsertHead("two")
	l.InsertHead("three")

	val, ok := l.PeekTail()
	assert.Equal(t, "one", val)
	assert.True(t, ok)
}

func TestPeekTailEmptyListSinglyList(t *testing.T) {
	var l list.SinglyLinkedList

	val, ok := l.PeekTail()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// Testing doubly linked list:
//
// test InsertHead
func TestInsertHeadDoublyList(t *testing.T) {
	var l list.DoublyLinkedList

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
func TestInsertDoublyListSuccess(t *testing.T) {
	var l list.DoublyLinkedList

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

func TestInsertOutOfRangeDoublyList(t *testing.T) {
	var l list.DoublyLinkedList

	ok := l.Insert(0, "one")
	assert.True(t, ok)

	ok = l.Insert(5, "two")
	assert.False(t, ok)
}

func TestInsertOutOfRangeEmptyListDoublyList(t *testing.T) {
	var l list.DoublyLinkedList

	ok := l.Insert(5, "one")
	assert.False(t, ok)

	ok = l.Insert(-1, "two")
	assert.False(t, ok)
}

// test InsertTail
func TestInsertTailDoublyList(t *testing.T) {
	var l list.DoublyLinkedList

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
