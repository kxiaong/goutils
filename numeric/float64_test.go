package util

import (
	"testing"
)

func TestToFixed(t *testing.T) {
	type args struct {
		num       float64
		precision int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "保留2位小数，第3位上取整",
			args: args{num: 123.456, precision: 2},
			want: 123.46,
		},
		{
			name: "保留2位小数，第3位下取整",
			args: args{num: 123.654, precision: 2},
			want: 123.65,
		},
		{
			name: "保留1位小数，第2位上取整",
			args: args{num: 123.456, precision: 1},
			want: 123.5,
		},
		{
			name: "保留1位小数，第2位下取整",
			args: args{num: 123.546, precision: 1},
			want: 123.5,
		},
		{
			name: "保留0位小数，第1位下取整",
			args: args{num: 123.456, precision: 0},
			want: 123,
		},
		{
			name: "保留0位小数，第1位上取整",
			args: args{num: 123.546, precision: 0},
			want: 124,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFixed(tt.args.num, tt.args.precision); got != tt.want {
				t.Errorf("ToFixed() = %v, want %v", got, tt.want)
			}
		})
	}
}
