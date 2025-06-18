package set

import (
	"testing"
)

func TestSliceSubtract(t *testing.T) {
	type args[T comparable] struct {
		a []T
		b []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int64]{
		{
			name: "Case 1: 测试a-b, a b 不为空",
			args: args[int64]{
				a: []int64{1, 2, 3},
				b: []int64{1, 2},
			},
			want: []int64{3},
		},
		{
			name: "Case 2: 测试a-b, a 为空",
			args: args[int64]{
				a: []int64{},
				b: []int64{1, 2},
			},
			want: []int64{},
		},
		{
			name: "Case 3: 测试a-b, b 为空",
			args: args[int64]{
				a: []int64{1, 2},
				b: []int64{},
			},
			want: []int64{1, 2},
		},
		{
			name: "Case 4: 测试a-b, a b 都为空",
			args: args[int64]{
				a: []int64{},
				b: []int64{},
			},
			want: []int64{},
		},
		{
			name: "Case 5: 测试a-b, a b 都不为空",
			args: args[int64]{
				a: []int64{1, 2, 3},
				b: []int64{1, 2, 3},
			},
			want: []int64{},
		},
		{
			name: "Case 6: 测试a-b, a b 都不为空",
			args: args[int64]{
				a: []int64{1, 2, 3},
				b: []int64{4, 5, 6},
			},
			want: []int64{1, 2, 3},
		},
		{
			name: "Case 7: 测试a-b, a b 都不为空",
			args: args[int64]{
				a: []int64{1, 2, 3},
				b: []int64{1, 2, 3, 4, 5, 6},
			},
			want: []int64{},
		},
		{
			name: "Case 8: 测试a-b, a b 都不为空",
			args: args[int64]{
				a: []int64{1, 2, 3, 4, 5, 6},
				b: []int64{1, 2, 3},
			},
			want: []int64{4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceSubtract(tt.args.a, tt.args.b); !DeepEqual[int64](got, tt.want) {
				t.Errorf("SliceSubtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceUnion(t *testing.T) {
	type args[T comparable] struct {
		a []T
		b []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[string]{
		{
			name: "Case 1: 测试a-b, a b 不为空",
			args: args[string]{
				a: []string{"1", "2", "3"},
				b: []string{"1", "2"},
			},
			want: []string{"1", "2", "3"},
		},
		{
			name: "Case 2: 测试a-b, a 为空",
			args: args[string]{
				a: []string{},
				b: []string{"1", "2"},
			},
			want: []string{"1", "2"},
		},
		{
			name: "Case 3: 测试a-b, b 为空",
			args: args[string]{
				a: []string{"1", "2"},
				b: []string{},
			},
			want: []string{"1", "2"},
		},
		{
			name: "Case 4: 测试a-b, a b 都为空",
			args: args[string]{
				a: []string{},
				b: []string{},
			},
			want: []string{},
		},
		{
			name: "Case 5: 测试a-b, a b 都不为空",
			args: args[string]{
				a: []string{"1", "2", "3"},
				b: []string{"1", "2", "3"},
			},
			want: []string{"1", "2", "3"},
		},
		{
			name: "Case 6: 测试a-b, a b 都不为空",
			args: args[string]{
				a: []string{"1", "2", "3"},
				b: []string{"4", "5", "6"},
			},
			want: []string{"1", "2", "3", "4", "5", "6"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceUnion(tt.args.a, tt.args.b); !DeepEqual[string](got, tt.want) {
				t.Errorf("SliceUnion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceIntersect(t *testing.T) {
	type args[T comparable] struct {
		a []T
		b []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[float64]{
		{
			name: "Case 1: 测试a-b, a b 不为空",
			args: args[float64]{
				a: []float64{1.0, 2.0, 3.0},
				b: []float64{1.0, 2.0},
			},
			want: []float64{1.0, 2.0},
		},
		{
			name: "Case 2: 测试a-b, a 为空",
			args: args[float64]{
				a: []float64{},
				b: []float64{1.0, 2.0},
			},
			want: []float64{},
		},
		{
			name: "Case 3: 测试a-b, b 为空",
			args: args[float64]{
				a: []float64{1.0, 2.0},
				b: []float64{},
			},
			want: []float64{},
		},
		{
			name: "Case 4: 测试a-b, a b 都为空",
			args: args[float64]{
				a: []float64{},
				b: []float64{},
			},
			want: []float64{},
		},
		{
			name: "Case 5: 测试a-b, a b 都不为空",
			args: args[float64]{
				a: []float64{1.0, 2.0, 3.0},
				b: []float64{1.0, 2.0, 3.0},
			},
			want: []float64{1.0, 2.0, 3.0},
		},
		{
			name: "Case 6: 测试a-b, a b 都不为空",
			args: args[float64]{
				a: []float64{1.0, 2.0, 3.0},
				b: []float64{4.0, 5.0, 6.0},
			},
			want: []float64{},
		},
		{
			name: "Case 7: 测试a-b, a b 都不为空",
			args: args[float64]{
				a: []float64{1.0, 2.0, 3.0},
				b: []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0},
			},
			want: []float64{1.0, 2.0, 3.0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceIntersect(tt.args.a, tt.args.b); !DeepEqual[float64](got, tt.want) {
				t.Errorf("SliceIntersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
