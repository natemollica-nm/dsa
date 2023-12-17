package maxheap

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
	MaxHeap: Complete Binary Tree with the properties --
		  	 - every node's key >= children's node keys
			 - root always has the max key in the entire tree
*/

type MaxHeap struct {
	heap []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func (h *MaxHeap) Insert(value int) {
	h.heap = append(h.heap, value)
	h.PercolateUp(len(h.heap) - 1)
}

func (h *MaxHeap) Remove() int {
	if len(h.heap) == 0 {
		return 0
	}

	root := h.heap[0]
	last := h.heap[len(h.heap)-1]
	h.heap[0] = last
	h.heap = h.heap[:len(h.heap)-1]

	h.PercolateDown(0)
	return root
}

func (h *MaxHeap) PercolateUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.heap[index] <= h.heap[parentIndex] {
			break
		}
		h.heap[index], h.heap[parentIndex] = h.heap[parentIndex], h.heap[index]
		index = parentIndex
	}
}

func (h *MaxHeap) PercolateDown(index int) {
	for {
		leftChildIdx := 2*index + 1
		rightChildIdx := 2*index + 2
		largest := index

		if leftChildIdx < len(h.heap) && h.heap[leftChildIdx] > h.heap[largest] {
			largest = leftChildIdx
		}
		if rightChildIdx < len(h.heap) && h.heap[rightChildIdx] > h.heap[largest] {
			largest = rightChildIdx
		}

		if largest == index {
			break
		}

		h.heap[index], h.heap[largest] = h.heap[largest], h.heap[index]
		index = largest
	}
}

func (h *MaxHeap) Sort() *MaxHeap {
	sorted := NewMaxHeap()
	for len(h.heap) > 0 {
		sorted.heap = append(sorted.heap, h.Remove())
	}
	return sorted
}

// visualize returns the lines that comprise a visual representation of the MaxHeap,
// with proper tree-like spacing.
func (h *MaxHeap) visualize(index int, prefix string) []string {
	if index >= len(h.heap) {
		return []string{}
	}

	lineValue := fmt.Sprintf("%d[%d]", h.heap[index], index)
	leftLines := h.visualize(2*index+1, prefix+"L ")
	rightLines := h.visualize(2*index+2, prefix+"R ")

	var result []string
	result = append(result, prefix+lineValue)

	maxLines := max(len(leftLines), len(rightLines))
	for i := 0; i < maxLines; i++ {
		var leftLine, rightLine string
		if i < len(leftLines) {
			leftLine = leftLines[i]
		}
		if i < len(rightLines) {
			rightLine = rightLines[i]
		}
		result = append(result, fmt.Sprintf("%s%s  %s", prefix, leftLine, rightLine))
	}

	return result
}

// prettyPrint prints the heap in a tree-like structure
func (h *MaxHeap) prettyPrint() {
	const space = 4
	var levels = int(math.Floor(math.Log2(float64(len(h.heap))))) + 1
	lastLevel := int(math.Pow(2, float64(levels-1)))
	index := 0
	for level := 0; level < levels; level++ {
		prefix := strings.Repeat(" ", (lastLevel-1)/(int(math.Pow(2, float64(level))))*space)
		line := prefix
		for i := 0; i < int(math.Pow(2, float64(level))); i++ {
			if index >= len(h.heap) {
				break
			}
			val := strconv.Itoa(h.heap[index])
			line += val + strings.Repeat(" ", (lastLevel*2)/(int(math.Pow(2, float64(level))))*space-len(val))
			index++
		}
		fmt.Println(line)
	}
}

func TestMaxHeap(values, insert []int, remove, heapsort bool) {
	heap := &MaxHeap{values}
	heap.prettyPrint()

	if heapsort {
		heap = NewMaxHeap()
		for _, value := range values {
			heap.Insert(value)
		}
		heap = heap.Sort()
	}
	fmt.Printf("Post HeapSort:\n")
	heap.prettyPrint()

	if remove {
		removed := heap.Remove()
		fmt.Printf("After removing the root: %d\n", removed)
		heap.prettyPrint()
	}

	if len(insert) > 0 {
		for _, i := range insert {
			heap.Insert(i)
		}
		fmt.Printf("After inserting: %v\n", insert)
		heap.prettyPrint()
	}

}
