package datastructures

import "fmt"

type DequeValue interface{}

// DequeNode represents an element in the deque.
type DequeNode struct {
	Value DequeValue
	Prev  *DequeNode
	Next  *DequeNode
}

// Deque represents a double-ended queue.
type Deque struct {
	Front *DequeNode
	Back  *DequeNode
	Size  int
}

// NewDeque creates a new empty deque.
func NewDeque() *Deque {
	return &Deque{}
}

// PushFront adds an element to the front of the deque.
func (d *Deque) PushFront(value DequeValue) {
	newDequeNode := &DequeNode{Value: value}
	if d.IsEmpty() {
		d.Front = newDequeNode
		d.Back = newDequeNode
	} else {
		newDequeNode.Next = d.Front
		d.Front.Prev = newDequeNode
		d.Front = newDequeNode
	}
	d.Size++
}

// PushBack adds an element to the back of the deque.
func (d *Deque) PushBack(value DequeValue) {
	newDequeNode := &DequeNode{Value: value}
	if d.IsEmpty() {
		d.Front = newDequeNode
		d.Back = newDequeNode
	} else {
		newDequeNode.Prev = d.Back
		d.Back.Next = newDequeNode
		d.Back = newDequeNode
	}
	d.Size++
}

// PopFront removes and returns the element from the front of the deque.
func (d *Deque) PopFront() DequeValue {
	if d.IsEmpty() {
		return nil
	}
	value := d.Front.Value
	if d.Front == d.Back {
		d.Front = nil
		d.Back = nil
	} else {
		d.Front = d.Front.Next
		d.Front.Prev = nil
	}
	d.Size--
	return value
}

// PopBack removes and returns the element from the back of the deque.
func (d *Deque) PopBack() DequeValue {
	if d.IsEmpty() {
		return nil
	}
	value := d.Back.Value
	if d.Front == d.Back {
		d.Front = nil
		d.Back = nil
	} else {
		d.Back = d.Back.Prev
		d.Back.Next = nil
	}
	d.Size--
	return value
}

// PeekFront returns the element at the front of the deque without removing it.
func (d *Deque) PeekFront() DequeValue {
	if d.IsEmpty() {
		return nil
	}
	return d.Front.Value
}

// PeekBack returns the element at the back of the deque without removing it.
func (d *Deque) PeekBack() DequeValue {
	if d.IsEmpty() {
		return nil
	}
	return d.Back.Value
}

// IsEmpty returns true if the deque is empty, false otherwise.
func (d *Deque) IsEmpty() bool {
	return d.Size == 0
}

// GetLength returns the number of elements in the deque.
func (d *Deque) GetLength() int {
	return d.Size
}

// Print prints the elements of the deque from front to back.
func (d *Deque) Print() {
	current := d.Front
	for current != nil {
		fmt.Printf("%v ", current.Value)
		current = current.Next
	}
	fmt.Println()
}

func TestDeque(initDeque []DequeValue) {
	deque := NewDeque()
	for _, val := range initDeque {
		deque.PushBack(val)
	}
	fmt.Printf("initial deque: ")
	deque.Print()

	peekerbacker := deque.PeekBack()
	fmt.Printf("peek back test -%v-> deque: ", peekerbacker)
	deque.Print()

	//deque.PushBack(90)
	//fmt.Printf("push back test -> deque: ")
	//deque.Print()
	//
	//deque.PushFront(70)
	//fmt.Printf("push front test -> deque: ")
	//deque.Print()
	//
	//deque.PushBack(52)
	//fmt.Printf("push back test -> deque: ")
	//deque.Print()

	popperbacker := deque.PopBack()
	fmt.Printf("pop back -%v-> deque: ", popperbacker)
	deque.Print()

	//deque.PushFront(37)
	//fmt.Printf("push front test -> deque: ")
	//deque.Print()
	//
	//popperfronter := deque.PopFront()
	//fmt.Printf("pop front test -%v-> deque: ", popperfronter)
	//deque.Print()

	fmt.Printf("deque length: %v\n", deque.GetLength())
	fmt.Printf("deque isEmpty: %v\n", deque.IsEmpty())
}
