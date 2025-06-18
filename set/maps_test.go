package set

import (
	"testing"
)

func DeepEqual[T comparable](x, y []T) bool {
	if len(x) != len(y) {
		return false
	}
	m := make(map[T]struct{}, len(x))
	for _, e := range x {
		m[e] = struct{}{}
	}
	for _, e := range y {
		if _, ok := m[e]; !ok {
			return false
		}
	}
	return true
}

func TestMapSubtract(t *testing.T) {
	type args[K comparable, V comparable] struct {
		a map[K]V
		b map[K]V
	}
	type testCase[K comparable, V comparable] struct {
		name string
		args args[K, V]
		want []K
	}
	tests := []testCase[int64, int64]{
		{
			name: "测试a-b, a b 不为空",
			args: args[int64, int64]{
				a: map[int64]int64{
					1: 0,
					2: 0,
					3: 0,
				},
				b: map[int64]int64{
					1: 0,
					2: 0,
				},
			},
			want: []int64{3},
		},
		{
			name: "测试a-b, a 为空",
			args: args[int64, int64]{
				a: map[int64]int64{},
				b: map[int64]int64{
					1: 0,
					2: 0,
				},
			},
			want: []int64{},
		},
		{
			name: "测试a-b, b 为空",
			args: args[int64, int64]{
				a: map[int64]int64{
					1: 0,
					2: 0,
					3: 0,
				},
				b: map[int64]int64{},
			},
			want: []int64{1, 2, 3},
		},
		{
			name: "测试a-b, a b 都为空",
			args: args[int64, int64]{
				a: map[int64]int64{},
				b: map[int64]int64{},
			},
			want: []int64{},
		},
		{
			name: "测试a-b, a b 都不为空",
			args: args[int64, int64]{
				a: map[int64]int64{
					1: 0,
					2: 0,
					3: 0,
				},
				b: map[int64]int64{
					4: 0,
					5: 0,
					6: 0,
				},
			},
			want: []int64{1, 2, 3},
		},
		{
			name: "测试a-b, a b 都不为空",
			args: args[int64, int64]{
				a: map[int64]int64{
					4: 0,
					5: 0,
					6: 0,
				},
				b: map[int64]int64{
					1: 0,
					2: 0,
					3: 0,
				},
			},
			want: []int64{4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapSubtract(tt.args.a, tt.args.b); !DeepEqual[int64](got, tt.want) {
				t.Errorf("MapSubtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapUnion(t *testing.T) {
	type args[K comparable, V comparable] struct {
		a map[K]V
		b map[K]V
	}
	type testCase[K comparable, V comparable] struct {
		name string
		args args[K, V]
		want []K
	}
	tests := []testCase[int64, struct{}]{
		{
			name: "测试a-b, a b 不为空",
			args: args[int64, struct{}]{
				a: map[int64]struct{}{
					1: {},
					2: {},
					3: {},
				},
				b: map[int64]struct{}{
					1: {},
					2: {},
				},
			},
			want: []int64{1, 2, 3},
		},
		{
			name: "测试a-b, a 为空",
			args: args[int64, struct{}]{
				a: map[int64]struct{}{},
				b: map[int64]struct{}{
					1: {},
					2: {},
				},
			},
			want: []int64{1, 2},
		},
		{
			name: "测试a-b, b 为空",
			args: args[int64, struct{}]{
				a: map[int64]struct{}{
					1: {},
					2: {},
					3: {},
				},
				b: map[int64]struct{}{},
			},
			want: []int64{1, 2, 3},
		},
		{
			name: "测试a-b, a b 都为空",
			args: args[int64, struct{}]{
				a: map[int64]struct{}{},
				b: map[int64]struct{}{},
			},
			want: []int64{},
		},
		{
			name: "测试a-b, a b 都不为空",
			args: args[int64, struct{}]{
				a: map[int64]struct{}{
					1: {},
					2: {},
					3: {},
				},
				b: map[int64]struct{}{
					4: {},
					5: {},
					6: {},
				},
			},
			want: []int64{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapUnion(tt.args.a, tt.args.b); !DeepEqual[int64](got, tt.want) {
				t.Errorf("MapUnion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapIntersect(t *testing.T) {
	type args[K comparable, V comparable] struct {
		a map[K]V
		b map[K]V
	}
	type testCase[K comparable, V comparable] struct {
		name string
		args args[K, V]
		want []K
	}
	tests := []testCase[int64, interface{}]{
		{
			name: "测试a&b, a b 不为空",
			args: args[int64, interface{}]{
				a: map[int64]interface{}{
					1: nil,
					2: nil,
					3: nil,
				},
				b: map[int64]interface{}{
					1: nil,
					2: nil,
				},
			},
			want: []int64{1, 2},
		},
		{
			name: "测试a&b, a 为空",
			args: args[int64, interface{}]{
				a: map[int64]interface{}{},
				b: map[int64]interface{}{
					1: nil,
					2: nil,
				},
			},
			want: []int64{},
		},
		{
			name: "测试a&b, b 为空",
			args: args[int64, interface{}]{
				a: map[int64]interface{}{
					1: nil,
					2: nil,
					3: nil,
				},
				b: map[int64]interface{}{},
			},
			want: []int64{},
		},
		{
			name: "测试a&b, a b 都为空",
			args: args[int64, interface{}]{
				a: map[int64]interface{}{},
				b: map[int64]interface{}{},
			},
			want: []int64{},
		},
		{
			name: "测试a&b, a b 都不为空",
			args: args[int64, interface{}]{
				a: map[int64]interface{}{
					1: nil,
					2: nil,
					3: nil,
				},
				b: map[int64]interface{}{
					4: nil,
					5: nil,
					6: nil,
				},
			},
			want: []int64{},
		},
		{
			name: "测试a&b, a b 都不为空",
			args: args[int64, interface{}]{
				a: map[int64]interface{}{
					4: nil,
					5: nil,
					6: nil,
				},
				b: map[int64]interface{}{
					1: nil,
					2: nil,
					3: nil,
				},
			},
			want: []int64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapIntersect(tt.args.a, tt.args.b); !DeepEqual[int64](got, tt.want) {
				t.Errorf("MapIntersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
