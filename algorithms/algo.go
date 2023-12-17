package algorithms

import (
	"fmt"
	"math"
)

type SortingAlgorithm struct {
	Name           string
	SortFunc       func(arr []int)
	TimeComplexity map[string]string
}

// HeapSort sorts an array using the HeapSort algorithm.
func HeapSort(arr []int) {
	n := len(arr)

	// Build a max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Extract elements from the heap one by one
	for i := n - 1; i > 0; i-- {
		// Swap the root (maximum element) with the last element
		arr[0], arr[i] = arr[i], arr[0]

		// Call heapify on the reduced heap
		heapify(arr, i, 0)
	}
}

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	// Find the largest element among the root, left child, and right child
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// If the largest element is not the root, swap them and continue to heapify
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

// MergeSort performs the merge sort algorithm to sort an integer slice.
//
// MergeSort Runtime (Time Complexity)
// The merge sort algorithm's runtime is O(N log N). Merge sort divides the input in half until
// a list of 1 element is reached, which requires log N partitioning levels. At each level, the
// algorithm does about N comparisons selecting and copying elements from the left and right partitions,
// yielding N * log N comparisons.
//
// Merge sort requires O(N) additional memory elements for the temporary array of merged elements.
// For the final merge operation, the temporary list has the same number of elements as the input.
// Some sorting algorithms sort the list elements in place and require no additional memory, but are
// more complex to write and understand.
//
// To allocate the temporary array, the Merge() function dynamically allocates the array. mergedNumbers
// is a pointer variable that points to the dynamically allocated array, and new int[mergedSize] allocates
// the array with mergedSize elements. Alternatively, instead of allocating the array within the Merge()
// function, a temporary array with the same size as the array being sorted can be passed as an argument.
func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	mid := len(arr) / 2
	left := make([]int, mid)
	right := make([]int, len(arr)-mid)

	copy(left, arr[:mid])
	copy(right, arr[mid:])

	MergeSort(left)
	MergeSort(right)

	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
}

// RadixSort :
//   - sorting algorithm designed specifically for integers
//   - makes use of "buckets" and is considered a type of bucket sort
//   - bucket => a collection of integer values that all share a particular digit value (i.e., 57, 97, 77, and 17 all have 7 as 1s digit)
//
// Sorting Signed Integers
// The above radix sort algorithm correctly sorts arrays of non-negative integers.
// But if the array contains negative integers, the above algorithm would sort by
// absolute value, so the integers are not correctly sorted. A small extension to
// the algorithm correctly handles negative integers.
//
// In the extension, before radix sort completes, the algorithm allocates two buckets,
// one for negative integers and the other for non-negative integers. The algorithm
// iterates through the array in order, placing negative integers in the negative
// bucket and non-negative integers in the non-negative bucket. The algorithm then
// reverses the order of the negative bucket and concatenates the buckets to yield a
// sorted array.
func RadixSort(arr []int) {
	maxDigits := RadixGetMaxLength(arr)

	// Separate negative and non-negative integers
	var negativeArr []int
	var nonNegativeArr []int
	for _, num := range arr {
		if num < 0 {
			negativeArr = append(negativeArr, -num)
		} else {
			nonNegativeArr = append(nonNegativeArr, num)
		}
	}
	// inline radixSort: sorts an integer slice of non-negative numbers using Radix Sort.
	radixSort := func(subArray []int, maxDigs int) {
		for i := 0; i < maxDigs; i++ {
			buckets := make([][]int, 10)

			for _, num := range subArray {
				digit := (num / int(math.Pow(10, float64(i)))) % 10
				buckets[digit] = append(buckets[digit], num)
			}

			index := 0
			for j := 0; j < 10; j++ {
				for _, num := range buckets[j] {
					subArray[index] = num
					index++
				}
			}
		}
	}

	// Sort negative and non-negative parts separately
	radixSort(negativeArr, maxDigits)
	radixSort(nonNegativeArr, maxDigits)

	// Combine sorted arrays back with negative values restored
	arrIndex := 0
	for i := len(negativeArr) - 1; i >= 0; i-- {
		arr[arrIndex] = -negativeArr[i]
		arrIndex++
	}
	for _, num := range nonNegativeArr {
		arr[arrIndex] = num
		arrIndex++
	}
}

