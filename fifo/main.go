package main

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

type Queue struct {
	Head *Node
	Tail *Node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(value interface{}) {
	newNode := &Node{Value: value}

	if q.Tail == nil {
		q.Head = newNode
		q.Tail = newNode
	} else {
		q.Tail.Next = newNode
		q.Tail = newNode
	}
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	dequeueNode := q.Head
	q.Head = dequeueNode.Next
	if q.Head == nil {
		q.Tail = nil
	}

	return dequeueNode.Value
}

func (q *Queue) IsEmpty() bool {
	return q.Head == nil
}

func main() {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	for !q.IsEmpty() {
		fmt.Println(q.Dequeue())
	}
}
