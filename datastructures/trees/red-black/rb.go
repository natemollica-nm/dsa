package red_black

import (
	"fmt"
	"strings"
)

/*
	Red-Black Tree: BST with 2 node types; red and black
					- has supporting functions to ensure tree is balanced upon node insertion/removal
					- O(logN) Height Maintained by following the below requirements:
					  - Adhere to BST Ordering Property (left child < parent < right child)
					  - Every node is colored red || black
					  - Root Node is Black
					  - Red Node Children are NOT red
					  - Null children are Black Leaf Nodes
					  - All paths from a node => any null leaf descendant node,
                        must have == num. of black nodes
*/

// Node represents a node in the red-black tree.
type Node struct {
	Value       int
	Left, Right *Node
	Color       bool // true for red, false for black
}

// RedBlackTree represents a red-black tree.
type RedBlackTree struct {
	Root *Node
}

const (
	red   = true
	black = false
)

// NewRedBlackTree creates a new red-black tree.
func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{}
}

// Insert inserts a value into the red-black tree while maintaining its properties.
func (t *RedBlackTree) Insert(value int) {
	t.Root = insert(t.Root, value)
	t.Root.Color = black // Ensure the root is black.
}

// insert recursively inserts a value into the red-black tree.
func insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{Value: value, Color: red}
	}

	if value < root.Value {
		root.Left = insert(root.Left, value)
	} else if value > root.Value {
		root.Right = insert(root.Right, value)
	}

	// Perform rotations and recoloring to maintain red-black tree properties.
	if isRed(root.Right) && !isRed(root.Left) {
		root = rotateLeft(root)
	}
	if isRed(root.Left) && isRed(root.Left.Left) {
		root = rotateRight(root)
	}
	if isRed(root.Left) && isRed(root.Right) {
		flipColors(root)
	}

	return root
}

// isRed returns true if a node is red, and false otherwise.
func isRed(node *Node) bool {
	if node == nil {
		return false // Null nodes are considered black.
	}
	return node.Color == red
}

// rotateLeft performs a left rotation on the node and returns the new root.
func rotateLeft(node *Node) *Node {
	x := node.Right
	node.Right = x.Left
	x.Left = node
	x.Color = node.Color
	node.Color = red
	return x
}

// rotateRight performs a right rotation on the node and returns the new root.
func rotateRight(node *Node) *Node {
	x := node.Left
	node.Left = x.Right
	x.Right = node
	x.Color = node.Color
	node.Color = red
	return x
}

// flipColors inverts the colors of a node and its children.
func flipColors(node *Node) {
	node.Color = !node.Color
	node.Left.Color = !node.Left.Color
	node.Right.Color = !node.Right.Color
}

// InorderTraversal returns the values of the red-black tree in sorted order.
func (t *RedBlackTree) InorderTraversal() []int {
	var result []int
	inorder(t.Root, &result)
	return result
}

// inorder recursively traverses the red-black tree in-order.
func inorder(node *Node, result *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, result)
	*result = append(*result, node.Value)
	inorder(node.Right, result)
}

// Print prints the red-black tree in a visually appealing way with colors.
func (t *RedBlackTree) Print() {
	lines, _, _ := t.Root.visualizeWithColor()
	fmt.Println("Red-Black Tree:")
	for _, line := range lines {
		fmt.Println(line)
	}
}

// visualizeWithColor returns the lines that comprise a visual representation of the red-black tree
// with color information, as well as the row and column indices of the root node in this representation.
func (n *Node) visualizeWithColor() ([]string, int, int) {
	if n == nil {
		return []string{}, 0, 0
	}

	colorTag := "B" // Black by default
	if n.Color == red {
		colorTag = "R"
	}

	lineValue := fmt.Sprintf("%v%s", n.Value, colorTag)
	leftLines, leftPos, leftWidth := n.Left.visualizeWithColor()
	rightLines, rightPos, rightWidth := n.Right.visualizeWithColor()

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

func TestRedBlackTree(values, insert []int) {
	tree := NewRedBlackTree()

	for _, value := range values {
		tree.Insert(value)
	}

	tree.Print()
	fmt.Printf("%v\n", tree.InorderTraversal())

	if len(insert) > 0 {
		for _, i := range insert {
			tree.Insert(i)
		}
		fmt.Printf("Inorder Traversal (insert => %v): %v", insert, tree.InorderTraversal())
	}
}