// RadixGetMaxLength returns the number of digits in the maximum element in the array.
func RadixGetMaxLength(arr []int) int {
	max := arr[0]
	for _, num := range arr {
		if num < 0 {
			num = -num // Make negative numbers positive for digit counting
		}
		if num > max {
			max = num
		}
	}
	return RadixGetLength(max)
}

// RadixGetLength returns the number of digits in a given number.
func RadixGetLength(num int) int {
	if num == 0 {
		return 1
	}
	return int(math.Log10(float64(num))) + 1
}

func ShellSort(arr []int) {
	n := len(arr)
	gap := (len(arr) - 1) / 2

	for gap > 0 {
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i

			// Move elements of arr[0..i-gap] that are greater than temp
			// to positions ahead of their current position
			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}

			// Place temp (the current element) in its correct position
			arr[j] = temp
		}
		// Reduce the gap for the next iteration
		gap /= 10
	}
}

func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		// Find index of smallest remaining element
		indexSmallest := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[indexSmallest] {
				indexSmallest = j
			}
		}
		// Swap numbers[i] and numbers[indexSmallest
		arr[i], arr[indexSmallest] = arr[indexSmallest], arr[i]
	}
}

// InsertionSort
// Runtime for nearly sorted input for each outer loop execution,
// if the element is already in sorted position, only a single comparison
// is made.
//
// Each element not in sorted position requires at most N comparisons.
// If there are a constant number, C, of unsorted elements, sorting the N - C
// sorted elements requires one comparison each, and sorting the C unsorted elements
// requires at most N comparisons each.
//
// The runtime for nearly sorted inputs is O((N - C) * 1 + C * N) = O(N).
func InsertionSort(arr []int) {
	n := len(arr)

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		// Move elements of arr[0..i-1] that are greater than key
		// to one position ahead of their current position
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		// Place the key in its correct position
		arr[j+1] = key
	}
}

// QuickSort
// Runtime is typically O(N log N).
// QuickSort has several partitioning levels, the first level dividing the
// input into 2 parts, the second into 4 parts, the third into 8 parts, etc.
// At each level, the algorithm does at most N comparisons moving the lowIndex and
// highIndex indices.
//
// If the pivot yields two equal-sized parts, then there will be log N levels,
// requiring the N * log N comparisons.
func QuickSort(arr []int) {
	low := 0
	high := len(arr) - 1
	if low < high {
		// Partition the array and get the index of the pivot element
		pivotIndex := partition(arr, low, high)

		// Recursively sort the elements in the left and right partitions
		QuickSort(arr[:pivotIndex])
		QuickSort(arr[pivotIndex+1 : high])
	}
}

func partition(arr []int, low int, high int) int {
	// Choose the middle element as the pivot
	mid := low + (high-low)/2
	pivot := arr[mid]
	var done bool
	for !done {
		// Increment low while numbers[lowIndex] < pivot
		for arr[low] < pivot {
			low++
		}
		// Decrement high while pivot < numbers[highIndex]
		for pivot < arr[high] {
			high--
		}
		// If zero or one elements remain, then all numbers are
		// partitioned. Return high.
		if low >= high {
			done = true
		} else {
			// Swap arr[lowIndex] and arr[highIndex]
			arr[low], arr[high] = arr[high], arr[low]
			// Finish out incrementing low and high indx
			low++
			high--
		}
	}
	return high
}

func Fibonacci(i int) int {
	if i == 0 {
		return 0
	} else if i == 1 {
		return 1
	}
	return Fibonacci(i-1) + Fibonacci(i-2)
}

// DynamicFibonacci
// Numbers are computed with an iterative approach that stores the 2 previous terms
// in memory, instead of making recursive calls that recompute the same term many times over.
//
// Benefit: each term is computed only once
// Example:
//
//	 3 loop iterations are needed to compute DynamicFibonacci(4).
//			==> The previous 2 terms are stored
//			==> no term is computed more than once.
func DynamicFibonacci(term int) int {
	if term == 0 {
		return 0
	}
	prev, current := 0, 1
	// iteration will not go beyond the input term here...
	for i := 1; i < term; i++ {
		next := prev + current
		prev = current
		current = next
	}
	return current
}

