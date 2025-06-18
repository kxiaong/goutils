package main

import (
	"fmt"

	"github.com/kxiaong/goutils/rbtree"
)

func main() {
	seq := []float64{23, 45, 33, 57, 43, 100, 61, 72, 49}
	T := 90.0
	e := 30.0
	fmt.Println("序列:", seq, "目标:", T, "容差:", e)
	fmt.Println("满足条件的子数组:")
	for _, sub := range rbtree.FindSubarrays(seq, T, e) {
		fmt.Printf("[%d, %d]\n", sub[0], sub[1])
	}
	fmt.Println()
}
