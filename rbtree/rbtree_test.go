package rbtree

import (
	"testing"
)

func TestRangeQuery(t *testing.T) {
	seq := []float64{23, 45, 33, 57, 43, 122, 61, 72, 49}

	S := make([]float64, len(seq)+1)
	S[0] = 0
	for i := 1; i <= len(seq); i++ {
		S[i] = S[i-1] + seq[i-1]
	}

	tree := NewRBTree()
	tree.Insert(0, 0)
}
