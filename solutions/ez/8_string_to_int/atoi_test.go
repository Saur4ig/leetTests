package __string_to_int

import (
	"testing"
)

func Test_myAtoi(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Example 1",
			s:    "  1 ",
			want: 1,
		},
		{
			name: "Example 2",
			s:    "692",
			want: 692,
		},
		{
			name: "Example 3",
			s:    "-912",
			want: -912,
		},
		{
			name: "Example 4",
			s:    "009",
			want: 9,
		},
		{
			name: "Example 5",
			s:    "9223372036854775808",
			want: 2147483647,
		},
		{
			name: "Example 6",
			s:    "18446744073709551617",
			want: 2147483647,
		},
		{
			name: "Example 7",
			s:    "  0000000000012345678",
			want: 12345678,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
