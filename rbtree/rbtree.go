package rbtree

import (
	"fmt"
)

type RBTree struct {
	root    *Node
	nilNode *Node
}

func NewRBTree() *RBTree {
	nilNode := &Node{color: BLACK}
	return &RBTree{
		root:    nilNode,
		nilNode: nilNode,
	}
}

func (t *RBTree) LeftRotate(x *Node) {
	y := x.right
	x.right = y.left
	if y.left != t.nilNode {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == t.nilNode {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func (t *RBTree) RightRotate(y *Node) {
	x := y.left
	y.left = x.right
	if x.right != t.nilNode {
		x.right.parent = y
	}
	x.parent = y.parent
	if y.parent == t.nilNode {
		t.root = x
	} else if y == y.parent.left {
		y.parent.left = x
	} else {
		y.parent.right = x
	}
	x.right = y
	y.parent = x
}

func (t *RBTree) Insert(key float64, index int) {
	parent := t.nilNode
	current := t.root
	for current != t.nilNode {
		parent = current
		if key < current.key {
			current = current.left
		} else if key > current.key {
			current = current.right
		} else {
			current.indexes = append(current.indexes, index)
			return
		}
	}

	newNode := &Node{
		key:     key,
		indexes: []int{index},
		left:    t.nilNode,
		right:   t.nilNode,
		parent:  parent,
		color:   RED,
	}

	if parent == t.nilNode {
		t.root = newNode
	} else if key < parent.key {
		parent.left = newNode
	} else {
		parent.right = newNode
	}

	t.insertFixup(newNode)
}

func (t *RBTree) insertFixup(z *Node) {
	for z.parent.color == RED {
		if z.parent == z.parent.parent.left {
			uncle := z.parent.parent.right
			if uncle.color == RED {
				z.parent.color = BLACK
				uncle.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					t.LeftRotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.RightRotate(z.parent.parent)
			}
		} else {
			uncle := z.parent.parent.left
			if uncle.color == RED {
				z.parent.color = BLACK
				uncle.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					t.RightRotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.LeftRotate(z.parent.parent)
			}
		}
	}
	t.root.color = BLACK
}

func (t *RBTree) RangeQuery(low, high float64) []int {
	result := []int{}
	t.rangeQueryRecursive(t.root, low, high, &result)
	return result
}

func (t *RBTree) rangeQueryRecursive(node *Node, low, high float64, result *[]int) {
	if node == t.nilNode {
		return
	}

	if node.key > low {
		t.rangeQueryRecursive(node.left, low, high, result)
	}

	if node.key > low && node.key < high {
		*result = append(*result, node.indexes...)
	}

	if node.key < high {
		t.rangeQueryRecursive(node.right, low, high, result)
	}
}

func FindSubarrays(seq []float64, T, e float64) [][2]int {
	n := len(seq)
	S := make([]float64, n+1)
	S[0] = 0
	for i := 1; i <= n; i++ {
		S[i] = S[i-1] + seq[i-1]
	}

	tree := NewRBTree()
	tree.Insert(0, 0)

	var results [][2]int

	for j := 1; j <= n; j++ {
		sum := S[j]
		low := sum - T - e
		high := sum - T + e

		indexes := tree.RangeQuery(low, high)
		for _, k := range indexes {
			results = append(results, [2]int{k + 1, j})
		}

		tree.Insert(sum, j)
	}

	return results
}

// 最终正确版本
func (t *RBTree) PrintTree() {
	if t.root == t.nilNode {
		fmt.Println("Empty Red-Black Tree")
		return
	}
	fmt.Println("Red-Black Tree Structure:")
	t.printNode(t.root, "", true)
}

func (t *RBTree) printNode(node *Node, prefix string, isTail bool) {
	if node == t.nilNode {
		return
	}

	newPrefix := prefix
	if isTail {
		fmt.Printf("%s└── ", prefix)
		newPrefix += "    "
	} else {
		fmt.Printf("%s├── ", prefix)
		newPrefix += "│   "
	}

	colorStr := "RED"
	if !node.color {
		colorStr = "BLACK"
	}
	parentKey := "nil"
	if node.parent != t.nilNode {
		parentKey = fmt.Sprintf("%.2f", node.parent.key)
	}
	fmt.Printf("[%.2f] %s (Parent: %s)\n", node.key, colorStr, parentKey)

	t.printNode(node.left, newPrefix, false)
	t.printNode(node.right, newPrefix, true)
}