// LongestCommonSubstring
//  1. Initialize a 2D matrix (dp) to store the lengths of common substrings.
//  2. Iterate through the input strings str1 and str2
//  3. Fill the 2D matrix based on character matches.
//     => if match found: we increment the length of the common substring
//     => if the len(common_string) > previously matched len(common_string), update as new longest.
//  4. Update the maxLength and endIndex accordingly.
//
// Time Complexity:  T(m, n) = O(m * n) | for each m chars of str1, n matrix entries are computed
// Space Complexity: S(m, n) = O(m * n) | m * n int matrix is built
func LongestCommonSubstring(str1, str2 string) string {
	// Initialize a 2D matrix to store the length of common substrings
	m, n := len(str1), len(str2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Variables to keep track of the length and ending position of the longest common substring
	maxLength, endIndex := 0, 0

	// Fill the DP matrix
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxLength {
					maxLength = dp[i][j]
					endIndex = i
				}
			} else {
				dp[i][j] = 0
			}
		}
	}

	// Extract the longest common substring, if at least one
	// substring was found; otherwise return empty string (no match).
	if maxLength == 0 {
		return ""
	}
	startIndex := endIndex - maxLength
	return str1[startIndex:endIndex]
}

// DynamicLongestCommonSubString
// Uses a 1D array dp of size m+1 to store the current and previous rows of the DP matrix.
// We iterate through the strings similarly to the previous version, but we only keep track of
// the current and previous row values, which significantly reduces the memory usage.
//
// The longest common substring algorithm can be implemented such that only the
// previously computed row and the largest matrix entry's location and value are
// stored in memory.
// With this optimization, the space complexity is reduced to O(m).
//
// Time Complexity: T(m,n):  O(m * n)
// Space Complexity: S(m,n): O(m) || O(min(m,n))
func DynamicLongestCommonSubString(str1, str2 string) string {
	// Ensure str1 is the shorter string
	if len(str1) > len(str2) {
		str1, str2 = str2, str1
	}
	m, n := len(str1), len(str2)

	// Initialize a 1D array to store the current and previous rows of the DP matrix
	dp := make([]int, m+1)

	// Variables to keep track of the length and ending position of the longest common substring
	maxLength := 0
	endIndex := 0

	// Fill the DP array
	for j := 1; j <= n; j++ {
		prev := 0
		for i := 1; i <= m; i++ {
			temp := dp[i]
			if str1[i-1] == str2[j-1] {
				dp[i] = prev + 1
				if dp[i] > maxLength {
					maxLength = dp[i]
					endIndex = i
				}
			} else {
				dp[i] = 0
			}
			prev = temp
		}
	}

	// Extract the longest common substring
	if maxLength == 0 {
		return ""
	}
	startIndex := endIndex - maxLength
	return str1[startIndex:endIndex]
}

func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func discreteVerify(n int) {
	sum := 1
	for x := 1; x <= n; x++ {
		sum = sum * x
	}
	fmt.Printf("!%d = %d * (%d-1) * (%d-2) ... = %d", n, n, n, n, sum)
}

func AnswerBell(bell string) int {
	switch bell {
	case "1/3":
		discreteVerify(1)
		return factorial(1)
	case "2/3":
		discreteVerify(2)
		return factorial(2)
	case "Standard":
		discreteVerify(3)
		return factorial(3)
	case "Full":
		discreteVerify(4)
		return factorial(4)
	case "Flank":
		discreteVerify(5)
		return factorial(5)
	default:
		return -1
	}
}

func LinearSearch(arr []int, target int) int {
	fmt.Printf("[linear-search] arr: %v | target: %d\n", arr, target)
	for i, val := range arr {
		fmt.Printf("[linear-search] checking: %d\n", val)
		if val == target {
			return i // Return the index where the target element is found
		}
	}
	return -1 // Return -1 if the target element is not found in the array
}

func BinarySearch(arr []int, target int) int {
	return binarySearchRecursive(arr, target, 0, len(arr)-1)
}

func binarySearchRecursive(arr []int, target int, left int, right int) int {
	if left > right {
		return -1 // Element not found
	}

	mid := left + (right-left)/2
	fmt.Printf("%v | %d | %d | %d\n", arr[mid], left, mid, right)
	if arr[mid] == target {
		fmt.Printf("found target at %d\n", mid)
		return mid // Element found at index mid
	} else if arr[mid] < target {
		fmt.Printf("right search: %d => %d\n", arr[mid+1], arr[right])
		return binarySearchRecursive(arr, target, mid+1, right) // Search right subarray
	} else {
		fmt.Printf("left search: %d => %d\n", arr[left], arr[mid-1])
		return binarySearchRecursive(arr, target, left, mid-1) // Search left subarray
	}
}
