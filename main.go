package main

import (
	"fmt"
)

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

type BinaryTree struct {
	root *Node
}

func newBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (t *BinaryTree) add(val int64) {
	if t.root == nil {
		node := Node{
			left:  nil,
			right: nil,
			val:   val,
		}

		t.root = &node
		return
	}

	t.root.addNode(val)

  if t.root.balanceFactor < -1 || t.root.balanceFactor > 1 {
    t.rebalance()
  }
}

func (t *BinaryTree) rebalance() {
  if t == nil || t.root == nil {
    return
  }

  t.root = t.root.rebalance()
}

func (node *Node) addNode(val int64) int64 {
	if node == nil {
		return 0
	}

	if val <= node.val {
		if node.left == nil {
			newNode := newNode()
			newNode.val = val
			node.left = newNode

			return 1
		}

		new_left_height := node.left.addNode(val)
		right_height := findHeight(node.right)

		// fmt.Printf("Left height: %d\n", new_left_height)
		// fmt.Printf("Right height: %d\n", right_height)

		// node.height = max(new_left_height, right_height)
		node.height = findHeight(node)

		node.balanceFactor = right_height - new_left_height
		// fmt.Printf("Node height: %d\n", node.height)

    // if node.balanceFactor < -1 || node.balanceFactor > 1 {
    //   node.rebalance()
    // }

		return node.height
	} else {

		if node.right == nil {
			newNode := newNode()
			newNode.val = val
			node.right = newNode
			return 1
		}

		new_right_height := node.right.addNode(val)
		left_height := findHeight(node.left)

		// fmt.Printf("Left height: %d\n", left_height)
		// fmt.Printf("Right height: %d\n", new_right_height)

		// node.balanceFactor = int32(math.Abs(float64(int32(new_right_height) - int32(left_height))))
		node.balanceFactor = (new_right_height - left_height)
		// node.height = max(left_height, new_right_height)
		node.height = findHeight(node)
		// fmt.Printf("Node height: %d\n", node.height)

    // if node.balanceFactor < -1 || node.balanceFactor > 1 {
    //   node.rebalance()
    // }

		return node.height
	}
}

func findHeight(node *Node) int64 {
	if node == nil || node.left == nil && node.right == nil {
		return 0
	}

	lh := findHeight(node.left)
	rh := findHeight(node.right)

	// if node.val == 7 {
	//   fmt.Printf("lh: %d, rh: %d\n", lh, rh)
	// }

	m := 1 + max(lh, rh)

	return m
}

func (n *Node) getHeight() int64 {
	if n == nil {
		return 0
	}

	return n.height
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
	fmt.Printf("rebalance %d\n", n.val)
	// n.Dump(0, "")
	switch {
	case n.balanceFactor < -1 && n.left.balanceFactor == -1:
		return n.rotateRight()
	case n.balanceFactor > 1 && n.right.balanceFactor == 1:
		return n.rotateLeft()
	case n.balanceFactor < -1 && n.left.balanceFactor == 1:
		return n.rotateLeftRight()
	case n.balanceFactor > 1 && n.right.balanceFactor == -1:
		return n.rotateRightLeft()
	}
	return n
}

func balanceTree(node *Node) {
}

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

// func (t *BinaryTree) prettyPrint() {
//   if t.root == nil {
//     return
//   }
//
// }
//
// func ppRecursive(n *Node, height int) {
//
// }

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

func main() {
	fmt.Println("Hello, World!")

	t := newBinaryTree()
	t.add(10)
	t.print()
	t.add(8)
	t.print()
	t.add(7)
	t.print()
	t.add(6)
	t.print()
	t.add(12)
	t.print()
	t.add(14)
	t.print()


  t.inorder()
	// fmt.Printf("%+v\n", t.find(6))
	//
	// t.remove(6)
	// t.remove(14)
	//
	// t.print()
}
