package binary

import (
	"fmt"
	"strings"
)

/*
	Binary Tree: Datastructure that contains "nodes" that by definition have 2 children (left/right child).

	Leaf Node: Tree node with no children
	Internal Node: Tree node with >= 1 (at least 1 child)
	Parent Node: Tree node relationship between a node, and it's child node. The non-child is the parent.
				 - Ancestors: A node's parent + parent's parent + .... all the way to the root node.
	Root Node: Tree node with no parent (the "top" node)

	Tree DataStructure Applications:
		- Filesystems
		- Binary Space Partitioning (BSP):
		  * Implements a BSP Tree
		  * BSP Tree Stores Information from the partitioning process
		  * The process is accomplished by repeatedly partitioning a region of space into 2 parts
		    and cataloging objects contained within each region.
		- Binary Search Tree: a special type of binary tree that has an ordering property where any
							  - node's left subtree keys <= the node's key
						      - node's right subtree keys >= the node's key
		  	=> BST enables fast searching for an item
			BST Successors/Predecessors
			A BST defines an ordering among nodes, from smallest to largest.
			A BST node's successor is the node that comes after in the BST ordering,
			so in A B C, A's successor is B, and B's successor is C.

			A BST node's predecessor is the node that comes before in the BST ordering.

			If a node has a right subtree, the node's successor is that right subtree's
			leftmost child: Starting from the right subtree's root, follow left children
			until reaching a node with no left child (maybe that subtree's root itself).

			If a node doesn't have a right subtree, the node's successor is the first ancestor
			having this node in a left subtree. Another section provides an algorithm for printing
			a BST's nodes in order.
*/

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func NewBinaryTree() *Tree {
	return &Tree{}
}

func (tree *Tree) Insert(value int) {
	newNode := &Node{Value: value}

	if tree.Root == nil {
		tree.Root = newNode
	} else {
		tree.insertNode(tree.Root, newNode)
	}
}

func (tree *Tree) insertNode(current *Node, newNode *Node) {
	if newNode.Value < current.Value {
		if current.Left == nil {
			current.Left = newNode
		} else {
			tree.insertNode(current.Left, newNode)
		}
	} else {
		if current.Right == nil {
			current.Right = newNode
		} else {
			tree.insertNode(current.Right, newNode)
		}
	}
}

// Search
// Binary Tree Searching Function:
// Searching a BST in the worst case requires H + 1 comparisons, meaning O(H) comparisons,
// where H is the tree height.
//
// Ex: A tree with a root node and one child has height 1;
// the worst case visits the root and the child: 1 + 1 = 2. A major BST benefit is
// that an N-node binary tree's height may be as small as O(logN), yielding extremely fast
// searches.
//
// Ex: A 10,000 node list may require 10,000 comparisons, but a 10,000 node BST may
// require only 14 comparisons.
//
// A binary tree's height can be minimized by keeping all levels full, except possibly the last
// level. Such an "all-but-last-level-full" binary tree's height is H = log2N.
//
// Time Complexity Analysis BST Search Algorithm:
// Perfect BST Worst Case:     T(n) = O(log2N) + 1
// Non-perfect BST Worst Case: T(n) = O(H) + 1; H = Height of Tree
func (tree *Tree) Search(value int) bool {
	fmt.Printf("searching for %d\n", value)
	return tree.searchNode(tree.Root, value)
}

func (tree *Tree) searchNode(node *Node, value int) bool {
	if node == nil {
		return false
	}

	if value == node.Value {
		return true
	} else if value < node.Value {
		fmt.Printf("%d < %d (search <=)\n", value, node.Value)
		return tree.searchNode(node.Left, value)
	} else {
		fmt.Printf("%d > %d (search =>)\n", value, node.Value)
		return tree.searchNode(node.Right, value)
	}
}

// Remove removes a node with the given key from the binary search tree
func (tree *Tree) Remove(key int) {
	var parent *Node
	current := tree.Root
	for current != nil {
		if current.Value == key { // Node found
			if current.Left == nil && current.Right == nil { // Remove leaf
				if parent == nil { // Node is root
					tree.Root = nil
				} else if parent.Left == current {
					parent.Left = nil
				} else {
					parent.Right = nil
				}
			} else if current.Right == nil { // Remove node with only left child
				if parent == nil { // Node is root
					tree.Root = current.Left
				} else if parent.Left == current {
					parent.Left = current.Left
				} else {
					parent.Right = current.Left
				}
			} else if current.Left == nil { // Remove node with only right child
				if parent == nil { // Node is root
					tree.Root = current.Right
				} else if parent.Left == current {
					parent.Left = current.Right
				} else {
					parent.Right = current.Right
				}
			} else { // Remove node with two children
				successor := current.Right
				for successor.Left != nil {
					successor = successor.Left
				}
				successorData := successor.Value
				tree.BSTRemove(successor.Value) // Remove the successor
				current.Value = successorData
			}
			return // Node found and removed
		} else if current.Value < key { // Search right
			parent = current
			current = current.Right
		} else { // Search left
			parent = current
			current = current.Left
		}
	}
	return // Node not found
}

// BSTRemove removes a node with the given key from the binary search tree recursively
func (tree *Tree) BSTRemove(key int) {
	tree.Root = removeNode(tree.Root, key)
}

