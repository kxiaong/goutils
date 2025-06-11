package rbtree

import (
	"reflect"
	"testing"
)

type fields struct {
	Nodes []*Node
}

func TestNewTree(t *testing.T) {
	tests := []struct {
		name   string
		fields *fields
		want   *Tree
	}{
		{
			name: "test",
			fields: &fields{
				Nodes: []*Node{
					{
						Key:    1,
						Value:  "one",
						Color:  Black,
						Left:   nil,
						Right:  nil,
						Parent: nil,
					},
					{
						Key:    2,
						Value:  "two",
						Color:  Red,
						Left:   nil,
						Right:  nil,
						Parent: nil,
					},
					{
						Key:    3,
						Value:  "three",
						Color:  Black,
						Left:   nil,
						Right:  nil,
						Parent: nil,
					},
					{
						Key:    4,
						Value:  "four",
						Color:  Red,
						Left:   nil,
						Right:  nil,
						Parent: nil,
					},
				},
			},
			want: &Tree{
				Root: &Node{
					Key:   1,
					Value: "one",
					Color: Black,
					Left: &Node{
						Key:   2,
						Value: "two",
						Color: Red,
						Left: &Node{
							Key:   3,
							Value: "three",
							Color: Black,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
					Right: &Node{
						Key:    3,
						Value:  "three",
						Color:  Black,
						Left:   nil,
						Right:  nil,
						Parent: nil,
					},
				},
				Size: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTree()
			for _, node := range tt.fields.Nodes {
				got.Insert(node.Key, node.Value)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Delete(t1 *testing.T) {
	type fields struct {
		Root *Node
		Size int
	}
	type args struct {
		key T
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tree{
				Root: tt.fields.Root,
				Size: tt.fields.Size,
			}
			t.Delete(tt.args.key)
		})
	}
}

func TestTree_Insert(t1 *testing.T) {
	type fields struct {
		Root *Node
		Size int
	}
	type args struct {
		key   T
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tree{
				Root: tt.fields.Root,
				Size: tt.fields.Size,
			}
			t.Insert(tt.args.key, tt.args.value)
		})
	}
}

func TestTree_Search(t1 *testing.T) {
	type fields struct {
		Root *Node
		Size int
	}
	type args struct {
		key T
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tree{
				Root: tt.fields.Root,
				Size: tt.fields.Size,
			}
			if got := t.Search(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_findNode(t1 *testing.T) {
	type fields struct {
		Root *Node
		Size int
	}
	type args struct {
		key T
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tree{
				Root: tt.fields.Root,
				Size: tt.fields.Size,
			}
			if got := t.findNode(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("findNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_fixDelete(t1 *testing.T) {
	type fields struct {
		Root *Node
		Size int
	}
	type args struct {
		node *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tree{
				Root: tt.fields.Root,
				Size: tt.fields.Size,
			}
			t.fixDelete(tt.args.node)
		})
	}
}

func TestTree_fixInsert(t1 *testing.T) {
	type fields struct {
		Root *Node
		Size int
	}
	type args struct {
		node *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tree{
				Root: tt.fields.Root,
				Size: tt.fields.Size,
			}
			t.fixInsert(tt.args.node)
		})
	}
}

func TestTree_getMin(t1 *testing.T) {
	type fields struct {
		Root *Node
		Size int
	}
	type args struct {
		node *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tree{
				Root: tt.fields.Root,
				Size: tt.fields.Size,
			}
			if got := t.getMin(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("getMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_rotateLeft(t1 *testing.T) {
	type fields struct {
		Root *Node
		Size int
	}
	type args struct {
		node *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tree{
				Root: tt.fields.Root,
				Size: tt.fields.Size,
			}
			t.rotateLeft(tt.args.node)
		})
	}
}

func TestTree_rotateRight(t1 *testing.T) {
	type fields struct {
		Root *Node
		Size int
	}
	type args struct {
		node *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tree{
				Root: tt.fields.Root,
				Size: tt.fields.Size,
			}
			t.rotateRight(tt.args.node)
		})
	}
}
