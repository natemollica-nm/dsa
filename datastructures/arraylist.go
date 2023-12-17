package datastructures

import "fmt"

type ArrayList struct {
	data     []ArrayListData
	capacity int
	size     int
}

type ArrayListData interface{}

func NewArrayList() *ArrayList {
	return &ArrayList{
		data:     make([]ArrayListData, 5), // Initial capacity of 10, you can change it as needed
		capacity: 5,
		size:     0,
	}
}

func (list *ArrayList) ArrayListAppend(value ArrayListData) {
	if list.size == list.capacity {
		list.ArrayListResize(list.capacity * 2)
	}

	list.data[list.size] = value
	list.size++
}

func (list *ArrayList) ArrayListResize(newCapacity int) {
	newData := make([]ArrayListData, newCapacity)
	copy(newData, list.data)
	list.data = newData
	list.capacity = newCapacity
}

func (list *ArrayList) ArrayListPrepend(value ArrayListData) {
	if list.size == list.capacity {
		list.ArrayListResize(list.capacity * 2)
	}

	copy(list.data[1:], list.data[:list.size])
	list.data[0] = value
	list.size++
}

func (list *ArrayList) ArrayListInsertAfter(value ArrayListData, index int) error {
	if index < 0 || index >= list.size {
		return fmt.Errorf("index out of range")
	}

	if list.size == list.capacity {
		list.ArrayListResize(list.capacity * 2)
	}

	copy(list.data[index+2:], list.data[index+1:list.size+1])
	list.data[index+1] = value
	list.size++
	return nil
}

func (list *ArrayList) ArrayListSearch(value ArrayListData) int {
	for i := 0; i < list.size; i++ {
		if list.data[i] == value {
			return i
		}
	}
	return -1
}

func (list *ArrayList) ArrayListRemoveAt(index int) error {
	if index < 0 || index >= list.size {
		return fmt.Errorf("index out of range")
	}

	copy(list.data[index:], list.data[index+1:list.size])
	list.data[list.size-1] = nil // Zero out the last element
	list.size--

	if list.size <= list.capacity/4 && list.capacity > 10 {
		list.ArrayListResize(list.capacity / 2)
	}

	return nil
}

func (list *ArrayList) Print() {
	fmt.Print("[")
	for i := 0; i < list.size; i++ {
		fmt.Print(list.data[i])
		if i < list.size-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}

func TestArrayList(initData []ArrayListData) {
	list := NewArrayList()
	for _, data := range initData {
		list.ArrayListAppend(data)
	}
	fmt.Printf("initial: ")
	list.Print()
	fmt.Printf("allocation size: %d\n", list.capacity)
	fmt.Printf("array list length: %d\n", list.size)

	indx := list.ArrayListSearch(26)
	fmt.Printf("arry search -> 26 ([%d]) -> list: ", indx)
	list.Print()

	_ = list.ArrayListRemoveAt(0)
	fmt.Printf("array remove -> index 0 -> list: ")
	list.Print()
	indx = list.ArrayListSearch(26)
	fmt.Printf("arry search -> 26 ([%d]) -> list: ", indx)
	list.Print()

	indx = list.ArrayListSearch(88)
	fmt.Printf("arry search -> 88 ([%d]) -> list: ", indx)
	list.Print()

}
