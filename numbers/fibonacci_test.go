package numbers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFib(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{
			name:  "Example 1",
			input: 2,
			want:  1,
		},
		{
			name:  "Example 2",
			input: 3,
			want:  2,
		},
		{
			name:  "Example 3",
			input: 4,
			want:  3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fib(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}
