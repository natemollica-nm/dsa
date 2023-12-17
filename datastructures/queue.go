package datastructures

import "fmt"

type Queue struct {
	items []QueueItem
}

type QueueItem interface{}

func NewQueue() *Queue {
	return &Queue{}
}

// Enqueue adds an item to the end of the queue.
func (q *Queue) Enqueue(item QueueItem) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the front item from the queue.
func (q *Queue) Dequeue() QueueItem {
	if len(q.items) == 0 {
		return nil
	}
	front := q.items[0]
	q.items = q.items[1:]
	return front
}

// Peek returns the front item of the queue without removing it.
func (q *Queue) Peek() QueueItem {
	if len(q.items) == 0 {
		return nil
	}
	return q.items[0]
}

// IsEmpty returns true if the queue is empty, false otherwise.
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// GetLength returns the number of items in the queue.
func (q *Queue) GetLength() int {
	return len(q.items)
}

// Print prints the items in the queue from front to back.
func (q *Queue) Print() {
	fmt.Printf("%v\n", q.items)
}

func TestQueue(initQueue []QueueItem) {
	queue := &Queue{items: initQueue}
	fmt.Printf("initial queue: ")
	queue.Print()

	//peaker := queue.Peek()
	//fmt.Printf("peak test -> peaker %v -> queue: ", peaker)
	//queue.Print()
	//
	queue.Enqueue(11)
	fmt.Printf("enqueue test -> queue: ")
	queue.Print()

	queue.Enqueue(34)
	fmt.Printf("enqueue test -> queue: ")
	queue.Print()

	removed := queue.Dequeue()
	fmt.Printf("dequeue test -> removed %v -> queue: ", removed)
	queue.Print()
	//
	//queue.Enqueue(81)
	//fmt.Printf("enqueue test -> queue: ")
	//queue.Print()
	//
	//peaker = queue.Peek()
	//fmt.Printf("peak test -> peaker %v -> queue: ", peaker)
	//queue.Print()

	removed = queue.Dequeue()
	fmt.Printf("dequeue test -> removed %v -> queue: ", removed)
	queue.Print()

	fmt.Println("queue length ->", queue.GetLength())
	fmt.Println("queue isEmpty ->", queue.IsEmpty())
}
