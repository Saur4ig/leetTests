package _9

import (
	"testing"
)

func Test_mySqrt(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{
			name: "Example 1",
			x:    16,
			want: 4,
		},
		{
			name: "Example 2",
			x:    4,
			want: 2,
		},
		{
			name: "Example 3",
			x:    8,
			want: 2,
		},
		{
			name: "Example 4",
			x:    6,
			want: 2,
		},
		{
			name: "Example 5",
			x:    36,
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
