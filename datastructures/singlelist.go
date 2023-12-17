package datastructures

import (
	"fmt"
	"reflect"
)

/*
	Data Structure:
		- way  of organizing, storing, and performing operations on data
          - operations => access/update (read/write) stored data, searching
						  data, inserting new data, and removing data.
	- Basic Datastructures:
      * Record:      ds that stores sub-items (fields) with an associated name
      * Array:       ds that stores an ordered list of items accessible by index
      * Linked List: ds that stores an ordered list of items in nodes, where each
                     node stores data and has a pointer to the next node.
      * Binary Tree: ds where each node stores data and has up to 2 children, known
                     as a left child and a right child.
      * Hash Table:  ds that stores unordered items by mapping (hashing)
                     each item in an array location.
      * Heap:
        * max-heap - tree structure that maintains the property that a nodes'
                     key is >= node's child keys
		* min-heap - tree structure that maintains the property that a node's
                     key is <= node's child keys
      * Graph: ds that represents connections among items, and consists of vertices
               connected by edges.
              - Vertex => item in graph
              -   Edge => connection between 2 vertices
*/

/*
	Abstract Datatypes: data types described by predefined user operations.
	Common ADTs:
	 * List: an ADT for holding ordered data
	 * Dynamic Array: an ADT for holding ordered data and allowing indexed access.
	 * Stack: an ADT where items are only inserted on or removed from top of stack
	 * Queue: an ADT where items are inserted at end of queue and remove from front.
	 * Deque: an ADT known as a double-ended queue (or deck) where items can be inserted
              or removed from the front and the back of the queue
     * Bag: an ADT for storing items in which the order doesn't matter
            duplicates are allowed
	 * Set: an ADT for collection of distinct items
     * Priority Queue: queue where each item has priority
                       high priority items => closer to front
                       lower priority items => closer to back
	 * Dictionary (Map): an ADT that associates (maps) keys => values
*/

// NodeData represents the data type for the linked list.
type NodeData interface{}

// Node represents a node in the singly-linked list
type Node struct {
	Data NodeData
	Next *Node
}

// LinkedList represents a singly-linked list
type LinkedList struct {
	Head *Node
	Tail *Node
}

// NewLinkedList creates a new empty linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Append adds a new node with the given data to the end of the list
func (l *LinkedList) Append(data NodeData) {
	newNode := &Node{Data: data}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}

	l.Tail.Next = newNode
	l.Tail = newNode
}

// ListPrepend inserts a new node at the beginning of the list.
func (l *LinkedList) Prepend(data NodeData) {
	newNode := &Node{Data: data}
	newNode.Next = l.Head
	l.Head = newNode
	if l.Tail == nil {
		l.Tail = newNode
	}
}

