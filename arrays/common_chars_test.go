package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommonChars(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "Example 1",
			input: []string{"bella", "label", "roller"},
			want:  []string{"e", "l", "l"},
		},
		{
			name:  "Example 2",
			input: []string{"cool", "lock", "cook"},
			want:  []string{"c", "o"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := commonChars(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}
