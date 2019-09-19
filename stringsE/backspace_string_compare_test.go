package stringsE

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBackspaceCompare(t *testing.T) {
	tests := []struct {
		name   string
		input1 string
		input2 string
		want   bool
	}{
		{
			name:   "Example 1",
			input1: "ab#c",
			input2: "ad#c",
			want:   true,
		},
		{
			name:   "Example 2",
			input1: "ab##",
			input2: "c#d#",
			want:   true,
		},
		{
			name:   "Example 3",
			input1: "a##c",
			input2: "#a#c",
			want:   true,
		},
		{
			name:   "Example 4",
			input1: "a#c",
			input2: "b",
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := backspaceCompare(tt.input1, tt.input2)
			require.Equal(t, tt.want, got)
		})
	}
}
