package datastructures

import "fmt"

type DoubleNode struct {
	data NodeData
	prev *DoubleNode
	next *DoubleNode
}

type DoublyLinkedList struct {
	head *DoubleNode
	tail *DoubleNode
	size int
}

// NewLinkedListDouble creates a new empty doubly linked list
func NewLinkedListDouble() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (dll *DoublyLinkedList) Append(data NodeData) {
	newNode := &DoubleNode{data: data, prev: dll.tail}
	if dll.size == 0 {
		dll.head = newNode
	} else {
		dll.tail.next = newNode
	}
	dll.tail = newNode
	dll.size++
}

func (dll *DoublyLinkedList) Prepend(data NodeData) {
	newNode := &DoubleNode{data: data, next: dll.head}
	if dll.size == 0 {
		dll.tail = newNode
	} else {
		dll.head.prev = newNode
	}
	dll.head = newNode
	dll.size++
}

func (dll *DoublyLinkedList) InsertAfter(target *DoubleNode, data NodeData) bool {
	if dll.head == nil {
		dll.head = &DoubleNode{}
		dll.tail = &DoubleNode{}
	}
	if target == nil {
		return false
	}
	newNode := &DoubleNode{data: data, prev: target, next: target.next}
	target.next = newNode
	if newNode.next != nil {
		newNode.next.prev = newNode
	} else {
		dll.tail = newNode
	}
	dll.size++
	return true
}

func (dll *DoublyLinkedList) Remove(target *DoubleNode) bool {
	if target == nil || dll.size == 0 {
		return false
	}
	if target == dll.head {
		dll.head = target.next
	} else {
		target.prev.next = target.next
	}
	if target == dll.tail {
		dll.tail = target.prev
	} else {
		target.next.prev = target.prev
	}
	target = nil
	dll.size--
	return true
}

func (dll *DoublyLinkedList) Search(targetData NodeData) *DoubleNode {
	current := dll.head
	for current != nil {
		if current.data == targetData {
			return current // Found the target data, return the node
		}
		current = current.next
	}
	return nil // Target data not found in the list
}

func (dll *DoublyLinkedList) InsertionSort() {
	// Start from the second node (the first node is already sorted).
	current := dll.head.next
	for current != nil {
		// Store the next node in the list before we rearrange curNode.
		next := current.next
		// Start searching for the correct position of curNode from the sorted part
		// of the list (before curNode).
		search := current.prev
		// Find the correct position by comparing data values.
		for search != nil && search.data.(int) > current.data.(int) {
			search = search.prev
		}
		// Remove curNode from its current position in the list.
		dll.Remove(current)

		// If searchNode is null, curNode should become the new head.
		if search == nil {
			current.prev = nil
			dll.Prepend(current)
		} else {
			dll.InsertAfter(search, current.data)
		}
		// Move to the next node in the unsorted part of the list.
		current = next
	}
}

func (dll *DoublyLinkedList) DisplayDouble() {
	if dll.head == nil {
		fmt.Println(" [head] -> <nil>")
		return
	}
	current := dll.head
	if dll.tail.data == nil {
		fmt.Printf(" [tail] -> <nil>\n")
	} else {
		fmt.Printf(" [tail] -> (%d)\n", dll.tail.data)
	}

	fmt.Printf(" [head] -> ")
	for current != nil {
		fmt.Print("(", current.data, ") <-> ")
		current = current.next
	}
	fmt.Println("nil")
}

func TestDoublyLinkedList(startList []NodeData) {
	list := NewLinkedListDouble()
	if len(startList) > 0 {
		fmt.Println("initial list:")
		for _, data := range startList {
			list.Append(data)
		}
	}
	list.DisplayDouble()
	removeTest := []NodeData{96, 59}
	for _, remove := range removeTest {
		fmt.Printf("\nremove test result (%d):\n", remove)
		list.Remove(list.Search(remove))
		list.DisplayDouble()
	}
}
