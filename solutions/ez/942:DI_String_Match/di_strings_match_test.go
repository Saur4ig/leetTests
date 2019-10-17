package _42_DI_String_Match

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiStringMatch(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []int
	}{
		{
			name:  "Example 1",
			input: "IDID",
			want:  []int{0, 4, 1, 3, 2},
		},
		{
			name:  "Example 2",
			input: "III",
			want:  []int{0, 1, 2, 3},
		},
		{
			name:  "Example 3",
			input: "DDI",
			want:  []int{3, 2, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diStringMatch(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}
