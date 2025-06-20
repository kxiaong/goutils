package main

import (
	"fmt"

	"github.com/kxiaong/goutils/rbtree"
)

func main() {
	seq := []float64{23, 45, 33, 57, 43, 100, 61, 72, 49}
	T := 90.0
	e := 30.0

	n := len(seq)
	S := make([]float64, n+1)
	S[0] = 0
	for i := 1; i <= n; i++ {
		S[i] = S[i-1] + seq[i-1]
	}
	fmt.Println(S)

	tree := rbtree.NewRBTree()
	tree.Insert(0, 0)
	for j := 1; j <= n; j++ {
		tree.Insert(S[j], j)
	}
	tree.PrintTree()
	fmt.Println("序列:", seq, "目标:", T, "容差:", e)
	fmt.Println("满足条件的子数组:")
	for _, sub := range rbtree.FindSubarrays(seq, T, e) {
		fmt.Printf("[%d, %d]\n", sub[0], sub[1])
	}
	fmt.Println()
}
