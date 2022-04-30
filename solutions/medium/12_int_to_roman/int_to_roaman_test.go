package _2_int_to_roman

import (
	"testing"
)

func Test_intToRoman(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want string
	}{
		{
			name: "Example 1",
			num:  2,
			want: "II",
		},
		{
			name: "Example 2",
			num:  58,
			want: "LVIII",
		},
		{
			name: "Example 3",
			num:  1994,
			want: "MCMXCIV",
		},
		{
			name: "Example 4",
			num:  94,
			want: "XCIV",
		},
		{
			name: "Example 5",
			num:  10,
			want: "X",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
