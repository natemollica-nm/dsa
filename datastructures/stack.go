package datastructures

import "fmt"

// Stack represents a basic stack data structure.
type Stack struct {
	items []StackItem
}

// StackItem represents a basic stack data item
type StackItem interface{}

func NewStack() *Stack {
	return &Stack{}
}

// Push adds an item to the top of the stack.
func (s *Stack) Push(item StackItem) {
	s.items = append([]StackItem{item}, s.items...)
}

// Pop removes and returns the top item from the stack. It returns nil if the stack is empty.
func (s *Stack) Pop() StackItem {
	if len(s.items) == 0 {
		return nil
	}

	topIndex := 0
	item := s.items[topIndex]
	s.items = s.items[1:]
	return item
}

// Peek returns the top item from the stack without removing it. It returns nil if the stack is empty.
func (s *Stack) Peek() StackItem {
	if len(s.items) == 0 {
		return nil
	}
	return s.items[0]
}

// IsEmpty checks if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack.
func (s *Stack) Size() int {
	return len(s.items)
}

// Print prints the current stack contents.
func (s *Stack) Print() {
	fmt.Printf("%v\n", s.items)
}

func TestStack(initStack []StackItem) {
	stack := &Stack{items: initStack}
	fmt.Printf("initial stack: ")
	stack.Print()

	//peaker := stack.Peek()
	//fmt.Printf("peek test -> peaker %v -> stack: ", peaker)
	//stack.Print()
	//
	//stack.Push(13)
	//fmt.Printf("push test -> stack: ")
	//stack.Print()

	popper := stack.Pop()
	fmt.Printf("pop test -> popped %v -> stack: ", popper)
	stack.Print()

	peaker := stack.Peek()
	fmt.Printf("peek test -> peaker %v -> stack: ", peaker)
	stack.Print()

	//stack.Push(11)
	//fmt.Printf("push test -> stack: ")
	//stack.Print()
	//
	//popper = stack.Pop()
	//fmt.Printf("pop test -> popped %v -> stack: ", popper)
	//stack.Print()
	//
	//peaker = stack.Peek()
	//fmt.Printf("peek test -> peaker %v -> stack: ", peaker)
	//stack.Print()

	fmt.Println("final stack length:", stack.Size())
	fmt.Println("empty:", stack.IsEmpty())

}
