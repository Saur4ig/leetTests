package _281_Subtract_sum_of_int

import (
	"testing"
)

func Test_subtractProductAndSum(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Example 1",
			n:    234,
			want: 15,
		},
		{
			name: "Example 2",
			n:    4421,
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtractProductAndSum(tt.n); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
