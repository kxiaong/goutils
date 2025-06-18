package set

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	type args[K comparable] struct {
		a []K
	}
	type testCase[K comparable] struct {
		name string
		args args[K]
		want []K
	}
	tests := []testCase[int64]{
		{
			name: "测试去重",
			args: args[int64]{
				a: []int64{1, 2, 3, 1, 2, 3},
			},
			want: []int64{1, 2, 3},
		},
		{
			name: "测试空数组",
			args: args[int64]{
				a: []int64{},
			},
			want: []int64{},
		},
		{
			name: "测试不重复",
			args: args[int64]{
				a: []int64{1, 2, 3},
			},
			want: []int64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueString(t *testing.T) {
	type args struct {
		a []string
	}
	type testCase struct {
		name string
		args args
		want []string
	}
	tests := []testCase{
		{
			name: "测试去重",
			args: args{
				a: []string{"1", "2", "3", "1", "2", "3"},
			},
			want: []string{"1", "2", "3"},
		},
		{
			name: "测试空数组",
			args: args{
				a: []string{},
			},
			want: []string{},
		},
		{
			name: "测试不重复",
			args: args{
				a: []string{"1", "2", "3"},
			},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueString() = %v, want %v", got, tt.want)
			}
		})
	}
}
