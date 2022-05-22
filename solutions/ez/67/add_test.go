package _7

import (
	"testing"
)

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				a: "11",
				b: "1",
			},
			want: "100",
		},
		{
			name: "Example 2",
			args: args{
				a: "1010",
				b: "1011",
			},
			want: "10101",
		},
		{
			name: "Example 3",
			args: args{
				a: "1",
				b: "1",
			},
			want: "10",
		},
		{
			name: "Example 4",
			args: args{
				a: "110010",
				b: "10111",
			},
			want: "1001001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
