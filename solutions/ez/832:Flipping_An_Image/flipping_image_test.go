package _32_Flipping_An_Image

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFlipAndInvertImage(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  [][]int
	}{
		{
			name:  "Example 1",
			input: [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}},
			want:  [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}},
		},
		{
			name:  "Example 2",
			input: [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}},
			want:  [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlipAndInvertImage(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}
