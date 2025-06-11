package util

import (
	"testing"
	"time"
)

func TestFormatDate(t *testing.T) {
	type args struct {
		t      time.Time
		layout string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "测试日期格式化: YYYY-MM-DD",
			args: args{
				t:      time.Date(2025, 1, 23, 15, 35, 40, 0, time.Local),
				layout: DateLayout,
			},
			want: "2025-01-23",
		},
		{
			name: "测试日期格式化: YYYY-MM-DD HH:mm:ss",
			args: args{
				t:      time.Date(2025, 1, 23, 15, 35, 40, 0, time.Local),
				layout: HourLayout,
			},
			want: "2025-01-23 15",
		},
		{
			name: "测试日期格式化: YYYY-MM-DD HH:mm",
			args: args{
				t:      time.Date(2025, 1, 23, 15, 35, 40, 0, time.Local),
				layout: MinuteLayout,
			},
			want: "2025-01-23 15:35",
		},
		{
			name: "测试日期格式化: YYYY-MM-DD HH:mm:ss",
			args: args{
				t:      time.Date(2025, 1, 23, 15, 35, 40, 0, time.Local),
				layout: SecondLayout,
			},
			want: "2025-01-23 15:35:40",
		},
		{
			name: "测试日期格式化: YYYYMMDD",
			args: args{
				t:      time.Date(2025, 1, 23, 15, 35, 40, 0, time.Local),
				layout: Y4M2D2,
			},
			want: "20250123",
		},
		{
			name: "测试日期格式化: YYYYMMDDHHmm",
			args: args{
				t:      time.Date(2025, 1, 23, 15, 35, 40, 0, time.Local),
				layout: Y4M2D2H2m2,
			},
			want: "202501231535",
		},
		{
			name: "测试日期格式化: YYYY年MM月DD日",
			args: args{
				t:      time.Date(2025, 1, 23, 15, 35, 40, 0, time.Local),
				layout: CNY4M2D2,
			},
			want: "2025年01月23日",
		},
		{
			name: "测试日期格式化: YYYY年MM月DD日HH点",
			args: args{
				t:      time.Date(2025, 1, 23, 15, 35, 40, 0, time.Local),
				layout: CNY2M2D2H2,
			},
			want: "2025年01月23日15点",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatDate(tt.args.t, tt.args.layout); got != tt.want {
				t.Errorf("FormatDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
