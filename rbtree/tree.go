package rbtree

type Tree struct {
	Root *Node
	Size int
}

func NewTree() *Tree {
	return &Tree{
		Root: nil,
		Size: 0,
	}
}

func (t *Tree) Insert(key T, value interface{}) {
	node := &Node{
		Key:   key,
		Value: value,
		Color: Red,
	}

	if t.Root == nil {
		t.Root = node
		t.Root.Color = Black
		t.Size++
		return
	}

	cur := t.Root
	var parent *Node
	for cur != nil {
		parent = cur
		if key < cur.Key {
			cur = cur.Left
		} else if key > cur.Key {
			cur = cur.Right
		} else {
			// 如果key已存在，更新value
			cur.Value = value
			return
		}
	}

	node.Parent = parent
	if key < parent.Key {
		parent.Left = node
	} else {
		parent.Right = node
	}

	t.Size++
	t.fixInsert(node)
}

func (t *Tree) rotateLeft(node *Node) {
	right := node.Right
	node.Right = right.Left
	if right.Left != nil {
		right.Left.Parent = node
	}
	right.Parent = node.Parent
	if node.Parent == nil {
		t.Root = right
	} else if node == node.Parent.Left {
		node.Parent.Left = right
	} else {
		node.Parent.Right = right
	}
	right.Left = node
	node.Parent = right
	node.Color = Red
	right.Color = Black
}

func (t *Tree) rotateRight(node *Node) {
	left := node.Left
	node.Left = left.Right
	if left.Right != nil {
		left.Right.Parent = node
	}
	left.Parent = node.Parent
	if node.Parent == nil {
		t.Root = left
	} else if node == node.Parent.Left {
		node.Parent.Left = left
	} else {
		node.Parent.Right = left
	}
	left.Right = node
	node.Parent = left
	node.Color = Red
	left.Color = Black
}

func (t *Tree) fixInsert(node *Node) {
	for node.Parent != nil && node.Parent.Color == Red {
		if node.Parent == node.Parent.Parent.Left {
			uncle := node.Parent.Parent.Right
			if uncle != nil && uncle.Color == Red {
				node.Parent.Color = Black
				uncle.Color = Black
				node.Parent.Parent.Color = Red
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Right {
					node = node.Parent
					t.rotateLeft(node)
				}
				node.Parent.Color = Black
				node.Parent.Parent.Color = Red
				t.rotateRight(node.Parent.Parent)
			}
		} else {
			uncle := node.Parent.Parent.Left
			if uncle != nil && uncle.Color == Red {
				node.Parent.Color = Black
				uncle.Color = Black
				node.Parent.Parent.Color = Red
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Left {
					node = node.Parent
					t.rotateRight(node)
				}
				node.Parent.Color = Black
				node.Parent.Parent.Color = Red
				t.rotateLeft(node.Parent.Parent)
			}
		}
	}
	t.Root.Color = Black
}

func (t *Tree) Search(key T) interface{} {
	cur := t.Root
	for cur != nil {
		if key < cur.Key {
			cur = cur.Left
		} else if key > cur.Key {
			cur = cur.Right
		} else {
			return cur.Value
		}
	}
	return nil
}

func (t *Tree) Delete(key T) {
	node := t.findNode(key)
	if node == nil {
		return
	}

	var child *Node
	if node.Left == nil || node.Right == nil {
		child = node.Left
		if child == nil {
			child = node.Right
		}
	} else {
		// 找到后继节点
		successor := t.getMin(node.Right)
		node.Key = successor.Key
		node.Value = successor.Value
		node = successor
		child = node.Right
	}

	if node.Color == Black {
		t.fixDelete(node)
	}

	// 更新父节点引用
	if node.Parent == nil {
		t.Root = child
	} else if node == node.Parent.Left {
		node.Parent.Left = child
	} else {
		node.Parent.Right = child
	}

	if child != nil {
		child.Parent = node.Parent
	}

	t.Size--
}

func (t *Tree) findNode(key T) *Node {
	cur := t.Root
	for cur != nil {
		if key < cur.Key {
			cur = cur.Left
		} else if key > cur.Key {
			cur = cur.Right
		} else {
			return cur
		}
	}
	return nil
}

func (t *Tree) fixDelete(node *Node) {
	for node != t.Root && node.Color == Black {
		if node == node.Parent.Left {
			brother := node.Parent.Right
			if brother == nil {
				break
			}

			if brother.Color == Red {
				brother.Color = Black
				node.Parent.Color = Red
				t.rotateLeft(node.Parent)
				brother = node.Parent.Right
			}

			if brother == nil {
				break
			}

			leftBlack := brother.Left == nil || brother.Left.Color == Black
			rightBlack := brother.Right == nil || brother.Right.Color == Black

			if leftBlack && rightBlack {
				brother.Color = Red
				node = node.Parent
			} else {
				if rightBlack {
					if brother.Left != nil {
						brother.Left.Color = Black
					}
					brother.Color = Red
					t.rotateRight(brother)
					brother = node.Parent.Right
				}

				if brother != nil {
					brother.Color = node.Parent.Color
					node.Parent.Color = Black
					if brother.Right != nil {
						brother.Right.Color = Black
					}
					t.rotateLeft(node.Parent)
				}
				node = t.Root
			}
		} else {
			brother := node.Parent.Left
			if brother == nil {
				break
			}

			if brother.Color == Red {
				brother.Color = Black
				node.Parent.Color = Red
				t.rotateRight(node.Parent)
				brother = node.Parent.Left
			}

			if brother == nil {
				break
			}

			leftBlack := brother.Left == nil || brother.Left.Color == Black
			rightBlack := brother.Right == nil || brother.Right.Color == Black

			if leftBlack && rightBlack {
				brother.Color = Red
				node = node.Parent
			} else {
				if leftBlack {
					if brother.Right != nil {
						brother.Right.Color = Black
					}
					brother.Color = Red
					t.rotateLeft(brother)
					brother = node.Parent.Left
				}

				if brother != nil {
					brother.Color = node.Parent.Color
					node.Parent.Color = Black
					if brother.Left != nil {
						brother.Left.Color = Black
					}
					t.rotateRight(node.Parent)
				}
				node = t.Root
			}
		}
	}
	if node != nil {
		node.Color = Black
	}
}

func (t *Tree) getMin(node *Node) *Node {
	for node.Left != nil {
		node = node.Left
	}
	return node
}
