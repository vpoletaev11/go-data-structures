package queue

import (
	"fmt"
	"testing"
)

func TestLen(t *testing.T) {
	var queue Queue

	queue.Enqueue("one")
	queue.Enqueue("two")
	queue.Enqueue("three")

	if queue.Len() != 3 {
		t.Error("Enqueued 3 elements, but length of queue ==", queue.Len())
	}
}

func TestEnqueue(t *testing.T) {
	var queue Queue

	queue.Enqueue("one")

	if queue.items[0] != "one" {
		t.Error("Enqueued element \"one\", but queue item ==", queue.items[0])
	}
}

func TestDequeueSuccess(t *testing.T) {
	var queue Queue

	queue.items = append(queue.items, "one")
	queue.items = append(queue.items, "two")

	elem, ok := queue.Dequeue()
	if ok != true {
		t.Error("Dequeue status should be true, but returns false")
	}
	if elem != "one" {
		t.Error("Enqueued value \"one\", but dequeued:", elem)
	}

	elem, ok = queue.Dequeue()
	if ok != true {
		t.Error("Dequeue status should be true, but returns false")
	}
	if elem != "two" {
		t.Error("Enqueued value \"one\", but dequeued:", elem)
	}
}

func TestDequeueEmptyQueue(t *testing.T) {
	var queue Queue

	elem, ok := queue.Dequeue()
	if elem != "" {
		t.Error("Nothing enqueued, but dequeued value ==", elem)
	}
	if ok != false {
		t.Error("Nothing enqueued, but dequeue status == true")
	}
}

func TestPeekSuccess(t *testing.T) {
	var queue Queue

	queue.items = append(queue.items, "one")

	elem, ok := queue.Peek()
	if ok != true {
		t.Error("Peek should return true, but returns false")
	}
	if elem != "one" {
		t.Error("Peeked value \"one\", but dequeued:", elem)
	}
}

func TestPeekEmptyQueue(t *testing.T) {
	var queue Queue

	elem, ok := queue.Peek()
	if elem != "" {
		t.Error("Nothing enqueued, but peeked value ==", elem)
	}
	if ok != false {
		t.Error("Nothing enqueued, but peek status == true")
	}
}

// TestQueueDemo displays example of using the queue data structure
func TestQueueDemo(t *testing.T) {
	fmt.Println("Demonstration of using the queue data structure:")
	var queue Queue

	fmt.Println("Enqueued: one")
	queue.Enqueue("one")

	fmt.Println("Enqueued: two")
	queue.Enqueue("two")

	fmt.Println("Enqueued: three")
	queue.Enqueue("three")
	fmt.Print("\n")

	elem, ok := queue.Peek()
	if ok {
		fmt.Println("Peeked:", elem)
	}

	for queue.Len() > 0 {
		elem, ok := queue.Dequeue()
		if ok {
			fmt.Println("Dequeued:", elem)
		}
	}
}
