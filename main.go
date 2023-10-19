package main

import (
	"fmt"
	// "math/rand"
)

// //////////////////////////////////////////////////////////////////////////////
// Node
// //////////////////////////////////////////////////////////////////////////////
type Node struct {
	left          *Node
	right         *Node
	val           int64
	height        int64
	balanceFactor int64
}

func newNode() *Node {
	return &Node{
		left:          nil,
		right:         nil,
		val:           0,
		balanceFactor: 0,
		height:        0,
	}
}

func (node *Node) insert(val int64) *Node {
	if node == nil {
		return &Node{
			val:           val,
			left:          nil,
			right:         nil,
			balanceFactor: 0,
			height:        1,
		}
	}

	if node.val == val {
		return node
	}

	if node.val > val {
		node.left = node.left.insert(val)
	} else {
		node.right = node.right.insert(val)
	}

  fmt.Println("Testing height difference")
  fmt.Printf("%+v\n", node)
  fmt.Printf("%+v\n", node.right)
  fmt.Printf("Left Height fh: %d\n", findHeight(node.left))
  fmt.Printf("Left Height gh: %d\n", node.left.getHeight())
  fmt.Printf("Right Height fh: %d\n", findHeight(node.right))
  fmt.Printf("Right Height gh: %d\n", node.right.getHeight())
	node.height = 1 + max(findHeight(node.left), findHeight(node.right))
	// node.height = 1 + max(node.left.getHeight(), node.right.getHeight())

	fmt.Printf("Node balance factor: %d\n", node.bf())
	return node.rebalance()
}

func (node *Node) bf() int64 {
	return node.right.getHeight() - node.left.getHeight()
}

// Deprecated
// func (node *Node) addNode(val int64) int64 {
// 	if node == nil {
// 		return 0
// 	}
//
// 	if val == node.val {
// 		return 0
// 	}
//
// 	if val <= node.val {
// 		if node.left == nil {
// 			newNode := newNode()
// 			newNode.val = val
// 			node.left = newNode
//
// 			return 1
// 		}
//
// 		new_left_height := node.left.addNode(val)
// 		right_height := findHeight(node.right)
//
// 		node.height = findHeight(node)
//
// 		node.balanceFactor = right_height - new_left_height
//
// 		return node.height
// 	} else {
//
// 		if node.right == nil {
// 			newNode := newNode()
// 			newNode.val = val
// 			node.right = newNode
// 			return 1
// 		}
//
// 		new_right_height := node.right.addNode(val)
// 		left_height := findHeight(node.left)
//
// 		node.balanceFactor = (new_right_height - left_height)
// 		node.height = findHeight(node)
//
// 		return node.height
// 	}
// }

func (n *Node) getHeight() int64 {
	if n == nil {
		return 0
	}

	return n.height
}

func findHeight(node *Node) int64 {
	if node == nil {
		return 0
	}

  if node.left == nil && node.right == nil {
    return 1
  }

	lh := findHeight(node.left)
	rh := findHeight(node.right)

	m := 1 + max(lh, rh)

	return m
}

func (n *Node) rotateLeft() *Node {
	fmt.Printf("rotateLeft %d\n", n.val)

	r := n.right

	n.right = r.left

	r.left = n

	n.height = max(n.left.getHeight(), n.right.getHeight()) + 1
	r.height = max(r.left.getHeight(), r.right.getHeight()) + 1

	return r
}

func (n *Node) rotateRight() *Node {
	fmt.Printf("rotateRight %d\n", n.val)

	l := n.left

	n.left = l.right

	l.right = n

	n.height = max(n.left.getHeight(), n.right.getHeight()) + 1
	l.height = max(l.left.getHeight(), l.right.getHeight()) + 1

	return l
}

func (n *Node) rotateRightLeft() *Node {
	n.right = n.right.rotateRight()
	n = n.rotateLeft()
	n.height = max(n.left.getHeight(), n.right.getHeight()) + 1
	return n
}

func (n *Node) rotateLeftRight() *Node {
	n.left = n.left.rotateLeft()
	n = n.rotateRight()
	n.height = max(n.left.getHeight(), n.right.getHeight()) + 1
	return n
}

func (n *Node) rebalance() *Node {
	fmt.Printf("rebalance %d: %d\n", n.val, n.balanceFactor)
	switch {
	case n.bf() < -1 && n.left.bf() == -1:
		return n.rotateRight()
	case n.bf() > 1 && n.right.bf() == 1:
		return n.rotateLeft()
	case n.bf() < -1 && n.left.bf() == 1:
		return n.rotateLeftRight()
	case n.bf() > 1 && n.right.bf() == -1:
		return n.rotateRightLeft()
	}
	return n
}

// //////////////////////////////////////////////////////////////////////////////
// Tree
// //////////////////////////////////////////////////////////////////////////////
type BinaryTree struct {
	root *Node
}

func newBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (t *BinaryTree) insert(val int64) {
	t.root = t.root.insert(val)

	fmt.Printf("Root balanceFactor: %d\n", t.root.balanceFactor)
	if t.root.balanceFactor < -1 || t.root.balanceFactor > 1 {
		fmt.Println("Going to rebalance")
		t.rebalance()
	}
}

func (t *BinaryTree) rebalance() {
	if t == nil || t.root == nil {
		return
	}

	t.root = t.root.rebalance()
}

func (t *BinaryTree) find(val int64) *Node {
	if t == nil || t.root == nil {
		return nil
	}

	return findRecursive(t.root, val)
}

func findRecursive(node *Node, val int64) *Node {
	if node == nil {
		return nil
	}

	if val == node.val {
		return node
	} else if val < node.val {
		return findRecursive(node.left, val)
	} else {
		return findRecursive(node.right, val)
	}
}

func (t *BinaryTree) remove(val int64) {
	if t == nil || t.root == nil {
		return
	}

	removeRecursive(t.root, val)

	return
}

func removeRecursive(node *Node, val int64) *Node {
	if node == nil {
		return nil
	}

	if node.val == val {
		return nil
	}

	if val < node.val {
		node.left = removeRecursive(node.left, val)
		return node
	} else {
		node.right = removeRecursive(node.right, val)
		return node
	}
}

func (t *BinaryTree) print() {
	if t == nil || t.root == nil {
		println("Empty tree")
		return
	}

	counter := 1
	printRecursive(t.root, int64(counter))
}

func printRecursive(node *Node, height int64) {
	if node == nil {
		return
	}

	for i := 0; i < int(height); i++ {
		fmt.Print("-")
	}
	fmt.Printf("%v\n", node.val)

	height += 1

	printRecursive(node.left, height)
	printRecursive(node.right, height)
}

func (t *BinaryTree) inorder() {
	if t == nil || t.root == nil {
		return
	}

	inorderRecursive(t.root)
}

func inorderRecursive(n *Node) {
	if n == nil {
		return
	}

	inorderRecursive(n.left)
	fmt.Printf("%d ", n.val)
	inorderRecursive(n.right)
}

// //////////////////////////////////////////////////////////////////////////////
// Utils
// //////////////////////////////////////////////////////////////////////////////
func max(a int64, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a int64, b int64) int64 {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println("Hello, World!")

	t := newBinaryTree()

	for i := 1; i <= 100; i++ {
		t.insert(int64(i))
	}

	// t.insert(1)
	// t.insert(2)
	// t.insert(3)

	t.print()
	t.inorder()

	// fmt.Println("")
	// fmt.Printf("%d\n", findHeight(t.root))
}
