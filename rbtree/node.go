package rbtree

type T int

type Color int

const (
	Red Color = iota
	Black
)

type Node struct {
	Key   T
	Value interface{}
	Color Color
	Left  *Node
	Right *Node
}
