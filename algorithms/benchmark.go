package algorithms

import (
	"fmt"
	"math/rand"
	"time"
)

func prettyPrintMap(m map[string]string) {
	for key, value := range m {
		fmt.Printf("  · %s: %s\n", key, value)
	}
}

func generateRandomArray(size int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(1000000) // Adjust the range as needed
	}
	return arr
}

func benchmarkSortAlgorithm(arrSize int, sortFunc func([]int)) time.Duration {
	arr := generateRandomArray(arrSize)
	startTime := time.Now()
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sortFunc(sortedArr)
	return time.Since(startTime)
}

func BenchmarkSortAlgorithms() {
	sortingAlgorithms := []SortingAlgorithm{
		{
			Name:     "QuickSort",
			SortFunc: QuickSort,
			TimeComplexity: map[string]string{
				"Best Case":  "Ω(nlogn)",
				"Avg Case":   "θ(nlogn)",
				"Worst Case": "O(n²)",
			},
		},
		{
			Name:     "HeapSort",
			SortFunc: HeapSort,
			TimeComplexity: map[string]string{
				"Best Case":  "Ω(nlogn)",
				"Avg Case":   "θ(nlogn)",
				"Worst Case": "O(nlogn)",
			},
		},
		{
			Name:     "RadixSort",
			SortFunc: RadixSort,
			TimeComplexity: map[string]string{
				"Best Case":  "Ω(nk)",
				"Avg Case":   "θ(nk)",
				"Worst Case": "O(nk)",
				"k":          "largest num of digits",
			},
		},
		{
			Name:     "MergeSort",
			SortFunc: MergeSort,
			TimeComplexity: map[string]string{
				"Best Case":  "Ω(nlogn)",
				"Avg Case":   "θ(nlogn)",
				"Worst Case": "O(nlogn)",
			},
		},
		{
			Name:     "ShellSort",
			SortFunc: ShellSort,
			TimeComplexity: map[string]string{
				"Best Case":  "Ω(nlog²n)",
				"Avg Case":   "θ(nlog²n) <= between => θ(n²)",
				"Worst Case": "O(n²)",
			},
		},
		{
			Name:     "InsertionSort",
			SortFunc: InsertionSort,
			TimeComplexity: map[string]string{
				"Best Case":  "Ω(n²)",
				"Avg Case":   "θ(n²)",
				"Worst Case": "O(n²)",
			},
		},
		{
			Name:     "SelectionSort",
			SortFunc: SelectionSort,
			TimeComplexity: map[string]string{
				"Best Case":  "Ω(n²)",
				"Avg Case":   "θ(n²)",
				"Worst Case": "O(n²)",
			},
		},
	}

	inputSizes := []int{1000, 10000, 50000, 100000, 150000, 200000, 250000, 300000}

	// Run the benchmarks
	for _, algo := range sortingAlgorithms {
		fmt.Println("Algorithm:", algo.Name)
		fmt.Println("Time Complexity:")
		prettyPrintMap(algo.TimeComplexity)
		for _, size := range inputSizes {
			duration := benchmarkSortAlgorithm(size, algo.SortFunc)
			fmt.Printf("Input Size %d: %s\n", size, duration)
		}
		fmt.Println()
	}
}
