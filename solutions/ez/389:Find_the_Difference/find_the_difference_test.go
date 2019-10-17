package _89_Find_the_Difference

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindTheDifference(t *testing.T) {
	tests := []struct {
		name   string
		input1 string
		input2 string
		want   byte
	}{
		{
			name:   "Example 1",
			input1: "abcd",
			input2: "abcde",
			want:   []byte("e")[0],
		},
		{
			name:   "Example 1",
			input1: "abcd",
			input2: "abcda",
			want:   []byte("a")[0],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findTheDifference(tt.input1, tt.input2)
			require.Equal(t, tt.want, got)
		})
	}
}
