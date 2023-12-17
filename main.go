package main

import algo "dsa/algorithms"

func main() {
	//algo.BenchmarkSortAlgorithms()
	// listData := []ds.NodeData{96, 12, 59}
	// ds.TestDoublyLinkedList(listData)
	//testItems := []datastructures.DequeValue{46, 74}
	//datastructures.TestDeque(testItems)
	// testData := []datastructures.ArrayListData{19, 26, 47}
	// datastructures.TestHashTable()
	// var search []int
	// var insert []int
	// remove := []int{20}
	// binary.TestBST([]int{40, 20, 50, 10, 30, 60}, search, insert, remove)
	//inputs := map[string][]int{
	//	"init-unsorted-array": {51, 26, 22, 34, 56, 30, 42},
	//}
	// keys := []int{22, 51, 98, 34, 85, 23, 13}
	// avl.TestAVL(keys, insert)
	// red_black.TestRedBlackTree(keys, insert)
	//for key, arr := range inputs {
	//	fmt.Printf("%s:\n", key)
	//	maxheap.TestMaxHeap(arr, []int{}, false, true)
	//}
	algo.TestDijkstraAlgorithm()
}
