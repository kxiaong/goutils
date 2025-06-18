package rbtree

const (
	RED   = true
	BLACK = false
)

type Node struct {
	key     float64
	indexes []int
	color   bool
	left    *Node
	right   *Node
	parent  *Node
}
