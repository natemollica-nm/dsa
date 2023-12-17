package datastructures

import (
	"fmt"
	"github.com/ryanuber/columnize"
	"strings"
)

// HashNode represents a key-value pair stored in the hash table.
type HashNode struct {
	key   int
	value HashNodeValue
	next  *HashNode
}

type HashNodeValue interface{}

// HashTable represents a basic hash table with separate chaining.
type HashTable struct {
	table    []*HashNode
	capacity int
	emptied  []bool // Tracks emptied buckets
}

// NewHashTable creates a new hash table with a specified capacity.
func NewHashTable(capacity int) *HashTable {
	return &HashTable{
		table:    make([]*HashNode, capacity),
		capacity: capacity,
		emptied:  make([]bool, capacity),
	}
}

// HashInsert inserts a key-value pair into the hash table.
func (ht *HashTable) HashInsert(key int, value HashNodeValue) {
	hash := key % ht.capacity
	node := &HashNode{key: key, value: value}

	if ht.table[hash] == nil {
		ht.table[hash] = node
	} else {
		// Handle collision by chaining.
		current := ht.table[hash]
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
}

// HashInsertLinearProbing inserts a key-value pair into the hash table with linear probing.
func (ht *HashTable) HashInsertLinearProbing(key int, value HashNodeValue) {
	index := key % ht.capacity

	// Linear probing until an empty slot is found.
	for ht.table[index] != nil {
		index = (index + 1) % ht.capacity
	}

	ht.table[index] = &HashNode{key: key, value: value}
}

// HashSearch searches for a key in the hash table and returns its value (if found).
// It also returns the number of elements checked during the search.
func (ht *HashTable) HashSearch(key int) (HashNodeValue, bool, int) {
	hash := key % ht.capacity
	current := ht.table[hash]
	elementsChecked := 0

	for current != nil {
		elementsChecked++
		if current.key == key {
			return current.value, true, elementsChecked
		}
		current = current.next
	}

	return nil, false, elementsChecked
}

func (ht *HashTable) HashSearchLinearProbe(key int) (HashNodeValue, bool, int, []int) {
	index := key % ht.capacity
	current := ht.table[index]
	elementsChecked := 0
	bucketsProbed := []int{index}

	for (current != nil || !ht.emptied[index]) && (len(bucketsProbed) < ht.capacity) {
		elementsChecked++
		if current.key == key {
			return current.value, true, elementsChecked, bucketsProbed
		}
		bucketsProbed = append(bucketsProbed, index) // Record the bucket probed
		// If the bucket is empty after removal or has been empty since the start, continue probing the next bucket.
		if ht.emptied[index] {
			index = (index + 1) % ht.capacity
			current = ht.table[index]
		} else {
			return nil, false, elementsChecked, bucketsProbed
		}
	}

	return nil, false, elementsChecked, bucketsProbed
}

// HashRemove removes a key-value pair from the hash table.
func (ht *HashTable) HashRemove(key int) {
	hash := key % ht.capacity
	current := ht.table[hash]
	var prev *HashNode

	for current != nil {
		if current.key == key {
			if prev == nil {
				ht.table[hash] = current.next
			} else {
				prev.next = current.next
			}
			return
		}
		prev = current
		current = current.next
	}
}

// HashRemoveLinearProbe removes a key-value pair from the hash table with linear probing.
func (ht *HashTable) HashRemoveLinearProbe(key int) {
	index := key % ht.capacity
	current := ht.table[index]
	var prev *HashNode

	for current != nil {
		if current.key == key {
			if prev == nil {
				ht.table[index] = current.next
				ht.emptied[index] = true // Mark the bucket as emptied
			} else {
				prev.next = current.next
				ht.emptied[index] = true // Mark the bucket as emptied
			}
			return
		}
		prev = current
		current = current.next
	}
}

// Print prints the contents of the hash table.
func (ht *HashTable) Print() {
	result := []string{"\x1fBucket\x1fValues\x1fEmptied\x1f"}
	for i, node := range ht.table {
		bucketID := i
		var keyValues []string
		for node != nil {
			keyValues = append(keyValues, fmt.Sprintf("(%v)", node.value))
			node = node.next
		}
		if len(keyValues) == 0 {
			keyValues = append(keyValues, fmt.Sprintf("nil"))
		}
		bucketValues := strings.Join(keyValues, " -> ")
		result = append(result, fmt.Sprintf("\x1f%d\x1f%s\x1f%v\x1f", bucketID, bucketValues, ht.emptied[i]))
	}
	output := columnize.Format(result, &columnize.Config{Delim: string([]byte{0x1f}), Glue: " "})
	fmt.Println(output)
}

func TestHashTable() {
	ht := NewHashTable(10)
	values := []int{20, 12, 95, 47, 57}
	for _, v := range values {
		ht.HashInsertLinearProbing(v, v)
	}
	ht.emptied[1] = true
	ht.emptied[4] = true
	ht.emptied[8] = true
	fmt.Println("initial: ")
	ht.Print()

	ht.HashRemoveLinearProbe(95)
	fmt.Println("HashRemoveLinearProbe (95): ")
	ht.Print()
	//test := []int{13, 37}
	//for _, val := range test {
	//	ht.HashInsertLinearProbing(val, val)
	//}
	//ht.Print()

	// Perform searches and count comparisons.
	searchItems := []int{95, 57}
	for _, searchKey := range searchItems {
		value, found, elements, buckets := ht.HashSearchLinearProbe(searchKey)
		if found {
			fmt.Printf("(found) Key (%d) -> (%d) Value. Comparisons: %d | Buckets: %d\n", searchKey, value, elements, buckets)
		} else {
			fmt.Printf("(not found) Key (%d) -> (%v) returned. Comparisons: %d | Buckets: %d\n", searchKey, value, elements, buckets)
		}
	}
}