// removeNode recursively removes a node with the given key from the binary search tree
func removeNode(root *Node, key int) *Node {
	if root == nil {
		return nil
	}

	if key < root.Value {
		root.Left = removeNode(root.Left, key)
	} else if key > root.Value {
		root.Right = removeNode(root.Right, key)
	} else { // Node found
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		successor := findMin(root.Right)
		root.Value = successor.Value
		root.Right = removeNode(root.Right, successor.Value)
	}

	return root
}

// findMin finds the node with the minimum key in the binary search tree
func findMin(root *Node) *Node {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

func printTitle(center int) {
	title := "Binary Search Tree"
	indent := pow(2, (center+1)) - (len(title) / 2)
	fmt.Print(strings.Repeat(" ", indent))
	fmt.Println(title)
	fmt.Print(strings.Repeat(" ", indent))
	fmt.Println(strings.Repeat("-", len(title)))
}

func (tree *Tree) PrettyPrint() {
	levels := make(map[int][]string)
	maxLevel := tree.Root.fillLevels(0, levels)
	printTitle(maxLevel)
	for i := 0; i <= maxLevel; i++ {
		printLevel(levels[i], i, maxLevel)
	}
}

func (n *Node) fillLevels(currentLevel int, levels map[int][]string) int {
	if n == nil {
		return currentLevel
	}

	levels[currentLevel] = append(levels[currentLevel], fmt.Sprintf("(%d)", n.Value))

	leftMax := n.Left.fillLevels(currentLevel+1, levels)
	rightMax := n.Right.fillLevels(currentLevel+1, levels)

	if leftMax > rightMax {
		return leftMax
	}
	return rightMax
}

func printLevel(level []string, currentLevel, maxLevel int) {
	// Indentation (indent): This is the space before the first node in each level.
	// The goal is to have the leftmost node of each level appear approximately in the
	// middle of its parent node above it.
	//
	// For a complete binary tree, the number of nodes at the deepest level is
	// 2^h where h is the height of the tree (or the depth of the deepest node).
	// The initial indentation for the first level (root node) would be half of
	// that minus one for the root itself. As we go deeper into the tree, the indentation
	// decreases by half each time.
	//
	// The maxLevel - currentLevel gives the number of levels below the current one.
	// Adding 1 accounts for the current level itself. Subtracting 1 eliminates the
	// space taken by the node itself.
	indent := pow(2, (maxLevel-currentLevel+1)) - 1

	// Space Between Nodes (spaceBetween)
	// This is the space between nodes at the same level.
	// It is larger than the indentation because it needs to also
	// account for the space taken up by the nodes themselves and the
	// space needed to visually separate siblings from each other and
	// from the 'subtrees' beneath them.
	//
	// Add an extra power of 2 because we want to include the space for the 'subtree'
	// that can exist beneath the siblings at the current level.
	spaceBetween := pow(2, (maxLevel-currentLevel+2)) - 1 - 2 // Subtract 2 for parentheses

	// Print initial indent
	fmt.Print(strings.Repeat(" ", indent))

	for i, v := range level {
		fmt.Print(v)
		if i < len(level)-1 {
			// Print space between nodes at this level
			fmt.Print(strings.Repeat(" ", spaceBetween))
		}
	}
	fmt.Println()
}

func pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

// Print prints the binary tree in a visually appealing way
func (tree *Tree) Print() {
	lines, _, _ := tree.Root.visualize()
	for _, line := range lines {
		fmt.Println(line)
	}
}

// visualize returns the lines that comprise a visual representation of the binary tree,
// as well as the row and column indices of the root node in this representation.
func (n *Node) visualize() ([]string, int, int) {
	if n == nil {
		return []string{}, 0, 0
	}

	lineValue := fmt.Sprintf("%v", n.Value)
	leftLines, leftPos, leftWidth := n.Left.visualize()
	rightLines, rightPos, rightWidth := n.Right.visualize()

	middle := max(len(lineValue), leftPos+rightWidth+1+rightPos)
	pos := leftPos + middle/2

	firstLine := fmt.Sprintf("%*s%*s%*s", leftPos, "", middle-leftWidth, lineValue, rightPos, "")
	if len(leftLines) < len(rightLines) {
		leftLines = append(leftLines, make([]string, len(rightLines)-len(leftLines))...)
	} else if len(rightLines) < len(leftLines) {
		rightLines = append(rightLines, make([]string, len(leftLines)-len(rightLines))...)
	}

	var belowLines []string
	for i := 0; i < len(leftLines); i++ {
		leftPadding := abs(middle - leftWidth - len(leftLines[i]))
		rightPadding := abs(rightWidth - rightPos - len(rightLines[i]))
		connection := strings.Repeat(" ", leftPadding) + strings.Repeat(" ", leftWidth) +
			strings.Repeat(" ", rightPadding)
		belowLines = append(belowLines, leftLines[i]+connection+rightLines[i])
	}

	return append([]string{firstLine}, belowLines...), pos, middle
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func TestBST(treeNodeValues, search, insert, remove []int) {
	tree := NewBinaryTree()
	for _, v := range treeNodeValues {
		tree.Insert(v)
	}
	tree.Print()

	if len(search) > 0 {
		fmt.Println("BST: search =>", search)
		for _, s := range search {
			tree.Search(s)
		}
	}
	if len(insert) > 0 {
		fmt.Println("BST: insert =>", insert)
		for _, i := range insert {
			tree.Insert(i)
		}
		tree.Print()
	}
	if len(remove) > 0 {
		fmt.Println("BST: remove =>", remove)
		for _, r := range remove {
			tree.BSTRemove(r)
		}
		tree.Print()
	}
}
