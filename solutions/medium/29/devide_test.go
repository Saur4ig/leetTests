package _9

import (
	"testing"
)

func Test_divide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				dividend: 9,
				divisor:  3,
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				dividend: 10,
				divisor:  3,
			},
			want: 3,
		},
		{
			name: "Example 3",
			args: args{
				dividend: 7,
				divisor:  -3,
			},
			want: -2,
		},
		{
			name: "Example 4",
			args: args{
				dividend: -1,
				divisor:  -1,
			},
			want: 1,
		},
		{
			name: "Example 5",
			args: args{
				dividend: 1,
				divisor:  -1,
			},
			want: -1,
		},
		{
			name: "Example 6",
			args: args{
				dividend: -2147483648,
				divisor:  -1,
			},
			want: 2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.dividend, tt.args.divisor); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