// Delete removes the first occurrence of a node with the given data from the list
func (l *LinkedList) Delete(data NodeData) {
	if l.Head == nil {
		return
	}

	if l.Head.Data == data {
		if l.Head == l.Tail {
			l.Tail = nil
		}
		l.Head = l.Head.Next
		return
	}

	current := l.Head
	for current.Next != nil {
		if current.Next.Data == data {
			if current.Next == l.Tail {
				l.Tail = current
			}
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

// InsertAfter inserts a new node with the given data after the first occurrence
// of the specified existing data in the list.
func (l *LinkedList) InsertAfter(existingData, newData NodeData) {
	newNode := &Node{Data: newData}

	current := l.Head
	for current != nil {
		if current.Data == existingData {
			newNode.Next = current.Next
			current.Next = newNode
			if current == l.Tail {
				l.Tail = newNode
			}
			return
		}
		current = current.Next
	}
}

// RemoveAfter removes the node immediately after the first occurrence
// of the specified existing data in the list. If existingData is 0,
// it removes the head of the list.
func (l *LinkedList) RemoveAfter(existingData NodeData) {
	if l.Head == nil {
		// Empty list, nothing to remove
		return
	}

	if existingData == 0 {
		// Remove the head of the list
		if l.Head == l.Tail {
			l.Tail = nil
		}
		l.Head = l.Head.Next
		//if l.Head != nil {
		//	fmt.Printf("suc: %d\n", l.Head.Data)
		//}
		return
	}

	current := l.Head
	for current != nil && current.Next != nil {
		if current.Data == existingData {
			if current.Next == l.Tail {
				l.Tail = current
			}
			current.Next = current.Next.Next
			//if current.Next != nil {
			//	fmt.Printf("suc: %d\n", current.Next.Data)
			//}
			return
		}
		current = current.Next
	}
}

func (l *LinkedList) ListSearch(key NodeData) *Node {
	if l.Head == nil {
		return nil
	}
	current := l.Head
	for current != nil {
		if current.Data == key {
			return current
		}
		current = current.Next
	}
	return nil
}

// ListTraverse traverses the linked list and applies the given function to each node's data.
func (l *LinkedList) ListTraverse(fn func(data NodeData)) {
	current := l.Head
	for current != nil {
		fn(current.Data)
		current = current.Next
	}
}

// ListFindInsertionPosition finds the position to insert a node with the given data
// while keeping the list sorted.
func (l *LinkedList) ListFindInsertionPosition(data NodeData) *Node {
	current := l.Head
	var prev *Node

	for current != nil && data.(int) > current.Data.(int) {
		prev = current
		current = current.Next
	}

	return prev
}

// ListInsertionSortSinglyLinked sorts the linked list using the insertion sort algorithm.
func (l *LinkedList) ListInsertionSortSinglyLinked() {
	beforeCurrent := l.Head
	curNode := l.Head.Next

	for curNode != nil {
		next := curNode.Next
		position := l.ListFindInsertionPosition(curNode.Data)

		if position == beforeCurrent {
			beforeCurrent = curNode
		} else {
			l.RemoveAfter(beforeCurrent)
			if position == nil {
				l.Prepend(curNode.Data)
			} else {
				l.InsertAfter(position, curNode.Data)
			}
		}

		curNode = next
	}
}

// Display prints the elements of the linked list
func (l *LinkedList) Display() {
	if l.Head == nil {
		fmt.Println(" [head] -> nil")
		// Empty list, nothing to remove
		return
	}
	current := l.Head
	switch tail := l.Tail.Data.(type) {
	case int:
		fmt.Printf(" [tail]-> (%d)\n", tail)
	case rune:
		fmt.Printf(" [tail]-> (%c)\n", tail)
	}
	fmt.Printf(" [head]-> ")
	for current != nil {
		switch data := current.Data.(type) {
		case int:
			fmt.Printf("(%d) -> ", data)
		case rune:
			fmt.Printf("(%c) -> ", data)
		default:
			fmt.Printf("Unknown type: %v -> ", reflect.TypeOf(current.Data))
		}
		current = current.Next
	}
	fmt.Println("nil")
}

func TestLinkedList(startList []NodeData) {
	list := NewLinkedList()
	if len(startList) > 0 {
		fmt.Println("initial list:")
		for _, data := range startList {
			list.Append(data)
		}
	}
	list.Display()

	//removeAfterTest := []NodeData{49, 0}
	//for _, remove := range removeAfterTest {
	//	fmt.Printf("\nPost-RemoveAfter: %v\n", remove)
	//	list.RemoveAfter(remove)
	//	list.Display()
	//}
	searchTest := []NodeData{3, 6, 9}
	for _, search := range searchTest {
		fmt.Printf("\nListSearch(%v): ", search)
		found := list.ListSearch(search)
		if found != nil {
			if found.Next == nil {
				fmt.Printf("(%v) -> <nil>\n", found.Data)
			} else {
				fmt.Printf("(%v) -> (%v)\n", found.Data, found.Next.Data)
			}

		} else {
			fmt.Printf(" <nil> \n")
		}
	}
}
