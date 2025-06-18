package slices

import (
	"reflect"
	"testing"
)

func TestInt64InSlice(t *testing.T) {
	type args struct {
		target int64
		s      []int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "测试存在",
			args: args{
				target: 1,
				s:      []int64{1, 2, 3},
			},
			want: true,
		},
		{
			name: "测试不存在",
			args: args{
				target: 4,
				s:      []int64{1, 2, 3},
			},
			want: false,
		},
		{
			name: "测试空切片",
			args: args{
				target: 1,
				s:      []int64{},
			},
			want: false,
		},
		{
			name: "测试空切片",
			args: args{
				target: 1,
				s:      nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64InSlice(tt.args.target, tt.args.s); got != tt.want {
				t.Errorf("Int64InSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64SliceTo2D(t *testing.T) {
	type args struct {
		s      []int64
		arrLen int
	}
	tests := []struct {
		name string
		args args
		want [][]int64
	}{
		{
			name: "测试s为空",
			args: args{
				s:      []int64{},
				arrLen: 2,
			},
			want: [][]int64{},
		},
		{
			name: "测试s不为空",
			args: args{
				s:      []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
				arrLen: 3,
			},
			want: [][]int64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
		{
			name: "测试s不为空",
			args: args{
				s:      []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
				arrLen: 4,
			},
			want: [][]int64{{1, 2, 3, 4}, {5, 6, 7, 8}, {9}},
		},
		{
			name: "测试s不为空",
			args: args{
				s:      []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
				arrLen: 10,
			},
			want: [][]int64{{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64SliceTo2D(tt.args.s, tt.args.arrLen); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64SliceTo2D() = %v, want %v", got, tt.want)
			}
		})
	}
}
