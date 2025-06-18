package slices

import (
	"reflect"
	"testing"
)

func TestStringInSlice(t *testing.T) {
	type args struct {
		target string
		slice  []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "测试字符串是否在切片中",
			args: args{
				target: "apple",
				slice:  []string{"apple", "banana", "orange"},
			},
			want: true,
		},
		{
			name: "测试字符串不在切片中",
			args: args{
				target: "grape",
				slice:  []string{"apple", "banana", "orange"},
			},
			want: false,
		},
		{
			name: "测试空切片",
			args: args{
				target: "apple",
				slice:  []string{},
			},
			want: false,
		},
		{
			name: "测试空字符串",
			args: args{
				target: "",
				slice:  []string{"apple", "banana", "orange"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringInSlice(tt.args.target, tt.args.slice); got != tt.want {
				t.Errorf("StringInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringSliceTo2D(t *testing.T) {
	type args struct {
		s      []string
		arrLen int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "测试切片长度小于arrLen",
			args: args{
				s:      []string{"a", "b", "c"},
				arrLen: 5,
			},
			want: [][]string{{"a", "b", "c"}},
		},
		{
			name: "测试切片长度等于arrLen",
			args: args{
				s:      []string{"a", "b", "c", "d", "e"},
				arrLen: 5,
			},
			want: [][]string{{"a", "b", "c", "d", "e"}},
		},
		{
			name: "测试切片长度大于arrLen",
			args: args{
				s:      []string{"a", "b", "c", "d", "e", "f", "g"},
				arrLen: 3,
			},
			want: [][]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g"}},
		},
		{
			name: "测试arrLen为0",
			args: args{
				s:      []string{"a", "b", "c", "d", "e", "f", "g"},
				arrLen: 0,
			},
			want: nil,
		},
		{
			name: "测试arrLen为负数",
			args: args{
				s:      []string{"a", "b", "c", "d", "e", "f", "g"},
				arrLen: -1,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringSliceTo2D(tt.args.s, tt.args.arrLen); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringSliceTo2D() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomSelectFromStringSlice(t *testing.T) {
	type args struct {
		source []string
		cap    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "测试cap大于source长度",
			args: args{
				source: []string{"a", "b", "c", "d", "e"},
				cap:    10,
			},
			want: 5,
		},
		{
			name: "测试cap等于source长度",
			args: args{
				source: []string{"a", "b", "c", "d", "e"},
				cap:    5,
			},
			want: 5,
		},
		{
			name: "测试随机取3个元素",
			args: args{
				source: []string{"a", "b", "c", "d", "e"},
				cap:    3,
			},
			want: 3,
		},
		{
			name: "测试cap等于0",
			args: args{
				source: []string{"a", "b", "c", "d", "e"},
				cap:    0,
			},
			want: 0,
		},
		{
			name: "测试随机1个元素",
			args: args{
				source: []string{"a", "b", "c", "d", "e"},
				cap:    1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomSelectFromStringSlice(tt.args.source, tt.args.cap)
			if len(got) != tt.want {
				t.Errorf("RandomSelectFromStringSlice() = %v, want %v", got, tt.want)
			}
			for _, e := range got {
				if !StringInSlice(e, tt.args.source) {
					t.Errorf("RandomSelectFromStringSlice() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
