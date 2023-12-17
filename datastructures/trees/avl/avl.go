package avl

import (
	"fmt"
	"strings"
)

/*
	AVL: BST w/ Height Balancing Property
		 - Includes re-balancing functions for node removals/insertions
		 - Balanced == for any node, the heights of the <=|=> subtrees
                        differ by only 0 or 1.
		 - Balance Factor: (<= H) - (=> H) (i.e., 1, 0, or -1 in AVL tree)
						   leftsub - rightsub = factor
*/

// Tree represents an AVL tree.
type Tree struct {
	Root *Node // The root node of the AVL tree.
}

// Node represents a node in an AVL tree.
type Node struct {
	Key    int   // The key or value stored in this node.
	Left   *Node // The left child node.
	Right  *Node // The right child node.
	height int   // The height of the subtree rooted at this node.
}

// Height returns the height of the AVL node.
func (n *Node) Height() int {
	if n == nil {
		return 0
	}
	return n.height
}

// Bal
// returns the balance of a node’s subtrees: 0 for a balanced node, +n if the right subtree
// is n nodes taller than the left, -n if the left subtree is n nodes taller than the right.
func (n *Node) Bal() int {
	return n.Right.Height() - n.Left.Height()
}

// NewNode creates and returns a new Node with the specified key.
func NewNode(key int) *Node {
	return &Node{Key: key, height: 1}
}

// abs returns the absolute value of an integer.
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// rotateLeft performs a left rotation on the AVL node.
func (n *Node) rotateLeft() *Node {
	fmt.Println("rotateLeft", n.Key)
	r := n.Right
	n.Right = r.Left
	r.Left = n
	n.height = max(n.Left.Height(), n.Right.Height()) + 1
	r.height = max(r.Left.Height(), r.Right.Height()) + 1
	return r
}

// rotateRight performs a right rotation on the AVL node.
func (n *Node) rotateRight() *Node {
	fmt.Println("rotateRight", n.Key)
	l := n.Left
	n.Left = l.Right
	l.Right = n
	n.height = max(n.Left.Height(), n.Right.Height()) + 1
	l.height = max(l.Left.Height(), l.Right.Height()) + 1
	return l
}

func (n *Node) rotateRightLeft() *Node {
	n.Right = n.Right.rotateRight()
	n = n.rotateLeft()
	n.height = max(n.Left.Height(), n.Right.Height()) + 1
	return n
}

func (n *Node) rotateLeftRight() *Node {
	n.Left = n.Left.rotateLeft()
	n = n.rotateRight()
	n.height = max(n.Left.Height(), n.Right.Height()) + 1
	return n
}

func (n *Node) rebalance() *Node {
	fmt.Printf("%d's rebalance factor: %d\n", n.Key, n.Bal())
	switch {
	case n.Bal() < -1 && n.Left.Bal() == -1:
		return n.rotateRight()
	case n.Bal() > 1 && n.Right.Bal() == 1:
		return n.rotateLeft()
	case n.Bal() < -1 && n.Left.Bal() == 1:
		return n.rotateLeftRight()
	case n.Bal() > 1 && n.Right.Bal() == -1:
		return n.rotateRightLeft()
	}
	return n
}

// Insert inserts a key into the AVL tree while maintaining balance.
func (tree *Tree) Insert(key int) {
	fmt.Println("inserting:", key)
	tree.Root = tree.insert(tree.Root, key)
}

// insert inserts a key into the AVL tree rooted at the given node while maintaining balance.
func (tree *Tree) insert(root *Node, key int) *Node {
	if root == nil {
		return NewNode(key)
	}

	if key < root.Key {
		root.Left = tree.insert(root.Left, key)
	} else if key > root.Key {
		root.Right = tree.insert(root.Right, key)
	} else {
		return root
	}

	root.height = max(root.Left.Height(), root.Right.Height()) + 1
	return root.rebalance()
}

// Print prints the AVL tree in a visually appealing way.
func (tree *Tree) Print() {
	lines, _, _ := tree.Root.visualize()
	title := "AVL Tree"
	fmt.Println(title)
	fmt.Printf("%s\n", strings.Repeat("-", len(title)))
	for _, line := range lines {
		fmt.Println(line)
	}
}

// visualize returns the lines that comprise a visual representation of the AVL tree,
// as well as the row and column indices of the root node in this representation.
func (n *Node) visualize() ([]string, int, int) {
	if n == nil {
		return []string{}, 0, 0
	}

	lineKey := fmt.Sprintf("%v", n.Key)
	leftLines, leftPos, leftWidth := n.Left.visualize()
	rightLines, rightPos, rightWidth := n.Right.visualize()

	middle := max(len(lineKey), leftPos+rightWidth+1+rightPos)
	pos := leftPos + middle/2

	firstLine := fmt.Sprintf("%*s%*s%*s", leftPos, "", middle-leftWidth, lineKey, rightPos, "")
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

func TestAVL(nodes, insert []int) {
	avlTree := &Tree{}

	for _, key := range nodes {
		avlTree.Insert(key)
	}
	avlTree.Print()

	if len(insert) > 0 {
		fmt.Println("insert =>", insert)
		for _, i := range insert {
			avlTree.Insert(i)
		}
		avlTree.Print()
	}

}
